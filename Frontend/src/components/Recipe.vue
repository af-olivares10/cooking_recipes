<template>
  <div id="container-2">
    <h1><i class="fa fa-arrow-left back" @click="back()"></i> <span class = "recipe-title-main">{{recipe.title}}</span>
      <span class="pencil-icon" @click="showModal">
        <i class="far fa-edit"></i>
      </span>

    </h1>
    <div class = "recipe-content">
      <div class = "image-wrapper">
        <img class="image" :src = "recipe.image">
      </div>
      <div class = "recipe-content-title">
        Decription:
      </div>
      <div class = "description">
        <span  >{{recipe.description}} </span>
      </div>

      <div class = "recipe-content-title">
        Ingredients:
      </div>
      <div class = "ingredients-wrapper">
        <div v-for="(ingredient, index) in recipe.ingredients" :key = "index" class = "ingredient">
          <i class="fas fa-drumstick-bite drumstick"></i>
          <span class = "ingredient-title" >{{ingredient.quantity}} {{ingredient.unit}} of {{ingredient.name}}</span>
        </div>
      </div>
      <div class = "recipe-content-title">
        Directions:
      </div>
      <div class = "directions-wrapper">
        <div v-for="(direction, index) in recipe.directions" :key = "index" class = "ingredient">
          <i class="fas fa-drumstick-bite drumstick"></i>
            <span   class = "ingredient-title">{{direction.direction}}</span>
        </div>
      </div>
    </div>
    <b-modal
    title="Edit a recipe"
    ref="myModalRef"
    hide-footer
    size="lg"
    no-close-on-esc
    no-close-on-backdrop
    hide-header-close
    >
    <form  name = "form">
      <label class = "recipe-content-title-edit">Title:</label>

      <b-form-input
      type="text"
      placeholder="Enter a title for the recipe"
      style = "margin-bottom:10px;"
      class = "edit-field modal-item non-empty"
      v-model="recipe.title"
      @change="handleEdit">
    </b-form-input>

    <label class = "recipe-content-title-edit">Image:</label>

    <b-form-input
    type="text"
    placeholder="Image URL"
    style = "margin-bottom:10px;"
    class = "edit-field modal-item non-empty"
    v-model="recipe.image"
    @change="handleEdit"

    >
  </b-form-input>

    <label class = "recipe-content-title-edit">Description:</label>

    <b-form-input
    placeholder="Enter a description for the recipe"
    :rows="3"
    :max-rows="6"
    v-model="recipe.description"
    style = "margin-bottom:10px;"
    class = "form-control edit-field modal-item non-empty"
    @change="handleEdit"
    >
  </b-form-input>
  <label class = "recipe-content-title-edit">Ingredients:</label>
  <div class = "ingredients-wrapper modal-item">
    <div v-for="(ingredient, index) in recipe.ingredients" :key = "index" class = "ingredient">
      <span class = "ingredient-title" >
        {{index+1}}. Ingredient
        <b-form-input v-model = "recipe.ingredients[index].name" class ="ingredient-edit-item non-empty" @change="handleEdit"></b-form-input>
        Unit <b-form-input v-model = "recipe.ingredients[index].unit" class ="ingredient-edit-item non-empty"@change="handleEdit"></b-form-input>
        Amount <b-form-input v-model = "recipe.ingredients[index].quantity" type = "number" class ="ingredient-edit-item non-empty"@change="handleEdit"></b-form-input>
        <i class="far fa-times-circle" @click ="deleteIngredient(recipe.id, index)"></i>
      </span>
    </div>
    <i class="fas fa-plus-circle new-item-button" @click ="newIngredient(recipe.id)"></i>
  </div>
  <label class = "recipe-content-title-edit">Directions:</label>
  <div class = "ingredients-wrapper modal-item">
    <div v-for="(direction, index) in recipe.directions" :key = "index" class = "ingredient">
      <span class = "ingredient-title" >
        {{index+1}}.
        <b-form-input v-model = "recipe.directions[index].direction" class ="direction-edit-field non-empty" @change="handleEdit"></b-form-input>
        <i class="far fa-times-circle" @click ="deleteDirection(recipe.id, index)"></i>
      </span>

    </div>
    <div style = "width:100%; height:60px;">
      <i class="fas fa-plus-circle new-item-button" @click ="newDirection(recipe.id)"></i>
    </div>

  </div>
</form>
<div style = "color:red">{{errorMessage}}</div>
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
  name: 'Recipe',
  props: {
    recipe: Object
  },
  computed:{
  },
  mounted(){
  },
  data(){
    return {
      modified:false,
      errorMessage:"",
      loading:false
    }
  },
  methods:{
    back(){
      this.$emit("back")
    },
    showModal () {
      this.$refs.myModalRef.show()
    },
    handleEdit(){
      console.log("heme aqui");
      this.modified = true;
    },
    handleOk (evt) {
      this.loading = true;
      evt.preventDefault()
      for(let i of document.forms["form"]){
        if (i.value.trim()===""){
          this.loading = false;
          this.errorMessage = "No field can be empty";
          return;
        }
      }
      if(this.modified)
      {

      fetch(`http://localhost:8081/api/recipes/${this.recipe.id}`, {
        method: 'PUT',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(this.recipe)
      }).then(response => response.json()).then(json=>{
      this.$refs.myModalRef.hide()
      this.loading = false;
      this.errorMessage = "";
      this.$emit("handleOkEdit",json)
      this.modified = false;
      alert("Recipe edited!")
      }
    ).catch(err => console.log(err))
    }
    else{
      this.$refs.myModalRef.hide()
      this.loading = false;
      this.errorMesage = "";
    }
  },
  handleCancel(){
    this.$refs.myModalRef.hide()
    this.$emit("handleCancelEdit")
    this.modified = false

  },
  handleHide(e){
    e.preventDefault();
  },
  newIngredient(id){
    this.$emit("newIngredient",id)
  },
  newDirection(id){
    this.$emit("newDirection",id)
  },
  deleteIngredient(recipeId,index){
    this.$emit("deleteIngredient",recipeId,index)
  },
  deleteDirection(recipeId,index){
    this.$emit("deleteDirection",recipeId,index)
  }
}
}

</script>


<style >
.recipe-content{
  margin: 0;
  padding: 0;
}
.recipe-content-title{
  color: #666;
  font-size: 20px ;
  margin: 10px;
  border-bottom: 1px #d2d2d2 solid;
}
.recipe-content-title-edit{
  color: #666;
  font-size: 20px ;
  margin: 10px 30px;
}
.ingredients-wrapper{
  width: 80%;
}
.directions-wrapper{
  width: 100%;
}
.ingredient{
  margin: 10px;
}
.drumstick{
  color:#febe89;
  font-size:20px;
}
.ingredient-title{
  margin-left: 10px;
}
.description{
  margin: 20px;
}
.pencil-icon{
  text-align: center;
  color: white;
  display: inline-block;
  float: right;
}
.pencil-icon:hover{
  cursor: pointer;
}
.ingredient-edit-item{
  display:inline !important;
  width: 100px !important;
  margin: 0 5px;
}
.edit-field,.direction-edit-field{
  width: 80% !important;
  margin-left: 10px;
}
.direction-edit-field{
  display: inline !important;
  margin-right: 10px;
}
.modal-item{
  margin: 0 auto !important;
}
.new-item-button{
  color:#ff9665;
  font-size: 30px;
  float:right;
}
.new-item-button:hover{
  cursor: pointer;
}
.fa-times-circle{
  opacity: 0;
  font-size: 20px;
  color: #ff5e62;
}
.fa-times-circle:hover{
  cursor: pointer;
}
.ingredient-title:hover .fa-times-circle{
  opacity: 1;
}
.modal-buttons-wrapper{
  border-top: 1px solid #e9ecef;
  padding: 10px;
}
.modal-button{
  width: 150px !important;
  display: inline !important;
  margin: 0 !important;
  float: right;
  padding: 10px;
}
.lds-dual-ring {
  display: inline-block;
  width: 64px;
  height: 64px;
  float: right;
  margin-right: 100px;
}
.lds-dual-ring:after {
  content: " ";
  display: block;
  width: 46px;
  height: 46px;
  margin: 1px;
  border-radius: 50%;
  border: 5px solid #ff9665;
  border-color: #ff9665 transparent #ff9665 transparent;
  animation: lds-dual-ring 1.2s linear infinite;
}
.lds-dual-ring-main {
  display: inline-block;
  width: 64px;
  height: 64px;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}
.lds-dual-ring-main:after {
  content: " ";
  display: block;
  width: 46px;
  height: 46px;
  border-radius: 50%;
  border: 5px solid #fff;
  border-color: #fff transparent #fff transparent;
  animation: lds-dual-ring 1.2s linear infinite;
}
@keyframes lds-dual-ring {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
.image-wrapper{
  width: 400px;
  margin: 20px auto;
}
img{
  width: 100%;
}
</style>
