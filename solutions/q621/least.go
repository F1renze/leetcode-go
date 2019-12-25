package q621

import "sort"

func leastInterval(tasks []byte, n int) int {
	taskCnt := len(tasks)
	if taskCnt < 1 {
		return 0
	}
	slots := make([]int, 26)

	for i := range tasks {
		slots[tasks[i] - 'A']++
	}
	// 排序
	sort.Ints(slots)

	// 虚拟列 = 最多任务数 * (n + 1)
	maxCnt := slots[25] -1
	idles := maxCnt * n

	for i := 24; i >= 0 && slots[i] > 0; i-- {
		// min 等同于将数量并列最大的任务数 - 1，因安排到最后一行
		idles -= min(slots[i], maxCnt)
	}

	if idles <= 0 {
		return taskCnt
	}
	return taskCnt + idles
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


