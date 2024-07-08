package response

type MonnifyWebhookResponse struct {
	Event     string    `json:"eventType"`
	EventData EventData `json:"eventData"`
}

type EventData struct {
	Reference string `json:"paymentReference"`
}
