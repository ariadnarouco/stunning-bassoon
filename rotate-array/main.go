package main

import (
	"fmt"
)

func main() {

	fmt.Println("Problem: Rotate an array of n elements to the right by k steps. \nFor example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4]. \nHow many different ways do you know to solve this problem?\n*****************")
	fmt.Println("Begin")
	array := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(array)
	k := 3
	rotateArray(array, n, k)
	fmt.Printf("Result of Rotate with O(n) space, O(n) time: %v \n", array)

	rotateArrayAlternative(array, n, k)
	fmt.Printf("Result of Rotate with O(1) space, O(n*k) time: %v \n", array)

	fmt.Println("End")

}

func rotateArray(arrayA []int, n, k int) {
	arrayB := make([]int, n)
	for i := 0; i < n-k; i++ {
		arrayB[i] = arrayA[i]
	}

	j := n - k
	for i := 0; i < k; i++ {

		arrayA[i] = arrayA[j]
		j++
	}

	j = k
	for i := 0; i < n-k; i++ {
		arrayA[j] = arrayB[i]
		j++
	}
}

func rotateArrayAlternative(array []int, n, k int) {
	for i := 0; i < k; i++ {
		last := array[n-1]
		for j := n - 1; j > 0; j-- {
			array[j] = array[j-1]
		}
		array[0] = last
	}
}
