package niu

import "net/http"
//import "fmt"
import "niu/router"

type Application struct {
	//router
	router *router.Router

}

func New() *Application {
	app := &Application{}

	//inject router
	app.router = router.NewRouter()

	return app
}

func (a *Application)Get(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.router.Get(pattern, handleFunc)
}

func (a *Application)Post(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.router.Post(pattern, handleFunc)
}
func (a *Application)Put(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.router.Put(pattern, handleFunc)
}
func (a *Application)Delete(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.router.Delete(pattern, handleFunc)
}
func (a *Application)Head(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.router.Head(pattern, handleFunc)
}

func (a *Application) Run(addr string) {
	http.ListenAndServe(addr, a.router)
}


