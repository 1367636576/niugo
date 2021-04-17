package context

import "net/http"

type Context struct  {
	responseWriter http.ResponseWriter
	request *http.Request

}

func NewContext(responseWriter http.ResponseWriter, request *http.Request) Context {
	return Context{responseWriter, request}
}

//output string
//content response data
func (c *Context) String(content string){
	c.responseWriter.Write([]byte(content))
}


