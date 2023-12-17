// package main

// まだmain fileから分離していない構造体達

import (
	"time"
)

// 予約システムのユーザー情報を格納する構造体
type User struct {
	// IDはUUIDになるから、stringでまとめてしまうのはいささか乱雑な気がする
	Id       string
	Name     string
	Birthday time.Time
}

// 予約システムの部屋情報を格納する構造体
type Room struct {
	Id       string
	Name     string
	Capacity int
}

// 予約システムの予約情報を格納する構造体
type Reservation struct {
	Id        string
	UserId    string
	RoomId    string
	StartTime time.Time
	EndTime   time.Time
}

type UsersResponse struct {
	Users User `json:"users"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
