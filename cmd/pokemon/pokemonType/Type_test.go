package pokemonType

import (
	"reflect"
	"testing"
)

func Test_노말은_효과적인_타입_없다(t *testing.T) {
	// given
	allTypes := TypeRegistry
	normalType := TypeNormal

	// when, then
	for _, otherType := range allTypes {
		if normalType.IsEffective(otherType) || normalType.IsNotEffective(otherType) {
			t.Error("노말은 효과적인 타입이 없어야 한다")
		}
	}
}

func Test_노말은_비효과적인_타입_없다(t *testing.T) {
	// given
	allTypes := TypeRegistry
	normalType := TypeNormal

	// when, then
	for _, otherType := range allTypes {
		if normalType.IsEffective(otherType) || normalType.IsNotEffective(otherType) {
			t.Error("노말은 비효과적인 타입이 없어야 한다")
		}
	}
}

func Test_물은_불_효과적이다(t *testing.T) {
	// given
	effectiveTypes := []PokemonType{
		TypeFire,
	}

	// when, then
	for _, otherType := range effectiveTypes {
		if !TypeAqua.IsEffective(otherType) {
			t.Error("물은 ", reflect.TypeOf(otherType), "에 효과적이어야 한다")
		}
	}
}

func Test_물은_풀_비효과적이다(t *testing.T) {
	// given
	notEffectiveTypes := []PokemonType{
		TypeGrass,
	}
	// when, then
	for _, otherType := range notEffectiveTypes {
		if !TypeAqua.IsNotEffective(otherType) {
			t.Error("물은 ", reflect.TypeOf(otherType), "에 비효과적이어야 한다")
		}
	}
}

func Test_불은_풀_효과적이다(t *testing.T) {
	// given
	effectiveTypes := []PokemonType{
		TypeGrass,
	}

	// when, then
	for _, otherType := range effectiveTypes {
		if !TypeFire.IsEffective(otherType) {
			t.Error("불은 ", reflect.TypeOf(otherType), "에 효과적이어야 한다")
		}
	}
}

func Test_불은_물_비효과적이다(t *testing.T) {
	// given
	notEffectiveTypes := []PokemonType{
		TypeAqua,
	}
	// when, then
	for _, otherType := range notEffectiveTypes {
		if !TypeFire.IsNotEffective(otherType) {
			t.Error("불은 ", reflect.TypeOf(otherType), "에 비효과적이어야 한다")
		}
	}
}

func Test_풀은_물_효과적이다(t *testing.T) {
	// given
	effectiveTypes := []PokemonType{
		TypeAqua,
	}

	// when, then
	for _, otherType := range effectiveTypes {
		if !TypeGrass.IsEffective(otherType) {
			t.Error("풀은 ", reflect.TypeOf(otherType), "에 효과적이어야 한다")
		}
	}
}

func Test_풀은_불_비효과적이다(t *testing.T) {
	// given
	notEffectiveTypes := []PokemonType{
		TypeFire,
	}
	// when, then
	for _, otherType := range notEffectiveTypes {
		if !TypeGrass.IsNotEffective(otherType) {
			t.Error("풀은 ", reflect.TypeOf(otherType), "에 비효과적이어야 한다")
		}
	}
}
