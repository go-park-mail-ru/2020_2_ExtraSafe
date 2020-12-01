package middlewares

import (
	"github.com/labstack/echo"
	"net/http"
	"testing"
	"time"
)

func TestMiddlew_CORS(t *testing.T) {
	m := NewMiddleware()
	e := echo.New()
	e.Debug()
	e.Use(CORS())
	e.Get("/", test)
	go e.Run(":63246")
	time.Sleep(1 * time.Second)
	res, err := http.Get("http://127.0.0.1:63246")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("Wrong status code")
	}
	if res.Header.Get("Access-Control-Allow-Origin") != "*" {
		t.Fatal("Access-Control-Allow-Origin header != *")
	}
}
