package matrix

func Matrix(size int) [][]int {
	arrarr := make([][]int, size)
	val := 1
	for i := 0; i < size; i++ {
		arr := make([]int, size)
		for x := 0; x < size; x++ {
			arr[x] = val
			val++
		}
		arrarr[i] = arr
	}
	return arrarr
}

func SpiralMatrix(n int) [][]int {

	// initialize empty matrix
	var arrarr = make([][]int, n)
	for i := 0; i < n; i++ {
		arrarr[i] = make([]int, n)
	}

	counter := 1
	startCol := 0
	endCol := n - 1
	startRow := 0
	endRow := n - 1

	for startCol <= endCol && startRow <= endRow {
		// populates the top row
		for i := startCol; i <= endCol; i++ {
			arrarr[startRow][i] = counter
			counter++
		}
		startRow++
		// populates last (right) column for the rest of the rows
		for i := startRow; i <= endRow; i++ {
			arrarr[i][endCol] = counter
			counter++
		}
		endCol--
		// bottom row
		for i := endCol; i >= startCol; i-- {
			arrarr[endRow][i] = counter
			counter++
		}
		endRow--
		// first column
		for i := endRow; i >= startRow; i-- {
			arrarr[i][startCol] = counter
			counter++
		}
		startCol++
	}
	return arrarr
}
