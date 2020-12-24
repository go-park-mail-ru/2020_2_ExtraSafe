package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/cmd/handlers/mock"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/errorworker"
	logger2 "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/tools/logger"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMiddlew_CORS(t *testing.T) {
	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardService := mock.NewMockServiceBoard(ctrlBoard)

	ctrlAuth := gomock.NewController(t)
	defer ctrlAuth.Finish()
	mockAuthService := mock.NewMockServiceAuth(ctrlAuth)

	ctrlTransport := gomock.NewController(t)
	defer ctrlTransport.Finish()
	mockTransportService := mock.NewMockTransportAuth(ctrlTransport)

	worker := errorworker.NewErrorWorker()

	zeroLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
	logger := logger2.NewLogger(&zeroLogger)

	m := NewMiddleware(worker, mockAuthService, mockTransportService, mockBoardService, logger)
	e := echo.New()
	e.Use(m.CORS())
	e.GET("/", test)
	go e.Start(":63246")
	time.Sleep(1 * time.Second)
	res, err := http.Get("http://127.0.0.1:63246")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("Wrong status code")
	}
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test123")
}
