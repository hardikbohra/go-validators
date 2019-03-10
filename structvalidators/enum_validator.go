package structvalidators

import (
	"fmt"

	constants "github.com/techbrosource/go-validators/constants"
)

// EnumValidator performs enum validation.
type EnumValidator struct {
	Name string
}

// Validate performs validation on given int value
func (v EnumValidator) Validate(val interface{}) (bool, error) {
	value := val.(string)

	if value == constants.Empty {
		return false, fmt.Errorf(constants.EmptyString)
		// } else {
		// 	switch v.Name {
		// 	case constants.DeviceType:
		// 		if !enum.IsValidDeviceType(value) {
		// 			return false, fmt.Errorf(constants.InvalidEnumString)
		// 		}
		// 	case constants.Source:
		// 		if !enum.IsValidSource(value) {
		// 			return false, fmt.Errorf(constants.InvalidEnumString)
		// 		}
		// 	}
	}
	return true, nil
}

// Validate1 performs validation on given int value
func (v EnumValidator) Validate1(val interface{}) (bool, interface{}) {
	value := val.(string)

	if value == constants.Empty {
		return false, fmt.Errorf(constants.EmptyString)
		// } else {
		// 	switch v.Name {
		// 	case constants.DeviceType:
		// 		if !enum.IsValidDeviceType(value) {
		// 			return false, fmt.Errorf(constants.InvalidEnumString)
		// 		}
		// 	case constants.Source:
		// 		if !enum.IsValidSource(value) {
		// 			return false, fmt.Errorf(constants.InvalidEnumString)
		// 		}
		// 	}
	}
	return true, nil
}