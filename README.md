# sendflare-sdk-go  
The sdk for sendflare service written in Go.  

## Requirements
> Go Version >= 1.18  

## Installation  
```shell
go get -u github.com/sendflare/sendflare-sdk-go
```  

## Code Examples   
```go
req := SendEmailReq{
    From:    "test@example.com",
    To:      "to@example.com",
    Subject: "test",
    Body:    "test email",
}

NewSendflare("this-is-my-token").SendEmail(req)
```  

## Documentation  
[https://docs.sendflare.io](https://docs.sendflare.io)  

## LICENSE  
[MIT](./LICENSE)
