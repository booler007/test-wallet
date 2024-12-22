package controller

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var (
	ErrBinding     = errors.New("validation error")
	ErrBindingJSON = errors.New("problems with parsing data. Check JSON")
)

type Servicer interface {
	ProcessTheTransaction(walletID, operation string, amount float64) error
}

type APIController struct {
	Service Servicer
}

type inputTransaction struct {
	WalletID      string  `json:"wallet_id" binding:"required,uuid"`
	OperationType string  `json:"operation_type" binding:"required,oneof=DEPOSIT WITHDRAW"`
	Amount        float64 `json:"amount" binding:"required,numeric,gt>0"`
}

func NewAPIController(s Servicer) *APIController {
	return &APIController{s}
}

func (c *APIController) SetupRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")

	apiv1.POST("/wallets", c.ProcessTheTransaction)
	apiv1.GET("/wallets/:id", c.GetBalance)
}

func (c *APIController) ProcessTheTransaction(ctx *gin.Context) {
	var input inputTransaction
	meta, err := bindingInput(ctx, &input)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind).SetMeta(meta)
		return
	}

	err = c.Service.ProcessTheTransaction(input.WalletID, input.OperationType, input.Amount)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *APIController) GetBalance(ctx *gin.Context) {}

type out struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func bindingInput(ctx *gin.Context, input any) ([]out, error) {
	if err := ctx.ShouldBindJSON(&input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			output := make([]out, len(validationErrors))
			for i, fieldError := range validationErrors {
				output[i].Field = fieldError.Field()
				output[i].Message = getErrorMsg(fieldError)
			}
			return output, ErrBinding
		}

		return nil, ErrBindingJSON
	}

	return nil, nil
}

func getErrorMsg(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "this field is required"
	case "uuid":
		return "must be in UUID format"
	case "oneof":
		return "operation must be DEPOSIT or WITHDRAW"
	case "numeric":
		return "must be numeric"
	case "gt":
		return "amount must be greater than 0"
	default:
		return "Unknown error"
	}
}
