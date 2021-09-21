package main

import (
	"fmt"
	"github.com/MitchDresdner/secure-props/internal/security"
	"github.com/magiconair/properties"
	"log"
	"os"
	"strings"
)

const NONCE_SIZE = 8

func main() {

	app, args, err := parseArgs()
	if err != nil {
		fmt.Println("main::parseArgs", err)
		os.Exit(-1)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if args["encode"] != "" {
		encStr, err := security.Encode(args["secret"], args["encode"])
		if err != nil {
			os.Exit(-1)
		}

		fmt.Println("Encoded: ", encStr)

		//msg, err := security.Decode(args["secret"], encStr)
		//if err != nil {
		//	os.Exit(-1)
		//}
		//
		//fmt.Println("Decoded: ", msg)
	}

	if args["decode"] != "" {
		//msg, err := security.Decode(args["secret"], "SR4kG7IsTOMyo8/4kB+gF1LRWUKmQyodoV8XOw==")
		msg, err := security.Decode(args["secret"], args["decode"])
		if err != nil {
			os.Exit(-1)
		}

		fmt.Println("Secret is:", msg)
	}
}

// props playing with properties
func props() {

	loc, _ := os.Getwd()
	myEnv := getEnv()
	fmt.Println("Running in ", myEnv, loc)

	p := properties.MustLoadFile("./config/dev.properties", properties.UTF8)
	p.SetComment("db.host", "Database hostname")
	fmt.Println("-- Props ", p.GetString("db.host", "unk host"), p.GetComment("db.host"),p.GetInt("db.port", -99))

	p.SetValue("my.prop", "hello prop")
	fmt.Println("-- Props ", p.GetString("my.prop", "unk my.prop"))

	for k, v := range p.Map() {
		fmt.Println( k, ": ",v, p.GetComment(k) )
		if strings.HasPrefix(v, "![") && strings.HasSuffix(v, "]") {
			fmt.Println("Shhh ...")
		}
	}
}

// getEnv sets the default environment
//  See others:
//		https://github.com/joho/godotenv
//		https://github.com/spf13/viper
func getEnv() string{
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}
	return env
}
