package flatten

func Flatten(list interface{}) []interface{} {
	switch list.(type) {
	case []interface{}:
		result := []interface{}{}
		for _, item := range list.([]interface{}) {
			for _, element := range Flatten(item) {
				result = append(result, element)
			}
		}
		return result
	case nil:
		return []interface{}{}
	}
	return []interface{}{list}
}
