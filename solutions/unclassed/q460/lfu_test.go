package q460

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestLfu(t *testing.T) {
	input := `
["LFUCache","put","put","get","put","get","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
[[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]
["LFUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
`
	output := `
[null,null,null,1,null,-1,3,null,-1,3,4]
[null,null,null,2,1,2,null,null,-1,2,1,4]
[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
`
	tests := strings.Split(input, "\n")
	wants := strings.Split(output, "\n")
	var (
		lfu LFUCache
		ops []string
		kvs [][]int
		expects []int
	)
	pos := 0
	for i := 0; i < len(tests); i++ {
		if len(tests[i]) < 1 {
			continue
		}
		if len(wants[pos]) < 1 {
			pos++
		}
		_ = json.Unmarshal([]byte(tests[i]), &ops)
		_ = json.Unmarshal([]byte(tests[i+1]), &kvs)
		_ = json.Unmarshal([]byte(wants[pos]), &expects)
		//fmt.Printf("%+q\n", ops)
		//fmt.Printf("%+v\n", expects)
		lfu = Constructor(kvs[0][0])
		fmt.Println("\ncap", kvs[0][0])
		got := 0
		for j := 1; j < len(ops); j++ {
			switch ops[j] {
			case "put":
				lfu.Put(kvs[j][0], kvs[j][1])
				t.Logf("put k: %d, v: %d", kvs[j][0], kvs[j][1])
			case "get":
				got = lfu.Get(kvs[j][0])
				t.Logf("get k: %d", kvs[j][0])
				if got != expects[j] {
					t.Errorf("get key %d, expects: %d, got %d", kvs[j][0], expects[j], got)
				}
			}
		}
		i++
		pos++
	}
}
