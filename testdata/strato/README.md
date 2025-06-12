# Solver testdata directory

To provide your credentials for the test suite modify the `testdata/strato/config.json`.

### `config.json`

```json
{
  "username": "<username>",
  "password": "<password>",
  "totpSecret": "<secret>",
  "totpDeviceName": "<deviceName>",
  "waitingTime": <seconds>
}
```

| Property         | Type     | Optional | Default | Comment    |
|------------------|----------|----------|---------|------------|
| `apiUrl`         | `string` | YES      |         |            |
| `username`       | `string` | NO       |         |            |
| `password`       | `string` | NO       |         |            |
| `totpSecret`     | `string` | YES      |         |            |
| `totpDeviceName` | `string` | YES      |         |            |
| `waitingTime`    | `int`    | YES      | `0`     | In seconds |

See also [strato-certbot README.md#setup](https://github.com/Buxdehuda/strato-certbot/blob/master/README.md#setup).

> ⚠️ **Cautition:** The property names have changed. 
