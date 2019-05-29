package hetzner

type SupportRequest struct {
	ServerIP string `json:"server_ip"`
	Type     string `json:"type"`
	Message  string `json:"message"`
}
