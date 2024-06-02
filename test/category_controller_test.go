package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/mnabil1718/go-restful-api/app"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/middleware"
	"github.com/mnabil1718/go-restful-api/model/domain"
	"github.com/mnabil1718/go-restful-api/model/web"
	"github.com/mnabil1718/go-restful-api/repository"
	"github.com/mnabil1718/go-restful-api/service"
	"github.com/stretchr/testify/assert"
)

// empty table and resets id increments
func truncateTable(db *sql.DB, tableName string) {
	db.Exec("TRUNCATE " + tableName)
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	apiV1Path := "/api/v1"
	router := app.NewRouter(categoryController, app.RootPath(apiV1Path))

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := app.NewDB(app.Test)
	truncateTable(db, "categories")
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Gardening","is_active": true}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(actualResponse["code"].(float64)))
	assert.Equal(t, "OK", actualResponse["status"])
	assert.Equal(t, "Gardening", actualResponse["data"].(map[string]interface{})["name"])
	assert.Equal(t, true, actualResponse["data"].(map[string]interface{})["is_active"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := app.NewDB(app.Test)
	truncateTable(db, "categories")
	router := setupRouter(db)

	requestObject := web.CategoryCreateRequest{
		Name:     "",
		IsActive: false,
	}

	jsonBytes, err := json.Marshal(requestObject)
	helper.PanicIfError(err)

	requestBody := strings.NewReader(string(jsonBytes))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, http.StatusBadRequest, int(actualResponse["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", actualResponse["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, domain.Category{
		Name:     "Gardening",
		IsActive: true,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Garden & Park","is_active":false}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/categories/"+strconv.Itoa(int(category.Id)), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]any
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(actualResponse["code"].(float64)))
	assert.Equal(t, "OK", actualResponse["status"])
	assert.Equal(t, int(category.Id), int(actualResponse["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Garden & Park", actualResponse["data"].(map[string]interface{})["name"])
	assert.Equal(t, false, actualResponse["data"].(map[string]interface{})["is_active"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := app.NewDB(app.Test)
	truncateTable(db, "categories")
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Fitness","is_active":true}`) // validator error comes first before not found
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/categories/10", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(actualResponse["code"].(float64)))
	assert.Equal(t, "category not found", actualResponse["data"].(string))
}

func TestDeleteCategorySuccess(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, domain.Category{
		Name:     "Gardening",
		IsActive: true,
	})
	tx.Commit() // don't use defer

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/categories/"+strconv.Itoa(int(category.Id)), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]any
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(actualResponse["code"].(float64)))
	assert.Equal(t, "OK", actualResponse["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/categories/10", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	actualResponse := web.WebResponse{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(actualResponse.Code))
	if actualData, ok := actualResponse.Data.(string); ok {
		assert.Equal(t, "category not found", actualData)
	}
}

func TestGetCategorySuccess(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, domain.Category{
		Name:     "Gardening",
		IsActive: true,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/categories/"+strconv.Itoa(int(category.Id)), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(actualResponse["code"].(float64)))
	assert.Equal(t, "OK", actualResponse["status"])
	assert.Equal(t, int(category.Id), int(actualResponse["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, actualResponse["data"].(map[string]interface{})["name"])
	assert.Equal(t, category.IsActive, actualResponse["data"].(map[string]interface{})["is_active"])
}

func TestGetCategoryFailed(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/categories/10", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(actualResponse["code"].(float64)))
	assert.Equal(t, "NOT FOUND", actualResponse["status"])
	assert.Equal(t, "category not found", actualResponse["data"])
}

func TestListCategoriesSuccess(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category1 := repository.Save(context.Background(), tx, domain.Category{
		Name:     "Gardening",
		IsActive: true,
	})
	category2 := repository.Save(context.Background(), tx, domain.Category{
		Name:     "Sport",
		IsActive: false,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", middleware.ApiKey)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	categories := actualResponse["data"].([]interface{})

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(actualResponse["code"].(float64)))
	assert.Equal(t, "OK", actualResponse["status"])

	assert.Equal(t, category1.Id, int64(categories[0].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category1.Name, categories[0].(map[string]interface{})["name"])
	assert.Equal(t, category1.IsActive, categories[0].(map[string]interface{})["is_active"])

	assert.Equal(t, category2.Id, int64(categories[1].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category2.Name, categories[1].(map[string]interface{})["name"])
	assert.Equal(t, category2.IsActive, categories[1].(map[string]interface{})["is_active"])
}

func TestUnauthorized(t *testing.T) {

	db := app.NewDB(app.Test)
	truncateTable(db, "categories")
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var actualResponse map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&actualResponse)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.Equal(t, http.StatusUnauthorized, int(actualResponse["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", actualResponse["status"])
}
