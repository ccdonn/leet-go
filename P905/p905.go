package main

func sortArrayByParity(A []int) []int {
	i := 0
	j := 0
	for ; j < len(A); j++ {
		if A[j]%2 == 0 {
			A[i] += A[j]
			A[j] = A[i] - A[j]
			A[i] -= A[j]
			i++
		}
	}
	return A
}
