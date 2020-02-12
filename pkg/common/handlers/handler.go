package common

import (
	"encoding/json"
	"log"
	"net/http"
	"garlic-mayo/internal/service"
)

//Handler construct
type Handler struct {
	svc service.ZombieService
}

//WakeupZombieHandler handler for wakeup zombie
func (h *Handler) WakeupZombieHandler(w http.ResponseWriter, req *http.Request) {
	var request Request
	err := json.Unmarshal(req.Body, &request)
	if err != nil {
		log.Println(err)
		return
	}

	score, hErr:=h.svc.WakeupZombieWalker(&request)
	if hErr!=nil {
		log.Println(hErr)
		return
	}

	

}
//NewHandlers return new Handlers
func NewHandlers() *Handlers {
	return &Handler{
		svc: service.NewZombieService(),
	}
}
