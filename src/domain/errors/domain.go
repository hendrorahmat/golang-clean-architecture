package errors

type DomainError interface {
	Error() string
	GetTitle() string
	GetCode() uint
	GetStatusCode() int
}
