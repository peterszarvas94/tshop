// https://github.com/peterszarvas94/configloader

package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func Load(variables any, keys ...string) error {
	v := reflect.ValueOf(variables)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Load expects a pointer to a struct")
	}

	structValue := v.Elem()

	for i := range structValue.NumField() {
		field := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		if field.Kind() != reflect.String {
			return fmt.Errorf("Field %s must be of type string", fieldType.Name)
		}

		envVarName := strings.ToUpper(fieldType.Name)
		envVarValue, found := os.LookupEnv(envVarName)

		if !found {
			return fmt.Errorf("Environment variable %s not found", envVarName)
		}

		field.SetString(envVarValue)
	}

	return nil
}
