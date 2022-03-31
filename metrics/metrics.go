package metrics

import (
	log "github.com/cihub/seelog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"os"
	"time"
)

var DefaultRegistry *prometheus.Registry

func init() {
	DefaultRegistry = prometheus.NewRegistry()
}

var pusher *push.Pusher

func Init(gateway string, jobName string) {
	hostname, _ := os.Hostname()
	pusher = push.New(gateway, jobName).
		Gatherer(DefaultRegistry).
		Grouping("instance", hostname)

	go func() {
		ticker := time.Tick(time.Second * 5)
		for {
			select {
			case <-ticker:
				if err := pushMetrics(); err != nil {
					log.Errorf("push metrics error %v", err)
				}
			}
		}
	}()
}

func pushMetrics() error {
	if err := pusher.Add(); err != nil {
		return err
	}
	return nil
}
