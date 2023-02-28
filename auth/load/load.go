package load

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"service/auth/databases"
	g "service/auth/global"
	"service/build"
	iconfig "service/config"
	"service/pkg/colors"
	"service/pkg/config"
	"service/pkg/logging"

	migrate "github.com/rubenv/sql-migrate"
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

	if _, ok := cfg.Microservices[g.Name]; ok {
		cfg.CurrentMicroservice = cfg.Microservices[g.Name]
	} else {
		log.Fatalf("Microservice definition for %s not found", g.Name)
	}
	g.SecretKey = []byte(cfg.SecretKey)
	g.CFG = cfg
}

// Run dbs
func initialDBs() {
	err := databases.Setup(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	var ok bool = false
	if !g.CFG.Debug {
		_, ok = g.AllCons["main"]
		if !ok {
			log.Fatalln(errors.New("'main' db is not defined (required)"))
		}
	} else {
		_, ok = g.AllCons["test"]
		if !ok {
			log.Fatalln(errors.New("'test' db is not defined"))
		}
	}
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

func migrateLatestChanges(db *sql.DB) {
	mainOrTest := "test"
	if !g.CFG.Debug {
		mainOrTest = "main"
	}
	migrations := &migrate.FileMigrationSource{
		Dir: fmt.Sprintf("migrations/%s/", mainOrTest),
	}

	n, err := migrate.Exec(db, g.CFG.CurrentMicroservice.Databases[mainOrTest].Type, migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	if n > 0 {
		fmt.Printf("\n%s==%sMigrations%s==%s\n\n", colors.Cyan, colors.Green, colors.Cyan, colors.Reset)
		fmt.Printf("Applied %s%d%s migrations!\n", colors.Red, n, colors.Reset)
	}
}

func init() {
	setPwd()
	initializeConfigs()
	initialDBs()
	migrateLatestChanges(g.DB)
	initialLogger()
}
