package structmemoryvisualization

import (
	"reflect"
	"testing"
)

func BenchmarkDetail(b *testing.B) {
	testCases := []struct {
		name       string
		structType reflect.Type
	}{
		{
			name: "standard",
			structType: reflect.TypeOf(struct {
				Field1 string
				Field2 int32
				Field3 byte
				Field4 string
				Field5 float64
			}{}),
		},
		{
			name: "empty",
			structType: reflect.TypeOf(struct {
			}{}),
		},
		{
			name: "nested",
			structType: reflect.TypeOf(struct {
				Field1 string
				Field2 struct {
					Field3 int32
					Field4 struct {
						Field5 float64
					}
				}
			}{}),
		},
	}

	for _, testCase := range testCases {
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				detail(testCase.structType)
			}
		})
	}
}
