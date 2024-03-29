package duo

import (
	"net/url"

	duoapi "github.com/duosecurity/duo_api_golang"

	"github.com/authelia/authelia/v4/internal/middlewares"
)

// API interface wrapping duo api library for testing purpose.
type API interface {
	Call(values url.Values, ctx *middlewares.AutheliaCtx) (*Response, error)
}

// APIImpl implementation of DuoAPI interface.
type APIImpl struct {
	*duoapi.DuoApi
}

// Response response coming from Duo API.
type Response struct {
	Response struct {
		Result        string `json:"result"`
		Status        string `json:"status"`
		StatusMessage string `json:"status_msg"`
	} `json:"response"`
	Code          int    `json:"code"`
	Message       string `json:"message"`
	MessageDetail string `json:"message_detail"`
	Stat          string `json:"stat"`
}
