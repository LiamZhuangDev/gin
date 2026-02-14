Viper flow:

```
define keys  // Define or register keys, tell Viper these configuration extries exist
    |        // You can do it via SetDefault/config file/BindEnv
    |        // After registration, Viper knows these keys is part of the config tree
    +
load sources // Tell Viper where values may come from
    |        // e.g. viper.AutomaticEnv()  
    |
    +
values override // Viper handles priority/overrides automatically
    |           // e.g. ENV > config file > default
    |
    +
Unmarshal    // This copies the final values into Go fields
```