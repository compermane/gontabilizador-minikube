package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/compermane/gontabilizador/types"
)

type PageHandler struct {
	RitmistaStore types.RitmistaStore
	EnsaioStore   types.EnsaioStore
	PresencaStore types.PresencaStore
}

type DataForPresenca struct {
	Ritmistas []*types.Ritmista
	Ensaios   []*types.Ensaio
    SelectedEnsaioID  int
    PresencaMap       map[int]bool 
}

func NewPageHandler(rs types.RitmistaStore, es types.EnsaioStore, ps types.PresencaStore) *PageHandler {
	return &PageHandler{
		RitmistaStore: rs,
		EnsaioStore: es,
		PresencaStore: ps,
	}
}

func (ph *PageHandler) Home(w http.ResponseWriter, r *http.Request) {
	ritmistas, err := ph.RitmistaStore.GetAllRitmistas()
	if err != nil {
		http.Error(w, "Erro ao buscar ritmistas", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("static", "templates", "home.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Erro ao carregar template:", err)
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ritmistas)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
	}
}

func (ph *PageHandler) Ensaios(w http.ResponseWriter, r *http.Request) {
	ensaios, err := ph.EnsaioStore.GetAllEnsaios()
	if err != nil {
		http.Error(w, "Erro ao buscar ensaios", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("static", "templates", "ensaios.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Erro ao carregar template:", err)
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ensaios)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
	}
}

func (ph *PageHandler) Presencas(w http.ResponseWriter, r *http.Request) {
	ritmistas, err := ph.RitmistaStore.GetAllRitmistas()
	if err != nil {
		log.Println("[Presencas] erro ao buscar ritmistas")
		http.Error(w, "Erro ao buscar ritmistas", http.StatusInternalServerError)
		return
	}

	ensaios, err := ph.EnsaioStore.GetAllEnsaios()
	if err != nil {
		log.Println("[Presencas] erro ao buscar ensaios")
		http.Error(w, "Erro ao buscar ensaios", http.StatusInternalServerError)
		return
	}

	ensaioID, _  := strconv.Atoi(r.URL.Query().Get("ensaio_id"))
	presencas, _ := ph.PresencaStore.ListPresencasPorEnsaio(ensaioID)
	
	presencaMap := map[int]bool{}
	for _, p := range presencas {
		presencaMap[p] = true
	}

	data := DataForPresenca{Ritmistas: ritmistas, Ensaios: ensaios, SelectedEnsaioID: ensaioID, PresencaMap: presencaMap}

	tmplPath := filepath.Join("static", "templates", "presencas.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("[Presencas] Erro ao carregar template:", err)
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
		return
	}

	log.Println("[Presencas] SelectedEnsaioID:", ensaioID)
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("[Presencas] Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
		return
	}
}