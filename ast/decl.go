package ast

import "github.com/post04/go-fast/token"

type (
	FunctionDeclaration struct {
		Function *FunctionLiteral
	}

	ClassDeclaration struct {
		Class *ClassLiteral
	}

	VariableDeclaration struct {
		Idx     Idx
		Token   token.Token
		List    VariableDeclarators
		Comment string
	}

	VariableDeclarators []VariableDeclarator

	VariableDeclarator struct {
		Target      *BindingTarget
		Initializer *Expression `optional:"true"`
	}
)

func (*FunctionDeclaration) _stmt() {}
func (*ClassDeclaration) _stmt()    {}
func (*VariableDeclaration) _stmt() {}
