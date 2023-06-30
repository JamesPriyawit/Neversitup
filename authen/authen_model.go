package authen

type LoginRes struct {
	Token		string `json:"token"`
	UserId    string `json:"userId"`
}