# cooking_recipes
Recipes application made with Vue and Golang
The app is deployed in the following URL: https://recipesolivares.herokuapp.com/


## Features

**Public server:**
- Heroku.

**Back - End:**
- CockroachDB
- Golang

**Front - End:**
- Vuejs


## Deployment

**prerequisites:**
* Install Golang and set $GOPATH/bin to your $Path
* Install godep
* Set environment variables (DBHOST,DBPORT,DBNAME,DBUSER,DBPASS,PORT)

* Download the repository and run on the root 
  ```
  > godep go install ./…
  ```
* Run the app
  ```
  > go run man.go
  ```  

And you are all set up! Access the app on http://localhost:$PORT/
