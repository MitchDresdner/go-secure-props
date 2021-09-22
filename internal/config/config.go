package config

import "github.com/MitchDresdner/secure-props/internal/driver"

type AppConfig struct {
	DB           *driver.DB
	InProduction bool
	Environment  string
	Domain       string
	//MonitorMap    map[int]cron.EntryID
	PreferenceMap map[string]string
	Version       string
	Identifier    string
}
