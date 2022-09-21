package option

import (
    "flag"
    "fmt"
    "os"
)

// ConfigOption Stores the value for the configuration
type ConfigOption struct {
    Value   *string
    Default string
    EnvVar  string
    Flag    string
    Usage   string
}

// Config Represent all the configuration options there are for the application
type Config struct {
    // Port The port for the application
    Port ConfigOption

    // Host The host address for the application
    Host ConfigOption

    // Mode The running mode for the application
    Mode ConfigOption

    // KeyDir The path to the keys for JWT auth
    ReleaseDir ConfigOption
}

const (
    // DebugMode A config mode allowing for the most verbose logging. Used heavily for development.
    DebugMode = "debug"

    // ReleaseMode A config mode allowing for minimal logging. Intended to be used for production.
    ReleaseMode = "release"
)

var (
    DefaultConfig = Config{
        Port: ConfigOption{
            Default: "8180",
            EnvVar:  "SRO_UPDATER_PORT",
            Flag:    "port",
            Usage:   "The port for the application",
        },
        Host: ConfigOption{
            Default: "",
            EnvVar:  "SRO_UPDATER_HOST",
            Flag:    "host",
            Usage:   "The host address for the application",
        },
        Mode: ConfigOption{
            Default: DebugMode,
            EnvVar:  "SRO_UPDATER_MODE",
            Flag:    "mode",
            Usage:   "The running mode for the application",
        },
        ReleaseDir: ConfigOption{
            Default: "/etc/sro/releases",
            EnvVar:  "SRO_RELEASES_DIR",
            Flag:    "keys",
            Usage:   "The path to the releases directory for patches and base versions",
        },
    }
)

func NewConfig() Config {
    config := DefaultConfig
    config.readFlags()
    config.readEnvs()
    return config
}

// Address Gets the full address for the HTTP server
func (c *Config) Address() string {
    return fmt.Sprintf("%s:%s", c.Host.GetValue(), c.Port.GetValue())
}

// GetValue Gets the value for the Config option if it's set, otherwise it gets the default value
func (co *ConfigOption) GetValue() string {
    if co.Value == nil {
        return co.Default
    }
    return *co.Value
}

// IsRelease Returns true if the application mode is release
func (c *Config) IsRelease() bool {
    return c.Mode.GetValue() == ReleaseMode
}

func (c *Config) readFlags() {
    c.Port.readFlag()
    c.Host.readFlag()
    c.Mode.readFlag()
    c.ReleaseDir.readFlag()
    flag.Parse()
}

func (co *ConfigOption) readFlag() {
    co.Value = flag.String(co.Flag, co.Default, co.Usage)
}

func (c *Config) readEnvs() {
    c.Port.readEnv()
    c.Host.readEnv()
    c.Mode.readEnv()
    c.ReleaseDir.readEnv()
}

func (co *ConfigOption) readEnv() {
    env, found := os.LookupEnv(co.EnvVar)
    if found && !isFlagPassed(co.Flag) {
        co.Value = &env
    }
}

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}
