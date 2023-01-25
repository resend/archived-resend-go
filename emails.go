package sdk

import (
	"context"
	"fmt"
	"github.com/resendlabs/resend-go/pkg/models/operations"
	"github.com/resendlabs/resend-go/pkg/models/shared"
	"github.com/resendlabs/resend-go/pkg/utils"
	"net/http"
	"strings"
)

type Emails struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewEmails(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Emails {
	return &Emails{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// SendEmail - Send an email
func (s *Emails) SendEmail(ctx context.Context, request operations.SendEmailRequest) (*operations.SendEmailResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/emails"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s._language, s._sdkVersion, s._genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s._securityClient

	retryConfig := request.Retries
	if retryConfig == nil {
		retryConfig = &utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 5000,
				MaxInterval:     60000,
				Exponent:        1.5,
				MaxElapsedTime:  3600000,
			},
			RetryConnectionErrors: true,
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"500",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.SendEmailResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SendEmailResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SendEmailResponse = out
		}
	}

	return res, nil
}
