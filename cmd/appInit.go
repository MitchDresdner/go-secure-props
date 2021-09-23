package main

import (
	"fmt"
	"github.com/MitchDresdner/secure-props/internal/config"
	"github.com/MitchDresdner/secure-props/internal/driver"
	"github.com/MitchDresdner/secure-props/internal/handlers"
	"github.com/MitchDresdner/secure-props/internal/security"
	"github.com/magiconair/properties"
	"log"
	"os"
	"strings"
)

func appInit(args map[string]string) (*string, error) {

	p, _ := getProps(args)

	dbUser := p.GetString("db.user", "")
	dbPass := p.GetString("db.pass", "")
	dbHost := p.GetString("db.host", "localhost")
	dbPort := p.GetString("db.port", "")
	dbName := p.GetString("db.name", "")
	dbSsl := p.GetString("db.ssl", "")

	// DB Connection
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		fmt.Println("Missing required fields.")
		os.Exit(1)
	}

	log.Println("Connecting to database....")
	dsnString := ""

	// when developing locally, we often don't have a db password
	if dbPass == "" {
		dsnString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			dbHost,
			dbPort,
			dbUser,
			dbName,
			dbSsl)
	} else {
		dsnString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			dbHost,
			dbPort,
			dbUser,
			dbPass,
			dbName,
			dbSsl)
	}

	// Use stub db to show concepts
	db, err := driver.ConnectStub(dsnString)
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}

	// define application configuration
	a := config.AppConfig{
		DB:           db,
		Domain:       "localhost",   // pass as flag or envvar: *domain,
		Identifier:   "secureProps", // pass as unique instance id: *identifier,
		InProduction: false,         // pass as flag: *inProduction,
		Properties:   p,
		Version:      securePropsVersion,
	}

	app = a

	repo = handlers.NewPostgresqlHandlers(db, &app)
	handlers.NewHandlers(repo, &app)

	log.Println("Setting preferences...")
	preferenceMap := make(map[string]string)
	preferenceMap["fix-me"] = "foo"
	preferenceMap["fix-me-too"] = "bar"

	app.PreferenceMap = preferenceMap

	return nil, nil
}

func getProps(args map[string]string) (*properties.Properties, error) {

	loc, _ := os.Getwd()
	myEnv := getEnv()
	fmt.Println("Running in ", myEnv, loc)

	p := properties.MustLoadFile("./config/dev.properties", properties.UTF8)
	//p.SetComment("db.host", "Database hostname")
	//fmt.Println("-- Props ", p.GetString("db.host", "unk host"), p.GetComment("db.host"),p.GetInt("db.port", -99))

	//p.SetValue("my.prop", "hello prop")
	//fmt.Println("-- Props ", p.GetString("my.prop", "unk my.prop"))

	for k, v := range p.Map() {
		if strings.HasPrefix(v, "![") && strings.HasSuffix(v, "]") {
			msg, err := security.Decode(args["secret"], v[2:len(v)-1])
			if err != nil {
				fmt.Println("Error", err)
				os.Exit(-1)
			}

			p.SetValue(k, msg)
		}
	}

	return p, nil
}

// getEnv sets the default environment
//  See others:
//		https://github.com/joho/godotenv
//		https://github.com/spf13/viper
func getEnv() string {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}
	return env
}
