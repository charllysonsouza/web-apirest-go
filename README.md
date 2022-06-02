<h1 align="center">Bank Account API</h1>  

### Aplicação criada com intuito de simular operações básicas a se realizar com uma conta bancária, como: depósito, saque e transferência.

As sequintes restrições foram estabelecidas ao realizar as operações:

- Depósito: 1% sobre o valor depositado 
- Saque: R$ 4,00 sobre o valor do saque 
- Transferência: R$ 1,00 

## Ferramentas utilizadas
- `gin-gonic`
- `gorm-io`
- `postgresSQL`
- `pgAdmin4`
          
## Banco de dados
A seguinte modelagem de dados foi realizada para guardar as informações que necessitamos persistir.
<p align="center">
   <img src="https://user-images.githubusercontent.com/78328120/171493575-cb6d0e6b-c03a-42af-be43-655f8229d21b.png">  
</p>

Para esta aplicação foi utilizado o `gorm.io` como ferramenta de ORM para facilitar a criação das tabelas. É utilizado o banco `PostgreSQL` para persistêcida dos dados.

## Iniciando a aplicação

Precisamos iniciar o docker com as imagens do `pgAdmin` e `postgress`. Portanto, após clonar o repositório, na raiz do projeto execute o comando

```shell
docker-compose up
```
Em um outro terminal (ainda na pasta raiz), basta executar o seguinte comando para iniciar a execução da aplicação
```shell
go run main.go
```

## Endpoints disponíveis
- `/create-client`
- `/create-account/{client_id}`
- `/deposit`
- `/withdraw`
- `/transfer`
- `/extract`

## Cadastrando um cliente
Para cadastrar um cliente basta enviar uma *requisição* `POST` para o endpoint `/create-client` informando os seguintes dados
```json
{
	"Name": string, 
	"Document": string,
	"Email": string,    
	"Telephone": string
}
```
O endpoint deve retornar uma *response* com os seguintes dados
```json
{
	"id": number,
	"name": string,
	"document": string,
	"type": string,
	"email": string,
	"telephone": string
}
```
## Cadastrando um conta
Para gerar uma conta para um cliente específico basta realizarmos uma requisição `POST` passando o id do cliente na url da *requisição*. Nesse caso, uma conta será gerada para o cliente que corresponde ao id=2
```url
/create-account/2
```
O endpoint deve retornar uma *response* com os dados da conta
```json
{
	"accountNumber": string,
	"agencyNumber": string,
	"balance": number
}
```
## Realizando um depósito
Para realizar um depósito basta enviar uma *requisição* `POST` para o endpoint `/deposit` informando os seguintes dados
```json
{
	"NumberAccount": string,
	"Amount": number
}
```
O endpoint deve retornar uma *response* com a seguinte mensagem
```json
{
	"message": "deposit succeed!"
}
```
## Realizando um saque
Para realizar um saque basta enviar uma *requisição* `POST` para o endpoint `/withdraw` informando os seguintes dados
```json
{
	"NumberAccount": string,
	"Amount": number
}
```
O endpoint deve retornar uma *response* com a seguinte mensagem
```json
{
	"message": "withdraw succeed!"
}
```
## Realizando uma transferência
Para realizar uma transferência basta enviar uma *requisição* `POST` para o endpoint `/transfer` informando os seguintes dados
```json
{
	"NumberAccountOrigin": string,
	"NumberAccountTarget": string,
	"Amount": number
}
```
O endpoint deve retornar uma *response* com a seguinte mensagem
```json
{
	"message": "transfer succeed!"
}
```
## Emitindo um extrato
Para emitir um extrato com todas as movimentações basta enviar uma *requisição* `GET` para o endpoint `/extract` informando o seguinte dado
```json
{
	"accountNumber": string
}
```
O endpoint deve retornar uma *response* com os seguintes dados 
```json
{
	"accountNumber": string,
	"agencyNumber": string,
	"balance": number,
	"transactions": [
		{
			"transactionsAsOrigin": [
				{
					"transactionId": string,
					"value": number,
					"type": string,
					"date": string
				},
				{
					"transactionId": string,
					"value": number,
					"type": string,
					"date": string
				}
			],
			"transactionsAsDestination": [
				{
					"transactionId": string,
					"value": number,
					"type": string,
					"date": string
				},
				{
					"transactionId": string,
					"value": number,
					"type": string,
					"date": string
				}
			]
		}
	]
}
```

PS: As responses acima descritas correspondem ao cenário feliz, ou seja, que tudo deu certo. Caso contrário, erros de validação serão lançados pela a aplicação.
