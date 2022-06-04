package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int32  `json:"follow_count"`
	FollowerCount int32  `json:"follower_count"`
}
