package pkg

import "fmt"

type WalkStrategy struct {
}

func (wS *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk A:[%d] to B[%d] Avg Speed:[%d] Total:[%d] TotalTime:[%d] min\n", startPoint, endPoint, avgSpeed,
		total, totalTime)
}
