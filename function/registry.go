/**
 * function/registry.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package function
import (
  "github.com/jluchiji/go-fiz/base"
  "github.com/jluchiji/go-fiz/function/builtin"
)


var registry map[string]base.Func


func Register(name string, f base.Func) bool {
  _, ok := registry[name]
  if (ok) {
    return false
  }
  registry[name] = f
  return true
}


func Find(name string) (f base.Func, ok bool) {
  f, ok = registry[name]
  return
}


func Init() {
  registry = make(map[string]base.Func)

  /* Register builtin functions */
  Register("inc", NewNative([]string{ "x" }, builtin.Inc))
  Register("dec", NewNative([]string{ "x" }, builtin.Dec))
  Register("ifz", NewNative([]string{ "c", "t", "f" }, builtin.Ifz))
  Register("halt", NewNative([]string{ }, builtin.Halt))
}
