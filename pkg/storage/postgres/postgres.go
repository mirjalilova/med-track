package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mirjalilova/med-track/config"
)

type Postgres struct {
	Db *sql.DB
}

func DatabaseConnection() (*Postgres, error) {
	cfg := config.Load()

	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Db: db,
	}, nil
}

func (p *Postgres) Close() {
	p.Db.Close()
}
