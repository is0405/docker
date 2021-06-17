package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"regexp"
	"errors"

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

type CreateResponseOKRecipes struct {
	Message  string          `json:"message"`
	Recipe   *model.Recipes `json:"recipe"`
}

type CreateResponseNGRecipes struct {
	Message  string `json:"message"`
	Required string `json:"required"`
}

func (a *Recipes) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	rawRecipe := &model.Recipes{}

	res := CreateResponseNGRecipes{
		Message: "Recipe creation failed!",
		Required: "title, making_time, serves, ingredients, cost",
	}
	
	if err := json.NewDecoder(r.Body).Decode(&rawRecipe); err != nil {
		return http.StatusBadRequest, res, err
	}

	//Title
	if rawRecipe.Title == "" {
		return http.StatusUnprocessableEntity, res, errors.New("required parameter is missing:title")
	}

	if !StringCheck(rawRecipe.Title) {
		return http.StatusBadRequest, res, errors.New("invalid is missing:title")
	}

	//MakingTime
	if rawRecipe.MakingTime == "" {
		return http.StatusUnprocessableEntity, res, errors.New("required parameter is missing:MakingTime")
	}

	if !StringCheck(rawRecipe.MakingTime) || !MakingTimeCheck(rawRecipe.MakingTime)  {
		return http.StatusBadRequest, res, errors.New("invalid is missing:MakingTime")
	}
	
	//Serves
	if rawRecipe.Serves == "" {
		return http.StatusUnprocessableEntity, res, errors.New("required parameter is missing:Serves")
	}

	if !StringCheck(rawRecipe.Serves) || !ServesCheck(rawRecipe.Serves){
		return http.StatusBadRequest, res, errors.New("invalid is missing:Serves")
	}
	
	//Ingridients
	if rawRecipe.Ingridients == "" {
		return http.StatusUnprocessableEntity, res, errors.New("required parameter is missing:Ingridient")
	}
	
	if !IngridientsCheck(rawRecipe.Ingridients) {
		return http.StatusBadRequest, res, errors.New("invalid is missing:Ingridient")
	}

	//Cost
	if rawRecipe.Cost <= 0 {
		return http.StatusUnprocessableEntity, res, errors.New("required parameter is missing:Cost")
	}

	Service := service.NewRecipes(a.db)
	id, err := Service.Create(rawRecipe)
	if err != nil {
		return http.StatusInternalServerError, res, err
	}

	response, err := repository.GetRecipe(a.db, int(id))
	if err != nil {
		return http.StatusInternalServerError, res, err
	}
	// res = CreateResponseOKRecipes{
	// 	Message: "Recipe successfully created!",
	// 	Recipe: [
	// 		{
	// 			"id": id,
	// 			"title": rawRecipe.Title,
	// 			"making_time": rawRecipe.MakingTime,
	// 			"serves": rawRecipe.Serves,
	// 			"ingredients": rawRecipe.Ingridients,
	// 			"cost": rawRecipe.Cost,
	// 			"created_at": rawRecipe.CreatedAt,
	// 			"updated_at": rawRecipe.UpdatedAt
	// 		}],
	// }

	OKres := CreateResponseOKRecipes{
		Message: "Recipe successfully created!",
		Recipe: response,
	}
	
	return http.StatusOK, OKres, nil
}

func (a *Recipes) Update(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusNotFound, nil, nil
}

func (a *Recipes) Destroy(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusNotFound, nil, nil
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

