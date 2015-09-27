package array

func InMap(sNeedle string, aHackstack map[string]string) bool {
	for _, sValue := range aHackstack {
		if sValue == sNeedle {
			return true
		}
	}
	return false
}

func CloneMap(aMap map[string]string) map[string]string {
	aMapReturn := make(map[string]string)
	for sKey, sValue := range aMap {
		aMapReturn[sKey] = sValue
	}
	return aMapReturn
}
