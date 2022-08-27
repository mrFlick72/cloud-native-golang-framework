package application

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/mrflick72/cloud-native-golang-framework/configuration"
	heath "github.com/mrflick72/cloud-native-golang-framework/health"
	"github.com/mrflick72/cloud-native-golang-framework/middleware/security"
	"github.com/mrflick72/cloud-native-golang-framework/web"
	"sync"
)

var manager = configuration.GetConfigurationManagerInstance()

func newWebServer() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	return app
}

func NewApplicationServer() *iris.Application {
	app := newWebServer()

	if "true" == manager.GetConfigFor("security.oauth2.enabled") {
		security.SetUpOAuth2(app, security.Jwk{
			Url:    manager.GetConfigFor("security.jwk-uri"),
			Client: web.New(),
		}, manager.GetConfigFor("security.allowed-authority"))
	}

	return app
}

func StartApplicationServer(wg *sync.WaitGroup, app *iris.Application) {
	app.Listen(fmt.Sprintf(":%v", manager.GetConfigFor("server.port")))
	wg.Done()
}

func NewActuatorServer(wg *sync.WaitGroup) {
	app := newWebServer()
	endpoints := heath.HealthEndpoint{}
	endpoints.ResgisterEndpoints(app)
	app.Listen(fmt.Sprintf(":%v", manager.GetConfigFor("management.port")))
	wg.Done()
}
