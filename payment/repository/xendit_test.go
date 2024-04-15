package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/burubur/fastcampus/payment/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestXenditPayment_SendPaymentRequest_IncompleteRequestData(t *testing.T) {}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_ButGot500(t *testing.T) {}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_Got200_ButEmptyResponseBody(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_ButGotBrokenResponseBody(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_SuccessResponse(t *testing.T) {

}

func TestXenditPayment_SendPaymentRequest(t *testing.T) {
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
	host := "http://mock.server"

	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))

	xenditClient := NewXenditClient(httpClientMock, host)
	paymentID, err := xenditClient.SendPaymentRequest(context.Background())
	assert.Error(t, err, "it should not return error")

	assert.Empty(t, paymentID, "it should return a valid created paymentID")
}
