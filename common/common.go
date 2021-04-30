package common

import (
	"errors"
	"fmt"
	"reflect"
)

// UserName is readonly and returns the read value
func UserName(user interface{}) (string, error) {
	ot := reflect.TypeOf(user)
	typeName := ot.Name()

	isUserType := typeName == "RegularUser" || typeName == "PrivilegedUser"
	if !isUserType {
		return "", fmt.Errorf("passed interface is not of user type: %v", typeName)
	}

	ov := reflect.ValueOf(user)
	sf := ov.FieldByName("Name")

	return sf.Interface().(string), nil
}

// SetUserName works ONLY if the "user inteface{}" argument is given from the caller as a pointer to an underlying struct
func SetUserName(user interface{}, name string) error {
	ov := reflect.ValueOf(user)
	if ov.Kind() != reflect.Ptr {
		return errors.New("passed value must be a pointer to a user type")
	}

	ot := ov.Elem().Type()
	typeName := ot.Name()

	isUserType := typeName == "RegularUser" || typeName == "PrivilegedUser"
	if !isUserType {
		return fmt.Errorf("passed interface is not of user type: %v", typeName)
	}

	ov.Elem().FieldByName("Name").SetString(name)
	return nil
}
