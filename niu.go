package niu

import "net/http"
//import "fmt"
import "niu/router"
import "niu/context"

type Application struct {
	//router
	Router *router.Router
}

func New() *Application {
	app := &Application{}

	//inject router
	app.Router = router.NewRouter()

	//inject Middleware Ctx
	//app.MiddlewareCtx = context.NewContext()

	return app
}

func (a *Application)Get(pattern string, handleFunc func( *context.Context))  {
	a.Router.Get(pattern, handleFunc)
}

func (a *Application)Post(pattern string, handleFunc func(*context.Context))  {
	a.Router.Post(pattern, handleFunc)
}
func (a *Application)Put(pattern string, handleFunc func(*context.Context))  {
	a.Router.Put(pattern, handleFunc)
}
func (a *Application)Delete(pattern string, handleFunc func(*context.Context))  {
	a.Router.Delete(pattern, handleFunc)
}
func (a *Application)Head(pattern string, handleFunc func(*context.Context))  {
	a.Router.Head(pattern, handleFunc)
}

func (a *Application) Use(middleware func(*context.Context))  {
	a.Router.AddHandle(middleware)
}

func (a *Application) Run(addr string) {

	http.ListenAndServe(addr, a.Router)

}


