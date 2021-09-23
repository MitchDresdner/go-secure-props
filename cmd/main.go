package main

import (
	"fmt"
	"github.com/MitchDresdner/secure-props/internal/config"
	"github.com/MitchDresdner/secure-props/internal/handlers"
	"github.com/MitchDresdner/secure-props/internal/security"
	"log"
	"os"
	"runtime"
)

var app config.AppConfig
var repo *handlers.DBRepo

const securePropsVersion = "1.0.0"
const maxWorkerPoolSize = 5
const maxJobMaxWorkers = 5

func init() {
	//gob.Register(models.User{})
	_ = os.Setenv("TZ", "America/NewYork")
}

func main() {

	args := cmdLine()
	appInit(args)

	fmt.Println("Done")
}

func cmdLine() map[string]string {
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

	return args
}

// banner display
func banner() {
	// print info
	log.Printf("******************************************")
	log.Printf("** %sSecureProps%s v%s built in %s", "\033[31m", "\033[0m", securePropsVersion, runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")
}
