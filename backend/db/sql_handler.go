package db

import (
	"encoding/json"
	"log"
)

func (client *DBClient) getUserAccountDatas(id uuid) Userdata, error {
	// データベースからデータを取得する
	// TODO: URLのクエリーによって取得するデータを変えるようにする
	// 現在はUSERSテーブルの全てのデータを取得しているからよろしくない
	rows, err := client.DB.Query("SELECT * FROM USERS WHERE ID = $1", id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
			log.Print(err)
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	jsonresponse, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error:Json marshal %v", err)
		return nil, err
	}
	return jsonresponse, nil
}
