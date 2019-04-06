package helpers

import (
	"go/ast"
)

// EmptyS returns a empty statement AST.
func EmptyS() []ast.Stmt {
	return []ast.Stmt{}
}

// S converts passed statements to a collection of statements.
func S(stmts ...ast.Stmt) []ast.Stmt {
	return stmts
}

// EmptyE returns an empty expression AST
func EmptyE() []ast.Expr {
	return []ast.Expr{}
}

// E converts passed expressions to a collection of expressions.
func E(exprs ...ast.Expr) []ast.Expr {
	return exprs
}

// EmptyI returns an empty identity AST
func EmptyI() []*ast.Ident {
	return []*ast.Ident{}
}

// I converts passed identities to a collection of identities.
func I(ident *ast.Ident) []*ast.Ident {
	return []*ast.Ident{ident}
}
