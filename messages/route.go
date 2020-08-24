package messages

// Route for handling
type Route struct {
	Method  string
	Path    string
	Handler func(request *Request, response *Response)
}
