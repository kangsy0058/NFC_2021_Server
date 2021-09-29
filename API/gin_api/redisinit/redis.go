package redisinit

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "hoseolab420",
		DB:       0,
	})

	return rdb
}

// 웨어러블 시리얼 넘버를 이용하여 캐싱된 유저데이터가 있는지 확인후 가져옴
func GetUserData(sn string) (string, error) {
	cli := RedisClient()
	ctx := context.Background()

	// redis 검색 결과확인후 반환
	status, err := cli.Get(ctx, sn).Result()
	switch {
	case err == redis.Nil:
		return "", errors.New("key dose not exist")
	case err != nil:
		return "", err
	case status == "":
		return "", errors.New("value is empty")
	default:
		return status, nil
	}
}

// 웨어러블 시리얼 넘버와 상태정보를 캐싱정보에 넣음
func SetUserData(sn string, status string) error {
	cli := RedisClient()
	ctx := context.Background()

	err := cli.Set(ctx, sn, status, 0)
	if err != nil {
		return errors.New("data not inserted")
	}
	return nil
}
