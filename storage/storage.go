package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

func New(cfg *Config) *Storage {
	return &Storage{
		config: cfg,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Database connection seccessfully!")
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) User() *UserRepository {
	return &UserRepository{
		storage: s,
	}
}

func (s *Storage) Article() *ArticleRepository {
	return &ArticleRepository{
		storage: s,
	}
}
