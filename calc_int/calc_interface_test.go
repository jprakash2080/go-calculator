package calc_interface
import(
  "testing"
  //"fmt"
)
//Subtraction
type AdditionInput struct {
    input Addition
    expected float32
}
type SubtractionInput struct {
    input Subtraction
    expected float32
}
type MulitplicationInput struct {
    input Mulitplication
    expected float32
}
type DivisionInput struct {
    input Division
    expected float32
}

var additionTests = []AdditionInput{
  AdditionInput{Addition{4,4}, 1},
  AdditionInput{Addition{8, 8}, 1},
  AdditionInput{Addition{4.5, 2.2}, 2.0454545},
  AdditionInput{Addition{3, 8}, 0.375},
  AdditionInput{Addition{15, 20}, 0.75},
}

var divisionTest = []DivisionInput{
  DivisionInput{Division{4, 4}, 1},
  DivisionInput{Division{8, 8}, 1},
  DivisionInput{Division{4.5, 2.2}, 2.0454545},
  DivisionInput{Division{3, 8}, 0.375},
  DivisionInput{Division{15, 20}, 0.75},
}

var mulitplicationTest = []MulitplicationInput{
    MulitplicationInput{Mulitplication{4, 4}, 16},
    MulitplicationInput{Mulitplication{8, 8}, 64},
    MulitplicationInput{Mulitplication{2.3, 2.3}, 5.9},
    MulitplicationInput{Mulitplication{20, 15}, 300},
}

var subtractionTests = []SubtractionInput{
    SubtractionInput{Subtraction{4, 4}, 0},
    SubtractionInput{Subtraction{8, 8}, 0},
    SubtractionInput{Subtraction{5.5, 3.5}, 2},
    SubtractionInput{Subtraction{8.3, 5}, 3.3},
    SubtractionInput{Subtraction{20, 15}, 5},
}


func FinalTest(t *testing.T){
    //Addition Test cases...
    for _, test := range additionTests{
        if output := GetResult(test.input); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }
    //Subtraction Test cases...
    for _, test := range subtractionTests{
        if output := GetResult(test.input); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }

    //Multiplication Test cases...
    for _, test := range mulitplicationTest{
        if output := GetResult(test.input); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }

    //Division Test cases...
    for _, test := range divisionTest{
        if output := GetResult(test.input); output != test.expected {
            t.Errorf("Output %f not equal to expected %f", output, test.expected)
        }
    }
}
