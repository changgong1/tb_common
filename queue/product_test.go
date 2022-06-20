package queue

import (
	"fmt"
	"testing"

	"github.com/tokenbankteam/tb_common/cache"
)

func TestSendMsg(t *testing.T) {
	mp := map[string]cache.RedisInstConfig{}
	mp["master"] = cache.RedisInstConfig{
		Password: "",
		Url:      "127.0.0.1:6379",
	}
	config := cache.RedisConfig{
		Instances: mp,
	}
	topic, err := RegisterTopic("queue_test", &config, "master")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	p := topic.CreateSetProducer()
	p.SendMsg("this is test content")

	c := topic.CreateConsumer()
	for i := 0; i < 1; i++ {
		fmt.Println(c.OnceConsumeMsgWaitTime())
	}
}
