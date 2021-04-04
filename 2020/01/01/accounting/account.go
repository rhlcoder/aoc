package accounting

import "errors"

var ErrNoPairFound = errors.New("no pair of numbers satisfies the requirements")

func SumObjetive(slice []int, objetive int) (x, y, iterations int, e error) {
	l := len(slice)

	for i := 0; i < l; i++ {
		if slice[i] > objetive {
			continue
		}

		for j := i + 1; j < l; j++ {
			if slice[i]+slice[j] == objetive {
				return slice[i], slice[j], iterations, nil
			}
			iterations++
		}
	}

	return x, y, iterations, ErrNoPairFound
}
