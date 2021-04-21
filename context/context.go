package context

import "net/http"


type Context struct  {
	responseWriter http.ResponseWriter
	request *http.Request


	middlewareIndex int
	middlewareList []func(*Context)

}

func NewContext(responseWriter http.ResponseWriter, request *http.Request, middleFuncList []func(*Context)) *Context {
	return &Context{responseWriter:responseWriter, request:request,middlewareIndex:-1, middlewareList:middleFuncList}

}


//output string
//content response data
func (c *Context) String(content string){
	c.responseWriter.Write([]byte(content))
}

func (c *Context)  Next(){
	c.middlewareIndex ++
	middlewareListLen := len(c.middlewareList)
	if middlewareListLen > c.middlewareIndex {
		c.middlewareList[c.middlewareIndex](c)
		c.middlewareIndex ++
	}
}

// add a middleware
func (c *Context)AddMiddlewareList(fun func(*Context))  {
	c.middlewareList = append(c.middlewareList, fun)
}




