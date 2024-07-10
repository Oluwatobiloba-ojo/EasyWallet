package response

type PaystackWebhookResponse struct {
	Event string      `json:"event"`
	Data  WebhookData `json:"data"`
}

type WebhookData struct {
	Reference string `json:"reference"`
}
