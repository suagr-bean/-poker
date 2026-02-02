package Utils

//处理好手牌 过map 同时把已知手牌过滤掉
// 使用预计算的 MapHand，并只遍历需要的前缀（由 chance 决定）
func RangeToMap(chance int, cards []int) AllHand {
	all := RangeTable()
	limit := int((len(all.HandIndex) * chance) / 100)
	if limit <= 0 {
		return AllHand{}
	}
	res := AllHand{HandIndex: make([]Hands, 0, limit)}

OuterLoop:
	for i := 0; i < limit; i++ {
		h := all.HandIndex[i]
		mh := h.MapHand
		// 如果没有预计算（兼容性兜底），再计算一次
		if mh == nil {
			mh = DealMap(h.Hand)
		}
		for _, j := range cards {
			if mh[0] == j || mh[1] == j {
				continue OuterLoop
			}
		}
		h.MapHand = mh
		res.HandIndex = append(res.HandIndex, h)
	}
	return res
}