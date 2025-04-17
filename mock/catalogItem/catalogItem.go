package catalogItem

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CatalogItem struct {
	FieldA string `json:"field_a" validate:"required,min=1"`
	FieldB string `json:"field_b"`
	FieldC string `json:"field_c" validate:"required,oneof=R B G"`
	FieldD int    `json:"field_d" validate:"required,min=0"`
	FieldE string `json:"field_e" validate:"required,email"`
}


func New(fieldA string, fieldB string, fieldC string, fieldD int, fieldE string) (CatalogItem, error) {
	
	c := CatalogItem{
		FieldA: fieldA,
		FieldB: fieldB,
		FieldC: fieldC,
		FieldD: fieldD,
		FieldE: fieldE,
	}
	
	_, err := c.Validate()
	if err != nil {
		return CatalogItem{}, err
	} else {
		return c, nil
	}
}

func (c CatalogItem) Validate() (bool, error) {
	// return validator.New().Struct(c)
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(c)

	if err != nil {
		var validateErr validator.ValidationErrors
		if errors.As(err, &validateErr) {
			for _, err := range validateErr {
				fmt.Println(err.Namespace(), err.Field(), err.StructNamespace(), err.StructField(), err.Tag(), err.ActualTag())
			}
		}
		return false, err
	}
	return true, nil
}