package pokemon

import "testing"

func Test_초기_랭크는_0이다(t *testing.T) {
	// given
	rank := NewRank()

	// when
	result := rank.value

	// then
	expected := 0
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_랭크를_증가시킨다(t *testing.T) {
	// given
	rank := NewRank()

	// when
	rank.increase()
	result := rank.value

	// then
	expected := 1
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_랭크를_감소시킨다(t *testing.T) {
	// given
	rank := NewRank()

	// when
	rank.decrease()
	result := rank.value

	// then
	expected := -1
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_최대_랭크는_2이다(t *testing.T) {
	// given
	rank := NewRank()

	// when
	rank.increase()
	rank.increase()
	rank.increase()
	result := rank.value

	// then
	expected := 2
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_최소_랭크는_마이너스2이다(t *testing.T) {
	// given
	rank := NewRank()

	// when
	rank.decrease()
	rank.decrease()
	rank.decrease()
	result := rank.value

	// then
	expected := -2
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_랭크에_따른_계수를_반환한다(t *testing.T) {
	t.Run("0 => 1", func(t *testing.T) {
		// given
		rank := NewRank()

		// when
		result := rank.getCoefficient()

		// then
		expected := float32(1)
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
	t.Run("1 => 1.5", func(t *testing.T) {
		// given
		rank := NewRank()
		rank.increase()

		// when
		result := rank.getCoefficient()

		// then
		expected := float32(1.5)
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
	t.Run("2 => 2", func(t *testing.T) {
		// given
		rank := NewRank()
		rank.increase()
		rank.increase()

		// when
		result := rank.getCoefficient()

		// then
		expected := float32(2)
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
	t.Run("-1 => 0.66", func(t *testing.T) {
		// given
		rank := NewRank()
		rank.decrease()

		// when
		result := rank.getCoefficient()

		// then
		expected := float32(2.0 / 3.0)
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
	t.Run("-2 => 0.5", func(t *testing.T) {
		// given
		rank := NewRank()
		rank.decrease()
		rank.decrease()

		// when
		result := rank.getCoefficient()

		// then
		expected := float32(0.5)
		if result != expected {
			t.Error("expected: ", expected, ",", "actual: ", result)
		}
	})
}
