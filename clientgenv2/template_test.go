package clientgenv2

import (
	"go/types"
	"testing"
)

// TestReturnTypeName tests the returnTypeName function with various types.
func TestReturnTypeName(t *testing.T) {
	tests := []struct {
		name     string
		input    types.Type
		nested   bool
		expected string
	}{
		{
			name:     "Basic",
			input:    types.Typ[types.String],
			nested:   false,
			expected: "string",
		},
		{
			name:     "Pointer",
			input:    types.NewPointer(types.Typ[types.Int]),
			nested:   false,
			expected: "*int",
		},
		{
			name:     "Slice",
			input:    types.NewSlice(types.Typ[types.Float64]),
			nested:   false,
			expected: "[]float64",
		},
		{
			name:     "Named",
			input:    types.NewNamed(types.NewTypeName(0, nil, "MyType", nil), nil, nil),
			nested:   false,
			expected: "*MyType",
		},
		{
			name:     "Named with package",
			input:    types.NewNamed(types.NewTypeName(0, types.NewPackage("github.com/Yamashou/hoge", "hoge"), "MyType", nil), nil, nil),
			nested:   false,
			expected: "*MyType",
		},
		{
			name:     "Interface",
			input:    types.NewInterfaceType(nil, nil).Complete(),
			nested:   false,
			expected: "interface{}",
		},
		{
			name:     "Map",
			input:    types.NewMap(types.Typ[types.Int], types.Typ[types.Bool]),
			nested:   false,
			expected: "map[int]bool",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := returnTypeName(test.input, test.nested)
			if output != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, output)
			}
		})
	}
}
