package router

import(
	"net/http"
)

// Route contains the information about a registered Route.
type Route struct  {
	//route list
	RList map[string] func (http.ResponseWriter, *http.Request)

}

//register a router
func (r *Route) Add(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	r.RList[pattern] = handleFunc
}


func (r *Route) ServeHTTP(w http.ResponseWriter, request *http.Request){
	reqUri := request.RequestURI
	for pattern, handleFun := range r.RList {
		if pattern == reqUri {
			handleFun(w, request)
			return
		}
	}
	w.Write([]byte("404"))

}
