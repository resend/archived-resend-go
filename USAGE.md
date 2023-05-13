<!-- Start SDK Example Usage -->
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
        Bcc: resend.String("corrupti"),
        Cc: resend.String("provident"),
        From: "distinctio",
        HTML: resend.String("quibusdam"),
        ReplyTo: resend.String("unde"),
        Subject: "nulla",
        Text: resend.String("corrupti"),
        To: "illum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->