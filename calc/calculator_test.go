package calculator
import(
  "testing"
  "fmt"
)

type CalculateInput struct {
    num1 float32
    num2 float32
    expected float32
}

var divisionTest = []CalculateInput{
  CalculateInput{4, 4, 1},
  CalculateInput{8, 8, 1},
  CalculateInput{4.5, 2.2, 2.0454545},
  CalculateInput{3, 8, 0.375},
  CalculateInput{15, 20, 0.75},
}

var mulitplicationTest = []CalculateInput{
    CalculateInput{4, 4, 16},
    CalculateInput{8, 8, 64},
    CalculateInput{2.3, 2.3, 5.9},
    CalculateInput{20, 15, 300},
}


var subtractionTests = []CalculateInput{
    CalculateInput{4, 4, 0},
    CalculateInput{8, 8, 0},
    CalculateInput{5.5, 3.5, 2},
    CalculateInput{8.3, 5, 3.3},
    CalculateInput{20, 15, 5},
}

var additionTests = []CalculateInput{
    CalculateInput{4, 4, 8},
    CalculateInput{8, 8, 16},
    CalculateInput{3.3, 8.5, 11.8},
    CalculateInput{20, 15, 35},
}

func FinalTest(t *testing.T){


    //Addition Test cases...
    for _, test := range additionTests{
        if output := Addition(test.num1, test.num2); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }

    //Subtraction Test cases...
    for _, test := range subtractionTests{
        if output := Subtraction(test.num1, test.num2); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }

    //Multiplication Test cases...
    fmt.Println("Multiplication Test cases...")
    for _, test := range mulitplicationTest{
        if output := Mulitplication(test.num1, test.num2); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }

    //Division Test cases...
    for _, test := range divisionTest{
        if output := Division(test.num1, test.num2); float32(output) != float32(test.expected) {
            t.Errorf("Output %f not equal to expected %f.", float32(output), float32(test.expected))
        }
    }

}
