package array

func InSlice(sNeedle string, aHackstack []string) bool {
	for _, sValue := range aHackstack {
		if sValue == sNeedle {
			return true
		}
	}
	return false
}
