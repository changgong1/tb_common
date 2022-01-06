package mapping

import (
	"github.com/bluele/gcache"
	"time"
)

type Mapping struct {
	model       *Model
	NCache      gcache.Cache
	UkCache     gcache.Cache
	cacheSwitch bool
}

func NewMapping(config *AddressConfig) (*Mapping, error) {
	m := &Mapping{}
	model, err := newModel(config)
	if err != nil {
		return nil, err
	}
	m.model = model
	if config.CacheSize > 0 {
		m.NCache = gcache.New(config.CacheSize).Expiration(time.Duration(config.Expire)).LRU().LoaderFunc(func(i interface{}) (interface{}, error) {
			return m.model.getUkByItemCheckExist(i.(string))
		}).Build()
		m.UkCache = gcache.New(config.CacheSize).Expiration(time.Duration(config.Expire)).LRU().LoaderFunc(func(i interface{}) (interface{}, error) {
			return m.model.getItemByUk(i.(int64))
		}).Build()
		m.cacheSwitch = true
	}
	return m, nil
}

func (s *Mapping) GetUkByItem(n string) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getUkByItem(n)
}

func (s *Mapping) GetUkByItemCheckExist(n string) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getUkByItemCheckExist(n)
}

func (s *Mapping) GetUkByItemList(ns []string) ([]*Item, error) {
	return s.model.getUkByItemList(ns)
}

func (s *Mapping) GetUkByItemListCheckExist(ns []string) ([]*Item, error) {
	return s.model.getUkByItemListCheckExist(ns)
}

func (s *Mapping) GetUkByItems(ns ...string) ([]*Item, error) {
	return s.model.getUkByItems(ns...)
}

func (s *Mapping) GetItemByUk(uk int64) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.UkCache.Get(uk)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getItemByUk(uk)
}

func (s *Mapping) GetItemByUkList(uks []int64) ([]*Item, error) {
	return s.model.getItemByUkList(uks)
}

func (s *Mapping) GetItemsByUks(uks ...int64) ([]*Item, error) {
	return s.model.getItemByUks(uks...)
}
