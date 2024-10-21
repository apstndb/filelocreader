package testdata

// 日本語があっても大丈夫？

type Expr interface {
	isExpr()
}

// 本当に？

type Ident struct{}

func (Ident) isExpr() {}

type Literal struct{}

func (Literal) isExpr() {}

type Statement struct{}
