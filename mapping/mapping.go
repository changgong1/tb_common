package mapping

import (
	"errors"
	log "github.com/cihub/seelog"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tokenbankteam/tb_common/gid"
)

type Item struct {
	N  string
	Uk int64
}

type Model struct {
	Session   *mgo.Session
	Database  *mgo.Database
	Col       *mgo.Collection
	GidServer *gid.Server
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
	return &Model{
		Session:   session,
		Database:  database,
		Col:       col,
		GidServer: gidServer,
	}, nil
}

func (s *Model) GetUkByItem(n string) (*Item, error) {
	item := new(Item)
	if err := s.Col.Find(bson.M{"n": n}).All(&item); err != nil {
		if err == mgo.ErrNotFound {
			result, err := s.GidServer.Get()
			if err != nil {
				return nil, err
			}
			if err = s.Col.Insert(bson.M{"n": n, "uk": result.Id}); err != nil {
				return nil, err
			}
			return &Item{N: n, Uk: result.Id}, nil
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
