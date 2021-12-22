package gid

import (
	"encoding/json"
	"errors"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var client = &http.Client{}

func init() {
	client.Timeout = time.Second
}

type Server struct {
	UrlPrefix string
}

type Result struct {
	Id        int64 `json:"id"`
	MachineId int64 `json:"machine-id"`
	Msb       int64 `json:"msb"`
	Sequence  int64 `json:"sequence"`
	Timestamp int64 `json:"time"`
}

func NewServer(urlPrefix string) *Server {
	if urlPrefix == "" {
		urlPrefix = "http://sonyflake.live.xunlei.com/"
	}
	var S = &Server{
		UrlPrefix: urlPrefix,
	}
	return S
}

func (s *Server) GetId() (int64, error) {
	var err error
	var result *Result
	var reties int64 = 6
	for {
		if reties <= 0 {
			break
		}
		reties--
		result, err = s.Get()
		if err != nil || result == nil {
			log.Errorf("get error %v", err)
			time.Sleep(time.Millisecond * 100)
		} else {
			return result.Id, nil
		}
	}
	return -1, err
}

func (s *Server) Get() (*Result, error) {
	req, err := http.NewRequest("POST", s.UrlPrefix, strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("resp is nil")
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if status := resp.StatusCode; status < 200 || status >= 300 {
		return nil, errors.New("status code is " + strconv.Itoa(status))
	}
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		ret := &Result{}
		if err = json.Unmarshal([]byte(body), ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
