package lib

import (
	"errors"
	"fmt"
)

// ReduceZZZ は、ZZZを引数として受け取るJavaScriptのreduceです.
func ReduceZZZ(values []ZZZ, f func(acc, cur ZZZ) ZZZ, initial ZZZ) (newValue ZZZ) {
	newValue = initial
	for _, value := range values {
		newValue = f(newValue, value)
	}
	return
}

// ReduceZZZSlice はZZZ Sliceを引数として受け取るJavaScriptのreduceです.
func ReduceZZZSlice(values [][]ZZZ, f func(acc ZZZ, cur []ZZZ) ZZZ, initial ZZZ) (newValue ZZZ) {
	newValue = initial
	for _, value := range values {
		newValue = f(newValue, value)
	}
	return
}

// CopyZZZ は、ZZZ Sliceをコピーします.
func CopyZZZ(values []ZZZ) []ZZZ {
	dst := make([]ZZZ, len(values))
	copy(dst, values)
	return dst
}

// CopyZZZSlice は、ZZZの2次元Sliceをコピーします.
func CopyZZZSlice(values [][]ZZZ) [][]ZZZ {
	dst := make([][]ZZZ, len(values))
	for i, value := range values {
		dst[i] = CopyZZZ(value)
	}
	return dst
}

// ReverseZZZ は、ZZZ Sliceの順序を反転させたコピーを返します.
func ReverseZZZ(values []ZZZ) []ZZZ {
	newValues := CopyZZZ(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}

// MapZZZ はZZZ Sliceを引数として受け取るJavaScriptのmapです.
func MapZZZ(values []ZZZ, f func(v ZZZ) ZZZ) (newValues []ZZZ) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

// MapZZZSlice はZZZの2次元Sliceを引数として受け取るJavaScriptのmapです.
func MapZZZSlice(values [][]ZZZ, f func(v []ZZZ) []ZZZ) (newValues [][]ZZZ) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

// ZipZZZ はZZZ Sliceを引数として受け取るrubyのzipです.
func ZipZZZ(valuesList ...[]ZZZ) (newValuesList [][]ZZZ, err error) {
	if len(valuesList) == 0 {
		return nil, errors.New("empty values list are given to ZipZZZ")
	}
	valuesLen := len(valuesList[0])
	for _, values := range valuesList {
		if (len(values)) != valuesLen {
			return nil, errors.New("different lengths values are given to ZipZZZ")
		}
	}

	for i := 0; i < valuesLen; i++ {
		var newValues []ZZZ
		for _, values := range valuesList {
			newValues = append(newValues, values[i])
		}
		newValuesList = append(newValuesList, newValues)
	}
	return
}

// SomeZZZ は、与えた条件式に一つでもマッチする値があればtrueを返します.
func SomeZZZ(values []ZZZ, f func(v ZZZ) bool) bool {
	for _, value := range values {
		if f(value) {
			return true
		}
	}
	return false
}

// SomeZZZSlice は、与えた条件式に一つでもマッチする値があればtrueを返します.
func SomeZZZSlice(valuesList [][]ZZZ, f func(v []ZZZ) bool) bool {
	for _, values := range valuesList {
		if f(values) {
			return true
		}
	}
	return false
}

// EveryZZZ は、与えた条件式に値が全てマッチすればtrueを返します.
func EveryZZZ(values []ZZZ, f func(v ZZZ) bool) bool {
	for _, value := range values {
		if !f(value) {
			return false
		}
	}
	return true
}

// EveryZZZSlice は、与えた条件式に値が全てマッチすればtrueを返します.
func EveryZZZSlice(valuesList [][]ZZZ, f func(v []ZZZ) bool) bool {
	for _, values := range valuesList {
		if !f(values) {
			return false
		}
	}
	return true
}

// ChunkZZZByBits は、bitがtrueになるごとにchunkを分割して返します.
func ChunkZZZByBits(values []ZZZ, bits []bool) (newValues [][]ZZZ, err error) {
	if len(values) != len(bits)+1 {
		return nil, errors.New(fmt.Sprintf("there are different length between values(%d) and bits(%d)", len(values), len(bits)))
	}

	var chunk []ZZZ
	for i, bit := range bits {
		chunk = append(chunk, values[i])
		if bit {
			newValues = append(newValues, chunk)
			chunk = []ZZZ{}
		}
	}
	chunk = append(chunk, values[len(values)-1])
	newValues = append(newValues, chunk)
	return
}

// UnsetZZZ は、i番目の要素を削除したSliceを返します.
func UnsetZZZ(values []ZZZ, i int) ([]ZZZ, error) {
	if i < 0 {
		return nil, fmt.Errorf("negative number index is given: %d", i)
	}

	if i >= len(values) {
		return nil, fmt.Errorf("index(%d) is larger than slice length(%d)", i, len(values))
	}

	if len(values) == 1 {
		return []ZZZ{}, nil
	}

	if i == len(values)-1 {
		return values[:len(values)-1], nil
	}

	newValues := make([]ZZZ, len(values))
	copy(newValues, values)
	return append(newValues[:i], newValues[i+1:]...), nil
}

// ZZZCombination は、与えられた値からr個を取り出す場合の全組み合わせを返します.
func ZZZCombination(values []ZZZ, r int) (combinations [][]ZZZ, err error) {
	if r == 1 {
		for _, value := range values {
			combinations = append(combinations, []ZZZ{value})
		}
		return
	}

	for i := range values {
		newValues := values
		for j := 0; j <= i; j++ {
			newValues = newValues[1:]
		}
		partialCombinations, err := ZZZCombination(newValues, r-1)
		if err != nil {
			return nil, err
		}

		for _, pc := range partialCombinations {
			newC := append(pc, values[i])
			combinations = append(combinations, newC)
		}
	}
	return
}

// ZZZPermutation は、与えらえれた値からr個を順序考慮して取り出す全組み合わせを返します。
func ZZZPermutation(values []ZZZ, r int) (permutations [][]ZZZ, err error) {
	if r == 1 {
		for _, value := range values {
			permutations = append(permutations, []ZZZ{value})
		}
		return
	}
	for i := range values {
		newValues, err := UnsetZZZ(values, i)
		if err != nil {
			return nil, err
		}
		perm, err := ZZZPermutation(newValues, r-1)
		if err != nil {
			return nil, err
		}

		for _, pc := range perm {
			newC := append(pc, values[i])
			permutations = append(permutations, newC)
		}
	}
	return
}

// TernaryOPZZZ は、okがtrueであればv1, falseであればv2を返します.
func TernaryOPZZZ(ok bool, v1, v2 ZZZ) ZZZ {
	if ok {
		return v1
	}
	return v2
}

// UnionFindZZZ はUnionFind木を表す構造体です.
type UnionFindZZZ struct {
	nodes map[ZZZ]ZZZ
}

// NewUnionFindZZZ は、与えられた値を初期値として持つUnionFind木の構造体を返します.
func NewUnionFindZZZ(values []ZZZ) *UnionFindZZZ {
	m := map[ZZZ]ZZZ{}
	for _, v := range values {
		m[v] = v
	}
	return &UnionFindZZZ{nodes: m}
}

// GetRoot は、与えられた値の根の値を返します.
func (u *UnionFindZZZ) GetRoot(value ZZZ) (ZZZ, int) {
	v := value
	newV := u.nodes[v]
	cnt := 0
	for newV != v {
		cnt++
		oldV := v
		v = newV
		newV = u.nodes[newV]
		u.nodes[oldV] = newV
	}
	return newV, cnt
}

// Unite は、v1とv2のグループをマージします.
func (u *UnionFindZZZ) Unite(v1, v2 ZZZ) (ZZZ, bool) {
	v1Root, v1HopNum := u.GetRoot(v1)
	v2Root, v2HopNum := u.GetRoot(v2)
	if v1Root == v2Root {
		return v1Root, false
	}
	if v1HopNum >= v2HopNum {
		u.nodes[v2Root] = v1Root
		return v1Root, true
	}
	u.nodes[v1Root] = v2Root
	return v2Root, true
}

// IsSameGroup は、v1とv2が同じグループに所属しているかを返します.
func (u *UnionFindZZZ) IsSameGroup(v1, v2 ZZZ) bool {
	v1Root, _ := u.GetRoot(v1)
	v2Root, _ := u.GetRoot(v2)
	return v1Root == v2Root
}

type MemoZZZ struct {
	M map[ZZZ]ZZZ
	F func(ZZZ, func(ZZZ) ZZZ) ZZZ
}

// MemoizeZZZ は、与えられた関数の戻り値をキャッシュし、2度目からはO(1)で結果を返すラッパーを返します.
func MemoizeZZZ(f func(ZZZ, func(ZZZ) ZZZ) ZZZ) func(ZZZ) ZZZ {
	m := &MemoZZZ{
		M: map[ZZZ]ZZZ{},
		F: f,
	}
	return m.call
}

func (m *MemoZZZ) call(v ZZZ) ZZZ {
	if cachedResult, ok := m.M[v]; ok {
		return cachedResult
	}
	res := m.F(v, m.call)
	m.M[v] = res
	return res
}

type CounterZZZ struct {
	M    map[ZZZ]int64
	Func func(values ...ZZZ) ZZZ
}

func NewCounterZZZ() *CounterZZZ {
	return &CounterZZZ{M: map[ZZZ]int64{}}
}

func (c CounterZZZ) CountBy(values ...ZZZ) {
	c.M[c.Func(values...)]++
}

// New2DimZZZSlice は、ZZZのゼロ値を持つ2次元sliceを返します.
func New2DimZZZSlice(row, col int) [][]ZZZ {
	ret := make([][]ZZZ, row)
	for r := 0; r < row; r++ {
		ret[r] = make([]ZZZ, col, col)
	}
	return ret
}

// NewZZZSliceWithInitialValue は、initialValueを値として持つZZZのSliceを返します.
func NewZZZSliceWithInitialValue(length int, initialValue ZZZ) []ZZZ {
	ret := make([]ZZZ, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	return ret
}

// New2DimZZZSliceWithInitialValue は、initialValueを値として持つZZZの2次元Sliceを返します.
func New2DimZZZSliceWithInitialValue(row, col int, initialValue ZZZ) [][]ZZZ {
	ret := make([][]ZZZ, row)
	values := NewZZZSliceWithInitialValue(col, initialValue)
	for r := 0; r < row; r++ {
		ret[r] = CopyZZZ(values)
	}
	return ret
}
