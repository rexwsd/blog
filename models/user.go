package models

import (
	"time"
)

type User struct {
	Id           int64     `gorose:"id"`            //数据库id
	Uid          string    `gorose:"snow_id"`       //全局唯一uid 雪花算法生成
	Pwd          string    `gorose:"-"`             //密码 md5
	Email        string    `gorose:"email"`         //邮箱
	Avatar       string    `gorose:"avatar"`        //头像
	Level        int8      `gorose:"level"`         //用户级别
	Name         string    `gorose:"name"`          //用户名
	RegisterDate string    `gorose:"register_date"` //注册日期
	CreatedAt    time.Time `gorose:"created_at"`    //数据生成时间
	UpdatedAt    time.Time `gorose:"updated_at"`    //数据最后一次更新时间
}

func (u *User) TableName() string {
	return "users"
}
