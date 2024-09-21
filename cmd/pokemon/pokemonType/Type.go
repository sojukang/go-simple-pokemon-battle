package pokemonType

type PokemonType interface {
	IsEffective(target PokemonType) bool
	IsNotEffective(target PokemonType) bool
}

var TypeRegistry = []PokemonType{
	TypeNormal,
	TypeAqua,
	TypeFire,
	TypeGrass,
}

var TypeNormal = Normal{}
var TypeAqua = Aqua{}
var TypeFire = Fire{}
var TypeGrass = Grass{}
