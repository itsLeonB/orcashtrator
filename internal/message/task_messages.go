package message

type ExpenseBillUploaded struct {
	URI string `json:"uri"`
}

func (ebu ExpenseBillUploaded) Type() string {
	return "expense-bill-uploaded"
}
