package calc_interface

type Calculator interface{
  Calculate() float32
}

type Addition struct{
  num1 float32
  num2 float32
}
type Subtraction struct{
  num1 float32
  num2 float32
}
type Mulitplication struct{
  num1 float32
  num2 float32
}
type Division struct{
  num1 float32
  num2 float32
}

func (input Addition) Calculate() float32 {
  return input.num1 + input.num2
}

func (input Subtraction) Calculate() float32 {
  return input.num1 - input.num2
}

func (input Mulitplication) Calculate() float32 {
  return input.num1 * input.num2
}

func (input Division) Calculate() float32 {
  return input.num1 / input.num2
}

func GetResult(calc Calculator) float32 {
  return calc.Calculate()
}
