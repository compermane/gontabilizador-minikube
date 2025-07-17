package ensaio

import (
	"log"
	"net/http"
	"time"

	"github.com/compermane/gontabilizador/types"
	"github.com/compermane/gontabilizador/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.EnsaioStore
}

func NewHandler(store types.EnsaioStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegistroRoutes(router *mux.Router) {
	router.HandleFunc("/registro-ensaio", h.handleCriarEnsaio).Methods("POST")
}

func (h *Handler) handleCriarEnsaio(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleCriarEnsaio] excuted")

	var payload types.RegisterEnsaioPayload
	if err := r.ParseForm(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data, _ := time.Parse("2006-01-02", r.FormValue("data"))
	payload = types.RegisterEnsaioPayload{
		Nome: r.FormValue("nome"),
		Data: data,
	}

	err := h.store.CreateEnsaio(types.Ensaio{
		Data: payload.Data,
		Nome: payload.Nome,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, "/ensaios", http.StatusSeeOther)
}