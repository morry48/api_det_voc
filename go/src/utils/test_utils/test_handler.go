package test_utils

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type HandlerTestCase struct {
	// testcase name
	Name string

	// prepare
	PrepareFunc func() (*gin.Context, *httptest.ResponseRecorder)

	// exec
	ExecHandlerFunc func(c *gin.Context)

	// assertion
	WantCode int
	WantFunc func(*httptest.ResponseRecorder)
}

func RunHandlerTest(t *testing.T, tests []HandlerTestCase) {
	// turn off warning
	gin.SetMode(gin.ReleaseMode)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			RunSingleHandlerTestCase(t, tt)
		})
	}
}

func RunSingleHandlerTestCase(t *testing.T, tt HandlerTestCase) {
	// turn off warning
	gin.SetMode(gin.ReleaseMode)

	// prepare
	c, w := tt.PrepareFunc()

	if c.Request.Header.Get("Content-Type") == "" {
		c.Request.Header.Set("Content-Type", "application/json")
	}

	// exec
	tt.ExecHandlerFunc(c)

	// assertion
	assert.Equal(
		t,
		tt.WantCode,
		w.Code,
		fmt.Sprintf("レスポンスコードがマッチしていません。want = %d, got = %d. Response Body(%v)", tt.WantCode, w.Code, w.Body.String()),
	)

	// custom assertion
	if tt.WantFunc != nil {
		tt.WantFunc(w)
	}
}
