package operations

import (
	"github.com/resendlabs/resend-go/pkg/models/shared"
	"github.com/resendlabs/resend-go/pkg/utils"
)

type SendEmailRequest struct {
	Retries *utils.RetryConfig
	Request shared.Email `request:"mediaType=application/json"`
}

type SendEmailResponse struct {
	ContentType       string
	SendEmailResponse *shared.SendEmailResponse
	StatusCode        int64
}
