package mapping

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tokenbankteam/tb_common/gid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

var m *Mapping

func init() {
	var err error
	m, err = NewMapping(&AddressConfig{
		MongoAddr: "mongodb://127.0.0.1:27017",
		Database:  "mapping",
		Col:       "addr",
		GidAddr:   "http://127.0.0.1:8082/",
		CacheSize: 1024,
	})
	if err != nil {
		fmt.Println(err)
	}
}
func TestInsert(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Errorf("mongo connect error %v", err)
	}
	database := client.Database("mapping")
	col := database.Collection("addr")
	gidServer := gid.NewServer("http://127.0.0.1:8082/")
	id, err := gidServer.GetId()
	if err != nil {
		t.Error(err)
	}
	insertResult, err := col.InsertOne(context.Background(), bson.M{"n": "test", "uk": id})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(insertResult)
	t.Log(string(bytes))
	item, err := m.GetUkByItem("test")
	if err != nil {
		t.Error(err)
	}
	marshal, _ := json.Marshal(item)
	fmt.Println(string(marshal))
}
func TestGetUkByItem(t *testing.T) {
	result, err := m.GetUkByItem("test0")
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetUkByItemCheckExist(t *testing.T) {
	result, err := m.GetUkByItemCheckExist("test")
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetUkByItemList(t *testing.T) {
	params := []string{"test1", "test", "test2"}
	result, err := m.GetUkByItems(params...)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetUkByItemListCheckExist(t *testing.T) {
	params := make(map[string]interface{}, 0)
	params["test"] = 0
	params["test"] = 0
	params["test"] = 0
	params["test"] = 0
	result, err := m.GetUkByItemListCheckExist(params)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetItemByUk(t *testing.T) {
	result, err := m.GetItemByUk(389102355586482177)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetItemByUkList(t *testing.T) {
	params := make(map[int64]interface{}, 0)
	params[389102355586482177] = 0
	params[389102378000842753] = 0
	params[389102386574000129] = 0
	params[389103705900711937] = 0
	result, err := m.GetItemByUkList(params)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}
