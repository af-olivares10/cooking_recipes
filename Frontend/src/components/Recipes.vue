<template>
  <div id="container-1">
    <h1>My recipes
      <b-form-input  class ="search" placeholder="Find a recipe" @keyup.native = "handleSearch" v-model = "searchKey"></b-form-input>
       <i class="fa fa-plus" @click="showModal"></i>
     </h1>


    <ul v-if= "recipes">
      <li v-for="(recipe, index) in recipes" :key = "index" >
        <span class = "recipe-title"  @click= "recipeSelected(recipe)">
          {{recipe.title}}
        </span>
        <span class="trash-icon"  @click= "deleteRecipe(recipe)"><i class="fa fa-trash "  ></i>
        </span>

      </li>
    </ul>
    <div v-else class="lds-dual-ring-main" ></div>
    <b-modal
    title="Delete a recipe"
    ref="deleteModal"
    hide-footer
    size="md"
    >
    <div>Do you want to delete the recipe {{recipeToDelete.title}}? </div>
    <div class = "modal-buttons-wrapper" >
      <div class = "modal-button">
        <b-button block style="background:#2980b9; border:none;"  @click="handleOkDelete" >Ok</b-button>
      </div>
      <div class = "modal-button">
        <b-button block   style="background:#ff5e62; border:none;" @click="handleCancelDelete" > Cancel </b-button>
      </div>
    </div>
  </b-modal>


  <b-modal
  title="Create a recipe"
  ref="myModalRef"
  hide-footer
  size="lg"
  no-close-on-esc
  no-close-on-backdrop
  hide-header-close
  >
  <form name = "form-new">
    <label class = "recipe-content-title-edit">Title:</label>

    <b-form-input
    type="text"
    placeholder="Enter a title for the recipe"
    style = "margin-bottom:10px;"
    class = "edit-field modal-item non-empty"
    v-model="recipe.title"
    >
  </b-form-input>

  <label class = "recipe-content-title-edit">Image:</label>

  <b-form-input
  type="text"
  placeholder="Image URL"
  style = "margin-bottom:10px;"
  class = "edit-field modal-item non-empty"
  v-model="recipe.image"
  >
</b-form-input>
  <label class = "recipe-content-title-edit">Description:</label>
  <b-form-input
  placeholder="Enter a description for the recipe"
  v-model="recipe.description"
  style = "margin-bottom:10px;"
  class = " edit-field modal-item non-empty"
  >
</b-form-input>
<label class = "recipe-content-title-edit">Ingredients:</label>
<div class = "ingredients-wrapper modal-item">
  <div v-for="(ingredient, index) in recipe.ingredients" :key = "index" class = "ingredient">
    <span class = "ingredient-title" >
      {{index+1}}. Ingredient
      <b-form-input v-model = "recipe.ingredients[index].name" class ="ingredient-edit-item non-empty" ></b-form-input>
      Unit <b-form-input v-model = "recipe.ingredients[index].unit" class ="ingredient-edit-item non-empty"></b-form-input>
      Amount <b-form-input v-model = "recipe.ingredients[index].quantity" type = "number" class ="ingredient-edit-item non-empty"></b-form-input>
      <i class="far fa-times-circle" @click ="deleteIngredient(index)"></i>
    </span>
  </div>
  <i class="fas fa-plus-circle new-item-button" @click ="newIngredient()"></i>
</div>
<label class = "recipe-content-title-edit">Directions:</label>
<div class = "ingredients-wrapper modal-item">
  <div v-for="(direction, index) in recipe.directions" :key = "index" class = "ingredient">
    <span class = "ingredient-title" >
      {{index+1}}.
      <b-form-input v-model = "recipe.directions[index].direction" class ="direction-edit-field non-empty" ></b-form-input>
      <i class="far fa-times-circle" @click ="deleteDirection( index)"></i>
    </span>

  </div>
  <div style = "width:100%; height:60px;">
    <i class="fas fa-plus-circle new-item-button" @click ="newDirection()"></i>
  </div>

</div>
</form>
<div style = "color:red; text-align:center;">{{errorMessage}}</div>
<div class = "modal-buttons-wrapper" v-if="!loading">
  <div class = "modal-button">
    <b-button block style="background:#2980b9; border:none;"  @click="handleOk" >Ok</b-button>
  </div>
  <div class = "modal-button">
    <b-button block   style="background:#ff5e62; border:none;" @click="handleCancel" > Cancel </b-button>
  </div>
</div>
<div v-else class="lds-dual-ring" ></div>
</b-modal>

</div>

</template>

<script>
export default {
  name: 'Recipes',
  props: {
    recipes: Array
  },
  computed:{
  },
  data(){
    return {
      recipe:{directions:[{direction:""}],ingredients:[{name:"",quantity:0,unit:""}]},
      loading:false,
      errorMessage:"",
      recipeToDelete: {},
      searchKey: ""
    }
  },
  methods:{
    recipeSelected(recipe){
      this.$emit("recipeSelected",recipe)
    },
    deleteRecipe(recipe){
      this.recipeToDelete = recipe;
      this.$refs.deleteModal.show();

    },
    handleOkDelete(){
      fetch(`http://localhost:8081/api/recipes/${this.recipeToDelete.id}`,{
        method: 'DELETE',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
      }).then(response=>{
        this.$emit("deleteRecipe",this.recipeToDelete.id);
        this.recipeToDelete = {};
        this.$refs.deleteModal.hide();

      }).catch(err=>console.error(err));
    },
    handleCancelDelete(){
      this.$refs.deleteModal.hide();
      this.recipeToDelete = {};
    },
    handleOk(evt){
      this.loading = true;
      evt.preventDefault()
      for(let i of document.forms["form-new"]){
        if (i.value.trim()===""){
          this.loading = false;
          this.errorMessage = "No field can be empty";
          return;
        }
      }
      fetch(`http://localhost:8081/api/recipes`, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(this.recipe)
      }).then( response => response.json()).then((json2)=>{
        this.$refs.myModalRef.hide()
        this.loading = false;
        this.errorMessage = "";
        this.$emit("newRecipe", json2);
      }
    ).catch(err => console.error(err))
  },
  showModal(){
    this.$refs.myModalRef.show()
  },
  handleCancel(){
    this.$refs.myModalRef.hide()
    this.recipe = {directions:[{direction:""}],ingredients:[{name:"",quantity:0,unit:""}]};
  },
  newIngredient(){
    this.recipe.ingredients.push({name:"",quantity:0,unit:""});
  },
  deleteIngredient(index){
    this.recipe.ingredients = this.recipe.ingredients.filter((ingredient,i) => i!== index);

  },
  newDirection(){
    this.recipe.directions.push({direction:""});
  },
  deleteDirection(index){
    this.recipe.directions= this.recipe.directions.filter((direction,i) => i!== index);
  },
  handleSearch(){
    this.$emit("handleSearch",this.searchKey);
  }
}
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
h1{
  background: #2980b9!important;
  color: white!important;
  margin: 0!important;
  padding: 10px 20px!important;
  text-transform: uppercase!important;
  font-size: 24px!important;
  font-weight: normal!important;
}
ul{
  list-style: none;
  margin: 0;
  padding: 0;
}
li{
  background: white;
  height: 40px;
  line-height: 40px;
  color: #666;
}
li:nth-child(2n){
  background: #f7f7f7;
}
input{
  color: #2980b9;
  font-size: 18px;
  background-color: #ebebeb;
  width: 100%;
  padding: 13px 13px 13px 20px;
  box-sizing: border-box; /*para que el width (y el height) incluya al padding (normalmente solo incluye al contenido)*/
  border: none;
  display: none;
}
input:focus{
  background: white;
  border: 3px solid #2980b9;
  outline: none;
}
#container-1{
  width: 30%;
}
#container-2{
  width: 50%;
  color: #666;
  padding-bottom: 10px;
}
#container-1,#container-2{
  margin: 100px auto;
  background: #f7f7f7;
  box-shadow: 0 0 3px rgba(0,0,0, 0.1);
}
.trash-icon:hover{
cursor: pointer;
}

.trash-icon{
  background: #e74c3c;
  z-index: 99;
  height: 40px;
  text-align: center;
  color: white;
  display: inline-block;
  opacity: 0;
  width: 0;
  transition: 0.2s linear;
  float: right;
}
li:hover .trash-icon{
  width: 40px;
  opacity: 1.0;
}
li:active .trash-icon{
  width: 40px;
  opacity: 1.0;
}

.recipe-title:hover {
  cursor:pointer;
}
.recipe-title{
  margin-left: 20px;
}
.recipe-title-main{
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}
.back:hover{
  cursor: pointer;
}
.fa-plus:hover{
  cursor: pointer;
}
.search{
  width:40% !important;
  display: inline-block!important;
  border-radius: 20px!important;
  margin: 0  0 0 40px ;
}
</style>
