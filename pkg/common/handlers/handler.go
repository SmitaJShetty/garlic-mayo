package handlers

import (
	"encoding/json"
	"garlic-mayo/internal/service"
	apperror "garlic-mayo/pkg/apperrors"
	"log"
	"net/http"
)

//Handler construct
type Handler struct {
	svc service.Apocalypse
}

//WakeupZombieHandler handler for wakeup zombie
func (h *Handler) WakeupZombieHandler(w http.ResponseWriter, req *http.Request) {
	var request service.Request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}

	validateErr:= request.Validate()
	if validateErr!=nil {
		log.Println(validateErr)
		return
	}

	score, hErr := h.svc.WakeupZombieWalker(&request)
	if hErr != nil {
		log.Println(hErr)
		return
	}

	SendResult(w, req, []byte(string(score)))
}

//NewHandlers return new Handlers
func NewHandlers() *Handler {
	return &Handler{
		svc: service.NewZombieService(),
	}
}

//SendResult sends result over http response
func SendResult(w http.ResponseWriter, r *http.Request, resultJSON []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

//SendErrorResponse sends error response
func SendErrorResponse(w http.ResponseWriter, r *http.Request, appErr *apperror.AppError) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json, _ := json.Marshal(appErr)
	w.Write(json)
}
