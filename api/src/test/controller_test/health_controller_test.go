//go:build integration
// +build integration

package controller_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"

	"golang-core/api/src/api/controller"
	"golang-core/api/src/api/route"
	"golang-core/api/src/client/response"
	echoutils "golang-core/api/src/common/echo_utils"
	"golang-core/api/src/common/logging"
	"golang-core/api/src/common/parser"
	"golang-core/api/src/infrastructure/database"
	"golang-core/api/src/infrastructure/repository"
	"golang-core/api/src/manager"
	"golang-core/api/src/test"
)

type HealthControllerTestSuite struct {
	suite.Suite
	Db *database.Db
	ec *echo.Echo
}

const (
	HealthCheckEndpoint = "/api/v1/health"
)

// Run once before all tests
func (s *HealthControllerTestSuite) SetupSuite() {
	var err error
	env := "test"
	cfg, err := parser.ParseAppConfig(env)
	if err != nil {
		panic(err)
	}
	logger := logging.NewLogging()

	s.Db, err = database.NewDbContext(&cfg.DatabaseConfig, logger)
	s.Require().NoError(err)

	repositories := repository.NewRopositories(s.Db)
	managers := manager.NewManagers(*repositories)
	controllers := controller.NewControllers(logger, managers)

	validators := echoutils.NewValidators()
	err = validators.Setup()
	s.Require().NoError(err)

	s.ec = route.NewRouter(controllers, validators)
}

// Run once after all tests
func (s *HealthControllerTestSuite) TearDownSuite() {
	s.Db.Close()
}

// Run once before each test
func (s *HealthControllerTestSuite) SetupTest() {}

// Run once after each test
func (s *HealthControllerTestSuite) TearDownTest() {}

// Run all tests
func TestHealthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerTestSuite))
}

func (s *HealthControllerTestSuite) TestHealthCheck_Success() {
	actual, code, err := test.RequestSuccess[response.HealthResponse](s.ec, http.MethodGet, HealthCheckEndpoint, nil)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.Equal(actual.Data.Status, "OK")
}
