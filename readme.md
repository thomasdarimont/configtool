# Usage examples

## Derive config location from Parameters 
```
configtool -a configtool -c configtool-home get setting1
```

## Derive config location from Environment Variables 
```
CONFIGTOOL_APP=configtool \
CONFIGTOOL_CONFIG=configtool-home \
configtool get setting1
```

## Use default values
```
configtool -a configtool -c configtool-home get setting2 -d default-value
```