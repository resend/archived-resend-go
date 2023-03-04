package operations

import (
	"github.com/resendlabs/resend-go/pkg/models/shared"
	"net/http"
)

type SendEmailRequest struct {
	Request shared.Email `request:"mediaType=application/json"`
}

type SendEmailResponse struct {
	ContentType       string
	SendEmailResponse *shared.SendEmailResponse
	StatusCode        int
	RawResponse       *http.Response
}
