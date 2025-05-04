//go:build integration
// +build integration

package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"

	"golang-core/api/src/api/controller"
	"golang-core/api/src/api/route"
	"golang-core/api/src/client/response"
	echoutils "golang-core/api/src/common/echo_utils"
	"golang-core/api/src/common/logging"
	"golang-core/api/src/common/orm"
	"golang-core/api/src/common/parser"
	"golang-core/api/src/infrastructure/database"
	"golang-core/api/src/infrastructure/repository"
	"golang-core/api/src/manager"
	"golang-core/api/src/test"
	"golang-core/api/src/test/mock"
)

type UserControllerTestSuite struct {
	suite.Suite
	Db *database.Db
	ec *echo.Echo
}

const (
	InsertUserEndpoint  = "/api/v1/user/insert"
	UpdatetUserEndpoint = "/api/v1/user/update"
	DeleteUserEndpoint  = "/api/v1/user/%s"
	FindByIdEndpoint    = "/api/v1/user/%s"
	ListByIdsEndpoint   = "/api/v1/user/list?offset=%d&limit=%d"
)

// Run once before all tests
func (s *UserControllerTestSuite) SetupSuite() {
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
func (s *UserControllerTestSuite) TearDownSuite() {
	s.Db.Close()
}

// Run once before each test
func (s *UserControllerTestSuite) SetupTest() {}

// Run once after each test
func (s *UserControllerTestSuite) TearDownTest() {
	err := test.TruncateTable("users", s.Db)
	s.Require().NoError(err)
}

// Run all tests
func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (s *UserControllerTestSuite) TestInsertUser_Success() {
	insertRequest := mock.MockUserInsertRequest()
	actual, code, err := test.RequestSuccess[response.UserInfoResponse](s.ec, http.MethodPost, InsertUserEndpoint, insertRequest)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.Equal(insertRequest.Email, *actual.Data.Email)
	s.Equal(insertRequest.Name, *actual.Data.Name)
}

func (s *UserControllerTestSuite) TestUpdateUserInfo_Success() {
	insertRequest := mock.MockUserInsertRequest()
	insertResponse, code, err := test.RequestSuccess[response.UserInfoResponse](s.ec, http.MethodPost, InsertUserEndpoint, insertRequest)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)

	updateRequest := mock.MockUserUpdateRequest(*insertResponse.Data.Id)
	updateResponse, code, err := test.RequestSuccess[response.UserInfoResponse](s.ec, http.MethodPost, UpdatetUserEndpoint, updateRequest)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.Equal(*updateRequest.Email, *updateResponse.Data.Email)
	s.Equal(*updateRequest.Name, *updateResponse.Data.Name)
}

func (s *UserControllerTestSuite) TestDeleteUser_Success() {
	insertRequest := mock.MockUserInsertRequest()
	insertResponse, code, err := test.RequestSuccess[response.UserInfoResponse](s.ec, http.MethodPost, InsertUserEndpoint, insertRequest)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)

	getResponse, code, err := test.RequestSuccess[response.UserInfoResponse](
		s.ec,
		http.MethodGet,
		fmt.Sprintf(FindByIdEndpoint, *insertResponse.Data.Id),
		nil,
	)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.NotNil(getResponse.Data.Id)

	deleteResponse, code, err := test.RequestSuccess[string](
		s.ec,
		http.MethodDelete,
		fmt.Sprintf(FindByIdEndpoint, *insertResponse.Data.Id),
		nil,
	)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.Equal(deleteResponse.Data, "Deleted")

	getResponse, code, err = test.RequestSuccess[response.UserInfoResponse](
		s.ec,
		http.MethodGet,
		fmt.Sprintf(FindByIdEndpoint, *insertResponse.Data.Id),
		nil,
	)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.Nil(getResponse.Data.Id)
}

func (s *UserControllerTestSuite) TestListUserByIds_Success() {
	expectedUserNum := 5
	for _ = range expectedUserNum {
		userInsertRequest := mock.MockUserInsertRequest()
		_, code, err := test.RequestSuccess[response.UserInfoResponse](s.ec, http.MethodPost, InsertUserEndpoint, userInsertRequest)
		s.Require().NoError(err)
		s.Equal(http.StatusOK, code)
	}

	listRes, code, err := test.RequestSuccess[orm.PaginationData[response.UserInfoResponse]](
		s.ec,
		http.MethodGet,
		fmt.Sprintf(ListByIdsEndpoint, 0, 100),
		nil,
	)
	s.Require().NoError(err)
	s.Equal(http.StatusOK, code)
	s.Equal(listRes.Data.Total, expectedUserNum)
}
