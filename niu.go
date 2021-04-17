package niu

import "net/http"
//import "fmt"
import "niu/router"

type Application struct {
	//router
	Router *router.Router

}

func New() *Application {
	app := &Application{}

	//inject router
	app.Router = router.NewRouter()

	return app
}

func (a *Application)Get(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.Router.Get(pattern, handleFunc)
}

func (a *Application)Post(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.Router.Post(pattern, handleFunc)
}
func (a *Application)Put(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.Router.Put(pattern, handleFunc)
}
func (a *Application)Delete(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.Router.Delete(pattern, handleFunc)
}
func (a *Application)Head(pattern string, handleFunc func(http.ResponseWriter, *http.Request))  {
	a.Router.Head(pattern, handleFunc)
}

func (a *Application) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
}


