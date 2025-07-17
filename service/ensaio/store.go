package ensaio

import (
	"database/sql"

	"github.com/compermane/gontabilizador/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateEnsaio(ensaio types.Ensaio) error {
	_, err := s.db.Exec("INSERT INTO ensaio (data_ensaio, nome) VALUES (?, ?)", 
						ensaio.Data, ensaio.Nome)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAllEnsaios() ([]*types.Ensaio, error) {
	rows, err := s.db.Query("SELECT * FROM ensaio")
	ensaios := make([]*types.Ensaio, 0)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ensaio, err := scanRowIntoEnsaio(rows)

		if err != nil {
			return nil, err
		}

		ensaios = append(ensaios, ensaio)
	}

	return ensaios, nil
}

func (s *Store) GetEnsaioByID(id int) (*types.Ensaio, error) {
	return nil, nil
}

func scanRowIntoEnsaio(row *sql.Rows) (*types.Ensaio, error) {
	ensaio := new(types.Ensaio)

	err := row.Scan(
		&ensaio.ID,
		&ensaio.Nome,
		&ensaio.Data,
	)

	if err != nil {
		return nil, err
	}

	return ensaio, nil
}