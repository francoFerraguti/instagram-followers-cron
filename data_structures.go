package main

import (
	"time"
)

type UserStruct struct {
	ID           int
	Username     string
	FullName     string
	InstagramID  int
	IsFollower   bool
	ImFollowing  bool
	FollowDate   time.Time
	UnfollowDate time.Time
}

func GetMockUser() UserStruct {
	return UserStruct{
		ID:           0,
		Username:     "test",
		FullName:     "Test Name",
		InstagramID:  19832193,
		IsFollower:   false,
		ImFollowing:  false,
		FollowDate:   nil,
		UnfollowDate: nil,
	}
}
