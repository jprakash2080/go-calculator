package calculator
import(
//  "fmt"
  "golang.org/x/exp/constraints"
)
type Number interface {
    constraints.Integer | constraints.Float
}

//Add
func Addition[N Number](num1, num2 N) N{
  return num1 + num2
}

//Subtract
func Subtraction[N Number](num1, num2 N) N{
  return num1 - num2
}

//Mulitplication
func Mulitplication[N Number](num1, num2 N) N{
  return num1 * num2
}

//Division
func Division[N Number](num1, num2 N) N{
  return num1 / num2
}
