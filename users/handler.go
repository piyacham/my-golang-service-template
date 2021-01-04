package users

import (
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

func (h *Handler) getUser(c echo.Context) error {
	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		log.Printf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res, err := h.svc.getUser(requestID, req)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(err.Error()))
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, res)
}

func (h *Handler) createUser(c echo.Context) error {
	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		log.Printf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res, err := h.svc.createUser(requestID, req)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(err.Error()))
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, res)
}

func (h *Handler) updateUser(c echo.Context) error {

	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		log.Printf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res, err := h.svc.updateUser(requestID, req)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(err.Error()))
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, res)
}

func (h *Handler) deleteUser(c echo.Context) error {

	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		log.Printf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res, err := h.svc.deleteUser(requestID, req)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(err.Error()))
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, res)
}
