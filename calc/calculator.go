package calculator

// type Number interface {
//     constraints.Integer | constraints.Float
// }

//Add
// func Addition1[N Number](num1, num2 N) N{
//   return num1 + num2
// }
func Addition(num1 float32, num2 float32) float32{
  return num1 + num2
}

//Subtract
// func Subtraction1[N Number](num1, num2 N) N{
//   return num1 - num2
// }
func Subtraction(num1 float32, num2 float32) float32{
  return num1 - num2
}

//Mulitplication
// func Mulitplication1[N Number](num1, num2 N) N{
//   return num1 * num2
// }
func Mulitplication(num1 float32, num2 float32) float32{
  return num1 * num2
}

//Division
// func Division[N Number](num1, num2 N) N{
//   f := num2 / num1
//   fmt.Println(reflect.TypeOf((f)))
//   return f
// }
func Division(num1 float32, num2 float32) float32{
  return num1 / num2
}
