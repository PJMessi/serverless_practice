package lambdahelper

import (
	"sharedlambdacode/internal/excep"
	"sharedlambdacode/internal/util/structutil"
	"sharedlambdacode/internal/util/validationutil"
)

func MapRequestBody(reqBodyStr string, reqBodyStruct interface{}) error {
	err := structutil.ConvertStrToStruct(reqBodyStr, reqBodyStruct)

	if err != nil {
		return excep.NewInvalidExcep("invalid json payload")
	}

	err = validationutil.ValidateStruct(reqBodyStruct)
	if err != nil {
		return err
	}

	return nil
}
