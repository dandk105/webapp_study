package schema

// まだmain fileから分離していない構造体達

import (
	"time"
)

// ユーザー情報を格納する構造体
type User struct {
	// IDはUUIDになるから、stringでまとめてしまうのはいささか乱雑な気がする
	ID          string
	Name        string
	Birthday    time.Time
	DeleateFlag bool
}

// 部屋情報を格納する構造体
type Room struct {
	ID       string
	Name     string
	Capacity int
}

// 予約情報を格納する構造体
type Reservation struct {
	ID string
	// User structのID
	UserID string
	// Room structのID
	RoomID    string
	StartTime time.Time
	EndTime   time.Time
}

// ホテル情報を格納する構造体
type Hotel struct {
	ID      string
	Name    string
	Address string
	RoomID  string
}

type UsersResponse struct {
	Users User `json:"users"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
