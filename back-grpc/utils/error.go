package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateFields(data interface{}) error {

	validate := validator.New()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	violations := []*errdetails.BadRequest_FieldViolation{}

	for _, err := range err.(validator.ValidationErrors) {
		st := reflect.TypeOf(data)
		var field string
		json_name, e := st.FieldByName(err.Field())
		if e {
			field = json_name.Tag.Get("json")
		} else {
			field = strings.ToLower(err.Field())
		}
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       field,
			Description: err.Tag(),
		})

	}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusDetails, err := statusInvalid.WithDetails(badRequest)

	if err != nil {
		return nil
	}

	return statusDetails.Err()

}
