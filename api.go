package sendflare_sdk_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	baseUrl        = "https://api.sendflare.io"
	requestTimeout = time.Second * 10
)

type sendKit struct {
	token string
}

type SendflareImpl interface {
	// SendEmail Send an email
	SendEmail(req SendEmailReq) (SendEmailResp, error)
	// GetContactList Get contact list
	GetContactList(req ListContactReq) (ListContactResp, error)
	// SaveContact Create or update contact
	SaveContact(req SaveContactReq) (SaveContactResp, error)
	// DeleteContact Delete a contact
	DeleteContact(req DeleteContactReq) (DeleteContactResp, error)
	makeHeaders() map[string]string
}

var _ SendflareImpl = &sendKit{}

// NewSendflare New sendflare client instance
func NewSendflare(token string) SendflareImpl {
	return &sendKit{token: token}
}

func (s *sendKit) SendEmail(req SendEmailReq) (resp SendEmailResp, err error) {
	path := "/v1/send"

	payload, err := json.Marshal(&req)
	if err != nil {
		return
	}
	rsq, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseUrl, path), bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	headers := s.makeHeaders()
	for k, v := range headers {
		rsq.Header.Set(k, v)
	}
	client := &http.Client{Timeout: requestTimeout}
	rsp, err := client.Do(rsq)
	if err != nil {
		return
	}
	defer func() {
		_ = rsp.Body.Close()
	}()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (s *sendKit) GetContactList(req ListContactReq) (resp ListContactResp, err error) {
	path := "/v1/contact"

	params := url.Values{}
	params.Add("appId", req.AppId)
	params.Add("page", strconv.Itoa(req.Page))
	params.Add("pageSize", strconv.Itoa(req.PageSize))

	rsq, err := http.NewRequest("GET", fmt.Sprintf("%s%s?%s", baseUrl, path, params.Encode()), nil)
	if err != nil {
		return
	}
	headers := s.makeHeaders()
	for k, v := range headers {
		rsq.Header.Set(k, v)
	}
	client := &http.Client{Timeout: requestTimeout}
	rsp, err := client.Do(rsq)
	if err != nil {
		return
	}
	defer func() {
		_ = rsp.Body.Close()
	}()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (s *sendKit) SaveContact(req SaveContactReq) (resp SaveContactResp, err error) {
	path := "/v1/contact"

	payload, err := json.Marshal(&req)
	if err != nil {
		return
	}
	rsq, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseUrl, path), bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	headers := s.makeHeaders()
	for k, v := range headers {
		rsq.Header.Set(k, v)
	}
	client := &http.Client{Timeout: requestTimeout}
	rsp, err := client.Do(rsq)
	if err != nil {
		return
	}
	defer func() {
		_ = rsp.Body.Close()
	}()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (s *sendKit) DeleteContact(req DeleteContactReq) (resp DeleteContactResp, err error) {
	path := "/v1/contact"

	params := url.Values{}
	params.Add("appId", req.AppId)
	params.Add("emailAddress", req.EmailAddress)
	rsq, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s?%s", baseUrl, path, params.Encode()), nil)
	if err != nil {
		return
	}
	headers := s.makeHeaders()
	for k, v := range headers {
		rsq.Header.Set(k, v)
	}
	client := &http.Client{Timeout: requestTimeout}
	rsp, err := client.Do(rsq)
	if err != nil {
		return
	}
	defer func() {
		_ = rsp.Body.Close()
	}()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (s *sendKit) makeHeaders() map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + s.token,
		"Content-Type":  "application/json",
	}
}
