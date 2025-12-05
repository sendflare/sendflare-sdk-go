# sendflare-sdk-go

The SDK for sendflare service written in Go.

## Requirements

> Go Version >= 1.18

## Installation

```shell
go get -u github.com/sendflare/sendflare-sdk-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    sendflare "github.com/sendflare/sendflare-sdk-go"
)

func main() {
    client := sendflare.NewSendflare("your-api-token")
    
    req := sendflare.SendEmailReq{
        From:    "test@example.com",
        To:      "to@example.com",
        Subject: "Hello",
        Body:    "Test email",
    }
    
    resp, err := client.SendEmail(req)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Email sent successfully: %v\n", resp.Success)
}
```

## Code Examples

### Basic Usage

```go
package main

import (
    "fmt"
    sendflare "github.com/sendflare/sendflare-sdk-go"
)

func main() {
    client := sendflare.NewSendflare("this-is-my-token")
    
    req := sendflare.SendEmailReq{
        From:    "test@example.com",
        To:      "to@example.com",
        Subject: "test",
        Body:    "test email",
    }
    
    resp, err := client.SendEmail(req)
    if err != nil {
        fmt.Printf("Failed to send email: %v\n", err)
        return
    }
    
    fmt.Printf("Email sent successfully: %v\n", resp.Success)
}
```

### Complete Example with All APIs

```go
package main

import (
    "fmt"
    "log"
    sendflare "github.com/sendflare/sendflare-sdk-go"
)

func main() {
    client := sendflare.NewSendflare("this-is-my-token")
    
    // Send an email
    emailReq := sendflare.SendEmailReq{
        From:    "test@example.com",
        To:      "to@example.com",
        Subject: "test",
        Body:    "test email",
    }
    
    emailResp, err := client.SendEmail(emailReq)
    if err != nil {
        log.Printf("Error sending email: %v", err)
    } else {
        fmt.Printf("Email sent: %v\n", emailResp.Success)
    }
    
    // Get contact list
    listReq := sendflare.ListContactReq{
        AppId: "your-app-id",
        PaginateReq: sendflare.PaginateReq{
            Page:     1,
            PageSize: 10,
        },
    }
    
    listResp, err := client.GetContactList(listReq)
    if err != nil {
        log.Printf("Error getting contacts: %v", err)
    } else {
        fmt.Printf("Total contacts: %d\n", listResp.TotalCount)
        for _, contact := range listResp.List {
            fmt.Printf("Contact: %s\n", contact.EmailAddress)
        }
    }
    
    // Save a contact
    saveReq := sendflare.SaveContactReq{
        AppId:        "your-app-id",
        EmailAddress: "contact@example.com",
        Data: map[string]string{
            "firstName": "John",
            "lastName":  "Doe",
        },
    }
    
    saveResp, err := client.SaveContact(saveReq)
    if err != nil {
        log.Printf("Error saving contact: %v", err)
    } else {
        fmt.Printf("Contact saved: %v\n", saveResp.Success)
    }
    
    // Delete a contact
    deleteReq := sendflare.DeleteContactReq{
        AppId:        "your-app-id",
        EmailAddress: "contact@example.com",
    }
    
    deleteResp, err := client.DeleteContact(deleteReq)
    if err != nil {
        log.Printf("Error deleting contact: %v", err)
    } else {
        fmt.Printf("Contact deleted: %v\n", deleteResp.Success)
    }
}
```

## API Reference

### NewSendflare

```go
func NewSendflare(token string) SendflareImpl
```

Create a new Sendflare client instance.

**Parameters:**
- `token` - Your Sendflare API token

**Returns:** A Sendflare client instance

### SendEmail

```go
func (s *sendKit) SendEmail(req SendEmailReq) (SendEmailResp, error)
```

Send an email.

**Parameters:**
- `req.From` - Sender email address
- `req.To` - Recipient email address
- `req.Subject` - Email subject
- `req.Body` - Email body content

**Returns:** 
- `SendEmailResp` - Response containing success status and message
- `error` - Error if the request fails

### GetContactList

```go
func (s *sendKit) GetContactList(req ListContactReq) (ListContactResp, error)
```

Get contact list with pagination.

**Parameters:**
- `req.AppId` - Application ID
- `req.Page` - Page number
- `req.PageSize` - Number of items per page

**Returns:**
- `ListContactResp` - Response containing contacts list and pagination info
- `error` - Error if the request fails

### SaveContact

```go
func (s *sendKit) SaveContact(req SaveContactReq) (SaveContactResp, error)
```

Create or update a contact.

**Parameters:**
- `req.AppId` - Application ID
- `req.EmailAddress` - Contact email address
- `req.Data` - Optional contact data fields (map[string]string)

**Returns:**
- `SaveContactResp` - Response containing success status
- `error` - Error if the request fails

### DeleteContact

```go
func (s *sendKit) DeleteContact(req DeleteContactReq) (DeleteContactResp, error)
```

Delete a contact.

**Parameters:**
- `req.AppId` - Application ID
- `req.EmailAddress` - Contact email address to delete

**Returns:**
- `DeleteContactResp` - Response containing success status
- `error` - Error if the request fails

## Data Types

### Request Types

- `SendEmailReq` - Send email request
- `ListContactReq` - Get contact list request (extends `PaginateReq`)
- `SaveContactReq` - Save contact request
- `DeleteContactReq` - Delete contact request
- `PaginateReq` - Pagination request parameters

### Response Types

- `SendEmailResp` - Send email response (type of `CommonResponse`)
- `ListContactResp` - Get contact list response (extends `PaginateResp`)
- `SaveContactResp` - Save contact response (type of `CommonResponse`)
- `DeleteContactResp` - Delete contact response (type of `CommonResponse`)
- `CommonResponse` - Common response structure
- `PaginateResp` - Pagination response structure
- `ContactItem` - Contact information structure

## Error Handling

All API methods return an `error` as the second return value. Always check for errors:

```go
resp, err := client.SendEmail(req)
if err != nil {
    // Handle error
    log.Printf("Error: %v", err)
    return
}

// Use response
fmt.Printf("Success: %v\n", resp.Success)
```

## Configuration

The SDK uses the following defaults:
- **Base URL**: `https://api.sendflare.io`
- **Request Timeout**: 10 seconds

## Building from Source

```bash
# Clone the repository
git clone https://github.com/sendflare/sendflare-sdk-go.git
cd sendflare-sdk-go

# Run tests
go test -v

# Build
go build
```

## Documentation

For more information, visit: [https://docs.sendflare.io](https://docs.sendflare.io)

## LICENSE

[MIT](./LICENSE)
