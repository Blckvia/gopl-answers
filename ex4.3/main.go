// ex3.4 reverse function uses a pointer to an array to reverse it
package main

func reverse(arr *[3]string) {
	for i, j := 0, len(arr) - 1; i< j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}