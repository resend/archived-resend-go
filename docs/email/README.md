# Email

## Overview

Email operations

### Available Operations

* [SendEmail](#sendemail) - Send an email

## SendEmail

Send an email

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/resendlabs/resend-go"
	"github.com/resendlabs/resend-go/pkg/models/shared"
)

func main() {
    s := resend.New(
        resend.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Email.SendEmail(ctx, shared.Email{
        Bcc: resend.String("vel"),
        Cc: resend.String("error"),
        From: "deserunt",
        HTML: resend.String("suscipit"),
        ReplyTo: resend.String("iure"),
        Subject: "magnam",
        Text: resend.String("debitis"),
        To: "ipsa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
}
```
