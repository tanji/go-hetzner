package hetzner

import (
	"fmt"
	"net/http"
)

type SupportService interface {
	CreateSupportRequest(req *SupportRequest) (*http.Response, error)
}

type SupportServiceImpl struct {
	client *Client
}

func (s *SupportServiceImpl) CreateSupportRequest(req *SupportRequest) (*http.Response, error) {
	path := fmt.Sprintf("/support/server/%v", req.ServerIP)

	resp, err := s.client.Call(http.MethodPost, path, req, nil, true)
	return resp, err
}
