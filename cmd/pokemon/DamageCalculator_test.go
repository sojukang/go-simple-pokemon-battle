package pokemon

import "testing"

func Test_상성_랭크_효과없이_계산(t *testing.T) {
	// given
	damageCalculator := DefaultDamageCalculator{}

	// when, then
	t.Run("공격, 방어 수치 같으면 기술 위력만큼이다", func(t *testing.T) {
		result := damageCalculator.calculate(
			100,
			35,
			100,
			1,
			1,
		)

		expected := 35
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
	t.Run("공격이 방어 수치 2배면 기술 위력 2배이다", func(t *testing.T) {
		result := damageCalculator.calculate(
			100,
			35,
			50,
			1,
			1,
		)

		expected := 70
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
	t.Run("방어가 공격 수치 2배이면 기술 위력 절반이다", func(t *testing.T) {
		result := damageCalculator.calculate(
			50,
			35,
			100,
			1,
			1,
		)

		expected := 17
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
}

func Test_상성_계수_반영(t *testing.T) {
	// given
	damageCalculator := DefaultDamageCalculator{}

	// when
	result := damageCalculator.calculate(
		100,
		35,
		100,
		2,
		1,
	)

	// then
	expected := 70
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_랭크_계수_반영(t *testing.T) {
	// given
	damageCalculator := DefaultDamageCalculator{}

	// when
	result := damageCalculator.calculate(
		100,
		35,
		100,
		1,
		2,
	)

	// then
	expected := 70
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_상성_랭크_계수_모두_반영(t *testing.T) {
	// given
	damageCalculator := DefaultDamageCalculator{}

	// when
	result := damageCalculator.calculate(
		100,
		35,
		100,
		2,
		2,
	)

	// then
	expected := 140
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}
