package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

func PopulateQueryParams(ctx context.Context, req *http.Request, queryParams interface{}) {
	queryParamsStructType := reflect.TypeOf(queryParams)
	queryParamsValType := reflect.ValueOf(queryParams)

	for i := 0; i < queryParamsStructType.NumField(); i++ {
		fieldType := queryParamsStructType.Field(i)
		valType := queryParamsValType.Field(i)

		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			continue
		}

		if qpTag.Serialization != "" {
			populateSerializedParams(req, qpTag, fieldType.Type, valType)
		} else {
			// TODO: support other styles
			switch qpTag.Style {
			case "deepObject":
				populateDeepObjectParams(req, qpTag, fieldType.Type, valType)
			case "form":
				populateFormParams(req, qpTag, fieldType.Type, valType)
			}
		}
	}
}

func populateSerializedParams(req *http.Request, tag *paramTag, objType reflect.Type, objValue reflect.Value) {
	if objType.Kind() == reflect.Pointer {
		if objValue.IsNil() {
			return
		}
		objType = objType.Elem()
		objValue = objValue.Elem()
	}
	if objValue.Interface() == nil {
		return
	}

	switch tag.Serialization {
	case "json":
		data, err := json.Marshal(objValue.Interface())
		if err != nil {
			fmt.Printf("error marshaling json: %v", err) // TODO support logging and returning error?
			return
		}
		queryParams := url.Values{}
		queryParams.Add(tag.ParamName, string(data))
		decoded, err := url.QueryUnescape(queryParams.Encode())
		if err != nil {
			fmt.Printf("error decoding query params: %v", err)
			return
		}

		req.URL.RawQuery += decoded
	}
}

func populateDeepObjectParams(req *http.Request, tag *paramTag, objType reflect.Type, objValue reflect.Value) {
	if objType.Kind() == reflect.Pointer {
		if objValue.IsNil() {
			return
		}
		objType = objType.Elem()
		objValue = objValue.Elem()
	}

	queryParams := url.Values{}

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

			queryParams.Add(fmt.Sprintf("%s[%s]", tag.ParamName, qpTag.ParamName), fmt.Sprintf("%v", valType.Interface()))
		}
	case reflect.Map:
		iter := objValue.MapRange()
		for iter.Next() {
			switch iter.Value().Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < iter.Value().Len(); i++ {
					queryParams.Add(fmt.Sprintf("%s[%s]", tag.ParamName, iter.Key().String()), fmt.Sprintf("%v", iter.Value().Index(i).Interface()))
				}
			default:
				queryParams.Add(fmt.Sprintf("%s[%s]", tag.ParamName, iter.Key().String()), fmt.Sprintf("%v", iter.Value().Interface()))
			}
		}
	}

	req.URL.RawQuery += queryParams.Encode()
}

func populateFormParams(req *http.Request, tag *paramTag, objType reflect.Type, objValue reflect.Value) {
	queryParams := populateForm(tag.ParamName, tag.Explode, objType, objValue, func(fieldType reflect.StructField) string {
		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			return ""
		}

		return qpTag.ParamName
	})

	decoded, err := url.PathUnescape(queryParams.Encode())
	if err != nil {
		fmt.Printf("error decoding query params: %v", err)
		return
	}

	req.URL.RawQuery += decoded
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
