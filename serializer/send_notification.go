package serializer

type ReceipientWrapper struct {
	Recipients []string `json:"recipients"`
}

func BuildRetrieveForNotificationResponse(input []string) Response {
	return Response{
		Data: ReceipientWrapper{Recipients: input},
	}
}
