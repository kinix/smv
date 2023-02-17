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

func compareSlice[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func TestDetail(t *testing.T) {
	testCases := []struct {
		name       string
		structType reflect.Type
		fieldList  []string
		memoryMap  []int
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
			fieldList: []string{"Field1", "Field2", "Field3", "Field4", "Field5"},
			memoryMap: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, -1, -1, -1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4},
		},
		{
			name: "empty",
			structType: reflect.TypeOf(struct {
			}{}),
			fieldList: []string{},
			memoryMap: []int{},
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
			fieldList: []string{"Field1", "Field2"},
			memoryMap: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, testCase := range testCases {
		fieldList, memoryMap := detail(testCase.structType)

		if !compareSlice(fieldList, testCase.fieldList) {
			t.Errorf("fieldList not matched for %s. got: %v, expected: %v", testCase.name, fieldList, testCase.fieldList)
		}

		if !compareSlice(memoryMap, testCase.memoryMap) {
			t.Errorf("memoryMap not matched for %s. got: %v, expected: %v", testCase.name, memoryMap, testCase.memoryMap)
		}
	}
}
