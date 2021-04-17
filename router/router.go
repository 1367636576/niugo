package router

import(
	"net/http"
)

//Pattern Handle Function Map without request method(GET:POST:DELETE,ect...)
type PatternHandlerMap map[string]func (http.ResponseWriter, *http.Request)



// Router contains the information about a registered Route.
type Router struct  {
	//method[GET,POST,DELETE,HEAD,PUT] route list
	RList map[string] PatternHandlerMap

}

//register a router
//method :request method
func (r *Router) add(method string ,pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	//invalid request method
	if !isValidMethod(method) {
		return
	}

	r.RList[method] = PatternHandlerMap{pattern: handleFunc}

}

//get request
func (r *Router) Get(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	r.add("GET", pattern, handleFunc)
}

//post request
func (r *Router) Post(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	r.add("POST", pattern, handleFunc)
}

//delete request
func (r *Router) Delete(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	r.add("DELETE", pattern, handleFunc)
}

//put request
func (r *Router) Put(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	r.add("PUT", pattern, handleFunc)
}

//head request
func (r *Router) Head(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	r.add("HEAD", pattern, handleFunc)
}



func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	reqUri := request.RequestURI
	reqMethod := request.Method

	for pattern, handleFun := range r.RList[reqMethod] {
		if pattern == reqUri {
			handleFun(w, request)
			return
		}
	}
	w.Write([]byte("404"))

}
