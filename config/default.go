package config

import (
	"time"

	"golang.org/x/exp/constraints"
)

var configs = New()

func SetConfigPath(path ...string) *Configuration {
	return configs.SetConfigPath(path...)
}

func AddConfigPath(path ...string) *Configuration {
	return configs.AddConfigPath(path...)
}

func SetConfigType(configType string) *Configuration {
	return configs.SetConfigType(configType)
}

func Load(name string, opts ...Option) error {
	return configs.Load(name, opts...)
}

func MustLoad(name string, opts ...Option) {
	configs.MustLoad(name, opts...)
}

func Has(key string) bool {
	return configs.Has(key)
}

func Get[T constraints.Ordered |
	[]T |
	[]any |
	map[string]any](key string, defaultValue ...any) T {
	return configs.Get(key, defaultValue...).(T)
}

func GetString(key string, defaultValue ...string) string {
	return configs.GetString(key, defaultValue...)
}

func GetBool(key string, defaultValue ...bool) bool {
	return configs.GetBool(key, defaultValue...)
}

func GetInt(key string, defaultValue ...int) int {
	return configs.GetInt(key, defaultValue...)
}

func GetInt32(key string, defaultValue ...int32) int32 {
	return configs.GetInt32(key, defaultValue...)
}

func GetInt64(key string, defaultValue ...int64) int64 {
	return configs.GetInt64(key, defaultValue...)
}

func GetUint(key string, defaultValue ...uint) uint {
	return configs.GetUint(key, defaultValue...)
}

func GetUint16(key string, defaultValue ...uint16) uint16 {
	return configs.GetUint16(key, defaultValue...)
}

func GetUint32(key string, defaultValue ...uint32) uint32 {
	return configs.GetUint32(key, defaultValue...)
}

func GetUint64(key string, defaultValue ...uint64) uint64 {
	return configs.GetUint64(key, defaultValue...)
}

func GetTime(key string, defaultValue ...time.Time) time.Time {
	return configs.GetTime(key, defaultValue...)
}

func GetDuration(key string, defaultValue ...time.Duration) time.Duration {
	return configs.GetDuration(key, defaultValue...)
}

func GetIntSlice(key string, defaultValue ...[]int) []int {
	return configs.GetIntSlice(key, defaultValue...)
}

func GetStringSlice(key string, defaultValue ...[]string) []string {
	return configs.GetStringSlice(key, defaultValue...)
}

func GetStringMap(key string, defaultValue ...map[string]any) map[string]any {
	return configs.GetStringMap(key, defaultValue...)
}

func GetStringMapString(key string, defaultValue ...map[string]string) map[string]string {
	return configs.GetStringMapString(key, defaultValue...)
}

func GetStringMapStringSlice(key string, defaultValue ...map[string][]string) map[string][]string {
	return configs.GetStringMapStringSlice(key, defaultValue...)
}
