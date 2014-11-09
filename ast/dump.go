package ast

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ncw/gpython/py"
)

// Dump an Ast node as a string
func AstDump(ast Ast) string {
	if ast == nil {
		return "<nil>"
	}
	name := ast.Type().Name
	astValue := reflect.Indirect(reflect.ValueOf(ast))
	astType := astValue.Type()
	args := make([]string, 0)
	for i := 0; i < astType.NumField(); i++ {
		fieldType := astType.Field(i)
		fieldValue := astValue.Field(i)
		fname := fieldType.Name
		if fieldValue.CanInterface() {
			v := fieldValue.Interface()
			switch x := v.(type) {
			case py.String:
				args = append(args, fmt.Sprintf("%s=%q", fname, string(x)))
			case ModBase:
			case StmtBase:
			case ExprBase:
			case SliceBase:
			case Pos:
			case Ast:
				args = append(args, fmt.Sprintf("%s=%s", fname, AstDump(x)))
			default:
				args = append(args, fmt.Sprintf("%s=%v", fname, x))
			}
		}
	}
	return fmt.Sprintf("%s(%s)", name, strings.Join(args, ","))
}
