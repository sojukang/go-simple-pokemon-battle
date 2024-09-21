package Pokemon

type Pokemon interface {
	getCode() string
}

type Bulbasaur struct {
	Code string
}

func (pokemon Bulbasaur) getCode() string {
	return pokemon.Code
}
