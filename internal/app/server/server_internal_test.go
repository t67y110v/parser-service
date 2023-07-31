package server

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/t67y110v/parser-service/internal/app/logging"
// 	"github.com/t67y110v/parser-service/internal/app/store/teststore"
// )

// func TestServer_HandleUserCreate(t *testing.T) {
// 	log := logging.GetLogger()
// 	s := newServer(teststore.NewPostgres(), teststore.NewMongo(), nil, log)
// 	testCases := []struct {
// 		name         string
// 		payload      interface{}
// 		expectedCode int
// 	}{
// 		{
// 			name: "valid",
// 			payload: map[string]string{
// 				"email":    "user@example.org",
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusCreated,
// 		},
// 		{
// 			name:         "invalid payload",
// 			payload:      "invalid",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "invalid params",
// 			payload: map[string]string{
// 				"email": "invalid",
// 			},
// 			expectedCode: http.StatusUnprocessableEntity,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			app := fiber.New()
// 			app.Post("/login")
// 			s.handlers.Login()

// 			rec := httptest.NewRecorder()
// 			b := &bytes.Buffer{}
// 			json.NewEncoder(b).Encode(tc.payload)
// 			req, _ := http.NewRequest(http.MethodPost, "/api/login", b)
// 			s.ServeHTTP(rec, req)
// 			assert.Equal(t, tc.expectedCode, rec.Code)
// 		})
// 	}

// }
