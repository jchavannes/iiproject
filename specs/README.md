# eId Specs

- Requests and responses are JSON objects
- By default, requests and responses are PGP encrypted and signed
  - Notable exception is `/id`, which returns the public key

## Examples

**Path**: `/id`

### Request
```json
{
    "name": "/get"
}
```

### Response
```json
{
    "public_key": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\n..."
}
```

**Path**: `/profile`

### Request
```json
{
    "name": "/get",
    "eid": "dev2:8252/u/cli"
}
```

### Response
```json
{
    "body": "-----Message-----\n\n..."
}
```
