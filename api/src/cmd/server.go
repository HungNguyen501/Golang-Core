package cmd

import (
	"fmt"

	"golang-core/api/src/api/controller"
	"golang-core/api/src/api/route"
	echoutils "golang-core/api/src/common/echo_utils"
	"golang-core/api/src/common/logging"
	"golang-core/api/src/common/parser"
	"golang-core/api/src/infrastructure/database"
	"golang-core/api/src/infrastructure/repository"
	"golang-core/api/src/manager"
)

func StartServer() {
	env := "local"
	cfg, err := parser.ParseAppConfig(env)
	if err != nil {
		panic(err)
	}
	logger := logging.NewLogging()

	db, err := database.NewDbContext(&cfg.DatabaseConfig, logger)
	if err != nil {
		logger.Error(err)
		return
	}
	defer db.Close()

	repositories := repository.NewRopositories(db)
	managers := manager.NewManagers(*repositories)
	controllers := controller.NewControllers(logger, managers)

	validators := echoutils.NewValidators()
	if err := validators.Setup(); err != nil {
		panic(err)
	}

	logger.Infow("starting api ...")
	router := route.NewRouter(controllers, validators)
	router.Start(fmt.Sprintf(":%d", cfg.ServerConfig.Port))
}
