package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-core/api/src/client/response"
	"golang-core/api/src/infrastructure/database"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func Request(ec *echo.Echo, method string, path string, body interface{}) *httptest.ResponseRecorder {
	var b bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&b).Encode(body)
	}
	req := httptest.NewRequest(method, path, &b)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ec.ServeHTTP(rec, req)
	return rec
}

func RequestSuccess[T any](ec *echo.Echo, method string, path string, body interface{}) (response.GeneralResponse[T], int, error) {
	var content response.GeneralResponse[T]
	res := Request(ec, method, path, body)
	err := json.Unmarshal(res.Body.Bytes(), &content)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return content, res.Code, err
}

func TruncateTable(name string, db *database.Db) error {
	stmt, err := db.Primary().Prepare(fmt.Sprintf("TRUNCATE TABLE %s;", name))
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	return err
}
