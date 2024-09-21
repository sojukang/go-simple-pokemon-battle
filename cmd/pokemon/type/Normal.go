package _type

type Normal struct {
	BaseType
}

func NewNormal() *Normal {
	return &Normal{BaseType: BaseType{
		effectiveTypes:    []Type{},
		notEffectiveTypes: []Type{},
	}}
}
