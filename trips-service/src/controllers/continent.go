package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	utilerrors "trips-service.com/src/errors"
	"trips-service.com/src/middleware"
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
	r.Post("/continent", c.Create, middleware.NewAuthMidleware)
}

type CreateContinentRequest struct {
	Name string `json:"name" validate:"required,min=4"`
	Code string `json:"code" validate:"required,min=2,max=2"`
}

func getValidationErrorResponse(err error) []*utilerrors.FieldError {
	var res []*utilerrors.FieldError

	for _, e := range err.(validator.ValidationErrors) {
		message := ""
		switch e.Field() {
		case "Name":
			message = "Name must be atleast 4 characters long"
		case "Code":
			message = "Code must exactly be 2 characters long"
		}

		res = append(res, &utilerrors.FieldError{
			Field:   e.Field(),
			Message: message,
		})
	}

	return res
}

func (c *ContinentController) Create(w http.ResponseWriter, req *http.Request, ctx *router.Conext) {
	var body CreateContinentRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		e := utilerrors.NewJSONErrorResponse(http.StatusBadRequest, "INVALID_REQUEST", "invalid request body provided")
		utilerrors.WriteErrorResponse(w, e)

		return
	}

	defer req.Body.Close()

	if err := ctx.Validator.Struct(&body); err != nil {
		e := utilerrors.NewJSONErrorResponse(http.StatusBadRequest, "INVALID_FIELD", "invalid value for field provided")
		e.AddFieldErrors(getValidationErrorResponse(err))
		utilerrors.WriteErrorResponse(w, e)

		return
	}

	continent := &models.Continent{
		Code: body.Code,
		Name: body.Name,
	}

	err := ctx.GormDB.Create(continent).Error
	var mysqlErr *mysql.MySQLError
	if err != nil && errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		w.WriteHeader(http.StatusBadRequest)
		fields := []*utilerrors.FieldError{
			{
				Field:   "code",
				Message: "Country already exists",
			},
		}

		e := utilerrors.NewJSONErrorResponse(http.StatusBadRequest, "INVALID_FIELD", "invalid value for field provided")
		e.AddFieldErrors(fields)
		e.AddFieldErrors(getValidationErrorResponse(err))
		utilerrors.WriteErrorResponse(w, e)

		return
	}

	if err != nil {
		e := utilerrors.NewJSONErrorResponse(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Something went wrong")
		utilerrors.WriteErrorResponse(w, e)
		ctx.Logger.Error(fmt.Sprintf("error while trying to create continent: %v", err))

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(continent)
}
