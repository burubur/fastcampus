package repository_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/burubur/fastcampus/payment/repository"
	"github.com/burubur/fastcampus/payment/repository/mock"
	"github.com/burubur/fastcampus/payment/vo"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type XenditPaymentTestSuite struct {
	suite.Suite
	ctx            context.Context
	httpClientMock *mock.MockHttpConector
	xenditClient   repository.XenditPayment
	xenditHost     string
	xenditAuthKey  string
}

func TestXenditPayment(t *testing.T) {
	suite.Run(t, new(XenditPaymentTestSuite))
}

func (s *XenditPaymentTestSuite) SetupSuite() {
	s.ctx = context.Background()
	ctrl := gomock.NewController(s.T())
	httpClientMock := mock.NewMockHttpConector(ctrl)

	s.xenditHost = "http://mock.server"
	s.xenditAuthKey = "supersecret"

	xenditClient := repository.NewXenditClient(httpClientMock, s.xenditHost, s.xenditAuthKey)

	s.httpClientMock = httpClientMock
	s.xenditClient = xenditClient
}

func (s *XenditPaymentTestSuite) SetupTest() {

}

func (s *XenditPaymentTestSuite) TestXenditPayment_SendPaymentRequest_EmptyAuthHeader() {
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(s.T()))
	host := "http://mock.server"
	authKey := ""
	xenditClient := repository.NewXenditClient(httpClientMock, host, authKey)
	ctx := context.Background()
	_, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequest{})
	s.Error(err, "it should return error due to empty auth key")
}

func (s *XenditPaymentTestSuite) TestXenditPayment_SendPaymentRequest_WithEmptyPayload() {
	_, gotErr := s.xenditClient.SendPaymentRequest(s.ctx, vo.XenditPaymentRequest{})
	s.Error(gotErr, "it should return error due to empty payment request payload")
}

func (s *XenditPaymentTestSuite) TestXenditPayment_SendPaymentRequest_ErrorWhileCreatingHttpRequest() {
	s.httpClientMock.EXPECT().Do(gomock.Any()).Return(&http.Response{}, errors.New("http failure"))
	paymentRequest := vo.XenditPaymentRequest{
		Currency: "IDR",
		PaymentMethod: vo.PaymentMethod{
			PaymentMethodType: "VIRTUAL_ACCOUNT",
			ReferenceID:       "random-id",
		},
	}
	_, gotErr := s.xenditClient.SendPaymentRequest(s.ctx, paymentRequest)
	s.Error(gotErr, "it should return an error due to http failure")
	s.ErrorContains(gotErr, "failure")
}
