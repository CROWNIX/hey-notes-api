package exception

type NotFound struct {
	Message string
}

type BadRequest struct {
	Message string
}

type InternalServer struct {
	Message string
}

type Unauthorized struct {
	Message string
}

func (e *NotFound) Error() string {
	return e.Message
}

func (e *BadRequest) Error() string {
	return e.Message
}

func (e *InternalServer) Error() string {
	return e.Message
}

func (e *Unauthorized) Error() string {
	return e.Message
}