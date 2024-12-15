package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	apiV1 "piyo-engine/api/v1"
	"piyo-engine/docs"
	"piyo-engine/internal/handler"
	"piyo-engine/internal/middleware"
	"piyo-engine/pkg/jwt"
	"piyo-engine/pkg/log"
	"piyo-engine/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	teamHandler *handler.TeamHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	setupRouter(logger, s, jwt, userHandler, teamHandler)

	return s
}

func setupRouter(logger *log.Logger, server *http.Server, jwt *jwt.JWT, userHandler *handler.UserHandler, teamHandler *handler.TeamHandler) {
	server.GET(
		"/swagger/*any", ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			// ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
			ginSwagger.DefaultModelsExpandDepth(-1),
			ginSwagger.PersistAuthorization(true),
		),
	)

	server.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		// middleware.SignMiddleware(log),
	)

	// Health routes
	server.GET(
		"/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			apiV1.HandleSuccess(
				ctx, map[string]interface{}{
					":)": "Thank you for using Piyo Engine!",
				},
			)
		},
	)

	// API v1 routes
	v1 := server.Group("/v1")
	{
		setupUserRoutes(logger, v1.Group("user"), jwt, userHandler)
		setupTeamRoutes(logger, v1.Group("team"), jwt, teamHandler)
	}
}

func setupUserRoutes(logger *log.Logger, r *gin.RouterGroup, jwt *jwt.JWT, userHandler *handler.UserHandler) {
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// Non-strict permission routing group
	noStrictAuthRouter := r.Use(middleware.NoStrictAuth(jwt, logger))
	{
		noStrictAuthRouter.GET("/", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := r.Use(middleware.StrictAuth(jwt, logger))
	{
		strictAuthRouter.PUT("/", userHandler.UpdateProfile)
	}
}

func setupTeamRoutes(logger *log.Logger, r *gin.RouterGroup, jwt *jwt.JWT, teamHandler *handler.TeamHandler) {
	strictRoleAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger))
	{
		strictRoleAuthRouter.POST("/", teamHandler.CreateTeam)
		strictRoleAuthRouter.GET("/:team_id", teamHandler.GetTeamProfile)
		strictRoleAuthRouter.GET("/:team_id/members", teamHandler.GetTeamMembers)
	}
}
