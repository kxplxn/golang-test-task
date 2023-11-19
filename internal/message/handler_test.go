package message

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPush(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		// ARRANGE
		req := httptest.NewRequest(
			http.MethodPost,
			"/message",
			bytes.NewReader([]byte(`{"nonField": "nonValue"}`)),
		)
		wantStatusCode := http.StatusBadRequest

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		// ACT
		Handle(ctx)

		// ASSERT
		resp := w.Result()
		if resp.StatusCode != wantStatusCode {
			t.Errorf(
				"mismatch status code - got: %d, want: %d",
				resp.StatusCode,
				wantStatusCode,
			)
		}
	})

	t.Run("OK", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPost,
			"/message",
			bytes.NewReader(
				[]byte(`{"sender": "foo", "receiver": "bar", "message": "baz"}`),
			),
		)
		wantStatusCode := http.StatusOK

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		// ACT
		Handle(ctx)

		// ASSERT
		resp := w.Result()
		if resp.StatusCode != wantStatusCode {
			t.Errorf(
				"mismatch status code - got: %d, want: %d",
				resp.StatusCode,
				wantStatusCode,
			)
		}
	})
}
