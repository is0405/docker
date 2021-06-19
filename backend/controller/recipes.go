package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"regexp"
	// "errors"
	// "fmt"
	"strconv"

	"github.com/is0405/docker/model"
	"github.com/is0405/docker/repository"
	"github.com/is0405/docker/service"
	"github.com/jmoiron/sqlx"
)

type Recipes struct {
	db *sqlx.DB
}

func NewRecipes(db *sqlx.DB) *Recipes {
	return &Recipes{db: db}
}

type ResponseOKRecipes struct {
	Message  string          `json:"message"`
	Recipe   []*model.Recipes `json:"recipe"`
}

type ResponseNGRecipes struct {
	Message  string `json:"message"`
	Required string `json:"required"`
}

type ResponseMessageRecipes struct {
	Message  string `json:"message"`
}

func (a *Recipes) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	rawRecipe := &model.Recipes{}

	res := ResponseNGRecipes{
		Message: "Recipe creation failed!",
		Required: "title, making_time, serves, ingredients, cost",
	}
	
	if err := json.NewDecoder(r.Body).Decode(&rawRecipe); err != nil {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}

	//Title
	if rawRecipe.Title == "" {
		// return http.StatusUnprocessableEntity, res, nil 
		return http.StatusOK, res, nil
	}

	if !StringCheck(rawRecipe.Title) {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}

	//MakingTime
	if rawRecipe.MakingTime == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	if !StringCheck(rawRecipe.MakingTime) || !MakingTimeCheck(rawRecipe.MakingTime)  {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}
	
	//Serves
	if rawRecipe.Serves == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	if !StringCheck(rawRecipe.Serves) || !ServesCheck(rawRecipe.Serves){
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}
	
	//Ingridients
	if rawRecipe.Ingridients == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}
	
	if !IngridientsCheck(rawRecipe.Ingridients) {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}

	//Cost
	if rawRecipe.Cost <= 0 {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	Service := service.NewRecipes(a.db)
	id, err := Service.Create(rawRecipe)
	if err != nil {
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	response, err := repository.GetRecipe(a.db, int(id))
	if err != nil {
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	OKres := ResponseOKRecipes{}
	OKres.Message = "Recipe successfully created!"
	var recipe []*model.Recipes
	recipe = append(recipe, response)
	
	OKres.Recipe = recipe;
	
	return http.StatusOK, OKres, nil
}

func (a *Recipes) GetRecipe(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	id, err := URLToID(r)
	if err != nil {
		id = 1
	}

	response, err := repository.GetRecipe(a.db, id)
	if err != nil {
		res := ResponseMessageRecipes{
			Message: "ID does not exist",
		}
		
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	OKres := ResponseOKRecipes{}
	OKres.Message = "Recipe details by id"
	var recipe []*model.Recipes
	recipe = append(recipe, response)
	
	OKres.Recipe = recipe;
	
	return http.StatusOK, OKres, nil
}

type GetAllRecipeResponse struct {
	Recipe   []*model.Recipes `json:"recipes"`
}

func (a *Recipes) GetAllRecipe(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	
	res, err := repository.GetRecipesList(a.db)
	if err != nil {
		msg := ResponseMessageRecipes{
			Message: "Recipe get failed!",
		}
		// return http.StatusInternalServerError, msg, nil
		return http.StatusOK, msg, nil
	}

	var response GetAllRecipeResponse
	response.Recipe = res
	return http.StatusOK, response, nil
}

func (a *Recipes) UpdateRecipe(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	rawRecipe := &model.Recipes{}

	res := ResponseMessageRecipes{
		Message: "Recipe updation failed!",
	}

	id, err := URLToID(r)
	if err != nil {
		id = 1
	}
	
	if err := json.NewDecoder(r.Body).Decode(&rawRecipe); err != nil {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}

	rawRecipe.ID = id

	//Title
	if rawRecipe.Title == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	if !StringCheck(rawRecipe.Title) {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}

	//MakingTime
	if rawRecipe.MakingTime == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	if !StringCheck(rawRecipe.MakingTime) || !MakingTimeCheck(rawRecipe.MakingTime)  {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}
	
	//Serves
	if rawRecipe.Serves == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	if !StringCheck(rawRecipe.Serves) || !ServesCheck(rawRecipe.Serves){
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}
	
	//Ingridients
	if rawRecipe.Ingridients == "" {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}
	
	if !IngridientsCheck(rawRecipe.Ingridients) {
		// return http.StatusBadRequest, res, nil
		return http.StatusOK, res, nil
	}

	//Cost
	if rawRecipe.Cost <= 0 {
		// return http.StatusUnprocessableEntity, res, nil
		return http.StatusOK, res, nil
	}

	Service := service.NewRecipes(a.db)
	_, err = Service.Update(rawRecipe)
	if err != nil {
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	response, err := repository.GetRecipe(a.db, id)
	if err != nil {
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	OKres := ResponseOKRecipes{}
	OKres.Message = "Recipe successfully updated!"
	var recipe []*model.Recipes
	recipe = append(recipe, response)
	
	OKres.Recipe = recipe;
	
	return http.StatusOK, OKres, nil
}

func (a *Recipes) DeleteRecipe(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	id, err := URLToID(r)
	var res ResponseMessageRecipes
	if id == 0 {
		id = 1
	}
	
	_, err = repository.GetRecipe(a.db, id)
	if err != nil {
		res.Message = "No Recipe found"
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	Service := service.NewRecipes(a.db)
	_, err = Service.Delete(id)
	
	if err != nil {
		res.Message = "No Recipe found"
		// return http.StatusInternalServerError, res, nil
		return http.StatusOK, res, nil
	}

	res.Message = "Recipe successfully removed!"
	
	return http.StatusOK, res, nil
}

func MakingTimeCheck(str string) bool {
	chars := []string{"時間", "分", "秒"}
    r := strings.Join(chars, "")
	symbol := regexp.MustCompile("[" + r + "]+")
	if symbol.Match([]byte(str)) {
		//上記が含まれている
		return true
	}
	return false
}

func ServesCheck(str string) bool {
	chars := []string{"人"}
    r := strings.Join(chars, "")
	symbol := regexp.MustCompile("[" + r + "]+")
	if symbol.Match([]byte(str)) {
		//上記が含まれている
		return true
	}
	return false
}

func StringCheck(str string) bool {
	chars := []string{"?", "!", "\\*","\\_", "\\#", "<", ">", "\\\\", "(", ")", "\\$", "\"", "%", "=", "~", "|", "[", "]", ";", "\\+", ":", "{", "}", "@", "\\`", "/", "；", "＠", "＋", "：", "＊", "｀", "「", "」", "｛", "｝", "＿", "？", "。", "、", "＞", "＜"}
    r := strings.Join(chars, "")
	symbol := regexp.MustCompile("[" + r + "]+")
	if symbol.Match([]byte(str)) {
		//上記が含まれている
		return false
	}
	return true
}

func IngridientsCheck(str string) bool {
	chars := []string{"?", "!", "\\*","\\_", "\\#", "<", ">", "\\\\", "(", ")", "\\$", "\"", "%", "=", "~", "|", "[", "]", ";", "\\+", ":", "{", "}", "@", "\\`", "/", "；", "＠", "＋", "：", "＊", "｀", "「", "」", "｛", "｝", "＿", "？", "。", "＞", "＜"}
    r := strings.Join(chars, "")
	symbol := regexp.MustCompile("[" + r + "]+")
	if symbol.Match([]byte(str)) {
		//上記が含まれている
		return false
	}
	return true
}

func URLToID(r *http.Request) (int, error) {
	url := r.URL.Path
	strID := url[strings.LastIndex(url, "/")+1:]
	id, err := strconv.Atoi( strID )
	if err != nil {
		return 0, err
	}
	
	return id, nil
}
