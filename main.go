package main

import (
	"boyzgenk/coba/api"
	"boyzgenk/coba/database"
	"boyzgenk/coba/user"
	"boyzgenk/coba/user/crud"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-redis/redis"

	"github.com/jmoiron/sqlx"
)

type cfg struct {
	dbconn      *sqlx.DB
	sessionconn *redis.Client
}

var config cfg

//this value usually stored somewhere, not hardcode
//this is just for testing purpose
var (
	dbType  = "postgres"
	connStr = "user=postgres dbname=postgres password=example host=127.0.0.1 port=1234 sslmode=disable"
	addr    = "127.0.0.1:6379"
	pass    = ""
)

//Assign all key value here
func init() {
	dbconn, err := database.NewPostgreDB(dbType, connStr)
	if err != nil {
		panic(err)
	}

	sesconn, err := database.NewRedisCache(addr, pass)
	if err != nil {
		panic(err)
	}

	config.dbconn = dbconn
	config.sessionconn = sesconn
}

func main() {
	//for testing purpose, we will create table user if not exist
	dep := crud.NewUserCRUD(config.dbconn)
	err := dep.CreateTable()
	if err != nil {
		panic(err)
	}

	u := user.NewUser(config.dbconn, config.sessionconn)

	a := api.NewAPI(u, dep)

	//Only for testing purpose
	//So we can test it easily
	r := mux.NewRouter()
	r.HandleFunc("/login/{email}/{password}", a.Login)
	r.HandleFunc("/logout/{id}", a.Logout)
	r.HandleFunc("/user/get/{email}/{password}", a.GetUser)
	r.HandleFunc("/user/insert/{email}/{address}/{password}", a.InsertUser)
	r.HandleFunc("/user/update/{email}/{address}/{password}/{id}", a.UpdateUser)
	r.HandleFunc("/user/delete/{id}", a.DeleteUser).Methods("GET")

	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
