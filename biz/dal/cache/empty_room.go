package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type EmptyRoomCache struct {
	RoomName []string `json:"roomName"`
	LastTime int64    `json:"lastTime"`
}

func SetEmptyRoomCache(ctx context.Context, key string, emptyRoomList []string) {
	emptyRoomCache := &EmptyRoomCache{
		RoomName: emptyRoomList,
		LastTime: time.Now().Unix(),
	}
	emptyRoomJson, err := json.Marshal(emptyRoomCache)
	// 10分钟过期
	err = RedisClient.Set(ctx, key, emptyRoomJson, time.Minute*10).Err()
	if err != nil {
		hlog.Error(err)
	}
}
func GetEmptyRoomCache(ctx context.Context, key string) (emptyRoomList []string, err error) {
	data, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	emptyRoomCache := new(EmptyRoomCache)
	err = json.Unmarshal([]byte(data), &emptyRoomCache)
	if err != nil {
		return nil, err
	}
	emptyRoomList = emptyRoomCache.RoomName
	return
}
func IsExistRoomInfo(ctx context.Context, key string) (exist int64, err error) {
	exist, err = RedisClient.Exists(ctx, key).Result()
	return
}
