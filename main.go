package main

import(
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  _"github.com/lib/pq"
  "database/sql"
  "os"
  handler "github.com/af-olivares10/cooking_recipes_api/handler"

)
var db *sql.DB

func initDb() {

  // Connect to the database.
  var err error
  fmt.Println("Connecting to data base...")
  config := dbConfig()

  db, err = sql.Open("postgres", "postgresql://"+config["DBUSER"]+":"+config["DBPASS"]+"@"+config["DBHOST"]+":"+config["DBPORT"]+"/"+config["DBNAME"])

  if err != nil {
    log.Fatal("error connecting to the database: ", err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }

    // Create the "ingridents" table.
    // if _, err := db.Exec(
    //   "DROP TABLE  ingredients "); err != nil {
    //     log.Fatal(err)
    //   }
    //   // Create the "directions" table.
    //   if _, err := db.Exec(
    //     "DROP TABLE directions "); err != nil {
    //       log.Fatal(err)
    //     }
    //     if _, err := db.Exec(
    //       "DROP TABLE  recipes"); err != nil {
    //         log.Fatal(err)
    //       }
  if _, err := db.Exec(
    "CREATE TABLE IF NOT EXISTS recipes (id SERIAL PRIMARY KEY, description STRING, title STRING, image STRING)"); err != nil {
      log.Fatal(err)
    }
    // Create the "ingridents" table.
    if _, err := db.Exec(
      "CREATE TABLE IF NOT EXISTS ingredients (id SERIAL PRIMARY KEY, name STRING, quantity STRING, unit STRING, recipe INT NOT NULL REFERENCES recipes (id) ON DELETE CASCADE ON UPDATE CASCADE)"); err != nil {
        log.Fatal(err)
      }
      // Create the "directions" table.
      if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS directions (id SERIAL PRIMARY KEY, direction STRING, recipe INT NOT NULL REFERENCES recipes (id) ON DELETE CASCADE ON UPDATE CASCADE)"); err != nil {
          log.Fatal(err)
        }
        fmt.Println("Successfully connected!")

}
func main(){
  r:= mux.NewRouter()
  initDb()
  defer db.Close()
  handler.SetDB(db)
  r.HandleFunc("/api/recipes",handler.GetRecipes).Methods("GET","OPTIONS")
  r.HandleFunc("/api/recipes",handler.CreateRecipe).Methods("POST","OPTIONS")
  r.HandleFunc("/api/recipes/{id}",handler.UpdateRecipe).Methods("PUT","OPTIONS")
  r.HandleFunc("/api/recipes/{id}",handler.DeleteRecipe).Methods("DELETE","OPTIONS")
  buildHandler := http.FileServer(http.Dir("static"))
  r.PathPrefix("/").Handler(buildHandler)
  fmt.Println("Server running!")

  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"),r))
}
func dbConfig() map[string]string {
    conf := make(map[string]string)
    user, ok := os.LookupEnv("DBUSER")
    if !ok {
        panic("DBUSER environment variable required but not set")
    }
    password, ok := os.LookupEnv("DBPASS")
    if !ok {
        panic("DBPASS environment variable required but not set")
    }
    host, ok := os.LookupEnv("DBHOST")
    if !ok {
        panic("DBHOST environment variable required but not set")
    }
    port, ok := os.LookupEnv("DBPORT")
    if !ok {
        panic("DBPORT environment variable required but not set")
    }
    name, ok := os.LookupEnv("DBNAME")
    if !ok {
        panic("DBNAME environment variable required but not set")
    }
    conf["DBUSER"] = user
    conf["DBPASS"] = password
    conf["DBHOST"] = host
    conf["DBPORT"] = port
    conf["DBNAME"] = name

    return conf
}
