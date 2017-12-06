# Config Reader
This activity reads a JSON Configuration file. The configuration can then be used in following activities.
This activity can read String, Int, Float and Boolean configuration values.
Default value can be set
Configuration can be cached (and not read at each execution) for better performances

## Installation

```bash
flogo add activity github.com/philippegabert/configreader/flogo-contrib/activity/configreader
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "configFile",
      "type": "string",
      "required": "true"
    },
    {
      "name": "configName",
      "type": "string",
      "required": "true"
    },
    {
      "name": "configType",
      "type": "string",
      "allowed" : ["string", "int", "float", "bool"]
    },
    {
      "name": "defaultValue", 
      "type": "any"
    },
    {
      "name": "readEachTime",
      "type": "bool"
    }
  ],
  "outputs": [
  	{
      "name": "configValue",
      "type": "any"
    }
  ]
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| configFile        | The path to the configuration file |         
| configName        | The name of the configuration element to retrieve |
| configType        | The type of the configuration element (string, int, float or bool) |
| defaultValue        | The default value to set if the configuration was not read correctly |
| readEachTime        | If set to true, the configuration file will be read at each execution of the flow. If set to false, the file will be read only once |
| configValue        | The value of the configuration |

## Configuration Examples

![String Config](string_config.png| width=100)
- Configuration will be read from file **/home/pi/config.json**
- Name of the configuration element is **string_config**
- Type of the configuration element is **string**
- Default value to apply if the conf. element is no| width=100)t found: **Default**
- Setting the "readEachTime" value to **false**: File will be read only once