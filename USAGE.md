<!-- Start SDK Example Usage -->
```go
package main

import (
    "github.com/resendlabs/resend-go"
    "github.com/resendlabs/resend-go/pkg/models/shared"
    "github.com/resendlabs/resend-go/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.SendEmailRequest{
        Request: shared.Email{
            Bcc: "sit",
            Cc: "voluptas",
            From: "culpa",
            HTML: "expedita",
            React: "consequuntur",
            ReplyTo: "dolor",
            Subject: "expedita",
            Text: "voluptas",
            To: "fugit",
        },
    }
    
    res, err := s.Emails.SendEmail(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->