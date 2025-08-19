package credential

// VerifyTicketHandle verify ticket 接口
type VerifyTicketHandle interface {
	SetVerifyTicket(verifyTicket string) error
}
