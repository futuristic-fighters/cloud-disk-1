package cfg

type Err struct {
	code  int
	error string
}

func NewError(Code int) *Err {
	return &Err{
		code:  Code,
		error: Lang[Code],
	}
}

func (e *Err) Code() int {
	return e.code
}

func (e *Err) Error() string {
	return e.error
}

func (e *Err) SetCode(c int) *Err {
	e.code = c
	return e
}

func (e *Err) SetError(msg string) *Err {
	e.error = msg
	return e
}
