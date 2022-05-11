package gin

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/proxy"
)

func TestRender(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		res := &proxy.Response{
			IsComplete: true,
			Data: map[string]interface{}{
				"a": map[string]interface{}{
					"content": "supu",
				},
				"content": "tupu",
				"foo":     42,
			},
		}
		Render(c, res)
	})

	expected := `<?xml version='1.0' standalone='yes'?><doc><a><content>supu</content></a><content>tupu</content><foo>42</foo></doc>`

	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	defer w.Result().Body.Close()

	body, ioerr := ioutil.ReadAll(w.Result().Body)
	if ioerr != nil {
		t.Error("reading response body:", ioerr)
		return
	}

	content := string(body)
	if w.Result().Header.Get("Content-Type") != gin.MIMEXML {
		t.Error("Content-Type error:", w.Result().Header.Get("Content-Type"))
	}
	if w.Result().StatusCode != http.StatusOK {
		t.Error("Unexpected status code:", w.Result().StatusCode)
	}
	if content != expected {
		t.Error("Unexpected body:", content, "expected:", expected)
	}
}
