package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

const (
	drivername = "DRIVER"
	dbUser     = "DBUSER"
	dbPassword = "PASSWORD"
	dbPort     = "PORT"
	dbName     = "DBNAME"
	dbSSLmode  = "SSLMODE"
)

var (
	driver   string
	user     string
	password string
	port     string
	name     string
	sslmode  string
)

func loadEnv() {
	godotenv.Load(".env")
	driver = os.Getenv(drivername)
	user = os.Getenv(dbUser)
	password = os.Getenv(dbPassword)
	port = os.Getenv(dbPort)
	name = os.Getenv(dbName)
	sslmode = os.Getenv(dbSSLmode)
}

func Connect() {
	loadEnv()
	connString := fmt.Sprintf("user=%s port=%s dbname=%s password=%s sslmode=%s", user, port, name, password, sslmode)
	fmt.Println(connString)
	var err error
	Client, err = sql.Open(driver, connString)
	if err != nil {
		fmt.Println("Postgres Connectivity failed")
		panic(err)
	}
	if pingErr := Client.Ping(); err != nil {
		fmt.Println("Ping err from postgres")
		panic(pingErr)
	}
	fmt.Println("Postgres Connected successfully")
}
