package boards

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTransport_BoardChangeRead(t *testing.T) {

}

func TestTransport_BoardRead(t *testing.T) {
//	e.GET("/board/:ID/", middle.CookieSession(middle.CheckBoardAdminPermission(board.Board)))

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", 1)
	c.SetPath("/board/:ID/")
	c.SetParamNames("ID")
	c.SetParamValues("1")

	transport := &transport{}
	expectedRequest := models.BoardInput{
		UserID:  1,
		BoardID: 1,
	}

	request, err := transport.BoardRead(c)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(request, expectedRequest) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", expectedRequest, request)
		return
	}


/*	ctrlTransport := gomock.NewController(t)
	defer ctrlTransport.Finish()
	mockTransport := mocks.NewMockAuthTransport(ctrlTransport)

	ctrlService := gomock.NewController(t)
	defer ctrlService.Finish()
	mockService := mocks.NewMockAuthService(ctrlService)
*/
	/*h := &handler{
		authService: mockService,
		authTransport: mockTransport,
	}*/



/*	userInput := models.UserInputLogin{
		Email: "jon@labstack.com",
		Password: "lalala",
	}
	userID := uint64(1)*/

/*	mockTransport.EXPECT().LoginRead(c).Return(userInput, nil)
	mockService.EXPECT().Login(userInput).Return(userID, models.UserOutside{}, nil)
*/


	// Assertions
	/*if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}*/


}

func TestTransport_BoardWrite(t *testing.T) {

}

func TestTransport_CardChangeRead(t *testing.T) {

}

func TestTransport_CardWrite(t *testing.T) {

}

func TestTransport_TaskChangeRead(t *testing.T) {

}

func TestTransport_TasksOrderRead(t *testing.T) {

}

func TestTransport_TaskWrite(t *testing.T) {

}