package util

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e UnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported type: %v", e.Type)
}

type StringOrDictMarshaler struct {
	Value interface{}
}

//nolint:exhaustruct,exhaustive,funlen,cyclop
func (m StringOrDictMarshaler) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	val := reflect.ValueOf(m.Value)

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return UnsupportedTypeError{Type: val.Type()}
	}

	start = xml.StartElement{Name: xml.Name{Local: "dict"}}
	if err := encoder.EncodeToken(start); err != nil {
		return err
	}

	for i := range val.NumField() {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Tag.Get("xml")

		if fieldName == "" {
			fieldName = val.Type().Field(i).Name
		}

		if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "key"}}); err != nil {
			return err
		}

		if err := encoder.EncodeToken(xml.CharData([]byte(fieldName))); err != nil {
			return err
		}

		if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "key"}}); err != nil {
			return err
		}

		//nolint:nestif
		if field.Kind() == reflect.Struct || field.Kind() == reflect.Interface {
			if err := (StringOrDictMarshaler{Value: field.Interface()}.
				MarshalXML(encoder, xml.StartElement{Name: xml.Name{Local: "dict"}})); err != nil {
				return err
			}
		} else {
			var valueTag string

			switch field.Kind() {
			case reflect.String:
				valueTag = "string"
			case reflect.Float64:
				valueTag = "real"
			default:
				valueTag = "string"
			}

			if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: valueTag}}); err != nil {
				return err
			}

			var value string

			switch field.Kind() {
			case reflect.String:
				value = field.String()
			case reflect.Float64:
				value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Int:
				value = strconv.Itoa(int(field.Int()))
			default:
				value = fmt.Sprintf("%v", field.Interface())
			}

			if err := encoder.EncodeToken(xml.CharData([]byte(value))); err != nil {
				return err
			}

			if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: valueTag}}); err != nil {
				return err
			}
		}
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "dict"}}); err != nil {
		return err
	}

	return nil
}
