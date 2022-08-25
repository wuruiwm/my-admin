package response

type Ssl struct {
	Key        string `json:"key"`
	Pem        string `json:"pem"`
	ExpireTime string `json:"expire_time"`
}
