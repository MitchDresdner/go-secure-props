# A Secure Property manager written in Go
*Create secrets for inclusion in property files, etc*

### Usage

`propsec -h`
 
```
NAME:
   propsec - Encode/Decode application properties

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --secret value, -s value  password for decoding properties  [ or env var $C0NF_SECRET]
   --encode value, -e value  string to encode
   --decode value, -d value  base64 encoded string
   --help, -h                show help (default: false)
```

*Secret can be stored in env var*
