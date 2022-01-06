package gid

import (
	"fmt"
	"testing"
)

var S = &Server{
	UrlPrefix: "http://127.0.0.1:8082/",
}

func TestGetId(t *testing.T) {
	result, err := S.GetId()
	if err != nil {
		t.Fatalf("get userRoomStat error, %v", err)
	}
	fmt.Println(result)
}
