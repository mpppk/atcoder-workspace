package lib

type YYY2ToZZZCache map[YYY]map[YYY]ZZZ

func (c YYY2ToZZZCache) has(v1 YYY, v2 YYY) bool {
	if _, ok := c[v1]; !ok {
		return false
	}
	_, ok := c[v1][v2]
	return ok
}

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

func MemoizeYYY5ToZZZ(f func(v1, v2, v3, v4, v5 YYY) ZZZ) func(v1, v2, v3, v4, v5 YYY) ZZZ {
	cache := map[YYY]map[YYY]map[YYY]map[YYY]map[YYY]ZZZ{}

	return func(v1, v2, v3, v4, v5 YYY) ZZZ {
		if _, ok := cache[v1]; !ok {
			cache[v1] = map[YYY]map[YYY]map[YYY]map[YYY]ZZZ{}
		}
		if _, ok := cache[v1][v2]; !ok {
			cache[v1][v2] = map[YYY]map[YYY]map[YYY]ZZZ{}
		}
		if _, ok := cache[v1][v2][v3]; !ok {
			cache[v1][v2][v3] = map[YYY]map[YYY]ZZZ{}
		}
		if _, ok := cache[v1][v2][v3][v4]; !ok {
			cache[v1][v2][v3][v4] = map[YYY]ZZZ{}
		}
		if cachedValue, ok := cache[v1][v2][v3][v4][v5]; ok {
			return cachedValue
		}

		result := f(v1, v2, v3, v4, v5)
		cache[v1][v2][v3][v4][v5] = result
		return result
	}
}
