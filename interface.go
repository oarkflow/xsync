package xsync

type IMap[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
	GetOrSet(K, V) (V, bool)
	GetAndSet(K, V) (V, bool)
	GetOrCompute(K, func() V) (V, bool)
	Compute(K, func(V, bool) (V, bool)) (V, bool)
	GetAndDelete(K) (V, bool)
	Del(K)
	ForEach(func(K, V) bool)
	Clear()
	Size() int
	Stats() MapStats
	Keys() []K
	Values() []V
	AsMap() map[K]V
	Clone() *MapOf[K, V]
}
