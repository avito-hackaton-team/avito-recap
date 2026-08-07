package app

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/generated/recapapi"
	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/config"
	recapcontroller "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/controller/http/recap"
	activityrepo "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/repository/activity"
	profilerepo "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/repository/profile"
	recaprepo "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/repository/recap"
	profileusecase "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/usecase/profile"
	recapusecase "github.com/avito-hackaton-team/avito-recap/backend/recap/internal/usecase/recap"
)

const apiPrefix = "/api/v1"

func buildHTTPHandler(
	pool *pgxpool.Pool,
	cfg config.Config,
	logger *zap.Logger,
) (http.Handler, error) {
	activityRepository := activityrepo.New(
		pool,
		cfg.Repository.OperationTimeout,
	)

	profileRepository := profilerepo.New(
		pool,
		cfg.Repository.OperationTimeout,
	)

	recapRepository := recaprepo.New(
		pool,
		cfg.Repository.OperationTimeout,
	)

	recapService := recapusecase.NewRecapService(
		activityRepository,
		recapRepository,
		profileRepository,
	)

	profileService := profileusecase.NewProfileService(
		profileRepository,
	)

	recapHandler := recapcontroller.NewRecapServer(
		logger,
		recapService,
		profileService,
	)

	server, err := recapapi.NewServer(recapHandler)
	if err != nil {
		return nil, fmt.Errorf("create recap openapi server: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, server))

	return mux, nil
}
