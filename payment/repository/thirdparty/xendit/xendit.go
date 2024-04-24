package xendit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/burubur/fastcampus/payment/vo"
)

type Config struct {
	AuthKey string
}

type XenditPayment struct {
	config        Config
	host          string
	httpConnector HttpConector
}

func NewXenditClient(httpConnector HttpConector, host string, authKey string) XenditPayment {
	return XenditPayment{
		httpConnector: httpConnector,
		host:          host,
		config: Config{
			AuthKey: authKey,
		},
	}
}

func (x XenditPayment) SendPaymentRequest(ctx context.Context, paymentRequest vo.XenditPaymentRequest) (paymentID string, err error) {
	if x.config.AuthKey == "" {
		return "", errors.New("empty auth key")
	}

	if paymentRequest.PaymentMethod.ReferenceID == "" {
		return "", errors.New("empty/invalid reference_id")
	}

	// TODO: inject the http client
	// validate request payload
	// call xendit PaymentRequest API
	// construct request body
	// handle error response
	// handle success response
	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(paymentRequest)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf("%s%s", x.host, "/payment_requests")
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, reqBody)
	httpReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", x.config.AuthKey))
	httpReq.Header.Set("Content-Type", "application/json")

	res, err := x.httpConnector.Do(httpReq)
	if err != nil {
		return "", err
	}

	rawResponseBody, err := io.ReadAll(res.Body)
	strResponseBody := string(rawResponseBody)
	_ = strResponseBody

	_ = res
	return "", err
}
