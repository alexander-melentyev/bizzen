package main //nolint:funlen

import (
	"os"

	_orgHandler "github.com/alexander-melentyev/bizzen/internal/org/delivery/http"
	_orgRepo "github.com/alexander-melentyev/bizzen/internal/org/repository"
	_orgUseCase "github.com/alexander-melentyev/bizzen/internal/org/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title bizzen application API
// @version 1.0
// @description API server for bizzen application
// @contact.name Aleksandr Melentyev
// @contact.email aleksandr.melentyev@nexign.com
// @BasePath /api/v1
// @x-apigw {"scopes": "internal", "domains": "enterprise", "auth": true}
// @x-app_name "bizzen"
// @x-app_version "1.0.0"

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	conn, err := sqlx.Connect("pgx", os.Getenv("DSN"))
	if err != nil {
		log.Warn().Msg("can't connect to database: " + err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	prefix := r.Group("api")

	v1 := prefix.Group("v1")

	orgRepo := _orgRepo.NewRepository(conn)
	orgUseCase := _orgUseCase.NewUseCase(orgRepo)

	_orgHandler.NewHandler(v1, orgUseCase)

	if err := r.Run(); err != nil {
		log.Warn().Msg("Server start error: " + err.Error())
	}
}
