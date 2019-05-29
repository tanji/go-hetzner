package hetzner

type SupportRequest struct {
	ServerIP string `url:"server_ip"`
	Type     string `url:"type"`
	Message  string `url:"message"`
}
