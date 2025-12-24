package equations

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/rat"
	"github.com/mazzegi/adventofcode/slices"
)

// something like:
// 2*a1 - 5*a2 - a3 + 26
// -> constant = 26
// -> coeffs = [0,2,-5,-1]
type Expression struct {
	VariableCoeffs []rat.Number
	Constant       rat.Number
}

func (expr Expression) Format() string {
	first := true
	var sl []string
	for i, c := range expr.VariableCoeffs {
		if rat.IsZero(c) {
			continue
		}
		if first {
			first = false
		} else {
			sl = append(sl, "+")
		}
		sl = append(sl, rat.Format(c)+fmt.Sprintf("*a%d", i))
	}
	if !rat.IsZero(expr.Constant) {
		sl = append(sl, "+", rat.Format(expr.Constant))
	}

	return strings.Join(sl, " ")
}

// An Equation is always: <expr> = 0
// 2*a1 - 5*a2 - a3 + 26 = 0
type Equation struct {
	Expr Expression
}

// something like:
// a3 = 2*a1 - 5*a2 + 26
// -> varIdx = 4
// -> Expr = 2*a1 - 5*a2 + 26
type VariableDetermination struct {
	VarIdx int
	Expr   Expression
}

func (vd VariableDetermination) Format() string {
	return fmt.Sprintf("a%d = %s", vd.VarIdx, vd.Expr.Format())
}

func (expr Expression) firstNonZeroCoeff() (idx int, val rat.Number, found bool) {
	for i, c := range expr.VariableCoeffs {
		if !rat.IsZero(c) {
			return i, c, true
		}
	}
	return -1, rat.Zero, false
}

func (expr Expression) clone() Expression {
	return Expression{
		VariableCoeffs: slices.Clone(expr.VariableCoeffs),
		Constant:       expr.Constant,
	}
}

func (expr Expression) isZero() bool {
	for _, c := range expr.VariableCoeffs {
		if !rat.IsZero(c) {
			return false
		}
	}
	return true
}

func (expr Expression) zeroCoeff(idx int) Expression {
	c := expr.clone()
	c.VariableCoeffs[idx] = rat.Zero
	return c
}

func (expr Expression) negate() Expression {
	c := expr.clone()
	for i, v := range c.VariableCoeffs {
		c.VariableCoeffs[i] = rat.Neg(v)
	}
	c.Constant = rat.Neg(c.Constant)
	return c
}

func (expr Expression) mult(by rat.Number) Expression {
	c := expr.clone()
	for i, v := range c.VariableCoeffs {
		c.VariableCoeffs[i] = rat.Simplify(rat.Mult(v, by))
	}
	c.Constant = rat.Simplify(rat.Mult(c.Constant, by))
	return c
}

func (expr Expression) add(otherExpr Expression) Expression {
	added := expr.clone()
	for i := range len(added.VariableCoeffs) {
		added.VariableCoeffs[i] = rat.Simplify(rat.Add(added.VariableCoeffs[i], otherExpr.VariableCoeffs[i]))
	}
	added.Constant = rat.Simplify(rat.Add(added.Constant, otherExpr.Constant))
	return added
}

func (expr Expression) simplified() Expression {
	c := expr.clone()
	for i, v := range c.VariableCoeffs {
		c.VariableCoeffs[i] = rat.Simplify(v)
	}
	return c
}

func (eq Equation) NumNonZeroCoeffs() int {
	var num int
	for _, c := range eq.Expr.VariableCoeffs {
		if !rat.IsZero(c) {
			num++
		}
	}
	return num
}

type System struct {
	Dim       int
	Equations []Equation
}

func NewSystem(dim int, eqs ...Equation) (*System, error) {
	s := &System{
		Dim: dim,
	}
	for _, eq := range eqs {
		err := s.AddEquation(eq)
		if err != nil {
			return nil, fmt.Errorf("add_equation: %w", err)
		}
	}
	return s, nil
}

func (s *System) AddEquation(eq Equation) error {
	if len(eq.Expr.VariableCoeffs) != s.Dim {
		return fmt.Errorf("invalid equation dim %d, want %d", len(eq.Expr.VariableCoeffs), s.Dim)
	}
	s.Equations = append(s.Equations, eq)
	return nil
}

// should eliminate one variable
func (s *System) eliminateOne() VariableDetermination {
	if len(s.Equations) == 1 {
		errutil.FatalWhen(fmt.Errorf("only one equation left"))
	}

	// we have at least 2 equations
	sort.Slice(s.Equations, func(i, j int) bool {
		return s.Equations[i].NumNonZeroCoeffs() < s.Equations[j].NumNonZeroCoeffs()
	})
	//
	elimEq := s.Equations[0]
	//build var-determination
	varIdx, varCoeff, ok := elimEq.Expr.firstNonZeroCoeff()
	if !ok {
		errutil.FatalWhen(fmt.Errorf("all coeffs are zero"))
	}

	varDet := VariableDetermination{
		VarIdx: varIdx,
	}
	divideExprBy := varCoeff
	varDetExpr := elimEq.Expr.zeroCoeff(varIdx)
	varDetExpr = varDetExpr.negate()
	varDetExpr = varDetExpr.mult(rat.Inv(divideExprBy))
	varDet.Expr = varDetExpr.simplified()

	//remove equation
	s.Equations = slices.DeleteIdx(s.Equations, 0)

	// in remaining equations replace varCoeff by expression
	for i, eq := range s.Equations {
		varC := eq.Expr.VariableCoeffs[varIdx]
		newExpr := eq.Expr.zeroCoeff(varIdx)
		newExpr = newExpr.add(varDet.Expr.mult(varC))
		s.Equations[i].Expr = newExpr
	}
	return varDet
}

func (s *System) removeZeroEquations() {
	s.Equations = slices.DeleteFunc(s.Equations, func(t Equation) bool { return t.Expr.isZero() })
}

// type SolutionFunc func(vars []int) int

func (s *System) Solve() {
	s.removeZeroEquations()
	s.Dump(os.Stdout, nil)

	var varDets []VariableDetermination
	for len(s.Equations) > 1 {
		varDet := s.eliminateOne()
		varDets = append(varDets, varDet)
		s.removeZeroEquations()
		s.Dump(os.Stdout, varDets)
	}
}

func (s *System) Dump(w io.Writer, varDets []VariableDetermination) {
	var sl []string
	sl = append(sl, "------------------------")
	for _, vd := range varDets {
		sl = append(sl, fmt.Sprintf("With: %s", vd.Format()))
	}
	sl = append(sl, "\n")

	for _, eq := range s.Equations {
		sl = append(sl, eq.Expr.Format()+" = 0")
	}
	sl = append(sl, "\n")
	w.Write([]byte(strings.Join(sl, "\n")))
}
