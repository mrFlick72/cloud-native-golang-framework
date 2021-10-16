package application

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mrflick72/cloud-native-golang-framework/configuration"
	"github.com/mrflick72/cloud-native-golang-framework/health"
	"github.com/mrflick72/cloud-native-golang-framework/middleware/security"
	"github.com/mrflick72/cloud-native-golang-framework/middleware/security/oidc"
	"github.com/mrflick72/cloud-native-golang-framework/web"
	"os"
	"sync"
	"time"
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

	if "true" == manager.GetConfigFor("security.oidc.enabled") {
		sessionLifeTime, _ := time.ParseDuration(os.Getenv("SESSION_LIFE_TIME"))
		sess := sessions.New(sessions.Config{
			Cookie:       "go_session_id",
			AllowReclaim: true,
			Expires:      sessionLifeTime,
		})

		app.Use(sess.Handler())
		oidc.SetUpOIDC(app)
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
