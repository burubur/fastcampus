package repository

import (
	"context"
	"net/http"
	"testing"

	"github.com/burubur/fastcampus/payment/repository/mock"
	"github.com/burubur/fastcampus/payment/vo"
	"go.uber.org/mock/gomock"
)

// Negative test cases
// sending request with empty auth header - DONE
// sending request with invalid auth header
// sending request with valid auth header
// sending request with empty request body - IN PROGRESS
// "{\"error_code\":\"API_VALIDATION_ERROR\",\"message\":\"amount must be greater than 0\"}\n"
// "{\"error_code\":\"API_VALIDATION_ERROR\",\"message\":\"Only one of 'payment_method' or 'payment_method_id' should be present per request\"}\n"
// sending request with broken request body
// sending request with incomplete
// sending request with invalid request body (missing some request field)
// sending request with duplicate reference_id
// Positive test cases
// "{\"id\":\"pr-e290e51a-5256-40a6-8df7-3eba31b56152\",\"country\":\"ID\",\"amount\":100000,\"currency\":\"IDR\",\"business_id\":\"599bd7f1ccab55b020bb1147\",\"reference_id\":\"a046d420-0cea-4422-8c48-e746b7c1a447\",\"payment_method\":{\"id\":\"pm-3d7d1670-ddc3-4f9e-aa16-e8160ea93143\",\"type\":\"VIRTUAL_ACCOUNT\",\"reference_id\":\"7666a8de-ee69-41c6-b6f3-b770f75e0caa\",\"description\":null,\"created\":\"2024-04-15T10:34:46.759047314Z\",\"updated\":\"2024-04-15T10:34:47.239059629Z\",\"card\":null,\"ewallet\":null,\"direct_debit\":null,\"direct_bank_transfer\":null,\"over_the_counter\":null,\"virtual_account\":{\"amount\":100000,\"currency\":\"IDR\",\"channel_code\":\"BRI\",\"channel_properties\":{\"customer_name\":\"John Doe\",\"virtual_account_number\":\"262158018484744\",\"expires_at\":\"2055-04-15T10:34:47.066351Z\"}},\"qr_code\":null,\"metadata\":null,\"billing_information\":{\"city\":null,\"country\":\"\",\"postal_code\":null,\"province_state\":null,\"street_line1\":null,\"street_line2\":null},\"reusability\":\"ONE_TIME_USE\",\"status\":\"PENDING\"},\"description\":null,\"metadata\":{\"sku\":\"\"},\"customer_id\":null,\"capture_method\":\"AUTOMATIC\",\"initiator\":null,\"card_verification_results\":null,\"created\":\"2024-04-15T10:34:46.537775371Z\",\"updated\":\"2024-04-15T10:34:46.537775371Z\",\"status\":\"PENDING\",\"actions\":[],\"failure_code\":null,\"channel_properties\":null,\"shipping_information\":null,\"items\":null}\n"
// sending request with complete and valid request body
// sending request with inactive channel_code: BCA
// Edge cases
// TODO: find the edge cases

func TestXenditPayment_SendPaymentRequest_APIExploration(t *testing.T) {
	httpClient := &http.Client{}
	hostName := "https://api.xendit.co"
	// apiKey := "something"
	authKey := "eG5kX2RldmVsb3BtZW50X09vbUFmT1V0aCtHb3dzWTZMZUpPSHpMQ1p0U2o4NEo5a1hEbitSeGovbUhXK2J5aERRVnhoZz09Og=="
	xenditClient := NewXenditClient(httpClient, hostName, authKey)

	ctx := context.Background()
	paymentID, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequest{})
	if err != nil {
		t.Fatalf("it should not return any error, but got: %s", err.Error())
	}

	if paymentID == "" {
		t.Errorf("it should not return empty paymentID, but got: %s", paymentID)
	}
}

func TestXenditPayment_SendPaymentRequest_EmptyAuthHeader(t *testing.T) {
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
	host := "http://mock.server"
	authKey := ""
	xenditClient := NewXenditClient(httpClientMock, host, authKey)
	ctx := context.Background()
	_, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequest{})

	if err == nil {
		t.Fatal("it should return error due to empty auth key")
	}
}

func TestXenditPayment_SendPaymentRequest_WithEmptyPayload(t *testing.T) {
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
	host := "http://mock.server"
	authKey := "supersecret"
	xenditClient := NewXenditClient(httpClientMock, host, authKey)
	ctx := context.Background()
	paymentReq := vo.XenditPaymentRequest{}
	_, err := xenditClient.SendPaymentRequest(ctx, paymentReq)

	if err == nil {
		t.Fatal("it should return error due to empty auth key")
	}
}

func TestXenditPayment_SendPaymentRequest_ErrorWhileCreatingHttpRequest(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_IncompleteRequestData(t *testing.T) {

}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_ButGot500(t *testing.T) {}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_Got200_ButEmptyResponseBody(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_ButGotBrokenResponseBody(t *testing.T) {
}

func TestXenditPayment_SendPaymentRequest_SuccessResponse(t *testing.T) {

}

// func TestXenditPayment_SendPaymentRequest(t *testing.T) {
// 	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
// 	host := "http://mock.server"

// 	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))

// 	xenditClient := NewXenditClient(httpClientMock, host)
// 	paymentID, err := xenditClient.SendPaymentRequest(context.Background())
// 	assert.Error(t, err, "it should not return error")

// 	assert.Empty(t, paymentID, "it should return a valid created paymentID")
// }
