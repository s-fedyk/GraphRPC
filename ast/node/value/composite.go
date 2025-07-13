package value

import (
	"fmt"
	"strings"
	"graphrpc.com/ast/node/base"
)

// ListValue represents a list value
type ListValue struct {
	Values   []base.Value
	Location *base.Location
}

func (l *ListValue) String() string {
	var values []string
	for _, v := range l.Values {
		values = append(values, v.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}

func (l *ListValue) GetLocation() *base.Location { return l.Location }
func (l *ListValue) valueNode()             {}

// ObjectValue represents an object value
type ObjectValue struct {
	Fields   []*ObjectField
	Location *base.Location
}

func (o *ObjectValue) String() string {
	var fields []string
	for _, f := range o.Fields {
		fields = append(fields, f.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
}

func (o *ObjectValue) GetLocation() *base.Location { return o.Location }
func (o *ObjectValue) valueNode()             {}

// ObjectField represents a field in an object value
type ObjectField struct {
	Name     *base.Name
	Value    base.Value
	Location *base.Location
}

func (o *ObjectField) String() string {
	return fmt.Sprintf("%s: %s", o.Name.Value, o.Value.String())
}

func (o *ObjectField) GetLocation() *base.Location { return o.Location }