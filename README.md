# book-lender

Book-lender é uma aplicação Rest feita na linguagem Go com o intuito de gerenciar empréstimos de livros, para fins de demonstração de conhecimento e aplicação de conceitos.

## Rodando a aplicação:
## Via Docker
Para executar o book-lender via docker, basta executar os seguintes comandos:

Constrói a aplicação:
```
make build
```
Gera a imagem docker:
```
make image
```
Por fim, roda:
```
sh cmd/run-docker.sh
```
Depois é só ser feliz!
## Localmente
### wire
o Book-lender utiliza wire para injetar dependencias em tempo de compilação. Para utilizado-liga os passos a seguir:

**Primeira vez**

Baixe o wire:

```
go get github.com/google/wire/cmd/wire
```

**Sempre:**
Rode o seguinte comando na pasta raíz do projeto para que a injeção de dependencia seja feita:
```
wire
```

### Build
Para fazer o build da aplicação, basta executar o comando na pasta raíz:
```
go build .
```

### Run
Para rodar o Book-lender, basta executar o comando na pasta raíz:
```
go run .
```
*Obs: A aplicação rodará por padrão na porta 8080*

### Endpoints


| API                         | Path                             | Método  | 
| --------------------------  | -------------------------------- | ------- | 
| Cria usuário                | /user                            | POST    |
| Busca usuário pelo ID       | /user{id}                        | GET     | 
| Busca usuário todos usuários| /user{id}                        | GET     |
| Adiciona um livro ao usuário| /book                            | POST    |
| Empresta o livro            | /book/lend                       | PUT     |
| Devolve o livro             | /book/return                     | PUT     |

### Environment Variables
O Book-lender dispõe das seguintes variáveis de ambiente para sua configuração;

```SERVER_PORT=8080```\
```DB_HOST=localhost```\
```DB_PORT=3306```\
```DB_USER=root```\
```DB_PASSWORD=admin```\
```DB_NAME=BOOK_LENDER```\
```LOG_LEVEL=debug```


### Banco de dados
A aplicação se conecta com um banco de dados MySql. Para fins de desenvolvimento, é criado uma instancia do banco localmente, via docker. Basta executar:
```shell
sh cmd/database/create-mysql.sh
```
Para parar e subir o container basta executar:
```shell
sh cmd/database/stop-mysql.sh
```
```shell
sh cmd/database/start-mysql.sh
```
Ou utilizar o CLI do docker puramente.

#### Migração
Por utilizar um banco relacional, as tabelas precisam ser criadas e configuradas, para isso, rodamos a migração na primeira vez que rodado cada versão.
Para executar a migração basta executar o seguinte comando:
```shell
sh cmd/database/migration/migrate.sh
```
PS. A migração rodará com os valores definidos no arquivo .env
