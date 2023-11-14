package calculator_test

import (
    "arthicli/pkg/calculator"
    "testing"
)

func TestAdd(t *testing.T) {
    result := calculator.Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
