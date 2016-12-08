/**
 * base/ast.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package base



/**
 * Base type for all Fiz AST nodes
 */
type AstNode interface {
  Resolve(Func)
  Evaluate([]int) int
}
