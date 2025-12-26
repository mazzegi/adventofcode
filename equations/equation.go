package equations

import (
	"fmt"
	"io"
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

func (varDet VariableDetermination) IsCompletelyDetermined() bool {
	for _, c := range varDet.Expr.VariableCoeffs {
		if !rat.IsZero(c) {
			return false
		}
	}
	return true
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

func (expr Expression) replaceVar(varIdx int, withExpr Expression) Expression {
	varC := expr.VariableCoeffs[varIdx]
	withExpr = withExpr.mult(varC)
	newExpr := expr.zeroCoeff(varIdx)
	return newExpr.add(withExpr)
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
	// if len(s.Equations) == 1 {
	// 	errutil.FatalWhen(fmt.Errorf("only one equation left"))
	// }

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

type SolutionFunc func(vars []rat.Number) ([]rat.Number, error)

func (s *System) Solve() (sFnc SolutionFunc, idpVars []int, err error) {
	s.removeZeroEquations()
	//s.Dump(os.Stdout, nil)

	var varDets []*VariableDetermination
	for len(s.Equations) > 0 {
		varDet := s.eliminateOne()
		varDets = append(varDets, &varDet)
		s.removeZeroEquations()
		//s.Dump(os.Stdout, varDets)
	}
	if len(varDets) == 0 {
		return nil, nil, fmt.Errorf("no var. determs.")
	}
	//lastVarDet := varDets[len(varDets)-1]

	//independentVars := lastVarDet.Expr.VariableCoeffs
	// for i, v := range lastVarDet.Expr.VariableCoeffs {
	// 	if !rat.IsZero(v) {
	// 		idpVars = append(idpVars, i)
	// 	}
	// }
	// collect all vars, which are NOT determined
	existsVarDetForIdx := func(varIdx int) bool {
		for _, vd := range varDets {
			if vd.VarIdx == varIdx {
				return true
			}
		}
		return false
	}

	for _, vd := range varDets {
		for varIdx, c := range vd.Expr.VariableCoeffs {
			if rat.IsZero(c) {
				continue
			}
			if existsVarDetForIdx(varIdx) {
				continue
			}
			if !slices.Contains(idpVars, varIdx) {
				idpVars = append(idpVars, varIdx)
			}
		}
	}
	sort.Ints(idpVars)

	// build solver func - in go the independent vars - out go all vars
	return func(vars []rat.Number) ([]rat.Number, error) {
		if len(vars) != len(idpVars) {
			return nil, fmt.Errorf("invalid input")
		}
		// make a copy of vardets
		solverVardDets := make([]*VariableDetermination, len(varDets))
		for i, vd := range varDets {
			cvd := VariableDetermination{
				VarIdx: vd.VarIdx,
				Expr:   vd.Expr.clone(),
			}
			solverVardDets[i] = &cvd
		}

		//
		ress := make([]rat.Number, s.Dim)
		// add indep. vars to ress
		for i, idpVarIdx := range idpVars {
			ress[idpVarIdx] = vars[i]
		}

		// replace idp vars in all var-determs.
		for _, varDet := range solverVardDets {
			for i, idpVarIdx := range idpVars {
				idpV := vars[i]
				varDet.Expr = varDet.Expr.replaceVar(idpVarIdx, Expression{
					Constant:       idpV,
					VariableCoeffs: make([]rat.Number, len(varDet.Expr.VariableCoeffs)),
				})
			}
			//solverVardDets[di] = varDet
		}

		//s.Dump(os.Stdout, solverVardDets)
		maxIter := s.Dim
		currIter := 0
		for {
			//
			var determined []*VariableDetermination
			var undetermined []*VariableDetermination
			for _, varDet := range solverVardDets {
				if varDet.IsCompletelyDetermined() {
					determined = append(determined, varDet)
					// update ress
					ress[varDet.VarIdx] = varDet.Expr.Constant
				} else {
					undetermined = append(undetermined, varDet)
				}
			}
			if len(undetermined) == 0 {
				break
			}
			// replace determined in all undetermined
			var replaced bool
			for _, undetVar := range undetermined {
				for _, detVar := range determined {
					if undetVar.Expr.VariableCoeffs[detVar.VarIdx] != rat.Zero {
						undetVar.Expr = undetVar.Expr.replaceVar(detVar.VarIdx, detVar.Expr)
						replaced = true
					}
				}
			}
			if !replaced {
				return nil, fmt.Errorf("replaced nothing")
			}
			//s.Dump(os.Stdout, solverVardDets)
			currIter++
			if currIter > maxIter {
				return nil, fmt.Errorf("max iter reached for input: %v", vars)
			}
		}

		return ress, nil
	}, idpVars, nil
}

func ratSolutionToIntSolution(rsol []rat.Number) ([]int, bool) {
	sol := make([]int, len(rsol))
	for i, r := range rsol {
		n, ok := rat.ToInteger(r)
		if !ok {
			return nil, false
		}
		sol[i] = n
	}
	return sol, true
}

func solutionPositiveAndInBounds(sol []int, bounds []int) bool {
	for i, v := range sol {
		if v < 0 || v > bounds[i] {
			return false
		}
	}
	return true
}

func intsToRats(ns []int) []rat.Number {
	rats := make([]rat.Number, len(ns))
	for i, n := range ns {
		rats[i] = rat.R(n, 1)
	}
	return rats
}

func (sys *System) AllPositiveIntegerSolutions(allbounds []int) ([][]int, error) {
	solFunc, idpVars, err := sys.Solve()
	if err != nil {
		return nil, fmt.Errorf("solve: %w", err)
	}
	if len(idpVars) == 0 {
		// all vars are determined
		rsol, err := solFunc(nil)
		if err != nil {
			return nil, fmt.Errorf("solve_func: %w", err)
		}
		sol, ok := ratSolutionToIntSolution(rsol)
		if !ok {
			return [][]int{}, nil
		}
		if !solutionPositiveAndInBounds(sol, allbounds) {
			return [][]int{}, nil
		}
		return [][]int{sol}, nil
	}

	var solutions [][]int
	input := make([]int, len(idpVars)) // all are 0

	increaseInput := func() bool {
		for i := range input {
			input[i]++
			if input[i] <= allbounds[idpVars[i]] {
				return true
			}
			input[i] = 0
		}
		return false
	}

	for {
		rsol, err := solFunc(intsToRats(input))
		if err != nil {
			return nil, fmt.Errorf("solve_func: %w", err)
		}
		sol, ok := ratSolutionToIntSolution(rsol)
		if !ok {

		} else if !solutionPositiveAndInBounds(sol, allbounds) {

		} else {
			solutions = append(solutions, sol)
		}
		ok = increaseInput()
		if !ok {
			break
		}
	}

	return solutions, nil
}

func (s *System) Dump(w io.Writer, varDets []*VariableDetermination) {
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
