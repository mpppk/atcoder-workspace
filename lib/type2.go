package lib

import "fmt"

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

type YYYToZZZMap map[YYY]ZZZ
type YYYToZZZ2DMap map[YYY]map[YYY]ZZZ

var ZZZMapCapForYYYToZZZ2DMap = 0

func NewYYYToZZZMap(cap int) YYYToZZZMap {
	return make(map[YYY]ZZZ, cap)
}

func NewYYYToZZZ2DMap(cap, cap2 int) YYYToZZZ2DMap {
	ZZZMapCapForYYYToZZZ2DMap = cap2
	return make(map[YYY]map[YYY]ZZZ, cap)
}

func (m YYYToZZZ2DMap) Has(key1, key2 YYY) bool {
	v1, ok := m[key1]
	if !ok {
		return false
	}
	_, ok = v1[key2]
	if !ok {
		return false
	}
	return true

}

func (m YYYToZZZ2DMap) Set(key1, key2 YYY, value ZZZ) (isNewValue bool) {
	v1, ok := m[key1]
	if !ok {
		m[key1] = NewYYYToZZZMap(ZZZMapCapForYYYToZZZ2DMap)
		v1 = m[key1]
	}
	_, ok = v1[key2]
	v1[key2] = value
	return !ok
}

// MustGetは、指定したkeyの値を返します. 指定したkeyの値が存在しない場合panicします.
func (m YYYToZZZ2DMap) MustGet(key1, key2 YYY) ZZZ {
	v1, ok := m[key1]
	if !ok {
		panic(fmt.Sprintf("ivnalid key1 is specfied in AAAMap: %v", key1))
	}
	v2, ok := v1[key2]
	if !ok {
		panic(fmt.Sprintf("ivnalid key2 is specfied in AAAMap: %v", key2))
	}
	return v2
}

// GetOr は、指定したkeyの値が存在すればその値を、存在しなければdefaultValueを返します.
func (m YYYToZZZ2DMap) GetOr(key1, key2 YYY, defaultValue ZZZ) ZZZ {
	if !m.Has(key1, key2) {
		return defaultValue
	}
	return m[key1][key2]
}

// ChMax は、与えられた値が既に存在する値よりも大きれば代入します.
// 指定したkeyの値が存在しない場合も代入します. この場合、2つめの戻り値はfalseになります.
func (m YYYToZZZ2DMap) ChBy(key1, key2 YYY, f func(cur, new ZZZ) bool, value ZZZ, values ...ZZZ) (replaced bool, valueAlreadyExist bool) {
	maxBy := func(f func(cur, new ZZZ) bool, values ...ZZZ) ZZZ {
		max := values[0]
		for _, v := range values {
			if f(max, v) {
				max = v
			}
		}
		return max
	}
	max := maxBy(f, append(values, value)...)
	if m.Has(key1, key2) {
		if f(m[key1][key2], max) {
			m.Set(key1, key2, max)
			return true, true
		} else {
			return false, true
		}
	}
	m.Set(key1, key2, max)
	return true, false
}

func (m YYYToZZZ2DMap) GetMap(key YYY) (YYYToZZZMap, bool) {
	m1, ok := m[key]
	return m1, ok
}

func (m YYYToZZZ2DMap) MustGetMap(key YYY) YYYToZZZMap {
	m1, ok := m.GetMap(key)
	if !ok {
		panic(fmt.Sprintf("invalid key is given to MustGetMap: %v", key))
	}
	return m1
}
