package error

import (
	"fmt"
	"reflect"
)

func ProcessError(oErr interface{}) {
	if oErr != nil {
		panic(oErr)
	}
}

func Catch(fCallback func(oErr interface{}, aArgs map[string]interface{}), aArgs map[string]interface{}) {
	if e := recover(); e != nil {
		fCallback(e, aArgs)
	}
}

func ParseError(e interface{}) string {
	if _, ok := e.(error); ok {
		return e.(error).Error()
	} else if _, ok := e.(string); ok {
		return e.(string)
	} else {
		return fmt.Sprintf(
			"Unparseable error message for type <%s> and error <%s>",
			reflect.TypeOf(e).String(),
			fmt.Sprint(e),
		)
	}
}
