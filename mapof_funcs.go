package xsync

// Keys returns a slice of all keys in the map.
func (m *MapOf[K, V]) Keys() []K {
	keys := make([]K, 0, m.Size())
	m.ForEach(func(key K, value V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

// Values returns a slice of all values in the map.
func (m *MapOf[K, V]) Values() []V {
	values := make([]V, 0, m.Size())
	m.ForEach(func(key K, value V) bool {
		values = append(values, value)
		return true
	})
	return values
}

// AsMap returns the entire map as a standard Go map.
func (m *MapOf[K, V]) AsMap() map[K]V {
	stdMap := make(map[K]V, m.Size())
	m.ForEach(func(key K, value V) bool {
		stdMap[key] = value
		return true
	})
	return stdMap
}

// Clone creates a deep copy of the map.
func (m *MapOf[K, V]) Clone() *MapOf[K, V] {
	clone := NewMap[K, V]()
	m.ForEach(func(key K, value V) bool {
		clone.Set(key, value)
		return true
	})
	return clone
}
