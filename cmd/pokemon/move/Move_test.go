package move

import (
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
	"testing"
)

func Test_꼬리흔들기(t *testing.T) {
	move := MOVE_꼬리흔들기

	t.Run("꼬리흔들기 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeNormal {
			t.Error()
		}
		if move.GetPower() != 0 {
			t.Error()
		}
		if move.GetPP() != 30 {
			t.Error()
		}
		if move.name != "꼬리흔들기" {
			t.Error()
		}
	})
	t.Run("꼬리흔들기 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("꼬리흔들기 방어자에게 변화 효과 방어 -1", func(t *testing.T) {
		if move.EffectOnDefender() != DEFENSE_MINUS_1 {
			t.Error()
		}
	})
}

func Test_덩굴채찍(t *testing.T) {
	move := MOVE_덩굴채찍

	t.Run("덩굴채찍 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeGrass {
			t.Error()
		}
		if move.GetPower() != 35 {
			t.Error()
		}
		if move.GetPP() != 10 {
			t.Error()
		}
		if move.name != "덩굴채찍" {
			t.Error()
		}
	})
	t.Run("덩굴채찍 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("덩굴채찍 방어자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnDefender() != NONE {
			t.Error()
		}
	})
}

func Test_몸통박치기(t *testing.T) {
	move := MOVE_몸통박치기

	t.Run("몸통박치기 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeNormal {
			t.Error()
		}
		if move.GetPower() != 35 {
			t.Error()
		}
		if move.GetPP() != 35 {
			t.Error()
		}
		if move.name != "몸통박치기" {
			t.Error()
		}
	})
	t.Run("몸통박치기 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("몸통박치기 방어자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnDefender() != NONE {
			t.Error()
		}
	})
}

func Test_물대포(t *testing.T) {
	move := MOVE_물대포

	t.Run("물대포 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeAqua {
			t.Error()
		}
		if move.GetPower() != 40 {
			t.Error()
		}
		if move.GetPP() != 25 {
			t.Error()
		}
		if move.name != "물대포" {
			t.Error()
		}
	})
	t.Run("물대포 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("물대포 방어자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnDefender() != NONE {
			t.Error()
		}
	})
}

func Test_불꽃세례(t *testing.T) {
	move := MOVE_불꽃세례

	t.Run("불꽃세례 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeFire {
			t.Error()
		}
		if move.GetPower() != 40 {
			t.Error()
		}
		if move.GetPP() != 25 {
			t.Error()
		}
		if move.name != "불꽃세례" {
			t.Error()
		}
	})
	t.Run("불꽃세례 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("불꽃세례 방어자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnDefender() != NONE {
			t.Error()
		}
	})
}

func Test_울음소리(t *testing.T) {
	move := MOVE_울음소리

	t.Run("울음소리 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeNormal {
			t.Error()
		}
		if move.GetPower() != 0 {
			t.Error()
		}
		if move.GetPP() != 40 {
			t.Error()
		}
		if move.name != "울음소리" {
			t.Error()
		}
	})
	t.Run("울음소리 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("울음소리 방어자에게 변화 효과 방어 -1", func(t *testing.T) {
		if move.EffectOnDefender() != ATTACK_MINUS_1 {
			t.Error()
		}
	})
}

func Test_째려보기(t *testing.T) {
	move := MOVE_째려보기

	t.Run("째려보기 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeNormal {
			t.Error()
		}
		if move.GetPower() != 0 {
			t.Error()
		}
		if move.GetPP() != 30 {
			t.Error()
		}
		if move.name != "째려보기" {
			t.Error()
		}
	})
	t.Run("째려보기 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("째려보기 방어자에게 변화 효과 방어 -1", func(t *testing.T) {
		if move.EffectOnDefender() != DEFENSE_MINUS_1 {
			t.Error()
		}
	})
}

func Test_할퀴기(t *testing.T) {
	move := MOVE_할퀴기

	t.Run("할퀴기 초기값 확인", func(t *testing.T) {
		if move.moveType != pokemonType.TypeNormal {
			t.Error()
		}
		if move.GetPower() != 40 {
			t.Error()
		}
		if move.GetPP() != 35 {
			t.Error()
		}
		if move.name != "할퀴기" {
			t.Error()
		}
	})
	t.Run("할퀴기 공격자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnAttacker() != NONE {
			t.Error()
		}
	})
	t.Run("할퀴기 방어자에게 변화 효과 없다", func(t *testing.T) {
		if move.EffectOnDefender() != NONE {
			t.Error()
		}
	})
}
