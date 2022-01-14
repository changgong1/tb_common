package mapping

import (
	"errors"
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

func (s *Mapping) GetUkByItemList(params map[string]interface{}) ([]*Item, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	ns := make([]string, 0)
	for k := range params {
		ns = append(ns, k)
	}
	return s.model.getUkByItemList(ns)
}

func (s *Mapping) GetUkByItemListCheckExist(params map[string]interface{}) (map[string]int64, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	ns := make([]string, 0)
	for k := range params {
		ns = append(ns, k)
	}
	return s.model.getUkByItemListCheckExist(ns)
}

func (s *Mapping) GetUkByItems(ns ...string) (map[string]int64, error) {
	if len(ns) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
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

func (s *Mapping) GetItemByUkList(params map[int64]interface{}) ([]*Item, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	uks := make([]int64, 0)
	for k := range params {
		uks = append(uks, k)
	}
	return s.model.getItemByUkList(uks)
}

func (s *Mapping) GetItemsByUks(uks ...int64) (map[int64]string, error) {
	if len(uks) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	return s.model.getItemByUks(uks...)
}
