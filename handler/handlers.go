package handler
import(
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  _"github.com/lib/pq"
  "database/sql"
)

var db *sql.DB

func SetDB(conn *sql.DB){
  db = conn
}
//Recipe Struct (Model)
type Recipe struct{
  ID string `json:"id"`
  Description string `json:"description"`
  Title string `json:"title"`
  Ingredients []Ingredient `json:"ingredients"`
  Directions []Direction `json:"directions"`
  Image string `json:"image"`

}
//Ingredient Struct (Model)
type Ingredient struct{
  Name string `json:"name"`
  Quantity string `json:"quantity"`
  Unit string `json:"unit"`
  ID string `json:"id"`

}
type Direction struct{
  ID string `json:"id"`
  Direction string `json:"direction"`
}
type recipes struct {
  Recipes []Recipe

}


//Get recipes
func GetRecipes(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

  // json.NewEncoder(w).Encode(recipes)
  recipes:= recipes{}
  err := fetchRecipes(&recipes)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  out, err := json.Marshal(recipes.Recipes)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  fmt.Fprintf(w, string(out))
}

func fetchRecipes(resultRecipes *recipes) error {

  rows, err := db.Query(`SELECT id,  description, title, image FROM recipes`)
  if err != nil {
    return err
  }
  defer rows.Close()
  for rows.Next() {

    var resultRecipe = Recipe{Ingredients:[]Ingredient{},Directions:[]Direction{} }
    err = rows.Scan(
      &resultRecipe.ID,
      &resultRecipe.Description,
      &resultRecipe.Title,
      &resultRecipe.Image,

    )
    if err != nil {
      return err
    }

    //Get ingredients
    rowsIngredients, err := db.Query(`
      SELECT
      id,
      name,
      quantity,
      unit
      FROM ingredients
      WHERE recipe=$1

      `,	resultRecipe.ID)
    if err != nil {
      return err
    }
    defer rowsIngredients.Close()
    for rowsIngredients.Next() {
      var ingredient Ingredient
      err = rowsIngredients.Scan(
        &ingredient.ID,
        &ingredient.Name,
        &ingredient.Quantity,
        &ingredient.Unit,
      )
      if err != nil {
        return err
      }
      resultRecipe.Ingredients = append(resultRecipe.Ingredients, ingredient)
    }
    //Get Directions
    rowsDirections, err := db.Query(`
      SELECT
      id,
      direction
      FROM directions
      WHERE recipe=$1

      `,	resultRecipe.ID)
    if err != nil {
      return err
    }
    defer rowsDirections.Close()
    for rowsDirections.Next() {
      var direction Direction
      err = rowsDirections.Scan(
        &direction.ID,
        &direction.Direction,
      )
      if err != nil {
        return err
      }
      resultRecipe.Directions = append(resultRecipe.Directions, direction)
    }

    resultRecipes.Recipes = append(resultRecipes.Recipes, resultRecipe)
  }
  err = rows.Err()
  if err != nil {
    return err
  }
  return nil
}

//Create recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS");

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept");
  if (r.Method == "OPTIONS") {
    var recipe Recipe
    json.NewEncoder(w).Encode(recipe)
    } else {
      var recipe Recipe
      json.NewDecoder(r.Body).Decode(&recipe)
       err := db.QueryRow(`INSERT INTO recipes (description,title,image) values ($1,$2,$3) RETURNING id`, recipe.Description,recipe.Title,recipe.Image ).Scan(&recipe.ID)
      if err != nil {
        http.Error(w,  err.Error(), 500)
        return
      }
      var newIngredients [] Ingredient
      for _, ingredient := range recipe.Ingredients {
        var newIngredient Ingredient
        newIngredient = ingredient
         err := db.QueryRow(`INSERT INTO ingredients (recipe,name,quantity, unit) values ($1,$2,$3,$4) RETURNING id`, recipe.ID,ingredient.Name,ingredient.Quantity,ingredient.Unit ).Scan(&newIngredient.ID)
        if err != nil {
          http.Error(w,  err.Error(), 500)
          return
        }
        newIngredients = append(newIngredients,newIngredient)
      }
      recipe.Ingredients = newIngredients

      var newDirections [] Direction
      for _, direction := range recipe.Directions {
        var newDirection Direction
        newDirection = direction
        err := db.QueryRow(`INSERT INTO directions (recipe,direction) values ($1,$2) RETURNING id`, recipe.ID,direction.Direction).Scan(&newDirection.ID)
        if err != nil {
          http.Error(w, err.Error(), 500)
          return
        }
        newDirections = append(newDirections,newDirection)
      }
      recipe.Directions = newDirections

      json.NewEncoder(w).Encode(recipe)

    }

  }
  //Update recipe
  func UpdateRecipe(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS");

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept");
    if (r.Method == "OPTIONS") {
      var recipe1 Recipe
      json.NewEncoder(w).Encode(recipe1)
      }else{
        var recipe Recipe
        json.NewDecoder(r.Body).Decode(&recipe)
         _,err := db.Exec(`UPDATE recipes SET (description,title,image) = ($1,$2,$4) WHERE id = $3`, recipe.Description,recipe.Title, recipe.ID, recipe.Image )

        if err != nil {
          http.Error(w,  err.Error()+"a4", 500)
          return
        }
        var newIngredients [] Ingredient
        for _, ingredient := range recipe.Ingredients {
          if  len(ingredient.ID) > 0{
            _, err := db.Exec(`UPDATE ingredients SET (name,quantity,unit) = ($1,$2,$3) WHERE id = $4`, ingredient.Name,ingredient.Quantity,ingredient.Unit,ingredient.ID )
            if err != nil {
              http.Error(w,  err.Error()+"a2", 500)
              return
            }
            newIngredients = append(newIngredients,ingredient)
          }else{
              var newIngredient Ingredient
              newIngredient = ingredient
              err := db.QueryRow(`INSERT INTO ingredients (recipe,name,quantity, unit) values ($1,$2,$3,$4) RETURNING id`, recipe.ID,ingredient.Name,ingredient.Quantity,ingredient.Unit ).Scan(&newIngredient.ID)
              if err != nil {
                http.Error(w,  err.Error()+"a1", 500)
                return
              }
              newIngredients = append(newIngredients,newIngredient)
            }
        }
        recipe.Ingredients = newIngredients


        var newDirections [] Direction
        for _, direction := range recipe.Directions {
          if len(direction.ID) > 0 {

            _, err := db.Exec(`UPDATE directions SET (direction) = ($1) WHERE id = $2`, direction.Direction,direction.ID )
            if err != nil {
              http.Error(w, err.Error(), 500)
              return
            }
            newDirections = append(newDirections,direction)

          }else{
              var newDirection Direction
              newDirection = direction
              err := db.QueryRow(`INSERT INTO directions (recipe,direction) values ($1,$2) RETURNING id`, recipe.ID,direction.Direction).Scan(&newDirection.ID)
              if err != nil {
                http.Error(w, err.Error(), 500)
                return
              }
              newDirections = append(newDirections,newDirection)
            }
          }
          recipe.Directions = newDirections

          json.NewEncoder(w).Encode(recipe)
        }
      }

    //Delete recipe
    func DeleteRecipe(w http.ResponseWriter, r *http.Request){
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept");
      w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS");

      id := mux.Vars(r)["id"]
      _, err := db.Exec(`DELETE FROM recipes where ID= $1`, id)
      if err != nil{
        http.Error(w, err.Error(), 500)
        return
      }
      fmt.Fprintf(w, "Deletion complete")
    }
