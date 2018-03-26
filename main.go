package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gogateway/config"
	"gogateway/continar"
	"gogateway/route"
	"net/http"
	"time"
)

var (
	//单api接口服务最大请求数量，超出后返回
	maxRequestNum = 1000
)

func subServiceEvent(c redis.Conn) (err error) {
	psc := redis.PubSubConn{Conn: c}

	if err := psc.Subscribe("test"); err != nil {
		return err
	}

	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				continar.SetMsg(string(v.Data))
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				return
			}
		}
	}()

	//健康检查
	healthTicker := time.NewTicker(time.Minute)
	defer healthTicker.Stop()
loop:
	for err == nil {
		select {
		case <-healthTicker.C:
			if err = psc.Ping(""); err != nil {
				fmt.Print("redis conn broke")
				//要不要尝试重连
				break loop
			}
		}
	}

	return err
}

func main() {
	config.ParseConfig("config.yaml")
	//订阅服务事件,拉取服务
	c, err := redis.Dial(config.RConf.RedisNetWork, config.RConf.RedisAddr, redis.DialReadTimeout(time.Minute+10*time.Second),
		redis.DialWriteTimeout(10*time.Second))
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()
	go subServiceEvent(c)

	//自定义路由
	handler := route.Mapper{}

	server := &http.Server{
		Addr:         ":8765",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      handler,
	}

	server.ListenAndServe()
}
