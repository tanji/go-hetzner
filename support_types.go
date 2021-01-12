package hetzner

type SupportRequest struct {
	Type         string `url:"type"`
	ServerIP     string `url:"server_ip"`
	ServerNumber int    `url:"server_number"`
	Downtime     string `url:"30_minutes_test_downtime"`
	Description  string `url:"description"`
}
