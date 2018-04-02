package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jarsen/hax/ast"
)

const (
	FUNCTION_OBJ     = "FUNCTION"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
)

type (
	ObjectType string

	Object interface {
		Type() ObjectType
		Inspect() string
	}

	Function struct {
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Env        *Environment
	}

	ReturnValue struct {
		Value Object
	}

	Error struct {
		Message string
	}

	Integer struct {
		Value int64
	}

	Boolean struct {
		Value bool
	}

	Null struct {
	}
)

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// if you got line, column, etc. information from the lexer, you could add stack trace to error
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
