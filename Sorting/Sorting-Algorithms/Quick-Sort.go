package main

import "fmt"

func main() {
	var list = []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("unsorted:", list)
	quicksort(list)
	fmt.Println("sorted:  ", list)
}

func quicksort(list []int) {
	var pex func(int, int)
	pex = func(lower, upper int) {
		for {
			switch upper - lower {
			case -1, 0:
				return
			case 1:
				if list[upper] < list[lower] {
					list[upper], list[lower] = list[lower], list[upper]
				}
				return
			}
			bx := (upper + lower) / 2
			b := list[bx]
			lp := lower
			up := upper
		outer:
			for {
				for lp < upper && !(b < list[lp]) {
					lp++
				}
				for {
					if lp > up {
						break outer
					}
					if list[up] < b {
						break
					}
					up--
				}
				list[lp], list[up] = list[up], list[lp]
				lp++
				up--
			}
			if bx < lp {
				if bx < lp-1 {
					list[bx], list[lp-1] = list[lp-1], b
				}
				up = lp - 2
			} else {
				if bx > lp {
					list[bx], list[lp] = list[lp], b
				}
				up = lp - 1
				lp++
			}
			if up-lower < upper-lp {
				pex(lower, up)
				lower = lp
			} else {
				pex(lp, upper)
				upper = up
			}
		}
	}
	pex(0, len(list)-1)
}
