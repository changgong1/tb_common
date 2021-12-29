package mapping

import (
	log "github.com/cihub/seelog"
	"github.com/globalsign/mgo"
)

type Item struct {
	N  string
	Uk int64
}

type Model struct {
	Session  *mgo.Session
	Database *mgo.Database
	Col      *mgo.Collection
}

func NewMapping(config *AddressConfig) (*Model, error) {
	session, err := mgo.Dial(config.MongoAddr)
	if err != nil {
		log.Errorf("get session error: %v", err)
		return nil, err
	}
	database := session.DB(config.Database)
	col := database.C(config.Col)

	return &Model{
		Session:  session,
		Database: database,
		Col:      col,
	}, nil
}

func (s *Model) GetUkByItem(n string) (*Item, error) {

	return nil, nil
}

func (s *Model) GetUkByItemList(ns []string) ([]*Item, error) {

	return nil, nil
}

func (s *Model) GetItemByUk(uk int64) (*Item, error) {

	return nil, nil
}

func (s *Model) GetItemByUkList(uks []int64) ([]*Item, error) {

	return nil, nil
}

func (s *Model) Insert(item *Item) {

}
