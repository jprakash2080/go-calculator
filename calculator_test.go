package calculator
import(
  "testing"
  "fmt"
)

type DivisionTest struct {
    num1 int32
    num2 int32
    expected float32
}

type MulitplicationTest struct {
    num1, num2, expected int32
}

type SubtractionTest struct {
    num1, num2, expected int32
}

type AddtitionTest struct {
    num1, num2, expected int32
}

var divisionTest = []DivisionTest{
  DivisionTest{4, 4, 1},
  DivisionTest{8, 8, 1},
  DivisionTest{3, 8, 0.375},
  DivisionTest{15, 20, 0.75},
}

var mulitplicationTest = []MulitplicationTest{
    MulitplicationTest{4, 4, 16},
    MulitplicationTest{8, 8, 64},
    MulitplicationTest{8, 3, 24},
    MulitplicationTest{20, 15, 300},
}

var subtractionTests = []SubtractionTest{
    SubtractionTest{4, 4, 0},
    SubtractionTest{8, 8, 0},
    SubtractionTest{8, 3, 5},
    SubtractionTest{20, 15, 5},
}

var additionTests = []AddtitionTest{
    AddtitionTest{4, 4, 8},
    AddtitionTest{8, 8, 16},
    AddtitionTest{3, 8, 11},
    AddtitionTest{20, 15, 35},
}
func FinalTest(t *testing.T){


    //Addition Test cases...
    for _, test := range additionTests{
        if output := Addition(test.num1, test.num2); output != test.expected {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
    }

    //Subtraction Test cases...
    for _, test := range subtractionTests{
        if output := Subtraction(test.num1, test.num2); output != test.expected {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
    }

    //Multiplication Test cases...
    fmt.Println("Multiplication Test cases...")
    for _, test := range mulitplicationTest{
        if output := Mulitplication(test.num1, test.num2); output != test.expected {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
    }

    //Division Test cases...
    for _, test := range divisionTest{
        if output := Division(test.num1, test.num2); float32(output) != float32(test.expected) {
            t.Errorf("Output %f not equal to expected %f.", float32(output), float32(test.expected))
        }
    }

}
