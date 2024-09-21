# go-simple-pokemon-battle
## Overview
Go 학습을 위한 간단한 포켓몬 배틀 프로그램 

## Features
### 포켓몬
- 불, 물, 풀 3가지 타입만 있다.
- 스탯은 공격, 방어, 스피드가 있다.
  - 스피드가 빠른 포캣몬이 선공한다.
  - 각 포켓몬은 정해진 스피드를 가진다.
- 4개의 기술을 가진다. 

#### 이상해씨
|  공격  |  방어  | 스피드  |
|:----:|:----:|:----:|
|  65  |  65  |  45  |

- 풀 타입
- 기술
  - 몸통박치기
  - 울음소리
  - 덩굴채찍 

#### 파이리
|  공격  |  방어  | 스피드  |
|:----:|:----:|:----:|
|  60  |  50  |  65  |

- 불 타입
- 기술
  - 할퀴기
  - 울음소리
  - 불꽃세례
  - 째려보기



#### 꼬부기
|  공격  |  방어  | 스피드  |
|:----:|:----:|:----:|
|  50  |  65  |  43  |

- 물 타입
- 기술
  - 몸통박치기
  - 꼬리흔들기
  - 물대포

### 배틀
- 포켓몬 A, B가 있다.
- 턴을 번갈아가며 기술을 사용한다.
- 기술 사용시 pp 1 감소한다.
- 공격은 데미지만큼 상대의 hp를 깎는다.
- 강한 상성은 약한 상성에 1.5배 데미지를 주고 0.5배 데미지를 입는다.
- hp가 먼저 0이 되면 패배한다. 
- 
#### 데미지 계산 
- 배틀의 각 기술은 공격, 방어, 랭크와 기술 위력을 조합한 데미지를 입힌다.
- 데미지 = ((공격력 * 기술위력 * 랭크) / (방어력 * 랭크)) * 상성

### 기술 
- 공격, 변화 기술이 있다.
- 변화 기술은 자신이나 상대의 랭크를 변화시킨다.
- 불, 물, 풀, 노말 타입이 있다.
- 지정된 pp를 가진다.

#### 랭크 
- 최대 랭크 변화는 2이다. 

|  -   |     상승     |    하락     |
|:----:|:----------:|:---------:|
| 1랭크  | (3/2)150%  | (2/3)66%  |
| 2랭크  | (4/2)200%  | (2/4)50%  |

#### 기술 목록
|  기술명  | 타입 | pp | 위력 |    효과     |
|:-----:|:--:|:--:|:--:|:---------:|
|  할퀴기  | 노말 | 35 | 40 |     -     |
| 몸통박치기 | 노말 | 35 | 35 |     -     |
| 울음소리  | 노말 | 40 | 0  |  상대 공격-1  |
| 불꽃세례  | 불  | 25 | 40 |     -     |
| 덩굴채찍  | 풀  | 10 | 35 |     -     |
|  물대포  | 풀  | 25 | 40 |     -     |
| 째려보기  | 노말 | 30 | 0  |  상대 방어-1  |
| 꼬리흔들기 | 노말 | 30 | 0  |  상대 방어-1  |