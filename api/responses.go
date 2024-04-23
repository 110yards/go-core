package api

import (
	"net/http"

	"110yards.ca/libs/go/core/logger"
	"github.com/go-chi/render"
)

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func SetResponse(w http.ResponseWriter, r *http.Request, status int, content interface{}) {
	if content == nil {
		w.WriteHeader(status)
	} else {
		render.Status(r, status)
		render.JSON(w, r, content)
	}
}

func InternalServerError(r *http.Request, w http.ResponseWriter, content interface{}) {
	logger.Errorf("Internal server error: %s", content)
	SetResponse(w, r, http.StatusInternalServerError, content)
}

func Ok(r *http.Request, w http.ResponseWriter, content interface{}) {
	SetResponse(w, r, http.StatusOK, content)
}

func BadRequest(r *http.Request, w http.ResponseWriter, content interface{}) {
	logger.Errorf("Bad request: %s", content)
	SetResponse(w, r, http.StatusBadRequest, content)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
