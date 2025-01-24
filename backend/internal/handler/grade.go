package handler

import (
	"backend/internal/service"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

func (h *handlerImpl) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var body service.CreateGradeBody

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

	res, err := h.gradeService.CreateGrade(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}

func (h *handlerImpl) GetGrade(w http.ResponseWriter, r *http.Request) {
	// get student id from url param
	studentID := r.PathValue("studentid")

	resp, err := h.gradeService.GetAllGradesByStudent(r.Context(), studentID)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, resp.StatusCode, resp)
}
