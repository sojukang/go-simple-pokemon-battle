package _type

import "reflect"

type Type interface {
	IsEffective(target Type) bool
	IsNotEffective(target Type) bool
}

var TYPE_NORMAL = NewNormal()

type BaseType struct {
	effectiveTypes    []Type
	notEffectiveTypes []Type
}

func (baseType *BaseType) IsEffective(target Type) bool {
	return contains(baseType.effectiveTypes, target)
}

func (baseType *BaseType) IsNotEffective(target Type) bool {
	return contains(baseType.notEffectiveTypes, target)
}

func contains(types []Type, target Type) bool {
	for _, t := range types {
		if reflect.TypeOf(t) == reflect.TypeOf(target) {
			return true
		}
	}
	return false
}
