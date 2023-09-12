package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	DefaultConfigPaths = []string{"configs"}
	DefaultConfigType  = JSON
)

type Configuration struct {
	configs     map[string]*viper.Viper
	configPaths []string
	configType  string
}

func New() *Configuration {
	c := &Configuration{
		configs:     make(map[string]*viper.Viper),
		configPaths: DefaultConfigPaths,
		configType:  DefaultConfigType,
	}
	return c
}

func (c *Configuration) SetConfigPath(path ...string) *Configuration {
	c.configPaths = path
	return c
}

func (c *Configuration) AddConfigPath(path ...string) *Configuration {
	c.configPaths = append(c.configPaths, path...)
	return c
}

func (c *Configuration) SetConfigType(configType string) *Configuration {
	c.configType = configType
	return c
}

func (c *Configuration) Load(name string, opts ...Option) error {
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType(c.configType)
	for _, configPath := range c.configPaths {
		v.AddConfigPath(configPath)
	}
	for _, option := range opts {
		if err := option(v); err != nil {
			return err
		}
	}
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	c.configs[name] = v
	return nil
}

func (c *Configuration) MustLoad(name string, opts ...Option) {
	if err := c.Load(name, opts...); err != nil {
		panic(err)
	}
}

func (c *Configuration) Has(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) == 0 {
		return false
	}
	v := c.configs[keys[0]]
	if v == nil {
		return false
	}
	return v.IsSet(strings.Join(keys[1:], ","))
}

func (c *Configuration) Get(key string, defaultValue ...any) any {
	var value any
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	if len(keys) == 0 {
		return value
	}
	v := c.configs[keys[0]]
	if v == nil {
		return value
	}
	if val := v.Get(strings.Join(keys[1:], ".")); val != nil {
		value = val
	}
	return value
}

func (c *Configuration) GetString(key string, defaultValue ...string) string {
	var value string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetString(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetBool(key string, defaultValue ...bool) bool {
	var value bool
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetBool(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetInt(key string, defaultValue ...int) int {
	var value int
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetInt(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetInt32(key string, defaultValue ...int32) int32 {
	var value int32
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetInt32(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetInt64(key string, defaultValue ...int64) int64 {
	var value int64
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetInt64(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetUint(key string, defaultValue ...uint) uint {
	var value uint
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetUint(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetUint16(key string, defaultValue ...uint16) uint16 {
	var value uint16
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetUint16(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetUint32(key string, defaultValue ...uint32) uint32 {
	var value uint32
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetUint32(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetUint64(key string, defaultValue ...uint64) uint64 {
	var value uint64
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetUint64(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetTime(key string, defaultValue ...time.Time) time.Time {
	var value time.Time
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetTime(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetDuration(key string, defaultValue ...time.Duration) time.Duration {
	var value time.Duration
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetDuration(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetIntSlice(key string, defaultValue ...[]int) []int {
	var value []int
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetIntSlice(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetStringSlice(key string, defaultValue ...[]string) []string {
	var value []string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetStringSlice(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetStringMap(key string, defaultValue ...map[string]any) map[string]any {
	var value map[string]any
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetStringMap(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetStringMapString(key string, defaultValue ...map[string]string) map[string]string {
	var value map[string]string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetStringMapString(strings.Join(keys[1:], ","))
}

func (c *Configuration) GetStringMapStringSlice(key string, defaultValue ...map[string][]string) map[string][]string {
	var value map[string][]string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return value
	}
	keys := strings.Split(key, ".")
	return c.configs[keys[0]].GetStringMapStringSlice(strings.Join(keys[1:], ","))
}
