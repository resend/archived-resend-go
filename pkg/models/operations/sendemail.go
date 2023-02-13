package operations

import (
	"github.com/resendlabs/resend-go/pkg/models/shared"
	"github.com/resendlabs/resend-go/pkg/models/utils"
)

type SendEmailRequest struct {
	Request shared.Email `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type SendEmailResponse struct {
	ContentType       string
	SendEmailResponse *shared.SendEmailResponse
	StatusCode        int64
}
