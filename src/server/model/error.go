package model

import "errors"

var (
	MEMBER_NOT_EXIST = errors.New("member not exist")
	MEMBER_PWD_ERR   = errors.New("member pwd err")
)
