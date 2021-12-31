package mapping

import (
	"errors"
	log "github.com/cihub/seelog"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tokenbankteam/tb_common/gid"
	"time"
)

type Item struct {
	N  string
	Uk int64
}

type Model struct {
	Config   *AddressConfig // 配置信息
	Session  *mgo.Session
	Database *mgo.Database
	Col      *mgo.Collection
	IdPool   chan int64
}

func NewMapping(config *AddressConfig) (*Model, error) {
	session, err := mgo.Dial(config.MongoAddr)
	if err != nil {
		log.Errorf("get session error: %v", err)
		return nil, err
	}
	database := session.DB(config.Database)
	col := database.C(config.Col)
	gidServer := gid.NewServer(config.GidAddr)
	model1 := &Model{
		Session:  session,
		Database: database,
		Col:      col,
		IdPool:   make(chan int64, config.IdPoolSize),
	}
	go func() {
		for {
			if len(model1.IdPool) < config.IdPoolSize {
				result, err := gidServer.Get()
				if err != nil || result == nil {
					log.Errorf("gid get id error %v", err)
					time.Sleep(time.Millisecond * 10)
					continue
				}
				model1.IdPool <- result.Id
			} else {
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()
	return model1, nil
}

func (s *Model) GetUkByItem(n string) (*Item, error) {
	item := new(Item)
	if err := s.Col.Find(bson.M{"n": n}).All(&item); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		log.Errorf("get uk by item n %v error %v", n, err)
		return nil, err
	}
	return item, nil
}

func (s *Model) GetUkByItemCheckExist(n string) (*Item, error) {
	item := new(Item)
	if err := s.Col.Find(bson.M{"n": n}).All(&item); err != nil {
		if err == mgo.ErrNotFound {
			id := <-s.IdPool
			if err = s.Col.Insert(bson.M{"n": n, "uk": id}); err != nil {
				return nil, err
			}
			return &Item{N: n, Uk: id}, nil
		}
		log.Errorf("get uk by item n %v error %v", n, err)
		return nil, err
	}
	return item, nil
}

func (s *Model) GetUkByItemList(ns []string) ([]*Item, error) {
	if len(ns) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	list := make([]*Item, 0)
	if err := s.Col.Find(bson.M{"n": bson.M{"$in": ns}}).All(&list); err != nil {
		log.Errorf("get uk by item list %v error %v", ns, err)
		return nil, err
	}
	return list, nil
}

func (s *Model) GetItemByUk(uk int64) (*Item, error) {
	item := new(Item)
	if err := s.Col.Find(bson.M{"uk": uk}).All(&item); err != nil {
		log.Errorf("get item by uk error %v", err)
		return nil, err
	}
	return item, nil
}

func (s *Model) GetItemByUkList(uks []int64) ([]*Item, error) {
	if len(uks) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	list := make([]*Item, 0)
	if err := s.Col.Find(bson.M{"uk": bson.M{"$in": uks}}).All(&list); err != nil {
		log.Errorf("get item by uk list %v error %v", uks, err)
		return nil, err
	}
	return list, nil
}
