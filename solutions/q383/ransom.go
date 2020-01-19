package q383


func canConstruct(ransomNote string, magazine string) bool {
	arr := make([]int, 26)

	for i := 0; i < len(magazine); i++ {
		arr[magazine[i] - 'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i] - 'a']--
		if arr[ransomNote[i] - 'a'] < 0 {
			return false
		}
	}
	return true
}