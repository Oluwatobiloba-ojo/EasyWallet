package response

import "time"

type PaystackWebhookResponse struct {
	Event string      `json:"event"`
	Data  WebhookData `json:"data"`
}

type WebhookData struct {
	Id              int       `json:"id"`
	Domain          string    `json:"domain"`
	Status          string    `json:"status"`
	Reference       string    `json:"reference"`
	Amount          int       `json:"amount"`
	GatewayResponse string    `json:"gateway_response"`
	PaidAt          time.Time `json:"paid_at"`
	CreatedAt       time.Time `json:"created_at"`
	Channel         string    `json:"channel"`
	Currency        string    `json:"currency"`
	IpAddress       string    `json:"ip_address"`
	Metadata        int       `json:"metadata"`
}
