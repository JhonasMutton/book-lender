# book-lender

Book-lender é uma aplicação rest feita na linguagem Go com o intuito de gerenciar empréstimos de livros.

##Rodando a aplicação
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

###Build
Para fazer o build da aplicação, basta executar o comando na pasta raíz:
```
go build .
```

###Run
Para rodar o Book-lender, basta executar o comando na pasta raíz:
```
go run .
```
*Obs: A aplicação rodará por padrão na porta 8080*

###Environment Variables
O Book-lender dispõe das seguintes variáveis de ambiente para sua configuração;

```SERVER_PORT=8080```