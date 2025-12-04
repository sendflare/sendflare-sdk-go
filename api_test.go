package sendflare_sdk_go

import "testing"

func TestNewSendflare(t *testing.T) {
	t.Run("NewSendflare", func(t *testing.T) {
		t.Log(NewSendflare("this-is-my-token"))
	})
}

func TestSend(t *testing.T) {
	t.Run("Send", func(t *testing.T) {
		req := SendEmailReq{
			From:    "test@example.com",
			To:      "to@example.com",
			Subject: "test",
			Body:    "test email",
		}
		t.Log(req)
		t.Log(NewSendflare("this-is-my-token").SendEmail(req))
	})
}

func TestGetContactList(t *testing.T) {
	t.Run("GetContactList", func(t *testing.T) {
		req := ListContactReq{
			AppId: "test",
			PaginateReq: PaginateReq{
				PageSize: 10,
				Page:     1,
			},
		}
		t.Log(req)
		t.Log(NewSendflare("this-is-my-token").GetContactList(req))
	})
}

func TestSaveContact(t *testing.T) {
	t.Run("SaveContact", func(t *testing.T) {
		req := SaveContactReq{
			AppId:        "test",
			EmailAddress: "test@example.com",
			Data: map[string]string{
				"firstName": "John",
				"lastName":  "Doe",
			},
		}
		t.Log(req)
		t.Log(NewSendflare("this-is-my-token").SaveContact(req))
	})
}

func TestDeleteContact(t *testing.T) {
	t.Run("DeleteContact", func(t *testing.T) {
		req := DeleteContactReq{
			AppId:        "test",
			EmailAddress: "test@example.com",
		}
		t.Log(req)
		t.Log(NewSendflare("this-is-my-token").DeleteContact(req))
	})
}
