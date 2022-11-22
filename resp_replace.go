package respReplacer

import "reflect"

func replaceNil(v reflect.Value) {
	if v.Type().Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			replaceNil(v.Field(i))
		}
	} else if v.Kind() == reflect.Slice {
		if v.Len() == 0 {
			v.Set(reflect.MakeSlice(v.Type(), 0, 0))
		}
		for i := 0; i < v.Len(); i++ {
			replaceNil(v.Index(i))
		}
	}
}

func ReplaceResp(resp interface{}) {
	v := reflect.ValueOf(resp)
	if v.IsNil() {
		resp = struct{}{}
	} else {
		replaceNil(v.Elem())
	}
}
