package Pokemon

type BasePokemon struct {
	hp int
}

func (basePokemon *BasePokemon) TakeDamage(damage int) {
	basePokemon.hp -= damage
}
