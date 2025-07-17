package presenca

import (
	"database/sql"
	"log"

	"github.com/compermane/gontabilizador/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreatePresenca(presenca types.Presenca) error {
	log.Printf("[CreatePresenca] executed on ensaio %v and ritmista %v\n", presenca.IDEnsaio, presenca.IDRitmista)
	_, err := s.db.Query("INSERT INTO presenca (ritmista_id, ensaio_id, present) VALUES (?, ?, ?)",
						 presenca.IDRitmista,
						 presenca.IDEnsaio,
						 presenca.Presente)

	if err != nil {
		return err
	}

	return nil
}
func (s *Store) BuscarEnsaioPorID(ensaio_id int) (*types.Presenca, error) {
	rows, err := s.db.Query("SELECT * FROM ensaio WHERE ensaio_id = ?", ensaio_id)

	if err != nil {
		return nil, err
	}

	var p *types.Presenca = nil
	for rows.Next() {
		p, err = scanRowsIntoPresenca(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) BuscarPresencaPorEnsaioIDRitmistaID(ensaio_id, ritmista_id int) (*types.Presenca, error) {
	rows, err := s.db.Query("SELECT * FROM presenca WHERE ensaio_id = ? AND ritmista_id = ?", ensaio_id, ritmista_id)

	if err != nil {
		return nil, err
	}

	var p *types.Presenca = nil
	for rows.Next() {
		p, err = scanRowsIntoPresenca(rows)

		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) UpdatePresencaRitmista(ritmista_id int, presenca bool) error {
	_, err := s.db.Query("UPDATE presenca SET present = ? WHERE ritmista_id = ?", presenca, ritmista_id)

	if err != nil {
		return err
	}


	return nil
}

func (s *Store) ListPresencasPorEnsaio(ensaio_id int) ([]int, error) {
	rows, err := s.db.Query("SELECT ritmista_id FROM presenca WHERE ensaio_id = ? AND present = TRUE", ensaio_id)

	if err != nil {
		return nil, err
	}

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func scanRowsIntoPresenca(rows *sql.Rows) (*types.Presenca, error) {
	presenca := new(types.Presenca)

	err := rows.Scan(
		&presenca.IDEnsaio,
		&presenca.IDRitmista,
		&presenca.Presente,
		&presenca.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return presenca, nil
}