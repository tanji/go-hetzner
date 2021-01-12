package hetzner

import (
	"net/http"
)

type SupportService interface {
	CreateSupportRequest(req *SupportRequest) (*http.Response, error)
}

type SupportServiceImpl struct {
	client *Client
}

func (s *SupportServiceImpl) CreateSupportRequest(req *SupportRequest) (*http.Response, error) {
	path := "/support"

	resp, err := s.client.Call(http.MethodPost, path, req, nil, true)
	return resp, err
}
