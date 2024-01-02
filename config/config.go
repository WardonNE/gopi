package config

import (
	"io"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/wardonne/gopi/support/maps"
)

// Default config settings
var (
	DefaultConfigPaths = []string{"configs"}
	DefaultConfigType  = JSON
)

// Configuration is used to manage configs
type Configuration struct {
	configs     *maps.HashMap[string, *viper.Viper]
	configPaths []string
	configType  string
}

// New creates a new [Configuration] instance
func New() *Configuration {
	c := &Configuration{
		configs:     maps.NewHashMap[string, *viper.Viper](),
		configPaths: DefaultConfigPaths,
		configType:  DefaultConfigType,
	}
	return c
}

// SetConfigPath sets a list of config file directories for searching
func (c *Configuration) SetConfigPath(path ...string) *Configuration {
	c.configPaths = path
	return c
}

// AddConfigPath adds a new list of config file directories for searching
func (c *Configuration) AddConfigPath(path ...string) *Configuration {
	c.configPaths = append(c.configPaths, path...)
	return c
}

// SetConfigType sets the config file extension
func (c *Configuration) SetConfigType(configType string) *Configuration {
	c.configType = configType
	return c
}

// Load loads config file
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
	c.configs.Set(name, v)
	return nil
}

// MustLoad loads config file but panic when error
func (c *Configuration) MustLoad(name string, opts ...Option) {
	if err := c.Load(name, opts...); err != nil {
		panic(err)
	}
}

// LoadFromReader loads config from a reader and registers as specific name
func (c *Configuration) LoadFromReader(name string, r io.Reader, opts ...Option) error {
	v := viper.New()
	for _, option := range opts {
		if err := option(v); err != nil {
			return err
		}
	}
	if err := v.ReadConfig(r); err != nil {
		return err
	}
	c.configs.Set(name, v)
	return nil
}

func (c *Configuration) GetViper(key string) *viper.Viper {
	return c.configs.Get(key)
}

// Has checks whether the config key exists
func (c *Configuration) Has(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) == 0 {
		return false
	}
	v := c.configs.Get(keys[0])
	if v == nil {
		return false
	}
	return v.IsSet(strings.Join(keys[1:], "."))
}

// Get gets the config value by specific key
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
	v := c.configs.Get(keys[0])
	if v == nil {
		return value
	}
	if val := v.Get(strings.Join(keys[1:], ".")); val != nil {
		value = val
	}
	return value
}

// GetString gets a string config value by specific key
func (c *Configuration) GetString(key string, defaultValue ...string) *string {
	var value string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetString(strings.Join(keys[1:], "."))
	return &value
}

// GetBool gets a bool config value by specific key
func (c *Configuration) GetBool(key string, defaultValue ...bool) *bool {
	var value bool
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetBool(strings.Join(keys[1:], "."))
	return &value
}

// GetInt gets an int config value by specific key
func (c *Configuration) GetInt(key string, defaultValue ...int) *int {
	var value int
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetInt(strings.Join(keys[1:], "."))
	return &value
}

// GetInt32 gets an int32 config value by specific key
func (c *Configuration) GetInt32(key string, defaultValue ...int32) *int32 {
	var value int32
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetInt32(strings.Join(keys[1:], "."))
	return &value
}

// GetInt64 gets an int64 config value by specific key
func (c *Configuration) GetInt64(key string, defaultValue ...int64) *int64 {
	var value int64
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetInt64(strings.Join(keys[1:], "."))
	return &value
}

// GetUint gets an uint config value by specific key
func (c *Configuration) GetUint(key string, defaultValue ...uint) *uint {
	var value uint
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetUint(strings.Join(keys[1:], "."))
	return &value
}

// GetUint16 gets an uint16 config value by specific key
func (c *Configuration) GetUint16(key string, defaultValue ...uint16) *uint16 {
	var value uint16
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetUint16(strings.Join(keys[1:], "."))
	return &value
}

// GetUint32 gets an uint32 config value by specific key
func (c *Configuration) GetUint32(key string, defaultValue ...uint32) *uint32 {
	var value uint32
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetUint32(strings.Join(keys[1:], "."))
	return &value
}

// GetUint64 gets an uint64 config value by specific key
func (c *Configuration) GetUint64(key string, defaultValue ...uint64) *uint64 {
	var value uint64
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetUint64(strings.Join(keys[1:], "."))
	return &value
}

// GetTime gets a [time.Time] config value by specific key
func (c *Configuration) GetTime(key string, defaultValue ...time.Time) *time.Time {
	var value time.Time
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetTime(strings.Join(keys[1:], "."))
	return &value
}

// GetDuration gets a [time.Duration] config value by specific key
func (c *Configuration) GetDuration(key string, defaultValue ...time.Duration) *time.Duration {
	var value time.Duration
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetDuration(strings.Join(keys[1:], "."))
	return &value
}

// GetIntSlice get a []int config value by specific key
func (c *Configuration) GetIntSlice(key string, defaultValue ...[]int) *[]int {
	var value []int
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetIntSlice(strings.Join(keys[1:], "."))
	return &value
}

// GetStringSlice get a []string config value by specific key
func (c *Configuration) GetStringSlice(key string, defaultValue ...[]string) *[]string {
	var value []string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetStringSlice(strings.Join(keys[1:], "."))
	return &value
}

// GetStringMap get a map[string]any config value by specific key
func (c *Configuration) GetStringMap(key string, defaultValue ...map[string]any) *map[string]any {
	var value map[string]any
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetStringMap(strings.Join(keys[1:], "."))
	return &value
}

// GetStringMapString gets a map[string]string config value by specific key
func (c *Configuration) GetStringMapString(key string, defaultValue ...map[string]string) *map[string]string {
	var value map[string]string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetStringMapString(strings.Join(keys[1:], "."))
	return &value
}

// GetStringMapStringSlice gets a map[string][]string config value by specific key
func (c *Configuration) GetStringMapStringSlice(key string, defaultValue ...map[string][]string) *map[string][]string {
	var value map[string][]string
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	if !c.Has(key) {
		return &value
	}
	keys := strings.Split(key, ".")
	value = c.configs.Get(keys[0]).GetStringMapStringSlice(strings.Join(keys[1:], "."))
	return &value
}
