package utils

import (
	"github.com/MitchDresdner/secure-props/internal/config"
)

var app *config.AppConfig

// NewState manages application state
func NewState(a *config.AppConfig) {
	app = a
}

func GetSecret() string {
	return app.ClArgs["secret"]
}
