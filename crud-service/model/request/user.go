package request

type User struct {
	Id           string `json:"id"`
	Username     string `json:"username" validate:"required"`
	NamaDepan    string `json:"nama_depan" validate:"required"`
	NamaBelakang string `json:"nama_belakang"`
	Level        string `json:"level" validate:"required"`
	Password     string `json:"password"`
}
