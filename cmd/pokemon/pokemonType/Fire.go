package pokemonType

type Fire struct {
}

func (fire Fire) IsEffective(target PokemonType) bool {
	switch target.(type) {
	case Grass:
		return true
	default:
		return false
	}
}

func (fire Fire) IsNotEffective(target PokemonType) bool {
	switch target.(type) {
	case Aqua:
		return true
	default:
		return false
	}
}
