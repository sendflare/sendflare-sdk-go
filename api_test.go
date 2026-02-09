package sendflare_sdk_go

import "testing"

var (
	token = "live_NTQzOThjYzEtMzFhZS00ZWUyLWExMDYtZmNkMTRiZmNiNzQ1QlBQSVdJT0JCUUZZQldJQklETEhQUUZBTktESUVRRlVNUUtLV0dFTldJVUNXVkdLWVY"
	appId = "f5e3886a19674edc9a6b3abf41900ead"
)

func TestNewSendflare(t *testing.T) {
	t.Run("NewSendflare", func(t *testing.T) {
		t.Log(NewSendflare(token))
	})
}

func TestSend(t *testing.T) {
	t.Run("Send", func(t *testing.T) {
		req := SendEmailReq{
			From:    "test@example.com",
			To:      "receive@example.com",
			Subject: "test email available",
			Body:    "test email",
		}
		t.Log(req)
		t.Log(NewSendflare(token).SendEmail(req))
	})
}

func TestGetContactList(t *testing.T) {
	t.Run("GetContactList", func(t *testing.T) {
		req := ListContactReq{
			AppId: appId,
			PaginateReq: PaginateReq{
				PageSize: 10,
				Page:     1,
			},
		}
		t.Log(req)
		t.Log(NewSendflare(token).GetContactList(req))
	})
}

func TestSaveContact(t *testing.T) {
	t.Run("SaveContact", func(t *testing.T) {
		req := SaveContactReq{
			AppId:        appId,
			EmailAddress: "test@example.com",
			Data: map[string]string{
				"firstName": "John",
				"lastName":  "Doe",
			},
		}
		t.Log(req)
		t.Log(NewSendflare(token).SaveContact(req))
	})
}

func TestDeleteContact(t *testing.T) {
	t.Run("DeleteContact", func(t *testing.T) {
		req := DeleteContactReq{
			AppId:        appId,
			EmailAddress: "test@example.com",
		}
		t.Log(req)
		t.Log(NewSendflare(token).DeleteContact(req))
	})
}
