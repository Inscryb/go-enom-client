package response

type Response struct {
	ResponseCode int    `xml:"RRPCode" json:"response_code"`
	ResponseText string `xml:"RRPText" json:"response_text"`
	Command      string `xml:"Command" json:"command"`
	Language     string `xml:"Language" json:"language"`
	ErrCount     int    `xml:"ErrCount" json:"err_count"`
	ErrX         string `xml:"ErrX" json:"error_message"`
}
