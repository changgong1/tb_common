package mapping

import (
	"context"
	"errors"
	log "github.com/cihub/seelog"
	"github.com/tokenbankteam/tb_common/gid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type Item struct {
	N  string
	Uk int64
}

type Model struct {
	Col       *mongo.Collection
	GidServer *gid.Server
}

func NewMapping(config *AddressConfig) (*Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoAddr))
	if err != nil {
		log.Errorf("mongo connect error %v", err)
		return nil, err
	}
	database := client.Database(config.Database)
	col := database.Collection(config.Col)
	gidServer := gid.NewServer(config.GidAddr)
	return &Model{
		Col:       col,
		GidServer: gidServer,
	}, nil
}

func (s *Model) GetUkByItem(n string) (*Item, error) {
	item := new(Item)
	if err := s.Col.FindOne(context.Background(), bson.M{"n": n}).Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Errorf("get uk by item n %v error %v", n, err)
		return nil, err
	}
	return item, nil
}

func (s *Model) GetUkByItemCheckExist(n string) (*Item, error) {
	item := new(Item)
	if err := s.Col.FindOne(context.Background(), bson.M{"n": n}).Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			id, err := s.GidServer.GetId()
			if err != nil {
				return nil, err
			}
			if _, err = s.Col.InsertOne(context.Background(), bson.M{"n": n, "uk": id}); err != nil {
				if !mongo.IsDuplicateKeyError(err) {
					return nil, err
				}
				if err = s.Col.FindOne(context.Background(), bson.M{"n": n}).Decode(&item); err != nil {
					return nil, err
				}
				return item, err
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
	cursor, err := s.Col.Find(context.Background(), bson.M{"n": bson.M{"$in": ns}})
	if err != nil {
		log.Errorf("get uk by item list %v error %v", ns, err)
		return nil, err
	}
	if err = cursor.All(context.Background(), &list); err != nil {
		log.Errorf("get uk by item list %v error %v", ns, err)
		return nil, err
	}

	return list, nil
}

func (s *Model) GetUkByItemListCheckExist(ns []string) ([]*Item, error) {
	if len(ns) > 100 {
		return nil, errors.New("params maximum limit exceeded")
	}
	list := make([]*Item, 0)
	cursor, err := s.Col.Find(context.Background(), bson.M{"n": bson.M{"$in": ns}})
	if err != nil {
		log.Errorf("get uk by item list %v error %v", ns, err)
		return nil, err
	}
	if err = cursor.All(context.Background(), &list); err != nil {
		log.Errorf("get uk by item list %v error %v", ns, err)
		return nil, err
	}
	if len(list) != len(ns) {
		nmap := make(map[string]bool, 0)
		for _, v := range list {
			nmap[v.N] = true
		}
		num := len(ns) - len(list)
		ids, err := s.GidServer.GetIds(num)
		if err != nil {
			return nil, err
		}
		if num != len(ids) {
			return nil, errors.New("get ids error")
		}
		body := make([]interface{}, 0)
		ns1 := make([]string, 0)
		for i, v := range ns {
			if ok := nmap[v]; ok {
				continue
			}
			ns1 = append(ns1, v)
			body = append(body, bson.M{"n": strings.ToLower(v), "uk": ids[i]})
		}
		opts := options.InsertMany().SetOrdered(false)
		_, err = s.Col.InsertMany(context.Background(), body, opts)
		if err != nil && !mongo.IsDuplicateKeyError(err) {
			log.Errorf("insert many error %v", err)
			return nil, err
		}
		items, err := s.GetUkByItemList(ns1)
		if err != nil {
			return nil, err
		}
		list = append(list, items...)
	}
	return list, nil
}

func (s *Model) GetItemByUk(uk int64) (*Item, error) {
	item := new(Item)
	if err := s.Col.FindOne(context.Background(), bson.M{"uk": uk}).Decode(&item); err != nil {
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
	cursor, err := s.Col.Find(context.Background(), bson.M{"uk": bson.M{"$in": uks}})
	if err != nil {
		log.Errorf("get item by uk list %v error %v", uks, err)
		return nil, err
	}
	if err = cursor.All(context.Background(), &list); err != nil {
		log.Errorf("get item by uk list %v error %v", uks, err)
		return nil, err
	}
	return list, nil
}
