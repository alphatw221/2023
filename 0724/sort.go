package main // Required for a standalone executable.

import (
	"fmt"
)

func selection_sort(slice_address *[]int32) {

	slice := *slice_address
	for i := 0; i < len(slice); i++ {
		minIndex := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				temp := slice[j]
				slice[j] = slice[minIndex]
				slice[minIndex] = temp
			}
		}

	}

	return
}

func insertion_sort(slice_address *[]int32) {

	slice := *slice_address
	for i := 1; i < len(slice); i++ {

		numberIndex := i
		j := i - 1
		for j >= 0 {

			if slice[j] > slice[numberIndex] {
				temp := slice[j]
				slice[j] = slice[numberIndex]
				slice[numberIndex] = temp
				numberIndex = j
			}
			j -= 1
		}

	}

	return
}

func bubble_sort(slice_address *[]int32) {
	done := false
	slice := *slice_address

	for !done {
		done = true
		for i := 1; i < len(slice); i++ {
			if slice[i] < slice[i-1] {
				temp := slice[i]
				slice[i] = slice[i-1]
				slice[i-1] = temp
				done = false
			}

		}
	}
	return
}

func merge_sort(unsort_slice []int32) (sort_slice []int32) {
	if len(unsort_slice) == 1 {
		sort_slice = unsort_slice

	} else if len(unsort_slice) == 2 {
		if unsort_slice[0] > unsort_slice[1] {
			temp := unsort_slice[1]
			unsort_slice[1] = unsort_slice[0]
			unsort_slice[0] = temp
		}
		sort_slice = unsort_slice

	} else {
		half := int32(len(unsort_slice) / 2)
		slice_a := merge_sort(unsort_slice[:half])
		slice_b := merge_sort(unsort_slice[half:])

		i, j := 0, 0

		for i < len(slice_a) || j < len(slice_b) {
			if i == len(slice_a) {
				sort_slice = append(sort_slice, slice_b[j])
				j += 1
			} else if j == len(slice_b) {
				sort_slice = append(sort_slice, slice_a[i])
				i += 1
			} else if slice_a[i] < slice_b[j] {
				sort_slice = append(sort_slice, slice_a[i])
				i += 1
			} else {
				sort_slice = append(sort_slice, slice_b[j])
				j += 1
			}
		}

	}

	return

}

func main() {

	unsort_slice := []int32{1, 2, 3, 7, 45, 9, 5, 23, 86, 31, 478, 125, 46}
	bubble_sort(&unsort_slice)
	fmt.Println(unsort_slice)

	sort_slice := merge_sort(unsort_slice)
	fmt.Println(sort_slice)

}
