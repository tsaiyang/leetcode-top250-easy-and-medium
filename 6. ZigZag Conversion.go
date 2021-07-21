package main

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	charMatrix := make([][]byte, numRows)
	for i := range charMatrix {
		charMatrix[i] = make([]byte, len(s))
	}
	baseNum := 2*numRows - 2
	for i := range s {
		chunk := i / baseNum
		if i < chunk*baseNum+numRows {
			charMatrix[i%baseNum][chunk*(numRows-1)] = s[i]
		} else {
			n := i % (chunk*baseNum + numRows - 1)
			charMatrix[numRows-1-n][chunk*(numRows-1)+1] = s[i]
		}
	}
	result := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if charMatrix[i][j] != '0' {
				result = append(result, charMatrix[i][j])
			}
		}
	}
	return string(result)
}
