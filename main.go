package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"os"
)

type RedisWriter struct {
	conn *redis.Client
}

var rw RedisWriter

func (r RedisWriter) Write(in []byte) (int, error) {
	data := make(map[string]interface{})
	json.Unmarshal(in, &data)
	if data["level"] == "error" {
		os.Stdout.Write([]byte("error happened"))
	}
	r.conn.HMSet("logrus", data)

	return len(in), nil
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
	//log.SetOutput(os.Stdout)
	rw = RedisWriter{}
	rw.conn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	log.SetOutput(rw)
}

func main() {
	entry := log.WithFields(
		log.Fields{
			"name":  "Orange",
			"color": "orange",
		},
	)

	entry.Info("Hello Apple, I am Orange")
	entry.Warn("Orange alarm")
	entry.Error("Red alarm")
}
