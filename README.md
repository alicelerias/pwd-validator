# Validador de senha
API graphql que valida senhas

## Getting Started
1. Instalar dependencias: `go mod tidy`
2. Executar unit tests: `go test ./...`
3. Iniciar graphql server: `go run server.go`
4. Gerar codigo para novo schema: `go run github.com/99designs/gqlgen generate`


## Exemplo

Realize a consulta na rota `POST localhost:8080/graphql`.

Selecione "graphql" e passe na query o seguinte código:

```graphql
query {
  verify(
    password: "TesteSenhaForte!123&", 
    rules: [
      { name: "minSize",value: 8 }, 
      { name: "minSpecialChar", value: 2 },
      { name: "minLowercase",value: 2 },
      { name: "minUppercase",value: 2 },
      { name: "minDigit",value: 2 },
      { name: "noRepeted",value: 0 }
    ]
  ) {
  valid
  noMatch
  }
}
```

Como pode perceber, alterei alguns nomes, "rule" passou a se chamar "name", e o "verify" anterior ao "noMatch", passou a se chamar "valid". A ideia é que o servidor faça uma verificação e retorne se a senha é válida ou não. Caso não seja válida, noMatch será um array com os nomes das regras que não passaram. 

O objeto retornado dessa consulta será:

```json
{
    "data": {
        "verify": {
            "valid": true,
            "noMatch": []
        }
    }
}
```