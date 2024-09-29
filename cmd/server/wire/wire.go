//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"piyo-engine/internal/handler"
	"piyo-engine/internal/repository"
	"piyo-engine/internal/server"
	"piyo-engine/internal/service"
	"piyo-engine/pkg/app"
	"piyo-engine/pkg/jwt"
	"piyo-engine/pkg/log"
	"piyo-engine/pkg/server/http"
	"piyo-engine/pkg/sid"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	// repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewTeamRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewTeamService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewTeamHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(
		wire.Build(
			repositorySet,
			serviceSet,
			handlerSet,
			serverSet,
			sid.NewSid,
			jwt.NewJwt,
			newApp,
		),
	)
}
