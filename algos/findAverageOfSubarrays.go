package main

func findAverageOfSubarrays(k int, a []float64) []float64 {
	r := []float64{}
	wsu, wst := 0.0, 0
	for we := range a {
		wsu += a[we]
		if we >= k - 1 {
			r = append(r, wsu / float64(k))
			wsu -= a[wst]
			wst += 1
		}
	}

	return r
}
