package utils

import "fmt"

// ErrInErr 很适合一个err包含另一个err，并且提供附带数据的情况
type ErrInErr struct {
	ErrDesc   string
	ErrDetail error
	Data      interface{}
}

func (e *ErrInErr) String() string {
	if e.Data != nil {

		if e.ErrDetail != nil {
			return fmt.Sprintf("%s : %s, Data: %v", e.ErrDesc, e.ErrDetail.Error(), e.Data)

		}

		return fmt.Sprintf("%s , Data: %v", e.ErrDesc, e.Data)

	}
	if e.ErrDetail != nil {
		return fmt.Sprintf("%s : %s", e.ErrDesc, e.ErrDetail.Error())

	}
	return e.ErrDesc

}

func (e *ErrInErr) Error() string {
	return e.String()
}

func NewErr(desc string, e error) *ErrInErr {
	return &ErrInErr{
		ErrDesc:   desc,
		ErrDetail: e,
	}
}

func NewDataErr(desc string, e error, data interface{}) *ErrInErr {
	return &ErrInErr{
		ErrDesc:   desc,
		ErrDetail: e,
		Data:      data,
	}
}
