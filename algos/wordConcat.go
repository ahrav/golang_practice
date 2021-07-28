package main

func wordConcat(s string, a []string) []int {
	res := []int{}
	wc, wl := len(a), len(a[0])
	if wc == 0 || wl == 0 {
		return res
	}

	wf := make(map[string]int)
	for _, val := range a {
		wf[val]++
	}

	r := (len(s) - wc * wl) +1
	for i := 0; i < r; i++ {
		ws := make(map[string]int)
		for j := 0; j < wc; j++ {
			nwi := i + j * wl
			w := s[nwi:nwi + wl]
			if _, ok := wf[w]; ok {
				if _, ok := ws[w]; ok {
					ws[w]++
				} else {
					ws[w] = 0
				}
			} else {
				break
			}
			if ws[w] > wf[w] {
				break
			}
			if j + 1 == wc {
				res = append(res, i)
			}
		}
	}
	return res
}
