package pokemonType

type Aqua struct {
}

func (aqua Aqua) IsEffective(target PokemonType) bool {
	switch target.(type) {
	case Fire:
		return true
	default:
		return false
	}
}

func (aqua Aqua) IsNotEffective(target PokemonType) bool {
	switch target.(type) {
	case Grass:
		return true
	default:
		return false
	}
}
