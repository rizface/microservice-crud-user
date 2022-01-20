package response

type User struct {
	Id           string `json:"id"`
	Username     string `json:"username,omitempty"`
	NamaDepan    string `json:"nama_depan,omitempty"`
	NamaBelakang string `json:"nama_belakang,omitempty"`
	Level        string `json:"level,omitempty"`
	Password     string `json:"password,omitempty"`
}
