package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicMiddleware(t *testing.T) {
	expectedCode := http.StatusInternalServerError

	// create a handler to use as "next" which will purposely panic
	panicker := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// middlware should recover from this
		panic("panicking")
	})

	// create the handler to test, using our custom "next" handler
	handlerToTest := NewPanicMiddleware()(panicker)
	req := httptest.NewRequest(http.MethodGet, "/vendor/1234444", nil)
	w := httptest.NewRecorder()

	handlerToTest.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Result().StatusCode)
}
