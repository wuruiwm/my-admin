package response

type Ssl struct {
	Key        string   `json:"key"`
	Pem        string   `json:"pem"`
	Domain     []string `json:"domain"`
	ExpireTime string   `json:"expire_time"`
}
