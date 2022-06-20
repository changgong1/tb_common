package queue

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/tokenbankteam/tb_common/cache"
)

// Consumer a msg consumer
type Consumer struct {
	topic    string // a topic
	bodyKey  string // the topic content
	cache    *cache.Cache
	waitTime int
	Msg      chan string
}

// CreateConsumer create set type producer
func (t *Topic) CreateConsumer() *Consumer {
	sBodyKey := CreateSetBodyKey(t.topic)
	msg := make(chan string)
	return &Consumer{
		topic:   t.topic,
		bodyKey: sBodyKey,
		cache:   t.cache,
		Msg:     msg,
	}
}

// OnceConsumeMsgWaitTime wait a minute
func (c *Consumer) OnceConsumeMsgWaitTime() (string, error) {
	time.Sleep(TIME_WAIT * time.Millisecond)
	topic, err := c.cache.BLPop(time.Duration(0), c.topic).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	if len(topic) < 2 {
		return "", errors.New("get topic failed")
	}
	r, err := c.cache.SPop(c.bodyKey).Result()
	if err != nil {
		return "", err
	}
	return r, nil
}
