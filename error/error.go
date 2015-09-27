package error

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
