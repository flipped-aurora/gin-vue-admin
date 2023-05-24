package util

// SliceChunk 用于将字符串切片分块
func SliceChunk(src []string, chunkSize int) (chunks [][]string) {
	total := len(src)
	chunks = make([][]string, 0)
	if chunkSize < 1 {
		chunkSize = 1
	}
	if total == 0 {
		return
	}

	chunkNum := total / chunkSize
	if total%chunkSize != 0 {
		chunkNum++
	}

	chunks = make([][]string, chunkNum)

	for i := 0; i < chunkNum; i++ {
		for j := 0; j < chunkSize; j++ {
			offset := i*chunkSize + j
			if offset >= total {
				return
			}

			if chunks[i] == nil {
				actualChunkSize := chunkSize
				if i == chunkNum-1 && total%chunkSize != 0 {
					actualChunkSize = total % chunkSize
				}
				chunks[i] = make([]string, actualChunkSize)
			}

			chunks[i][j] = src[offset]
		}
	}

	return
}
