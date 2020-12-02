package errorWorker

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type ErrorWorker interface {
	RespError(c echo.Context, serveError error) (err error)
	TransportError(c echo.Context) (err error)
	TokenError(c echo.Context, token string) (err error)
}

type errorWorker struct {
}

func NewErrorWorker() ErrorWorker {
	return &errorWorker{}
}

type ResponseError struct {
	Status int      `json:"status"`
	Codes  []string `json:"codes"`
}

type ResponseTokenError struct {
	Status int      `json:"status"`
	Codes  []string `json:"codes"`
	Token  string   `json:"token"`
}

func (e errorWorker) RespError(c echo.Context, serveError error) (err error) {
	responseErr := new(ResponseError)
	responseErr.Codes = serveError.(models.ServeError).Codes
	responseErr.Status = 500
	return c.JSON(http.StatusBadRequest, responseErr)
}

func (e errorWorker) TransportError(c echo.Context) (err error) {
	responseErr := new(ResponseError)
	responseErr.Codes = nil
	responseErr.Status = 500
	return c.JSON(http.StatusBadRequest, responseErr)
}

func (e errorWorker) TokenError(c echo.Context, token string) (err error) {
	responseErr := new(ResponseTokenError)
	responseErr.Codes = nil
	responseErr.Status = 777
	responseErr.Token = token
	return c.JSON(http.StatusBadRequest, responseErr)
}

func ConvertErrorToStatus(servError models.ServeError, serviceName string) error {
	st := status.New(codes.InvalidArgument, serviceName)

	br := &errdetails.BadRequest{}

	for i, code := range servError.Codes {
		if code == models.ServerError {
			br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
				Field: code,
				Description: servError.OriginalError.Error(),
			})
			break
		}
		br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
			Field: code,
			Description: servError.Descriptions[i],
		})
	}

	st, err := st.WithDetails(br)
	if err != nil {
		// If this errored, it will always error
		// here, so better panic so we can figure
		// out why than have this silently passing.
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}

	return st.Err()
}

func ConvertStatusToError(inputError error) (servError models.ServeError) {
	st := status.Convert(inputError)

	servError.MethodName = st.Message()

	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				servError.Codes = append(servError.Codes, violation.GetField())
				servError.Descriptions = append(servError.Descriptions, violation.GetDescription())
			}
		}
	}

	return servError
}