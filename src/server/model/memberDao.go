package model

import (
	"awesomeProject/src/common/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type MemberDao struct {
	rdb *redis.Client
}

func NewMemberDao() *MemberDao {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "tomo.ink:6379",
		Password: "",
		DB:       0,
	})
	return &MemberDao{rdb: rdb}
}

func (receiver *MemberDao) info(mobile string) (member model.Member, err error) {
	memberStr, err := receiver.rdb.Get(ctx, mobile).Result()
	if err != nil {
		err = MEMBER_NOT_EXIST
	} else {
		err = json.Unmarshal([]byte(memberStr), &member)
	}
	return
}
func (receiver *MemberDao) Login(member *model.Member) (err error) {
	info, err := receiver.info(member.Mobile)
	if err != nil {
		return
	}
	if info.Pwd != member.Pwd {
		err = MEMBER_PWD_ERR
	}
	return
}
func (receiver *MemberDao) Register(member *model.Member) (err error) {
	memberStr, err := json.Marshal(*member)
	if err != nil {
		return
	}
	err = receiver.rdb.Set(ctx, member.Mobile, string(memberStr), 0).Err()
	return
}
