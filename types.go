package sendflare_sdk_go

// PaginateReq Pagination request entity
type PaginateReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// PaginateResp Pagination response entity
type PaginateResp struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
}

// CommonResponse Common response entity
type CommonResponse struct {
	RequestID string      `json:"requestId"`
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"ts"`
	Data      interface{} `json:"data"`
}

// SendEmailReq Send Email request entity
type SendEmailReq struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// SendEmailResp Send Email response entity
type SendEmailResp CommonResponse

// ListContactReq Get Contact list request entity
type ListContactReq struct {
	PaginateReq `json:",inline"`
	AppId       string `form:"appId"`
}

// ListContactResp Get Contact list response entity
type ListContactResp struct {
	PaginateResp `json:",inline"`
	List         []ContactItem `json:"data"`
}

// ContactItem Contact info entity
type ContactItem struct {
	Status       string `json:"status"`
	EmailAddress string `json:"emailAddress"`
	Language     string `json:"language"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PhoneNumber  string `json:"phoneNumber"`
	Birthday     string `json:"birthday"`
	Company      string `json:"company"`
	VIPLevel     int    `json:"vipLevel"`
	Amount       string `json:"amount"`
}

// SaveContactReq Save contact request entity
type SaveContactReq struct {
	AppId        string            `json:"appId"`
	EmailAddress string            `json:"emailAddress"`
	Data         map[string]string `json:"data,omitempty"`
}

// SaveContactResp Save contact response entity
type SaveContactResp CommonResponse

// DeleteContactReq Delete a contact request entity
type DeleteContactReq struct {
	EmailAddress string `json:"emailAddress"`
	AppId        string `json:"appId"`
}

type DeleteContactResp CommonResponse
