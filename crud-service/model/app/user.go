package app

import "time"

type User struct {
	Id           string `json:"id" gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Username     string `json:"username" gorm:"index:idx_username,unique"`
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nam_belakang"`
	Level        string `json:"level"`
	Password     string `json:"password"`
	CreatedAt    time.Time
}
