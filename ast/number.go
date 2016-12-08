/**
 * ast/number.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package ast
import "github.com/jluchiji/go-fiz/base"


type Number struct {
  value int
}


func (this *Number) Resolve(fn base.Func) {
  /* no-op */
}


func (this *Number) Evaluate(context []int) int {
  return this.value
}


func NewNumber(value int) *Number {
  return &Number{ value }
}
