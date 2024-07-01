package shared

import "reflect"

/*
PatchStruct
*/
func PatchStruct(target interface{}, source interface{}) {
	targetValue := reflect.ValueOf(target).Elem()
	sourceValue := reflect.ValueOf(source).Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceValue.Field(i)
		fieldName := sourceValue.Type().Field(i).Name

		// Check if the field is set and is not the default value
		if !isZeroValue(field) {
			targetField := targetValue.FieldByName(fieldName)
			if targetField.CanSet() {
				targetField.Set(field)
			}
		}
	}
}

func isZeroValue(v reflect.Value) bool {
	zeroValue := reflect.Zero(v.Type()).Interface()
	currentValue := v.Interface()

	return reflect.DeepEqual(currentValue, zeroValue)
}
