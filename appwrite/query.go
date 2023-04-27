package appwrite

import (
	"fmt"
	"reflect"
)

type Query struct {}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) parseValue(value interface{}) string {
	if val, ok := value.(string); ok {
		return fmt.Sprintf("\"%s\"", val)
	}
	return toString(value)
}

func (q *Query) parseQuery(attribute string, method string, value interface{}) string {
	stringValue := ""
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		arr := reflect.Indirect(reflect.ValueOf(value))
		for i := 0; i < arr.Len(); i++ {
			stringValue += q.parseValue(arr.Index(i).Interface())
			if i < arr.Len()-1 {
				stringValue += ","
			}
		}
	default:
		stringValue = q.parseValue(value)
	}
	return fmt.Sprintf("%s(\"%s\", [%s])", method, attribute, stringValue)
}

func (q *Query) Equal(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "equal", value)
}

func (q *Query) NotEqual(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "notEqual", value)
}

func (q *Query) LessThan(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "lessThan", value)
}

func (q *Query) LessThanEqual(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "lessThanEqual", value)
}

func (q *Query) GreaterThan(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "greaterThan", value)
}

func (q *Query) GreaterThanEqual(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "greaterThanEqual", value)
}

func (q *Query) Search(attribute string, value string) string {
	return q.parseQuery(attribute, "search", value)
}

func (q *Query) IsNull(attribute string) string {
	return fmt.Sprintf("isNull(\"%s\")", attribute)
}

func (q *Query) IsNotNull(attribute string) string {
	return fmt.Sprintf("isNotNull(\"%s\")", attribute)
}

func (q *Query) Between(attribute string, start, end interface{}) string {
	return q.parseQuery(attribute, "between", []interface{}{start, end})
}

func (q *Query) StartsWith(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "startsWith", value)
}

func (q *Query) EndsWith(attribute string, value interface{}) string {
	return q.parseQuery(attribute, "endsWith", value)
}

func (q *Query) Select(attributes []string) string {
	stringValue := ""
	for i, attribute := range attributes {
		stringValue += fmt.Sprintf("\"%s\"", attribute)
		if i < len(attributes)-1 {
			stringValue += ","
		}
	}
	return fmt.Sprintf("select([%s])", stringValue)
}

func (q *Query) OrderAsc(attribute string) string {
	return fmt.Sprintf("orderAsc(\"%s\")", attribute)
}

func (q *Query) OrderDesc(attribute string) string {
	return fmt.Sprintf("orderDesc(\"%s\")", attribute)
}

func (q *Query) CursorBefore(documentId string) string {
	return fmt.Sprintf("cursorBefore(\"%s\")", documentId)
}

func (q *Query) CursorAfter(documentId string) string {
	return fmt.Sprintf("cursorAfter(\"%s\")", documentId)
}

func (q *Query) Limit(limit int) string {
	return fmt.Sprintf("limit(%d)", limit)
}

func (q *Query) Offset(offset int) string {
	return fmt.Sprintf("offset(%d)", offset)
}