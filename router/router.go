package router

import(
	"net/http"
	"niu/context"
	"fmt"
)

//Pattern Handle Function Map without request method(GET:POST:DELETE,ect...)
type PatternHandlerMap map[string]func (context.Context)



// Router contains the information about a registered Route.
type Router struct  {
	//method[GET,POST,DELETE,HEAD,PUT] route list
	RList map[string] PatternHandlerMap

}


//register a router
//method :request method
func (r *Router) add(method string ,pattern string, handleFunc func(context.Context))  {
	//invalid request method
	if !isValidMethod(method) {
		return
	}

	patternHandlerMap := make(map[string]func (context.Context))

	for pattern, val := range r.RList[method] {
		patternHandlerMap[pattern] = val
	}
	patternHandlerMap[pattern] = handleFunc

	//r.RList[method] = PatternHandlerMap{pattern: handleFunc}

	r.RList[method] = patternHandlerMap
	fmt.Println(patternHandlerMap)


}

//get request
func (r *Router) Get(pattern string, handleFunc func(context.Context))  {
	r.add("GET", pattern, handleFunc)
}

//post request
func (r *Router) Post(pattern string, handleFunc func(context.Context))  {
	r.add("POST", pattern, handleFunc)
}

//delete request
func (r *Router) Delete(pattern string, handleFunc func(context.Context))  {
	r.add("DELETE", pattern, handleFunc)
}

//put request
func (r *Router) Put(pattern string, handleFunc func(context.Context))  {
	r.add("PUT", pattern, handleFunc)
}

//head request
func (r *Router) Head(pattern string, handleFunc func(context.Context))  {
	r.add("HEAD", pattern, handleFunc)
}



func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	reqUri := request.RequestURI
	reqMethod := request.Method

	ctx := context.NewContext(w, request)
	for pattern, handleFun := range r.RList[reqMethod] {
		if pattern == reqUri {
			handleFun(ctx)
			return
		}
	}
	ctx.String("404")

}

// new router
func NewRouter()  *Router {
	routerInstance := &Router{}
	routerInstance.RList = make(map[string]PatternHandlerMap)
	return routerInstance
}
