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
    s := resend.New(resend.WithSecurity(
        shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        },
    ))
    
    req := operations.SendEmailRequest{
        Request: shared.Email{
            From: "hello@resend.com",
            To: "thefuture@yourcompany.com",
            Subject: "Welcome to Resend!",
            Text: "Hello, World!",
        },
    }

    ctx := context.Background()
    res, err := s.Email.SendEmail(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
}
```