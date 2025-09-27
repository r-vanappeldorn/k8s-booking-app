package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"trips-service.com/src/models"
	"trips-service.com/src/router"
)

type ContinentController struct {
	r *router.Router
}

func NewContinentController(r *router.Router) *ContinentController {
	return &ContinentController{r}
}

func (c *ContinentController) Mount(r *router.Router) {
	r.Post("/continent", c.Create)
}

type CreateContinentRequest struct {
	Name string `json:"name" validate:"required,min=4"`
	Code string `json:"code" validate:"required,min=2,max=2"`
}

type CreateValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getValidationErrorResponse(err error) []CreateValidationErrorResponse {
	var res []CreateValidationErrorResponse

	for _, e := range err.(validator.ValidationErrors) {
		message := ""
		switch e.Field() {
		case "Name":
			message = "Name must be atleast 4 characters long"
		case "Code":
			message = "Code must exactly be 2 characters long"
		}

		res = append(res, CreateValidationErrorResponse{
			Field:   e.Field(),
			Message: message,
		})
	}

	return res
}

func (c *ContinentController) Create(w http.ResponseWriter, req *http.Request, ctx *router.Conext) {
	var body CreateContinentRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad request"})

		return
	}

	defer req.Body.Close()

	if err := ctx.Validator.Struct(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(getValidationErrorResponse(err))

		return
	}

	continent := &models.Continent{
		Code: body.Code,
		Name: body.Name,
	}

	err := ctx.GormDB.Create(continent).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		w.WriteHeader(http.StatusBadRequest)
		res := []CreateValidationErrorResponse{
			{
				Field:   "code",
				Message: "Country already exists",
			},
		}
		json.NewEncoder(w).Encode(&res)

		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Something went wrong"})
		ctx.Logger.Error(fmt.Sprintf("error while trying to create continent: %v", err))

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(continent)
}
