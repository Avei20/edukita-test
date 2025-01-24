package handler

import (
	"backend/internal/service"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

func (h *handlerImpl) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	var body service.CreateAssignmentBody

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

	res, err := h.assignmentService.CreateAssignment(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}

func (h *handlerImpl) GetAssignment(w http.ResponseWriter, r *http.Request) {
	filterBy := r.URL.Query().Get("filterby")

	res, err := h.assignmentService.GetAllAssignmentsByTeacher(r.Context(), &filterBy)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
