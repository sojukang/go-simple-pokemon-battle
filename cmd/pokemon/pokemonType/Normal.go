package pokemonType

type Normal struct {
}

func (normal Normal) IsEffective(target PokemonType) bool {
	return false
}

func (normal Normal) IsNotEffective(target PokemonType) bool {
	return false
}
