package cache

import "time"

type cache interface {
	Get(key string, dest interface{}) error
	Set(key string, value interface{}, ttl time.Duration) error
	DeleteByPrefix(prefix string) error
}
