package response

type User struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
	Level        string `json:"level"`
	Password     string
}
