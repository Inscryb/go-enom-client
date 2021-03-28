package response

type SetContact struct {
	Response

	PendingVerification bool `xml:"PendingVerification" json:"pending_verification"`
}
