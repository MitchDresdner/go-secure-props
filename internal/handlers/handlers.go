package handlers

import (
	"github.com/MitchDresdner/secure-props/internal/config"
	"github.com/MitchDresdner/secure-props/internal/driver"
	"github.com/MitchDresdner/secure-props/internal/repository"
	"github.com/MitchDresdner/secure-props/internal/repository/dbrepo"
)

//Repo is the repository
var Repo *DBRepo
var app *config.AppConfig

// DBRepo is the db repo
type DBRepo struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewHandlers creates the handlers
func NewHandlers(repo *DBRepo, a *config.AppConfig) {
	Repo = repo
	app = a
}

// NewPostgresqlHandlers creates db repo for postgres
func NewPostgresqlHandlers(db *driver.DB, a *config.AppConfig) *DBRepo {
	return &DBRepo{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}
