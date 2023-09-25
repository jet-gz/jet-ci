package utils

type IRequestDto interface {
	Validate() error
}
