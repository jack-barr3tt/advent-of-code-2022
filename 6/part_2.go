package main

import "os"

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	q := make([]byte, 0)

	for i := range data {
		q = append(q, data[i])
		if len(q) <= 14 {
			continue
		}else{
			q = q[1:]
		}

		same := false

		for j := range q {
			for k := range q {
				if j == k {
					continue
				}
				if q[j] == q[k] {
					same = true
					break
				}
			}
			if same {
				break
			}
		}

		if !same {
			println(i + 1)
			break
		}
	}
}