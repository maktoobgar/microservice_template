package g

import (
	"database/sql"
	_ "embed"

	"service/config"
	"service/pkg/logging"
)

//go:embed version
var Version string

//go:embed name
var Name string

// Config
var CFG *config.Config = nil

// Utilities
var Logger logging.Logger = nil

// AppSecret
var SecretKey []byte = nil

// Default DB
var DB *sql.DB = nil

// Connections
var PostgresCons = map[string]*sql.DB{}
var SqliteCons = map[string]*sql.DB{}
var MySQLCons = map[string]*sql.DB{}
var SqlServerCons = map[string]*sql.DB{}
var AllCons = map[string]*sql.DB{}
