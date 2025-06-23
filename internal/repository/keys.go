package repository

import (
	"database/sql"
	"errors"
	"github.com/MatheusVict/streamer-app/internal/model"
	"github.com/labstack/gommon/log"
)

var QueryErr = errors.New("Error to find the stream key")

type KeysRepository interface {
	FindStreamKey(name, key string) (*model.Keys, error)
}

type keysRepository struct {
	*sql.DB
}

func newKeyRepository(db *sql.DB) KeysRepository {
	return &keysRepository{
		db,
	}
}

func (r *keysRepository) FindStreamKey(name, key string) (*model.Keys, error) {
	keys := &model.Keys{}
	row := r.QueryRow(`SELECT * FROM "Lives" WHERE "name"=$1 AND "stream_key"=$2`, name, key)
	err := row.Scan(&keys.Name, &keys.Key)

	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Keys{}, nil
		}
		return &model.Keys{}, QueryErr
	}
	return keys, nil
}
