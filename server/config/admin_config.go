package config

type AdminConfig struct {
	Notice Notice `json:"notice"`
}

type Notice struct {
	Type              string `json:"type"`
	EmailServerHost   string `json:"email_server_host"`
	EmailUsername     string `json:"email_username"`
	EmailPassword     string `json:"email_password"`
	EmailIsEncrypt    int    `json:"email_is_encrypt,string"`
	EmailPort         int    `json:"email_port,string"`
	EmailReceiveUser  string `json:"email_receive_user"`
	GotifyServerUrl   string `json:"gotify_server_url"`
	GotifyServerToken string `json:"gotify_server_token"`
}
