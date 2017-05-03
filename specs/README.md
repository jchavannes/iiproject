# eId Specs

- Requests and responses are JSON objects
- By default, requests and responses are PGP encrypted and signed
  - Notable exception is _id_, which returns the public key

## Example Request
```json
{
    "date": "2017-05-02T08:40:07.968Z"
}
```
