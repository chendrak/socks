package socks

var (
	ErrNoProperRespone              = NewSocksError("Server does not respond properly.")
	ErrNoSocks5Support              = NewSocksError("Server does not support Socks 5.")
	ErrMethodNegotiationFailed      = NewSocksError("socks method negotiation failed.")
	ErrCantCompleteSocks5Connection = NewSocksError("Can't complete SOCKS5 connection.")
	ErrSocksRequestRejectedOrFailed = NewSocksError("Socks connection request rejected or failed.")
	ErrCannotConnectToIdentd        = NewSocksError("Socks connection request rejected becasue SOCKS server cannot connect to identd on the client.")
	ErrClientServerDifferentUserIds = NewSocksError("Socks connection request rejected because the client program and identd report different user-ids.")
	ErrUnknown                      = NewSocksError("Socks connection request failed, unknown error.")
	ErrIpv6NotSupported             = NewSocksError("IPv6 is not supported by SOCKS4.")
)

type (
	SocksError struct {
		Message string
	}
)

func (s *SocksError) Error() string {
	return s.Message
}

func NewSocksError(message string) *SocksError {
	return &SocksError{
		Message: message,
	}
}
