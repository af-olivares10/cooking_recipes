<template>


  <div id="app">
    <transition  name="slide-fade">
    <div v-if="!recipe" key =1>
      <Recipes @recipeSelected="recipeSelected" :recipes = "recipes" @newRecipe = "newRecipe" @deleteRecipe = "deleteRecipe" @handleSearch = "handleSearch"/>
    </div>
    <div v-else key =2>
      <Recipe :recipe = "recipe" @back="back" @newIngredient="newIngredient" @newDirection="newDirection"  @deleteIngredient="deleteIngredient" @deleteDirection="deleteDirection" @handleCancelEdit = "handleCancelEdit" @handleOkEdit = "handleOkEdit"/>
    </div>
  </transition>
  </div>
</template>

<script>
import Recipes from './components/Recipes.vue'
import Recipe from './components/Recipe.vue'

export default {
  name: 'app',
  components: {
    Recipes,
    Recipe
  },
  data(){
    return {
      recipe:null,
      recipes:null,
      recipeCopy: null,

    }
  },
  mounted(){
    fetch("http://localhost:8081/api/recipes").then(res =>res.json()).then(json=>{
      this.recipes = (json||[]);
      this.recipesCopy = this.recipes;
    }).catch(error=>console.error(error))
  },
  methods:{
    recipeSelected(recipe){
      this.recipe = recipe;
      this.recipeCopy = JSON.parse(JSON.stringify(this.recipe));
    },
    back(){
      this.recipe = null;
    },
    newIngredient(id){
      this.recipe.ingredients.push({name:"",quantity:0,unit:""});
    },
    newDirection(id){
      this.recipe.directions.push({direction:""});
    },
    deleteIngredient(id, index){
       this.recipe.ingredients = this.recipe.ingredients.filter((ingredient,i) => i!== index);
    },
    deleteDirection(id,index){
       this.recipe.directions= this.recipe.directions.filter((direction,i) => i!== index);
    },
    handleOkEdit(editedRecipe){
      this.recipes[this.recipes.findIndex(r => r.id==editedRecipe.id)] = editedRecipe;
      this.recipe = editedRecipe;
      this.recipeCopy = JSON.parse(JSON.stringify(this.recipe));

    },
    handleCancelEdit(){
      this.recipe = JSON.parse(JSON.stringify(this.recipeCopy));
    },
    newRecipe(recipe){
      this.recipesCopy.push(recipe);
      this.recipes = this.recipesCopy;
    },
    deleteRecipe(id){
      this.recipesCopy = this.recipes.filter((i) => i.id !== id);
      this.recipes = this.recipesCopy;
    },
    handleSearch(searchKey){
      this.recipes= this.recipesCopy.filter(r => r.title.toUpperCase().startsWith(searchKey.toUpperCase()) );
    }
  }
}
</script>

<style>
@import url('https://use.fontawesome.com/releases/v5.6.3/css/all.css');
@import url('https://fonts.googleapis.com/css?family=Varela+Round');
body{
  font-family: 'Varela Round', sans-serif !important;
  background: #ff9966;  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #ff5e62, #ff9966);  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to right, #ff5e62, #ff9966); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
  overflow-x: hidden;

}
.fa-plus{
  float: right;
}
.slide-fade-enter-active {
  transition: all .2s ease;
  transition-delay: .2s;

}
.slide-fade-leave-active {
  transition: all .2s ease;
}
 .slide-fade-leave-to
/* .slide-fade-leave-active below version 2.1.8 */ {
  transform: translateX(100px);
  opacity: 0;
}
.slide-fade-enter{
  transform: translateX(-100px);
  opacity: 0;
}
</style>
