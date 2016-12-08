/**
 * ast/variable.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package ast
import "github.com/jluchiji/go-fiz/base"


type Variable struct {
  name  string
  index int
}


func (this *Variable) Resolve(f base.Func) {
  /* Find the index of the variable in context */
  this.index = -1
  for i, v := range f.GetArgs() {
    if this.name == v { this.index = i }
  }
  /* If still -1, variable is not defined */
  if this.index < 0 {
    panic("FIZ_UNDEF_VAR")
  }
}


func (this *Variable) Evaluate(context []int) int {
  return context[this.index]
}


func NewVariable(name string) *Variable {
  return &Variable{
    name,
    -1,
  }
}
