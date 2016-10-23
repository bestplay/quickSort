/**
* Quick sort
*
* Use the first element as base value.
*
**/
package quickSort

// import "fmt"

// // func main() {
// // 	s := []int{6, 4, 5, 8, 9, 11, 20, 10, 2, 1, 3, 21, 13, 14, 16}

// // 	adustArr(s, 0, len(s)-1)
// // 	fmt.Println(s)
// // }

func adustArr(arr []int, l int, r int) {

	if l >= r {
		return
	}
	i, j := l, r
	x := arr[i]
	for i < j {
		for i < j && arr[j] >= x {
			j--
		}

		if i < j {
			arr[i] = arr[j]
		}

		for i < j && (arr[i] <= x) {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}

	}
	arr[i] = x
	adustArr(arr, l, i-1)
	adustArr(arr, i+1, r)
}

func Sort(arr []int) {
	adustArr(arr, 0, len(arr)-1)
}
