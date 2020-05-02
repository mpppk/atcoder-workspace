package lib

// MemoizeYYY2ToZZZ は、引数として与えた関数の戻り値をキャッシュするラッパーを返します.
func MemoizeYYY2ToZZZ(f func(v1, v2 YYY) ZZZ) func(v1, v2 YYY) ZZZ {
	cache := map[YYY]map[YYY]ZZZ{}

	return func(v1, v2 YYY) ZZZ {
		if v1Map, ok := cache[v1]; ok {
			if cachedValue, ok := v1Map[v2]; ok {
				return cachedValue
			}
		}

		result := f(v1, v2)
		_, v1Ok := cache[v1]
		if !v1Ok {
			cache[v1] = map[YYY]ZZZ{}
		}

		cache[v1][v2] = result
		return result
	}
}
