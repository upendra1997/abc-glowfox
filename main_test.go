package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetClasses(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/classes", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestCreateClasses(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"name":  "pilates", "start_date":  "2024-05-22", "end_date":  "2024-05-24", "capacity": 20 }`)
	req, _ := http.NewRequest("POST", "/api/classes", body)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMultipleCreateClasses(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := strings.NewReader(`{"name":  "pilates", "start_date":  "2024-05-22", "end_date":  "2024-05-24", "capacity": 20 }`)
	req, _ := http.NewRequest("POST", "/api/classes", body)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	body = strings.NewReader(`{"name":  "yoga", "start_date":  "2024-05-22", "end_date":  "2024-05-24", "capacity": 50}`)
	req, _ = http.NewRequest("POST", "/api/classes", body)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/classes", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[
    {
        "name": "pilates",
        "date": "2024-05-22",
        "capacity": 40
    },
    {
        "name": "pilates",
        "date": "2024-05-23",
        "capacity": 40
    },
    {
        "name": "pilates",
        "date": "2024-05-24",
        "capacity": 40
    },
    {
        "name": "yoga",
        "date": "2024-05-22",
        "capacity": 50
    },
    {
        "name": "yoga",
        "date": "2024-05-23",
        "capacity": 50
    },
    {
        "name": "yoga",
        "date": "2024-05-24",
        "capacity": 50
    }
]`, w.Body.String())
}

func TestBookClasses(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"name":  "pilates", "date":  "2024-05-22"}`)
	req, _ := http.NewRequest("POST", "/api/booking/upendra", body)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/classes", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[
    {
        "name": "pilates",
        "date": "2024-05-22",
        "capacity": 39
    },
    {
        "name": "pilates",
        "date": "2024-05-23",
        "capacity": 40
    },
    {
        "name": "pilates",
        "date": "2024-05-24",
        "capacity": 40
    },
    {
        "name": "yoga",
        "date": "2024-05-22",
        "capacity": 50
    },
    {
        "name": "yoga",
        "date": "2024-05-23",
        "capacity": 50
    },
    {
        "name": "yoga",
        "date": "2024-05-24",
        "capacity": 50
    }
]`, w.Body.String())
}

func TestGetBookings(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/booking", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[
    {
        "user_name": "upendra",
        "class_name": "pilates",
        "date": "2024-05-22"
    }
]`, w.Body.String())
}
