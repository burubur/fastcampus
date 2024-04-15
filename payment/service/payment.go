package service

type PaymentGatewayProvider interface {
	SendPaymentRequest() (paymentID string, err error)
}

type Payment struct {
	xendit PaymentGatewayProvider
}

func (p *Payment) Pay() error {
	// create transaction ID
	// insert into PostgreSQL
	// call third party API
	_, err := p.xendit.SendPaymentRequest()
	if err != nil {
		return err
	}

	return nil
}
