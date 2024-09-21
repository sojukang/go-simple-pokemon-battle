package move

type Move interface {
	GetPower() int
	EffectOnAttacker() int
	EffectOnDefender() int
	GetPP() int
}

var MOVE_꼬리흔들기 = new꼬리흔들기()
var MOVE_덩굴채찍 = new덩굴채찍()
var MOVE_몸통박치기 = new몸통박치기()
var MOVE_물대포 = new물대포()
var MOVE_불꽃세례 = new불꽃세례()
var MOVE_울음소리 = new울음소리()
var MOVE_째려보기 = new째려보기()
var MOVE_할퀴기 = new할퀴기()
