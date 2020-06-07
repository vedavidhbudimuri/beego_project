package controllers

type UserNameAlreadyTaken struct {
}

func (e *UserNameAlreadyTaken) Error() string {
	return "Username already taken"
}

type ExpectationFailed struct {
	msg string
}

func (e *ExpectationFailed) Error() string {
	return e.msg
}
