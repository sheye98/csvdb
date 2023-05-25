package column

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumn_Add(t *testing.T) {
	t.Run("should_return_nil_when_add_column", func(t *testing.T) {
		column := New("./testdata/test_1.csv")
		err := column.Add("test_column")
		assert.Nil(t, err)
	})
}

func TestColumn_Alter(t *testing.T) {
	t.Run("should_return_nil_when_alter_column", func(t *testing.T) {
		column := New("./testdata/test_1.csv")
		err := column.Alter("name", "name1")
		assert.Nil(t, err)
	})
}
