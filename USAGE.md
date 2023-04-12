<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/resendlabs/resend-go"
    "github.com/resendlabs/resend-go/pkg/models/shared"
    "github.com/resendlabs/resend-go/pkg/models/operations"
)

func main() {
    s := resend.New(
        resend.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.Email{
        Bcc: "corrupti",
        Cc: "provident",
        From: "distinctio",
        HTML: "quibusdam",
        ReplyTo: "unde",
        Subject: "nulla",
        Text: "corrupti",
        To: "illum",
    }

    res, err := s.Email.SendEmail(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->