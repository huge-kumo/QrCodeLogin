package dao

import (
	"QrCodeLogin/pkg/log"
	"context"
	rdx "github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitRedisInstance(confPath string) (err error) {
	// 获取Redis配置信息
	conf := struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	}{}

	body, err := ioutil.ReadFile(confPath + "/redis.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(body, &conf)
	if err != nil {
		return
	}

	// 初始化Redis对象
	redis = rdx.NewClient(&rdx.Options{
		Addr: conf.Address,
		//Password: conf.Password,
	})

	if err = redis.Ping(context.Background()).Err(); err != nil {
		return
	}
	log.Info("Redis数据库启动成功")
	return
}

func GetRedisInstance() (r *rdx.Client) {
	return redis
}
