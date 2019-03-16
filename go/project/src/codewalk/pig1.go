package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
游戏概览
Pig是由两个玩家掷6面骰子的游戏。在每一轮中，你可能掷骰或停留。
如果你掷出了1，就会失去你在这轮中的所有点数，并交由你的对手玩。 掷出的其它值都将计入你这轮的分数中。
如果你停留了，你这轮的分数就会计入你的总分中，并交由你的对手玩。
总点数先达到100的人胜出
 */

//分解、抽象、模式识别和算法

//1. A和B初始都是0分, 每轮摇N次, 获胜需要达到的分数S
//2. 随机选出一个人先开始第一轮(设为A)
//3. 第一轮, A摇色子, 如果摇到1, A本轮得分为0, 进入下一轮(轮到B玩), 摇到2-6, 计入A的总分, 本轮A最多摇N次
//4. A摇完后就到B摇, B摇完到A摇, 如此反复, 直到其中一个人得分为S, 则胜出

//角色(2人玩)/得分(2人份)
//当前轮的总分

//初始化A和B...
//摇一次
//摇一轮(最多摇N次)
//A和B轮流进行

func init(){
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

type player struct {
	name string
	score int
}

var (
	//每轮摇的次数
	roundCount = 5
	//胜出的最少分数
	winCount = 100
	//玩家
	playersLists = [2] player {
		{"A", 0},
		{"B", 0},
	}
)

//摇一次(0, 2, 3, 4, 5, 6)
func playTimes()(num int) {
	outcome := rand.Intn(6) + 1

	if outcome == 1 {
		num = 0
	}else{
		num = outcome
	}
	return
}

//摇一轮
//更新玩家分数
func playRound(p *player)(total int) {
	//临时存放摇一次的结果
	var num int

	//摇roundCount次
	for i:=0; i<roundCount; i++  {
		num = playTimes()
		if num == 0 {
			total = 0
			break
		}
		total += num
	}

	//加入到用户的总分
	p.score += total

	fmt.Println(p.name, "在这一轮中得分为", total, ", 总分为", p.score)
	return
}

//A和B轮流玩
func playTurn() {
	//玩家的总人数
	playerCount := len(playersLists)
	//当前玩家在玩家列表中的序号
	playerNum := rand.Intn(playerCount)
	//当前玩家的引用
	currentPlayers := &playersLists[playerNum]

	fmt.Println("本局比赛由", currentPlayers.name, "开始")

	for {
		//存放摇一轮的结果
		playRound(currentPlayers)
		if currentPlayers.score >= winCount {
			fmt.Println(currentPlayers.name, "赢得了比赛, 总分为", currentPlayers.score)
			break
		}

		//替换A和B
		if playerNum == playerCount - 1 {
			playerNum = 0
		}else{
			playerNum++
		}
		currentPlayers = &playersLists[playerNum]
	}
}

func main()  {
	playTurn()

	fmt.Println(playersLists)
}


