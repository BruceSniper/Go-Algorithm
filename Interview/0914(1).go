package main

func main() {

}

func RotateMatrix(matrix [][]int) [][]int {
	// write code here
	width := len(matrix)
	pMatrix := make([][]int, width)
	for i := 0; i < width; i++ {
		pMatrix[i] = make([]int, width)
	}

	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			pMatrix[i][j] = matrix[i][j]
		}
	}

	//

	for i := 0; i < width; i++ {
		for j := 0; j < width-i; j++ {
			pMatrix[i][j], pMatrix[width-j-1][width-i-1] = pMatrix[width-j-1][width-i-1], pMatrix[i][j]
		}
	}

	for i := 0; i < width; i++ {
		for j := 0; j < width/2; j++ {
			pMatrix[i][j], pMatrix[i][width-j-1] = pMatrix[i][width-j-1], pMatrix[i][j]
		}
	}

	return pMatrix
}
