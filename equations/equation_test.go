package equations

import (
	"fmt"
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

		sols, err := sys.AllPositiveIntegerSolutions([]int{20, 42})
		tx.AssertNoErr(err)
		for i, sol := range sols {
			fmt.Printf("solution %d: %v\n", i+1, sol)
		}
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

		// solFunc, idpVars, err := sys.Solve()
		// tx.AssertNoErr(err)
		// fmt.Printf("indep.-vars: %v\n", idpVars)

		// input := make([]rat.Number, len(idpVars))
		// for i := range len(input) {
		// 	input[i] = rat.R(1, 1)
		// }
		// ress, err := solFunc(input)
		// tx.AssertNoErr(err)
		// fmt.Printf("results: %v", ress)
		sols, err := sys.AllPositiveIntegerSolutions([]int{20, 30, 39, 30, 20, 20, 16})
		tx.AssertNoErr(err)
		for i, sol := range sols {
			fmt.Printf("solution %d: %v\n", i+1, sol)
		}
	}

}
