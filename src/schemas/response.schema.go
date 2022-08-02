package schemas

type Response[T any] struct {
	Status       int
	ErrorMessage string
	Data         T
}
