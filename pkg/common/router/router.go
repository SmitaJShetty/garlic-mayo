package common

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"garlic-mayo/pkg/common/handlers"
)

//Start starts router
func Start(listenAddress string) {
	r := GetRouter()
	h := NewHandlers()
	addRoutes(r, h)

	go func() {
		err := http.ListenAndServe(listenAddress, r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Listening on port:", listenAddress)
	}()

}

// addRoutes adds routes
func addRoutes(router *mux.Router, h *Handlers) {
	router.HandleFunc("/walkthezombie", h.InvalidateCache).Methods("POST")
}

// GetRouter gets router
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
