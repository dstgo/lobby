package maputil

// GetFbMap returns the value of the key, otherwise it returns value of the fallback.
func GetFbMap[M ~map[K]V, K comparable, V any](key, fbkey K, m M) V {
	if val, ok := m[key]; ok {
		return val
	}
	return m[fbkey]
}
