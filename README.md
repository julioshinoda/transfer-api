# Transfer-api

## Como utilizar

>Para rodar é preciso ter o docker e o docker-compose instalado

Abra o terminal e na raiz do projeto use o comando abaixo

`docker-compose up`

Em seguida, em outro terminal rode o comando abaixo

`make migration`

Agora que a API está rodando e o banco ja esta com as tabelas criadas pela migration, pode ser usado qualquer dos enpoints abaixo

### obtém a lista de contas

`GET http://localhost:9011/accounts
Content-Type: application/json`

### obtém o saldo da conta

`GET http://localhost:9011/accounts/{account_id}/balance`

### cria um Account

`POST http://localhost:9011/accounts
Content-Type: application/json
{
  "name":"user-2",
  "cpf":"12312312318",
  "ballance":10000
}`

### obtém a lista de transferencias

`GET http://localhost:9011/transfers
Content-Type: application/json`

### faz transferencia de um Account para outro

`POST http://localhost:9011/transfers
Content-Type: application/json

{
  "account_origin_id":1,
  "account_destination_id":3,
  "amount":5000
}`
