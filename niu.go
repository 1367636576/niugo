package niu

import "net/http"
import "fmt"
import "niu/router"

type Application struct {


}

func (a *Application) Run(addr string,router *router.Route) {
	http.ListenAndServe(addr, router)
}

func (a *Application) Handle(method string,uri string, handleFn func (w http.ResponseWriter, r *http.Request)) {

	r := &http.Request{}
	fmt.Println("me::"+r.Method)
	if method != r.Method {
		fmt.Println("invalid request method")
		//return
	}
	http.HandleFunc(uri, handleFn)
}
