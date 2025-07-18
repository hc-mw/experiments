package reflection

import "reflect"

func Walk(x interface{}, fn func(ip string)) {
	v := getValue(x)
	var getField func(int) reflect.Value
	numOfValues := 0

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		numOfValues = v.NumField()
		getField = v.Field
	case reflect.Slice, reflect.Array:
		numOfValues = v.Len()
		getField = v.Index
	case reflect.Map:
		for _, key := range v.MapKeys() {
			Walk(v.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if val, ok := v.Recv(); ok {
				Walk(val.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		res := v.Call(nil)
		for _, r := range res {
			Walk(r.Interface(), fn)
		}
	}

	for i := range numOfValues {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
