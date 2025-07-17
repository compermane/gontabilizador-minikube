package presenca

import (
	"log"
	"net/http"
	"strconv"

	"github.com/compermane/gontabilizador/types"
	"github.com/compermane/gontabilizador/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.PresencaStore
}

func NewHandler(store types.PresencaStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) PresencaRoutes(router *mux.Router) {
	router.HandleFunc("/registrar-presenca", h.handlePresenca).Methods("POST")
}

func (h *Handler) handlePresenca(w http.ResponseWriter, r *http.Request) {
	log.Println("[handlePresenca] executed")

	var payload types.RegisterPresencaPayload
	if err := r.ParseForm(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	id_ensaio, _ := strconv.Atoi(r.FormValue("ensaio_id"))
	id_ritmistas := r.Form["presentes"]

	if id_ensaio == 0 {
		return 
	}

	ritmistas_ensaio, err := h.store.ListPresencasPorEnsaio(id_ensaio)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ritmistas_presente := make(map[int][]bool, 0)
	for _, id := range ritmistas_ensaio {
		// {já existe, selecionado}
		ritmistas_presente[id] = []bool{true, false}
	}
	for _, id := range id_ritmistas {
		id_int, _ := strconv.Atoi(id)
		log.Printf("[handlePresenca] %v %v %v %v\n", r.Form["ensaio_id"], id_ensaio, id, id_int)

		if value, ok := ritmistas_presente[id_int]; (ok && value[1] == false) {
			value[1] = true
		} else {
			payload = types.RegisterPresencaPayload{
				IDEnsaio: id_ensaio,
				IDRitmista: id_int,
				Presente: true,
			}
			
			p, err := h.store.BuscarPresencaPorEnsaioIDRitmistaID(id_ensaio, id_int)

			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}

			if p == nil {
				err := h.store.CreatePresenca(types.Presenca{
					IDRitmista: payload.IDRitmista,
					IDEnsaio: payload.IDEnsaio,
					Presente: payload.Presente,
				})

				if err != nil {
					utils.WriteError(w, http.StatusInternalServerError, err)
					return
				}
			} else {
				err := h.store.UpdatePresencaRitmista(id_int, true) 
				if err != nil {
					utils.WriteError(w, http.StatusInternalServerError, err)
					return
				}
			}
		}
	}

	for key, value := range ritmistas_presente {
		if value[1] == false {
			err = h.store.UpdatePresencaRitmista(key, false)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}
		}
	}
	
	log.Printf("[handlePresenca] succesfully executed on ensaio %v and ritmista %v\n", payload.IDEnsaio, payload.IDRitmista)
	http.Redirect(w, r, "/presencas", http.StatusSeeOther)
}
