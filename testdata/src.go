package testdata

type Expr interface {
	isExpr()
}

type Ident struct{}

func (Ident) isExpr() {}

type Literal struct{}

func (Literal) isExpr() {}

type Statement struct{}
