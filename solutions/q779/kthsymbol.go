package q779

/**
O(N)
父位 = (k + 1) / 2
找到父位后选左右位即可
*/
func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	father := kthGrammar(N-1, (K+1)/2)
	if father == 0 {
		return 1 - K%2
	} else {
		return K % 2
	}
}
