package databases

import (
	"database/sql"
	"fmt"
	"strings"

	g "service/auth/global"
	"service/config"
	db "service/pkg/database"
)

func SetConnections(cons map[string]*sql.DB) error {
	mainOrTest := "test"
	if !g.CFG.Debug {
		mainOrTest = "main"
	}
	for k, v := range cons {
		dbName := strings.Split(k, ",")[0]
		dbType := strings.Split(k, ",")[1]
		if dbName == mainOrTest {
			g.DB = v
		}
		switch dbType {
		case "postgres":
			g.PostgresCons[dbName] = v
		case "sqlite3":
			g.SqliteCons[dbName] = v
		case "mysql":
			g.MySQLCons[dbName] = v
		case "mssql":
			g.SqlServerCons[dbName] = v
		default:
			return fmt.Errorf("%s database not supported", strings.Split(k, ",")[1])
		}
		g.AllCons[dbName] = v
	}

	return nil
}

func Setup(cfg *config.Config) error {
	cons, err := db.New(cfg.CurrentMicroservice.Databases)
	if err != nil {
		return err
	}

	err = SetConnections(cons)
	if err != nil {
		return err
	}

	return nil
}
