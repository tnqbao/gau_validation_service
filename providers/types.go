package providers

type RequestMail struct {
	Mail    *string `json:"mail"`
	Content *string `json:"content"`
}

type RequestSMS struct {
	Phone   *string `json:"phone"`
	Content *string `json:"content"`
}

type ResponseMail struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
