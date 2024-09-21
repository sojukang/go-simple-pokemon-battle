package pokemon

import "reflect"

type Type interface {
	isEffective(target Type) bool
	isNotEffective(target Type) bool
}

var TYPE_NORMAL = Normal{
	effectiveTypes:    []Type{},
	notEffectiveTypes: []Type{},
}

type Normal struct {
	effectiveTypes    []Type
	notEffectiveTypes []Type
}

func (normal Normal) isEffective(target Type) bool {
	return contains(normal.effectiveTypes, target)
}

func (normal Normal) isNotEffective(target Type) bool {
	return contains(normal.notEffectiveTypes, target)
}

func contains(types []Type, target Type) bool {
	for _, t := range types {
		if reflect.TypeOf(t) == reflect.TypeOf(target) {
			return true
		}
	}
	return false
}
