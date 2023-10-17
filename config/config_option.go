package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Option is used to add options to [viper.Viper]
type Option func(v *viper.Viper) error

// Type see [viper.Viper.SetConfigType]
func Type(configType string) Option {
	return func(v *viper.Viper) error {
		v.SetConfigType(configType)
		return nil
	}
}

// Path see [viper.Viper.Path]
func Path(configPath ...string) Option {
	return func(v *viper.Viper) error {
		for _, cp := range configPath {
			v.AddConfigPath(cp)
		}
		return nil
	}
}

// AutomaticEnv see [viper.Viper.AutomaticEnv]
func AutomaticEnv(configPath ...string) Option {
	return func(v *viper.Viper) error {
		v.AutomaticEnv()
		return nil
	}
}

// BindEnv see [viper.Viper.BindEnv]
func BindEnv(key ...string) Option {
	return func(v *viper.Viper) error {
		return v.BindEnv(key...)
	}
}

// BindFlagValue see [viper.Viper.BindFlagValue]
func BindFlagValue(key string, flag viper.FlagValue) Option {
	return func(v *viper.Viper) error {
		return v.BindFlagValue(key, flag)
	}
}

// BindFlagValues see [viper.Viper.FlagValueSet]
func BindFlagValues(flags viper.FlagValueSet) Option {
	return func(v *viper.Viper) error {
		return v.BindFlagValues(flags)
	}
}

// BindPFlag see [viper.Viper.BindPFlag]
func BindPFlag(key string, flag *pflag.Flag) Option {
	return func(v *viper.Viper) error {
		return v.BindPFlag(key, flag)
	}
}

// BindPFlags see [viper.Viper.BindPFlags]
func BindPFlags(flags *pflag.FlagSet) Option {
	return func(v *viper.Viper) error {
		return v.BindPFlags(flags)
	}
}

// RemoteProvider see [viper.Viper.AddRemoteProvider]
func RemoteProvider(provider, endpoint, path string) Option {
	return func(v *viper.Viper) error {
		return v.AddRemoteProvider(provider, endpoint, path)
	}
}

// SecureRemoteProvider see [viper.Viper.AddSecureRemoteProvider]
func SecureRemoteProvider(provider, endpoint, path, secretkeyring string) Option {
	return func(v *viper.Viper) error {
		return v.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
	}
}

// SetValue see [viper.Viper.Set]
func SetValue(key string, value any) Option {
	return func(v *viper.Viper) error {
		v.Set(key, value)
		return nil
	}
}

// SetValues see [viper.Viper.Set]
func SetValues(values map[string]any) Option {
	return func(v *viper.Viper) error {
		for key, value := range values {
			v.Set(key, value)
		}
		return nil
	}
}

// DefaultValue see [viper.Viper.SetDefault]
func DefaultValue(key string, value any) Option {
	return func(v *viper.Viper) error {
		v.SetDefault(key, value)
		return nil
	}
}

// DefaultValues see [viper.Viper.SetDefault]
func DefaultValues(values map[string]any) Option {
	return func(v *viper.Viper) error {
		for key, value := range values {
			v.SetDefault(key, value)
		}
		return nil
	}
}

// Watch see [viper.Viper.WatchConfig]
func Watch() Option {
	return func(v *viper.Viper) error {
		v.WatchConfig()
		return nil
	}
}

// WatchRemoteConfig see [viper.Viper.WatchRemoteConfig]
func WatchRemoteConfig() Option {
	return func(v *viper.Viper) error {
		return v.WatchRemoteConfig()
	}
}

// WatchRemoteConfigOnChannel see [viper.Viper.WatchRemoteConfigOnChannel]
func WatchRemoteConfigOnChannel() Option {
	return func(v *viper.Viper) error {
		return v.WatchRemoteConfigOnChannel()
	}
}
