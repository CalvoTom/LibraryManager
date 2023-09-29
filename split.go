package librarymanager

func Split(s, sep string) []string {
	var tabStringSeparate []string
	lastIndice := 0
	for indice, val := range s {
		if val == rune(sep[0]) && s[indice+len(sep)-1] == sep[len(sep)-1] {
			if lastIndice == indice {
				lastIndice = indice + len(sep)
			} else {
				tabStringSeparate = append(tabStringSeparate, s[lastIndice:indice])
				lastIndice = indice + len(sep)
			}
		}
		if indice == len(s)-1 {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:])
		}
	}
	if tabStringSeparate[len(tabStringSeparate)-1] == "" {
		return tabStringSeparate[:len(tabStringSeparate)-1]
	}
	return tabStringSeparate
}
