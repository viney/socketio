package conf

import (
	"os"

	"engine/conf/config"
)

const (
	envRoot  = "GOPATH"
	filename = "etc.cfg"
)

var cfg *config.Config

func init() {
	c, err := config.ReadDefault(EtcDir(filename))
	if err != nil {
		panic(err)
	}
	cfg = c
}

func EtcDir(filename string) string {
	env := os.Getenv(envRoot)
	if len(env) == 0 {
		panic(envRoot + "not exits")
	}

	return env + "/etc/" + filename
}

func ResDir(filename string) string {
	env := RootPath()

	return env + "/res/" + filename
}

func RootPath() string {
	env := os.Getenv(envRoot)
	if len(env) == 0 {
		panic(envRoot + "not exits")
	}

	return env
}

func String(session, key string) string {
	s, err := cfg.String(session, key)
	if err != nil {
		panic(err)
	}

	return s
}

func Int(session, key string) int {
	i, err := cfg.Int(session, key)
	if err != nil {
		panic(err)
	}

	return i
}
