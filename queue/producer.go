package queue

import (
	"errors"

	"github.com/tokenbankteam/tb_common/cache"
)

// SendMsg send a msg
type SendMsg interface {
	SendMsg(string) error
}

// Topic is topic
type Topic struct {
	topic string
	cache *cache.Cache
}

// GetTopicName return topic
func (t *Topic) GetTopicName() string {
	return t.topic
}

// RegisterTopic register a topic
func RegisterTopic(sTopic string, config *cache.RedisConfig, dbName string) (*Topic, error) {
	sInTopic := CreateTopic(sTopic)
	caches, err := cache.GetCaches(config)
	if err != nil {
		return nil, err
	}
	cache := caches.GetCache(dbName)
	return &Topic{
		topic: sInTopic,
		cache: cache,
	}, nil
}

// Producer a queue producer
type Producer struct {
	topic   string // a topic
	bodyKey string // the topic content
	cache   *cache.Cache
}

// SetProducer set type content
type SetProducer struct {
	Producer
}

// SendMsg is Set Proudcer send msg
func (s *SetProducer) SendMsg(content string) error {
	effect, err := s.cache.SAdd(s.bodyKey, content).Result()
	if err != nil {
		return err
	}
	if effect == EFFECT_ZERO {
		return errors.New("send msg failed")
	}
	effect, err = s.cache.LPush(s.topic, WRITE_ONE).Result()
	if err != nil {
		s.cache.SRem(s.bodyKey, content)
		return err
	}
	if effect == EFFECT_ZERO {
		s.cache.SRem(s.bodyKey, content)
		return errors.New("send msg failed")
	}
	return nil
}

// CreateSetProducer create set type producer
func (t *Topic) CreateSetProducer() *SetProducer {
	sBodyKey := CreateSetBodyKey(t.topic)
	return &SetProducer{
		Producer: Producer{
			topic:   t.topic,
			bodyKey: sBodyKey,
			cache:   t.cache,
		},
	}
}
