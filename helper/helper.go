package helper

import "fmt"

func SliceToSql[T comparable](slice []T, scope string) string {
	var prefix, suffix string
	switch scope {
	case "(":
		prefix = "("
		suffix = ")"
	case "{":
		prefix = "{"
		suffix = "}"
	case "[":
		prefix = "["
		suffix = "]"
	default:
		prefix = "("
		suffix = ")"
	}

	query := prefix
	for i, v := range slice {
		query += fmt.Sprintf("%v", v)
		if i < len(slice)-1 {
			query += ", "
		}
	}
	query += suffix
	return query
}

func SliceToString[T comparable](slice []T) string {
	str := ""
	for i, v := range slice {
		str += fmt.Sprintf("%v", v)
		if i < len(slice)-1 {
			str += ", "
		}
	}
	return str
}
