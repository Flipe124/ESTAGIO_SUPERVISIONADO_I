package query

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// Make makes and return the full query params in first return, your arguments in second and if has, your "skips" in third.
func Make(ctx *gin.Context, model any, skips ...string) (string, []any, bool) {
	var (
		query  string
		values []any
	)
	object := reflect.ValueOf(model).Elem()
	for index := 0; index < object.NumField(); index++ {
		field := strings.ToLower(object.Type().Field(index).Name)
		if param, hasParam := ctx.GetQuery(field); hasParam && !regexp.MustCompile(field).MatchString(strings.ToLower(strings.Join(skips, " "))) {
			query += fmt.Sprintf(" %s = ? AND", field)
			values = append(values, param)
		}
	}
	return strings.TrimSuffix(query, " AND"), values, len(values) > 0
}
