package gateway

import (
	"sync"
)

// GMap 是一个泛型封装的 sync.Map
type GMap[K comparable, V any] struct {
	internal sync.Map
}

// Store 存储键值对
func (m *GMap[K, V]) Store(key K, value V) {
	m.internal.Store(key, value)
}

// Load 读取键对应的值
func (m *GMap[K, V]) Load(key K) (V, bool) {
	val, ok := m.internal.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return val.(V), true
}

// LoadOrStore 如果 key 存在就返回存在值，否则保存新值并返回
func (m *GMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	val, loaded := m.internal.LoadOrStore(key, value)
	return val.(V), loaded
}

// Delete 删除 key
func (m *GMap[K, V]) Delete(key K) {
	m.internal.Delete(key)
}

// Range 遍历所有键值对
func (m *GMap[K, V]) Range(f func(key K, value V) bool) {
	m.internal.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}
