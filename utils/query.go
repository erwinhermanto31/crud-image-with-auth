package utils

import (
	"bytes"
	"reflect"
	"strings"
)

// Query struct binder for default query param
type Query struct {
	Sort   string `json:"sort"`
	Filter map[string]interface{}
}

const (
	defaultPage  = 1
	defaultCount = 15
	limitPage    = 9999
)

// NewQuery initiate query
func NewQuery(sort string, filter map[string]interface{}) *Query {
	return &Query{
		Sort:   sort,
		Filter: filter,
	}
}

// Operator string transalation
var Operator = map[string]string{
	"gt":      ">",    // elastic
	"lt":      "<",    // elastic
	"eq":      "=",    // elastic
	"ne":      "!=",   // elastic
	"gte":     ">=",   // elastic
	"lte":     "<=",   // elastic
	"like":    "like", // elastic - text search
	"in":      "in",   // elastic
	"notin":   "in",   // elastic
	"null":    "is null",
	"notnull": "is not null",
}

func (q *Query) Where() (string, []interface{}) {
	query := new(bytes.Buffer)
	var args []interface{}
	i := 0
	for k, v := range q.Filter {
		fields := strings.Split(k, "$")
		columnName := fields[0]
		opr := translateOperator(fields[1])
		isRequire := func(s string) bool {
			return s[len(s)-1:] == "!"
		}(fields[1])
		if i == 0 {
			isNull, _ := isArgNil(v)
			if isRequire || !isNull {
				switch opr {
				case Operator["null"], Operator["notnull"]:
					query.WriteString(` WHERE ` + columnName + ` ` + opr)
				case Operator["like"]:
					query.WriteString(` WHERE lower(` + columnName + `) ` + opr + ` concat('%',?,'%') `)
					tmpArgs, _ := v.(string)
					args = append(args, tmpArgs)
				case Operator["in"]:
					s := reflect.ValueOf(v)
					if s.Kind() == reflect.Slice {
						var smt string
						for j := 0; j < s.Len(); j++ {
							smt += `?,`
							switch s.Index(j).Kind() {
							case reflect.Int:
								args = append(args, s.Index(j).Int())
							case reflect.Int32:
								args = append(args, s.Index(j).Int())
							case reflect.Int64:
								args = append(args, s.Index(j).Int())
							case reflect.Float32:
								args = append(args, s.Index(j).Float())
							case reflect.Float64:
								args = append(args, s.Index(j).Float())
							default:
								args = append(args, s.Index(j).String())
							}
						}
						query.WriteString(` WHERE ` + columnName + ` ` + opr + `(` + smt[:len(smt)-1] + `)`)
					}
				default:
					query.WriteString(` WHERE ` + columnName + ` ` + opr + ` ? `)
					args = append(args, v)
				}
			} else {
				query.WriteString(` WHERE 1 = 1 `)
			}

		} else {
			isNull, _ := isArgNil(v)
			if isRequire || !isNull {
				switch opr {
				case Operator["null"], Operator["notnull"]:
					query.WriteString(` AND ` + columnName + ` ` + opr)
				case Operator["like"]:
					query.WriteString(` AND lower(` + columnName + `) ` + opr + ` concat('%',?,'%') `)
					tmpArgs, _ := v.(string)
					args = append(args, tmpArgs)
				case Operator["in"]:
					s := reflect.ValueOf(v)
					if s.Kind() == reflect.Slice {
						var smt string
						for j := 0; j < s.Len(); j++ {
							smt += `?,`
							switch s.Index(j).Kind() {
							case reflect.Int:
								args = append(args, s.Index(j).Int())
							case reflect.Int32:
								args = append(args, s.Index(j).Int())
							case reflect.Int64:
								args = append(args, s.Index(j).Int())
							case reflect.Float32:
								args = append(args, s.Index(j).Float())
							case reflect.Float64:
								args = append(args, s.Index(j).Float())
							default:
								args = append(args, s.Index(j).String())
							}
						}
						query.WriteString(` AND ` + columnName + ` ` + opr + `(` + smt[:len(smt)-1] + `)`)
					}
				default:
					query.WriteString(` AND ` + columnName + ` ` + opr + ` ? `)
					args = append(args, v)
				}
			} else {
				query.WriteString(` AND 1 = 1 `)
			}
		}
		i++
	}
	return query.String(), args
}

// Order generate string ordering query statement
func (q *Query) Order() string {
	if len(q.Sort) > 0 {
		field := strings.Split(q.Sort, ",")
		sort := ` ORDER BY `
		for _, v := range field {
			sortType := func(str string) string {
				if strings.HasPrefix(str, "-") {
					return `desc`
				}
				return `asc`
			}
			sort += strings.TrimPrefix(v, "-") + ` ` + sortType(v) + `,`
		}
		return sort[:len(sort)-1]
	}
	return ` `
}

func translateOperator(s string) string {
	operator := Operator[strings.ToLower(s)]
	if operator == "" {
		return Operator["eq"]
	}
	return operator
}

func isArgNil(i interface{}) (bool, reflect.Kind) {
	r := reflect.ValueOf(i)
	switch r.Kind() {
	case reflect.Slice:
		return r.Len() == 0, reflect.Slice
	case reflect.String:
		return r.String() == "", reflect.String
	case reflect.Int:
		return r.Int() == 0, reflect.Int
	case reflect.Int32:
		return r.Int() == 0, reflect.Int32
	case reflect.Int64:
		return r.Int() == 0, reflect.Int64
	case reflect.Float32:
		return r.Float() == 0, reflect.Float32
	case reflect.Float64:
		return r.Float() == 0, reflect.Float64
	default:
		return false, reflect.String
	}
}
