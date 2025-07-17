package ritmista

import (
	"database/sql"
	"fmt"

	"github.com/compermane/gontabilizador/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetRitmistaByName(nome string) (*types.Ritmista, error) {
	rows, err := s.db.Query("SELECT * FROM ritmista WHERE nome = ?", nome)

	if err != nil {
		return nil, err
	}

	u := new(types.Ritmista)
	for rows.Next() {
		u, err = scanRowIntoRitmista(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("Ritmista not found")
	}

	return u, nil
}

func (s *Store) GetAllRitmistas() ([]*types.Ritmista, error) {
	rows, err := s.db.Query("SELECT * FROM ritmista")
	if err != nil {
		return nil, err
	}

	ritmistas := make([]*types.Ritmista, 0)
	for rows.Next() {
		p, err := scanRowIntoRitmista(rows)
		if err != nil {
			return nil, err
		}

		ritmistas = append(ritmistas, p)
	}

	return ritmistas, nil
}

func (s *Store) GetRitmistaByID(id int) (*types.Ritmista, error) {
	return nil, nil
}

func (s *Store) CreateRitmista(user types.Ritmista) error {
	_, err := s.db.Query("INSERT INTO ritmista (nome, modulo, naipe) VALUES (?, ?, ?)",
						user.Nome,
						user.Modulo,
						user.Naipe)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoRitmista(row *sql.Rows) (*types.Ritmista, error) {
	ritmista := new(types.Ritmista)

	err := row.Scan(
		&ritmista.ID,
		&ritmista.Nome,
		&ritmista.Modulo,
		&ritmista.Naipe,
	)

	if err != nil {
		return nil, err
	}

	return ritmista, nil
}