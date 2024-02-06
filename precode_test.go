package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Запрос сформирован корректно, код ответа 200, тело ответа не пустое
func TestMainHandlerWhenEverythingIsOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)

	responceRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responceRecorder, req)

	assert.Equal(t, http.StatusOK, responceRecorder.Code)
	assert.NotEmpty(t, responceRecorder.Body)
}

// Город, передаваемый в параметре city, не cуществует
// Код ответа 400, в теле ответа ошибка wrong city value
func TestMainHandlerWhenCityDoesNotExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=samara&count=1", nil)

	responceRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responceRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responceRecorder.Code)
	assert.Equal(t, "wrong city value", responceRecorder.Body.String())
}

// В параметре count передано кол-во больше максимума
// В теле ответа totalCount элементов
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=10", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}
