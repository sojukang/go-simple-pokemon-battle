package pokemonType

type Grass struct {
}

func (grass Grass) IsEffective(target PokemonType) bool {
	switch target.(type) {
	case Aqua:
		return true
	default:
		return false
	}
}

func (grass Grass) IsNotEffective(target PokemonType) bool {
	switch target.(type) {
	case Fire:
		return true
	default:
		return false
	}
}
