package schemas

type Response struct {
	Status       int
	ErrorMessage string
	Data         interface{}
}
