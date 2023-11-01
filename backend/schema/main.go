package schema

// まだmain fileから分離していない構造体達

import (
	"time"
)

// 予約システムのユーザー情報を格納する構造体
type User struct {
	// IDはUUIDになるから、stringでまとめてしまうのはいささか乱雑な気がする
	ID       string
	Name     string
	Birthday time.Time
}

// 予約システムの部屋情報を格納する構造体
type Room struct {
	ID       string
	Name     string
	Capacity int
}

// 予約システムの予約情報を格納する構造体
type Reservation struct {
	ID        string
	UserID    string
	RoomID    string
	StartTime time.Time
	EndTime   time.Time
}

type UsersResponse struct {
	Users User `json:"users"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
