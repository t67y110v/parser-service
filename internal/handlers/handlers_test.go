package handlers_test

import (
	"net/http"
	"net/http/httptest"

	"testing"
)

func SetStatus(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusTeapot)
}
func TestParse(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/parse", nil)
	res := httptest.NewRecorder()
	SetStatus(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got wrong status ")
	}
}
