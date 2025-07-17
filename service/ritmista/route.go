package ritmista

import (
	"fmt"
	"log"
	"net/http"

	"github.com/compermane/gontabilizador/types"
	"github.com/compermane/gontabilizador/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.RitmistaStore
}

func NewHandler(store types.RitmistaStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegistroRoutes(router *mux.Router) {
	router.HandleFunc("/registro-ritmista", h.handleRegistro).Methods("POST")
	router.HandleFunc("/ritmistas", h.handleGetAllRitmistas).Methods(http.MethodGet)
}

func (h *Handler) handleGetAllRitmistas(w http.ResponseWriter, r *http.Request) {
	ritmistas, err := h.store.GetAllRitmistas()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return 
	}

	utils.WriteJSON(w, http.StatusOK, ritmistas)
}

func (h *Handler) handleRegistro(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleRegisto] executed")
	var payload types.RegisterRitmistaPayload

	if err := r.ParseForm(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	payload = types.RegisterRitmistaPayload{
		Nome:   r.FormValue("nome"),
		Modulo: r.FormValue("modulo"),
		Naipe:  r.FormValue("naipe"),
	}
	_, err := h.store.GetRitmistaByName(payload.Nome)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Ritmista com nome %v já existe\n", payload.Nome))
		return
	}


	err = h.store.CreateRitmista(types.Ritmista{
		Nome: payload.Nome,
		Modulo: payload.Modulo,
		Naipe: payload.Naipe,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// utils.WriteJSON(w, http.StatusCreated, nil)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}