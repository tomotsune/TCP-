package model

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type memberDao struct {
	rdb *redis.Client
}

func newMemberDao() *memberDao {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "tomo.ink:6379",
		Password: "",
		DB:       0,
	})
	return &memberDao{rdb: rdb}
}

func (receiver *memberDao) info(mobile string) (member Member, err error) {
	val, err := receiver.rdb.Get(ctx, mobile).Result()
	if err != nil {
		return
	} else {
		err = json.Unmarshal([]byte(val), &member)
		if err != nil {
			return
		}
		return
	}
}
func (receiver *memberDao) Login(mobile, pwd string) (err error) {
	info, err := receiver.info(mobile)
	if err != nil {
		return
	}
	if info.pwd != pwd {
		err = MEMBER_NOT_EXIST
	}
	return
}
