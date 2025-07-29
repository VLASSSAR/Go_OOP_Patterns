package pkg

import "fmt"

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("PublicTransport A:[%d] to B[%d] Avg Speed:[%d] Total:[%d] TotalTime:[%d] min\n", startPoint, endPoint, avgSpeed,
		total, totalTime)
}
