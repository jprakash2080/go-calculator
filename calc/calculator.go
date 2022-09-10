package calculator
import(
  //"fmt"
  //"reflect"
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
// func Division[N Number](num1, num2 N) N{
//   f := num2 / num1
//   fmt.Println(reflect.TypeOf((f)))
//   return f
// }
func Division(num1 int32, num2 int32) float32{
  f := float32(float32(num1) / float32(num2))
  //fmt.Println(reflect.TypeOf((f)))
  return f
}
