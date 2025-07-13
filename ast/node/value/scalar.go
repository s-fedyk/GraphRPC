package value

import (
	"fmt"
	"graphrpc.com/ast/node/base"
)

// Variable represents a GraphQL variable
type Variable struct {
	Name     *base.Name
	Location *base.Location
}

func (v *Variable) String() string {
	return fmt.Sprintf("$%s", v.Name.Value)
}

func (v *Variable) GetLocation() *base.Location { return v.Location }
func (v *Variable) valueNode()             {}

// IntValue represents an integer value
type IntValue struct {
	Value    string
	Location *base.Location
}

func (i *IntValue) String() string {
	return i.Value
}

func (i *IntValue) GetLocation() *base.Location { return i.Location }
func (i *IntValue) valueNode()             {}

// FloatValue represents a float value
type FloatValue struct {
	Value    string
	Location *base.Location
}

func (f *FloatValue) String() string {
	return f.Value
}

func (f *FloatValue) GetLocation() *base.Location { return f.Location }
func (f *FloatValue) valueNode()             {}

// StringValue represents a string value
type StringValue struct {
	Value    string
	Block    bool
	Location *base.Location
}

func (s *StringValue) String() string {
	if s.Block {
		return fmt.Sprintf(`"""%s"""`, s.Value)
	}
	return fmt.Sprintf(`"%s"`, s.Value)
}

func (s *StringValue) GetLocation() *base.Location { return s.Location }
func (s *StringValue) valueNode()             {}

// BooleanValue represents a boolean value
type BooleanValue struct {
	Value    bool
	Location *base.Location
}

func (b *BooleanValue) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

func (b *BooleanValue) GetLocation() *base.Location { return b.Location }
func (b *BooleanValue) valueNode()             {}

// NullValue represents a null value
type NullValue struct {
	Location *base.Location
}

func (n *NullValue) String() string {
	return "null"
}

func (n *NullValue) GetLocation() *base.Location { return n.Location }
func (n *NullValue) valueNode()             {}

// EnumValue represents an enum value
type EnumValue struct {
	Value    string
	Location *base.Location
}

func (e *EnumValue) String() string {
	return e.Value
}

func (e *EnumValue) GetLocation() *base.Location { return e.Location }
func (e *EnumValue) valueNode()             {}