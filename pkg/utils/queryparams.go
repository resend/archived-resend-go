package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

func PopulateQueryParams(ctx context.Context, req *http.Request, queryParams interface{}) error {
	queryParamsStructType := reflect.TypeOf(queryParams)
	queryParamsValType := reflect.ValueOf(queryParams)

	values := url.Values{}

	for i := 0; i < queryParamsStructType.NumField(); i++ {
		fieldType := queryParamsStructType.Field(i)
		valType := queryParamsValType.Field(i)

		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			continue
		}

		if qpTag.Serialization != "" {
			vals, err := populateSerializedParams(req, qpTag, fieldType.Type, valType)
			if err != nil {
				return err
			}
			for k, v := range vals {
				for _, vv := range v {
					values.Add(k, vv)
				}
			}
		} else {
			switch qpTag.Style {
			case "deepObject":
				vals := populateDeepObjectParams(req, qpTag, fieldType.Type, valType)
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			case "form":
				vals := populateFormParams(req, qpTag, fieldType.Type, valType)
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			default:
				return fmt.Errorf("unsupported style: %s", qpTag.Style)
			}
		}
	}

	req.URL.RawQuery = values.Encode()

	return nil
}

func populateSerializedParams(req *http.Request, tag *paramTag, objType reflect.Type, objValue reflect.Value) (url.Values, error) {
	values := url.Values{}

	if objType.Kind() == reflect.Pointer {
		if objValue.IsNil() {
			return values, nil
		}
		objValue = objValue.Elem()
	}
	if objValue.Interface() == nil {
		return values, nil
	}

	switch tag.Serialization {
	case "json":
		data, err := json.Marshal(objValue.Interface())
		if err != nil {
			return nil, fmt.Errorf("error marshaling json: %v", err)
		}
		values.Add(tag.ParamName, string(data))
	}

	return values, nil
}

func populateDeepObjectParams(req *http.Request, tag *paramTag, objType reflect.Type, objValue reflect.Value) url.Values {
	values := url.Values{}

	if objType.Kind() == reflect.Pointer {
		if objValue.IsNil() {
			return values
		}
		objType = objType.Elem()
		objValue = objValue.Elem()
	}

	switch objType.Kind() {
	case reflect.Struct:
		for i := 0; i < objType.NumField(); i++ {
			fieldType := objType.Field(i)
			valType := objValue.Field(i)

			if fieldType.Type.Kind() == reflect.Pointer {
				if valType.IsNil() {
					continue
				}
				valType = valType.Elem()
			}

			qpTag := parseQueryParamTag(fieldType)
			if qpTag == nil {
				continue
			}

			switch valType.Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < valType.Len(); i++ {
					values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, qpTag.ParamName), valToString(valType.Index(i).Interface()))
				}
			default:
				values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, qpTag.ParamName), valToString(valType.Interface()))
			}
		}
	case reflect.Map:
		iter := objValue.MapRange()
		for iter.Next() {
			switch iter.Value().Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < iter.Value().Len(); i++ {
					values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, iter.Key().String()), valToString(iter.Value().Index(i).Interface()))
				}
			default:
				values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, iter.Key().String()), valToString(iter.Value().Interface()))
			}
		}
	}

	return values
}

func populateFormParams(req *http.Request, tag *paramTag, objType reflect.Type, objValue reflect.Value) url.Values {
	return populateForm(tag.ParamName, tag.Explode, objType, objValue, func(fieldType reflect.StructField) string {
		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			return ""
		}

		return qpTag.ParamName
	})
}

type paramTag struct {
	Style         string
	Explode       bool
	ParamName     string
	Serialization string
}

func parseQueryParamTag(field reflect.StructField) *paramTag {
	return parseParamTag(queryParamTagKey, field, "form", true)
}
