package goconf

import "reflect"

func GetOrDefault(obj, defaultValue any, fieldNames... string) any {
	// 获取对象的反射值
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // 解引用指针
	}

	// 遍历字段路径
	for _, fieldName := range fieldNames {
		if val.Kind() != reflect.Struct {
			return defaultValue
		}

		// 查找字段
		val = val.FieldByName(fieldName)
		if !val.IsValid() || !val.CanInterface() {
			return defaultValue
		}

		// 如果字段值为指针，解引用
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				return defaultValue
			}
			val = val.Elem()
		}
	}

	// 如果最终值是零值，返回默认值
	if reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface()) {
		return defaultValue
	}

	// 返回字段值
	return val.Interface()
}