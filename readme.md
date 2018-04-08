# configtool
Simple command-line utility to read settings from configuration files

```
$ configtool
configtool reads configuration from multiple sources like parameters, files and env variables.

Usage:
  configtool [command]

Available Commands:
  get         Reads the value of the given setting from the detected sources
  help        Help about any command
  list        Lists the available configuration settings

Flags:
  -a, --app string      app name used to derive configuration paths
  -c, --config string   config name used to derive the configuration file name, defaults to 'config' (default "config")
  -h, --help            help for configtool

Use "configtool [command] --help" for more information about a command.
```

# Usage examples

## Read configuration settings from default locations

```
configtool get setting1
```

Tries to read the value for `setting1` from the following locations:
- `config.yaml`
- `~/config.yaml`
- `/etc/config.yaml`

## Read app specific configuration settings
Tries to read the value for `setting1` from the following locations:

```
configtool -a myapp get setting1
```

- `./.myapp/config.yaml`
- `~/.myapp/config.yaml`
- `/etc/myapp/config.yaml`

## Read app specific configuration settings

```
configtool -a myapp -c myconfig get setting1
```

Tries to read the value for `setting1` from the following locations:
- `./.myapp/myconfig.yaml`
- `~/.myapp/myconfig.yaml`
- `/etc/myapp/myconfig.yaml`


## Derive config location from environment variables 
```
CONFIGTOOL_APP=myapp \
CONFIGTOOL_CONFIG=custom-config \
configtool get setting1
```

## Use default values for settings
```
configtool -a myapp get setting2 -d default-value
```