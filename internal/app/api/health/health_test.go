package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-react-app/internal/app/generated"
)

func TestHealth_Get(t *testing.T) {

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, router := gin.CreateTestContext(resp)

	health := NewHealth()

	router.GET("/test", health.Get)

	ctx.Request = httptest.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(resp, ctx.Request)

	assert.Equal(t, resp.Code, 200)

	data := generated.HealthResponse{}

	err := json.Unmarshal(resp.Body.Bytes(), &data)
	if err != nil {
		t.Log(err.Error())
		t.Fatalf("Could not unmarshal response body. Got %s", resp.Body.String())
	}

	assert.Equal(t, data, generated.HealthResponse{Status: "OK"})
}
