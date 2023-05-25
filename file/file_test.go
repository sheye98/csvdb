package file

import (
	"fmt"
	"testing"
)

func TestData_WriteHeader(t *testing.T) {
	data := New("./testdata/test_1.csv")
	data.Header = []string{"test_column1"}
	err := data.WriteHeader()
	if err != nil {
		fmt.Println(err)
	}

}
