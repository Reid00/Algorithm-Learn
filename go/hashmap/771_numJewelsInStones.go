package hashmap

func numJewelsInStones(jewels string, stones string) int {
	// HashMap
	hmap := make(map[byte]int)

	for _, v := range stones {
		hmap[byte(v)] += 1
	}

	cnt := 0

	for _, v := range jewels {
		if val, ok := hmap[byte(v)]; ok {
			cnt += val
		}
	}
	return cnt
}
