package gorm

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

// Field model field definition
type Field struct {
	*StructField
	IsBlank bool
	Field   reflect.Value
}

// Set set a value to the field
func (field *Field) Set(value interface{}) (err error) {
	if !field.Field.IsValid() {
		return errors.New("field value not valid")
	}

	if !field.Field.CanAddr() {
		return ErrUnaddressable
	}

	// 确保: value为reflect.Value格式
	reflectValue, ok := value.(reflect.Value)
	if !ok {
		reflectValue = reflect.ValueOf(value)
	}

	fieldValue := field.Field
	if reflectValue.IsValid() {
		if reflectValue.Type().ConvertibleTo(fieldValue.Type()) {
			// 如果类型兼容，则默认转
			fieldValue.Set(reflectValue.Convert(fieldValue.Type()))
		} else {
			// 如果是指针，创建一个Element
			if fieldValue.Kind() == reflect.Ptr {
				if fieldValue.IsNil() {
					fieldValue.Set(reflect.New(field.Struct.Type.Elem()))
				}
				// 然后fieldValue切换到Elem
				fieldValue = fieldValue.Elem()
			}

			if reflectValue.Type().ConvertibleTo(fieldValue.Type()) {
				fieldValue.Set(reflectValue.Convert(fieldValue.Type()))
			} else if scanner, ok := fieldValue.Addr().Interface().(sql.Scanner); ok {
				// fieldValue来主动解析 interface{}
				err = scanner.Scan(reflectValue.Interface())
			} else {
				err = fmt.Errorf("could not convert argument of field %s from %s to %s", field.Name, reflectValue.Type(), fieldValue.Type())
			}
		}
	} else {
		// 如果无效，则设置为默认的Zero
		field.Field.Set(reflect.Zero(field.Field.Type()))
	}

	field.IsBlank = isBlank(field.Field)
	return err
}
