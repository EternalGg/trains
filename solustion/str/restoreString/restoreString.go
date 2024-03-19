package restoreString

func restoreString(s string, indices []int) string {
	qa := make([]uint8, len(indices))

	for i, index := range s {
		qa[indices[i]] = uint8(index)
	}

	return string(qa)
}
