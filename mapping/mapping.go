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
			return m.model.getItemByNCheckExist(i.(string))
		}).Build()
		m.UkCache = gcache.New(config.CacheSize).Expiration(time.Duration(config.Expire)).LRU().LoaderFunc(func(i interface{}) (interface{}, error) {
			return m.model.getItemByUk(i.(int64))
		}).Build()
		m.cacheSwitch = true
	}
	return m, nil
}

// Deprecated: Use GetItemByN instead.
func (s *Mapping) GetUkByItem(n string) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getItemByN(n)
}

func (s *Mapping) GetItemByN(n string) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getItemByN(n)
}

// Deprecated: Use GetItemByNCheckExist instead.
func (s *Mapping) GetUkByItemCheckExist(n string) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getItemByNCheckExist(n)
}

func (s *Mapping) GetItemByNCheckExist(n string) (*Item, error) {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item), nil
		}
	}
	return s.model.getItemByNCheckExist(n)
}

func (s *Mapping) GetUkByN(n string) int64 {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item).Uk
		}
	}
	item, err := s.model.getItemByN(n)
	if err != nil || item == nil {
		return 0
	}
	return item.Uk
}

func (s *Mapping) GetUkByNCheckExist(n string) int64 {
	if s.cacheSwitch {
		result, err := s.NCache.Get(n)
		if err == nil && result != nil {
			return result.(*Item).Uk
		}
	}
	item, err := s.model.getItemByNCheckExist(n)
	if err != nil || item == nil {
		return 0
	}
	return item.Uk
}

// Deprecated: Use GetItemListByNs instead.
func (s *Mapping) GetUkByItemList(params map[string]interface{}) ([]*Item, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	ns := make([]string, 0)
	for k := range params {
		ns = append(ns, k)
	}
	return s.model.getItemListByNs(ns)
}

func (s *Mapping) GetItemListByNs(params map[string]interface{}) ([]*Item, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	ns := make([]string, 0)
	for k := range params {
		ns = append(ns, k)
	}
	return s.model.getItemListByNs(ns)
}

// Deprecated: Use GetUksByNsCheckExist instead.
func (s *Mapping) GetUkByItemListCheckExist(params map[string]interface{}) (map[string]int64, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	ns := make([]string, 0)
	for k := range params {
		ns = append(ns, k)
	}
	return s.model.getUksByNsCheckExist(ns)
}

func (s *Mapping) GetUksByNsCheckExist(params map[string]interface{}) (map[string]int64, error) {
	if len(params) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	ns := make([]string, 0)
	for k := range params {
		ns = append(ns, k)
	}
	return s.model.getUksByNsCheckExist(ns)
}

// Deprecated: Use GetUksByNs instead.
func (s *Mapping) GetUkByItems(ns ...string) (map[string]int64, error) {
	if len(ns) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	return s.model.getUksByNs(ns...)
}

func (s *Mapping) GetUksByNs(ns ...string) (map[string]int64, error) {
	if len(ns) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	return s.model.getUksByNs(ns...)
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

func (s *Mapping) GetNByUk(uk int64) string {
	if s.cacheSwitch {
		result, err := s.UkCache.Get(uk)
		if err == nil && result != nil {
			return result.(*Item).N
		}
	}
	item, err := s.model.getItemByUk(uk)
	if err != nil || item == nil {
		return ""
	}
	return item.N
}

// Deprecated: Use GetItemListByUks instead.
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
	return s.model.getItemListByUks(uks)
}

func (s *Mapping) GetItemListByUks(params map[int64]interface{}) ([]*Item, error) {
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
	return s.model.getItemListByUks(uks)
}

// Deprecated: Use GetNsByUks instead.
func (s *Mapping) GetItemsByUks(uks ...int64) (map[int64]string, error) {
	if len(uks) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	return s.model.getNsByUks(uks...)
}

func (s *Mapping) GetNsByUks(uks ...int64) (map[int64]string, error) {
	if len(uks) > 500 {
		return nil, errors.New("params maximum limit exceeded")
	}
	return s.model.getNsByUks(uks...)
}
