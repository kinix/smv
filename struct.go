package structmemoryvisualization

import (
	"fmt"
	"reflect"
)

// The main function
func DrawMemory(sampleStruct any, fileName string) {
	structType := reflect.TypeOf(sampleStruct)
	if structType.Kind() != reflect.Struct {
		err := fmt.Errorf("non-struct type: %s", structType)
		fmt.Println(err)
		return
	}

	fieldList, memoryMap := detail(structType)
	draw(fieldList, memoryMap, fileName)
}

// Create field list and memory map for given struct type
func detail(structType reflect.Type) (fieldList []string, memoryMap []int) {
	// Create a map of runes for each byte of the struct
	fieldList = []string{}
	memoryMap = make([]int, structType.Size())
	for i := range memoryMap {
		memoryMap[i] = -1
	}

	fields := reflect.VisibleFields(structType)

	// Mark memory slots
	for _, field := range fields {
		fieldList = append(fieldList, field.Name)
		for i := field.Offset; i < field.Offset+field.Type.Size(); i++ {
			// Set as last character of the field name
			memoryMap[i] = len(fieldList) - 1
		}
	}

	return
}
