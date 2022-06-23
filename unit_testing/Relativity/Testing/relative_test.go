package Testing

import (
	"testing"
	"fmt"
)



type relativeTest struct {
	arg1,arg2,expected float64
}

var relativeTests = []relativeTest{
	relativeTest{5,6,7},
	relativeTest{1,2,3},
}

func TestRelative(t *testing.T) {
	i:=1
	for _, test := range relativeTests {
		if output := Relativity(test.arg1, test.arg2); output != test.expected {
			fmt.Println("Test Case :",i," Failed")
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
		i+=1
		
	}
}