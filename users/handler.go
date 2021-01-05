package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *Service
}

/*NewHandler
Method: NewHandler
Desc: Maping to business logic
*/
func NewHandler(svc *Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		log.Printf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res, err := h.svc.CreateUser(requestID, req)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(err.Error()))
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, res)
}

func (h *Handler) GetUser(c echo.Context) error {
	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res := h.svc.GetUser(requestID, req)
	/*res := Response{
		Code:    status_code.Success,
		Message: "Inquiry Success !",
		// Data:    ResponseData{},
	}*/

	respJson, err := json.Marshal(res)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(`
						{
							"code" : 99,
							"message" : "can't not marshal req on response"
						}
				`),
		)
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, respJson)
}
