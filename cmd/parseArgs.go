package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//func parseArgs() (*cli.App, error) {
func parseArgs() (*cli.App, map[string]string, error) {

	args := make(map[string]string)

	var secret, encode, decode string
	//var filePath string

	app := &cli.App {
		Name: "propsec",
		Usage: "Encode/Decode application properties",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:        "secret",
				Aliases: []string{"s"},
				Usage:       "password for decoding properties",
				Destination: &secret,
				EnvVars: []string{"C0NF_SECRET"},
				//Required: true,
			},
			&cli.StringFlag{
				Name:        "encode",
				Aliases: []string{"e"},
				Usage:       "string to encode",
				Destination: &encode,
				//Required: true,
			},
			&cli.StringFlag{
				Name:        "decode",
				Aliases: []string{"d"},
				Usage:       "base64 encoded string",
				Destination: &decode,
				//Required: true,
			},

			//&cli.StringFlag{
			//	Name: "filePath",
			//	Aliases: []string{"f"},
			//	Usage: "Configuration file path",
			//	Value: "./config",
			//	Destination: &filePath,
			//	//FilePath: "./config",
			//},
			//&cli.BoolFlag{Name: "decode", Usage: "Decode config file", Aliases: []string{"d"}},
			//&cli.BoolFlag{Name: "encode", Usage: "Encode config file", Aliases: []string{"e"}},
		},
		Action: func(c *cli.Context) error {
			//name := "someone"
			//if c.NArg() > 0 {
			//	name = c.Args().Get(0)
			//}
			if secret == "s3cr3t" {
				fmt.Println("Consider using a stronger secret than", secret)
			} // else {
			//	fmt.Println("Got secret", secret)
			//	fmt.Println("Aka", name)
			//	fmt.Println("argc", c.NArg())
			//}

			args["secret"] = secret
			//args["filePath"] = filePath
			//args["encode"] = strconv.FormatBool(c.Bool("encode"))
			//args["decode"] = strconv.FormatBool(c.Bool("decode"))
			args["encode"] = encode
			args["decode"] = decode

			return nil
		},
		//Commands: []*cli.Command{
		//	{
		//		Name:    "complete",
		//		Aliases: []string{"c"},
		//		Usage:   "complete a task on the list",
		//		Action:  func(c *cli.Context) error {
		//			return nil
		//		},
		//	},
		//	{
		//		Name:    "add",
		//		Aliases: []string{"a"},
		//		Usage:   "add a task to the list",
		//		Action:  func(c *cli.Context) error {
		//			return nil
		//		},
		//	},
		//},
	}

	return app, args, nil
}
