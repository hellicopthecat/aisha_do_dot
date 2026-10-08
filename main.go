package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hellicopthecat/aisha_do_dot/users"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	pg "github.com/lib/pq"
)

func main() {
	// echo init
	app := echo.New()

	// dev mod args
	dev_mod := os.Args[1]
	fmt.Printf("args %s", dev_mod)

	// env load
	if err := godotenv.Load(); err != nil {
		log.Println(".env Load is Failed")
	}

	var cfg pg.Config
	if dev_mod == "dev" {
		// db connect
		cfg = pg.Config{
			Host:           "localhost",
			Port:           5432,
			User:           "root",
			Password:       "1111",
			ConnectTimeout: 5 * time.Second,
		}
	} else {
		// db connect
		num, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
		if err != nil {
			log.Fatalf("DB PW NOT LOADED :: %s", err)

		}
		cfg = pg.Config{
			Host:           os.Getenv("DATABASE_HOST"),
			Port:           uint16(num),
			User:           os.Getenv("DATABASE_USER"),
			Password:       os.Getenv("DATABASE_PW"),
			ConnectTimeout: 5 * time.Second,
		}
	}
	c, err := pg.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatalf("DB Connect Is Failed %s", err)
	}
	db := sql.OpenDB(c)
	defer db.Close()

	// App Use

	// App Service
	users.InitUserModule(db)

	// Initialized App
	if err := app.Start(":8080"); err != nil {
		app.Logger.Error("Fail to Start Server", "Error", err)
	}
}
