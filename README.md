# Resume API

A simple GraphQL API to expose information relating to my professional experience.

## Running

Run the go code

```
go run ./*.go
```

## API

The GraphQL API runs at `/graphql` and accepts `POST` requests.

```sh
curl --request POST \
  --url 'http://localhost:8080/graphql?query=%7B__schema%7Btypes%7Bkind%20name%20description%7D%7D' \
  --header 'content-type: application/json' \
  --data '{"query":"{\n\tme {\n\t\tfirstName\n\t\tmiddleName\n\t\tlastName\n\t\twebsite\n\t}\n}","variables":{}}'
```

