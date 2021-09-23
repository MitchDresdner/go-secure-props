package config

import (
	"github.com/MitchDresdner/secure-props/internal/driver"
	"github.com/magiconair/properties"
)

type AppConfig struct {
	DB           *driver.DB
	InProduction bool
	Environment  string
	Domain       string
	//MonitorMap    map[int]cron.EntryID
	PreferenceMap map[string]string
	Properties    *properties.Properties
	Version       string
	Identifier    string
}
