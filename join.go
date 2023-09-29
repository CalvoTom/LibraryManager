package librarymanager

func Join(strs []string, sep string) string {
	var string_join string
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			string_join += strs[i] + sep
		}
		if i == len(strs)-1 {
			string_join += strs[i]
		}
	}
	return string_join
}
