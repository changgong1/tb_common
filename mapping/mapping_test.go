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

var m *Model

func init() {
	var err error
	m, err = NewMapping(&AddressConfig{
		MongoAddr: "mongodb://127.0.0.1:27017",
		Database:  "mapping",
		Col:       "addr",
		GidAddr:   "http://127.0.0.1:8082/",
	})
	if err != nil {
		fmt.Println(err)
	}
}
func TestGetUkByItem(t *testing.T) {
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

func TestGetUkByItemCheckExist(t *testing.T) {
	result, err := m.GetUkByItemCheckExist("test2")
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetUkByItemList(t *testing.T) {
	params := []string{"test", "test1", "test2"}
	result, err := m.GetUkByItemList(params)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetUkByItemListCheckExist(t *testing.T) {
	params := []string{"test", "test1", "test2", "test3", "test4"}
	result, err := m.GetUkByItemListCheckExist(params)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, _ := json.Marshal(result)
	t.Log(string(marshal))
}

func TestGetItemByUk(t *testing.T) {

}

func TestGetItemByUkList(t *testing.T) {

}
