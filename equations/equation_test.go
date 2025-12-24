package equations

import (
	"testing"

	"github.com/mazzegi/adventofcode/rat"
	"github.com/mazzegi/adventofcode/testutil"
)

func mkRat(v int) rat.Number {
	return rat.R(v, 1)
}

func mkCoeffs(vs ...int) []rat.Number {
	rns := make([]rat.Number, len(vs))
	for i, v := range vs {
		rns[i] = mkRat(v)
	}
	return rns
}

func mkExpr(constant int, vs ...int) Expression {
	return Expression{
		VariableCoeffs: mkCoeffs(vs...),
		Constant:       mkRat(constant),
	}
}

func mkEq(constant int, vs ...int) Equation {
	return Equation{mkExpr(constant, vs...)}
}

func TestSystemBase(t *testing.T) {
	tx := testutil.NewTx(t)

	{
		sys, err := NewSystem(2)
		tx.AssertNoErr(err)
		sys.AddEquation(mkEq(20, 2, 4))
		sys.AddEquation(mkEq(42, 3, 1))
		sys.Solve()
	}
	{
		sys, err := NewSystem(7)
		tx.AssertNoErr(err)
		sys.AddEquation(mkEq(-20, 0, 0, 0, 0, 1, 0, 1))
		sys.AddEquation(mkEq(-30, 0, 0, 0, 1, 1, 1, 0))

		sys.AddEquation(mkEq(-39, 0, 1, 0, 1, 1, 0, 1))
		sys.AddEquation(mkEq(-30, 1, 0, 1, 1, 1, 0, 0))

		sys.AddEquation(mkEq(-20, 0, 1, 0, 0, 0, 1, 0))
		sys.AddEquation(mkEq(-20, 0, 1, 0, 0, 0, 1, 0))
		sys.AddEquation(mkEq(-16, 1, 0, 0, 0, 0, 0, 1))
		sys.Solve()
	}

}
