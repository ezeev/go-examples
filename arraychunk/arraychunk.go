package arraychunk

func IntChunk(in []int, chunkSize int) [][]int {
	chunked := [][]int{}
	idx := 0
	for idx < len(in) {
		lastPos := idx + chunkSize
		if lastPos > len(in) {
			lastPos = len(in)
		}

		slice := in[idx:lastPos]
		chunked = append(chunked, slice)
		idx += chunkSize
	}

	return chunked
}
