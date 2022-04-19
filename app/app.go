package app

import "github.com/gin-gonic/gin"

// CreateApp but not run
func CreateApp(param *WebAppParam) *WebApplication {
	engine := gin.Default()
	app := &WebApplication{
		param:  param,
		engine: engine,
	}
	app.mount()
	return app
}

type WebApplication struct {
	param  *WebAppParam
	engine *gin.Engine
}

func (app *WebApplication) mount() {
	app.engine.GET("/health", app.health)

}

func (app *WebApplication) Run(addr string) error {
	return app.engine.Run(addr)
}
