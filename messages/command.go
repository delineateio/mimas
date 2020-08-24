package messages

// Command performs the required action for the service
type Command func(request *Request, response *Response)
