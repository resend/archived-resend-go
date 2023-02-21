package operations

import (
	"github.com/resendlabs/resend-go/pkg/models/shared"
)

type SendEmailRequest struct {
	Request shared.Email `request:"mediaType=application/json"`
}

type SendEmailResponse struct {
	ContentType       string
	SendEmailResponse *shared.SendEmailResponse
	StatusCode        int
}
