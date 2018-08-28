package main

import "fmt"

/**
	使用广度优先算法走迷宫
    @author good mood
    2018/7/18 下午5:07
*/

type myPoint struct {
	i int
	j int
}

type myGrid struct {
	point myPoint
	pass  bool
}

var maze = [...][6]myGrid{
	{{myPoint{0,0},true},{myPoint{0,1},false},{myPoint{0,2},true},{myPoint{0,3},true},{myPoint{0,4},true},{myPoint{0,5},true}},
	{{myPoint{1,0},true},{myPoint{1,1},true},{myPoint{1,2},true},{myPoint{1,3},true},{myPoint{1,4},false},{myPoint{1,5},true}},
	{{myPoint{2,0},true},{myPoint{2,1},false},{myPoint{2,2},true},{myPoint{2,3},false},{myPoint{2,4},true},{myPoint{2,5},false}},
	{{myPoint{3,0},false},{myPoint{3,1},true},{myPoint{3,2},true},{myPoint{3,3},true},{myPoint{3,4},true},{myPoint{3,5},true}},
	{{myPoint{4,0},true},{myPoint{4,1},true},{myPoint{4,2},false},{myPoint{4,3},true},{myPoint{4,4},false},{myPoint{4,5},true}},
}

var minMazeI = 0
var maxMazeI = 4
var minMazeJ = 0
var maxMazeJ = 5

var start = myPoint{0,0}
var end = myPoint{4,5}



func (p *myPoint) getCurrentPoint() myPoint {
	return *p
}

func nextPointCanPass(p myPoint) bool {
	if p.i < minMazeI || p.i > maxMazeI || p.j < minMazeJ || p.j > maxMazeJ {
		return false
	}

	for _, row := range maze{
		for _, col := range row{
			if p != col.point {
				continue
			}

			for _, v := range alreadyExploreQueue  {
				if p == v {
					return false
				}
			}

			if col.pass {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// 探索方向  上左下右
func (p *myPoint) searchNext() []myPoint {
	//var queue = make([]myPoint, 0)
	var queue []myPoint

	// 👆
	if nextPointCanPass(myPoint{p.i - 1, p.j}) {
		queue = append(queue, myPoint{p.i -1, p.j})
	}

	// 👈
	if nextPointCanPass(myPoint{p.i, p.j - 1}) {
		queue = append(queue, myPoint{p.i, p.j - 1})
	}

	// 👇
	if nextPointCanPass(myPoint{p.i + 1, p.j}) {
		queue = append(queue, myPoint{p.i + 1, p.j})
	}

	// 👉
	if nextPointCanPass(myPoint{p.i, p.j + 1}) {
		queue = append(queue, myPoint{p.i, p.j + 1})
	}

	return queue
}

func dupQueue(queue []myPoint) []myPoint {
	var dupQueue = make([]myPoint, 0)
	hasDup := false
	for _, p := range queue{
		hasDup = false
		for _, v := range dupQueue  {
			if v == p {
				hasDup = true
				break
			}
		}
		if !hasDup {
			dupQueue = append(dupQueue, p)
		}
	}
	return dupQueue
}

func exploreQueueMap(nextSearchMap, currentTrace map[myPoint][]myPoint, currentPoint myPoint) {

	trace := currentTrace[currentPoint]

	length := len(trace)

	for _, p := range nextSearchMap[currentPoint]{

		alreadyExploreQueue = append(alreadyExploreQueue, p)
		trace = append(trace[:length], p)
		traces[p] = trace
		currentTrace[p] = trace
		if p == end {
			return
		}
		s := p.searchNext()
		searchQueues[p] = s

		nextSearchMap[p] = s

		exploreQueueMap(nextSearchMap, currentTrace, p)
	}

}

var alreadyExploreQueue = make([]myPoint, 0)
var traces = make(map[myPoint][]myPoint)
var searchQueues = make(map[myPoint][]myPoint)

func main() {
	var currentPoint = start

	needExploreQueue1 := make([]myPoint, 0)

	needExploreQueue1 = append(needExploreQueue1, currentPoint)


	currentTrace := map[myPoint][]myPoint{currentPoint:{currentPoint}}
	nextSearchMap := map[myPoint][]myPoint{currentPoint:currentPoint.searchNext()}

	exploreQueueMap(nextSearchMap, currentTrace, currentPoint)

	fmt.Println(traces[end])

}
