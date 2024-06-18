package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cringe-link-shortener/pkg/configToml"
	"cringe-link-shortener/pkg/database"

	"cringe-link-shortener/internal/app"
	"cringe-link-shortener/internal/postgres"

	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

func main() {
	overrideTables := flag.Bool("override_tables", false, "Override tables in database")
	configPath := flag.String("config", "config.toml", "Config file path")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERR\t", log.Ldate|log.Ltime)

	config := struct {
		Addr string
		DB   postgres.Config
	}{}

	err := configToml.Get(*configPath, &config)
	if err != nil {
		errorLog.Fatal("error while getting config ", err)
	}

	var db *sql.DB
	if dsn := config.DB.Dsn(); dsn == "" {
		db, err = database.OpenViaEnvVars("postgres")
	} else {
		db, err = database.OpenViaDsn(dsn, "postgres")
	}
	if err != nil {
		errorLog.Fatal("error while connecting to the db ", err)
	}
	defer db.Close()

	DbHandler, err := postgres.Get(db, *overrideTables)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := app.CreateApp(DbHandler, errorLog)

	sigQuit := make(chan os.Signal, 2)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)
	eg := new(errgroup.Group)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		}
	})

	go app.Run(config.Addr)

	if err := eg.Wait(); err != nil {
		infoLog.Printf("gracefully shutting down the server: %v\n", err)
		app.Shutdown()
	}
}
