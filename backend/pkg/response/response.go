package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type DataTemplate struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type ErrorTemplate struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
}

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Api-Key")
	return w
}

func SetRawResponse(w http.ResponseWriter, httpStatus int, data any) {
	sentResponse(w, httpStatus, data)
}

func SetResponse(w http.ResponseWriter, httpStatus int, data interface{}, message string, success bool) {
	response := DataTemplate{
		Message: message,
		Success: success,
		Status:  http.StatusText(httpStatus),
		Data:    data,
	}

	sentResponse(w, httpStatus, response)
}

func SetError(w http.ResponseWriter, httpStatus int, err error) {
	log.Println(err)
	response := ErrorTemplate{
		Message: err.Error(),
		Success: false,
		Status:  http.StatusText(httpStatus),
	}

	sentResponse(w, httpStatus, response)
}

func ReturnInternalServerError(w http.ResponseWriter, err error) {
	log.Println(err)

	response := ErrorTemplate{
		Message: "Internal Server Error",
		Success: false,
		Status:  http.StatusText(http.StatusInternalServerError),
	}

	sentResponse(w, http.StatusInternalServerError, response)
}

func sentResponse(w http.ResponseWriter, header int, data any) {
	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("[ERROR][Util][Write Error] Marshalling Reponse Error, %+v\n", err)
		ReturnInternalServerError(w, err)
		return
	}

	w = setHeader(w)
	w.WriteHeader(header)
	w.Write(encoded)
}
