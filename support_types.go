package hetzner

type SupportRequest struct {
	ServerIP string `url:"server_ip"`
	Type     string `url:"type"`
	Test     string `url:"test"`
	Message  string `url:"message"`
}
