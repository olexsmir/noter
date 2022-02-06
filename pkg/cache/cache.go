package cache

type Cacher interface {
	Set(key, value interface{}, ttl int64) error
	Get(key interface{}) (interface{}, error)
}
