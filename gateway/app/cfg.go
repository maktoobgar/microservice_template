package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	g "service/gateway/global"
	"service/gateway/middleware"

	iconfig "service/config"

	"service/pkg/config"
	"service/pkg/logging"
	"service/pkg/translator"

	"service/build"

	"golang.org/x/text/language"
)

var (
	cfg       = &iconfig.Config{}
	languages = []language.Tag{language.English, language.Persian}
)

// Set Project PWD
func setPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for parent := pwd; true; parent = filepath.Dir(parent) {
		if _, err := os.Stat(filepath.Join(parent, "go.mod")); err == nil {
			cfg.PWD = parent
			break
		}
	}
	os.Chdir(cfg.PWD)
}

// Initialization for config files in configs folder
func initializeConfigs() {
	// Loads default config, you just have to hard code it
	if err := config.ParseYamlBytes(build.Config, cfg); err != nil {
		log.Fatalln(err)
	}

	if err1, err2 := config.Parse("../env.yaml", cfg, false), config.Parse("../env.yml", cfg, false); err1 != nil || err2 != nil {
		if err1 != nil {
			log.Fatalln(err1)
		} else if err2 != nil {
			log.Fatalln(err2)
		}
	}

	g.CFG = cfg
	middleware.AddHeaders(strings.Split(g.CFG.AllowHeaders, ","))
}

// Translator initialization
func initialTranslator() {
	t, err := translator.New(build.Translations, languages[0], languages[1:]...)
	if err != nil {
		log.Fatalln(err)
	}
	g.Translator = t
}

// Logger initialization
func initialLogger() {
	cfg.Logging.Path += "/" + g.Name
	k := cfg.Logging
	opt := logging.Option(k)
	l, err := logging.New(&opt, cfg.Debug)
	if err != nil {
		log.Fatalln(err)
	}
	g.Logger = l
}

// Finds and places microservices into g
func initialMicroservices() {
	if auth, ok := g.CFG.Microservices["auth"]; ok {
		g.AuthMic = &auth
	} else {
		log.Fatalln("auth microservice is not defined")
	}
}

// Server initialization
func init() {
	setPwd()
	initializeConfigs()
	initialTranslator()
	initialLogger()
	initialMicroservices()
}
