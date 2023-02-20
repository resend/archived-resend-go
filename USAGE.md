<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/resendlabs/resend-go"
    "github.com/resendlabs/resend-go/pkg/models/shared"
    "github.com/resendlabs/resend-go/pkg/models/operations"
)

func main() {
    opts := []resend.SDKOption{
        resend.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := resend.New(opts...)
    
    req := operations.SendEmailRequest{
        Request: shared.Email{
            Bcc: "unde",
            Cc: "deserunt",
            From: "porro",
            HTML: "nulla",
            ReplyTo: "id",
            Subject: "vero",
            Text: "perspiciatis",
            To: "nulla",
        },
    }
    
    res, err := s.Email.SendEmail(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->