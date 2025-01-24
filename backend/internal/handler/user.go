package handler

import (
	"backend/internal/service"
	"backend/pkg/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *handlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body service.CreateUserBody
	// Validato role
	raw, err := io.ReadAll(r.Body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(raw, &body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	if body.Role != "TEACHER" && body.Role != "STUDENT" {
		response.SetError(w, http.StatusBadGateway, fmt.Errorf("Role undefined"))
		return
	}

	res, err := h.userService.CreateUser(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}

func (h *handlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		response.SetError(w, http.StatusBadRequest, fmt.Errorf("Email is required"))
		return
	}

	res, err := h.userService.Login(r.Context(), email)

	if err != nil {
		if err.Error() == "User not Found" {
			response.SetError(w, http.StatusNotFound, err)
			return
		}
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
