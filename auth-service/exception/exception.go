package exception

type Exception struct {
	Code int
	err string
}

func(e *Exception) Error() string{
	return e.err
}
