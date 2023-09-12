package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Option func(v *viper.Viper) error

func ConfigType(configType string) Option {
	return func(v *viper.Viper) error {
		v.SetConfigType(configType)
		return nil
	}
}

func ConfigPath(configPath ...string) Option {
	return func(v *viper.Viper) error {
		for _, cp := range configPath {
			v.AddConfigPath(cp)
		}
		return nil
	}
}

func AutomaticEnv(configPath ...string) Option {
	return func(v *viper.Viper) error {
		v.AutomaticEnv()
		return nil
	}
}

func BindEnv(key ...string) Option {
	return func(v *viper.Viper) error {
		return v.BindEnv(key...)
	}
}

func BindFlagValue(key string, flag viper.FlagValue) Option {
	return func(v *viper.Viper) error {
		return v.BindFlagValue(key, flag)
	}
}

func BindFlagValues(flags viper.FlagValueSet) Option {
	return func(v *viper.Viper) error {
		return v.BindFlagValues(flags)
	}
}

func BindPFlag(key string, flag *pflag.Flag) Option {
	return func(v *viper.Viper) error {
		return v.BindPFlag(key, flag)
	}
}

func BindPFlags(flags *pflag.FlagSet) Option {
	return func(v *viper.Viper) error {
		return v.BindPFlags(flags)
	}
}

func RemoteProvider(provider, endpoint, path string) Option {
	return func(v *viper.Viper) error {
		return v.AddRemoteProvider(provider, endpoint, path)
	}
}

func SecureRemoteProvider(provider, endpoint, path, secretkeyring string) Option {
	return func(v *viper.Viper) error {
		return v.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
	}
}

func SetValue(key string, value any) Option {
	return func(v *viper.Viper) error {
		v.Set(key, value)
		return nil
	}
}

func SetValues(values map[string]any) Option {
	return func(v *viper.Viper) error {
		for key, value := range values {
			v.Set(key, value)
		}
		return nil
	}
}

func DefaultValue(key string, value any) Option {
	return func(v *viper.Viper) error {
		v.SetDefault(key, value)
		return nil
	}
}

func DefaultValues(values map[string]any) Option {
	return func(v *viper.Viper) error {
		for key, value := range values {
			v.SetDefault(key, value)
		}
		return nil
	}
}

func Watch() Option {
	return func(v *viper.Viper) error {
		v.WatchConfig()
		return nil
	}
}

func WatchRemoteConfig() Option {
	return func(v *viper.Viper) error {
		return v.WatchRemoteConfig()
	}
}

func WatchRemoteConfigOnChannel() Option {
	return func(v *viper.Viper) error {
		return v.WatchRemoteConfigOnChannel()
	}
}
