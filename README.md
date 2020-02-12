#Meus estudos em golang

# Índice
1. [Basico](#como-criar-um-programa-minimo-em-go)
2. [Lógica de Programação](#example2)
3. [Testes unitários](#third-example)
4. [Comandos go](#comandos-go)
5. [Comandos linux](#comandos-linux)
6. [Comandos Git](#fourth-examplehttpwwwfourthexamplecom)

## como criar um programa minimo em go?

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world!")
}
```

## Comandos go

como inicializar um programa com go module

```shell script
go mod init github.com/mlfg/golernen
``` 

## Comandos linux

listar arquivos 

```shell script
ls -lah
```

encontrar todos arquivos go

```shell script
find . -type f -inname "*.go"
```

como criar uma pasta

```shell script
mkdir "dirname"
```

como copiar um arquivo 

```shell script
cp main.go dirname/subpasta/
```

como mover um arquivo

```shell script
mv main.go dirname/subpasta/
```

## Comandos Git

Como saber que arquivo foi alterado ou adicionado a minha pasta?

comando: ***git status***

```shell script
git status
```

como ver a modificação de um arquivo?
comando:***git diff*** ou ***gitk***
````shell script
git diff <nome do arquvo>
ex.:
git diff README.md
````

Como adicionar o arquivo para ser salvo?

comando: ***git add***

```shell script
git add <caminho_do_arquivo>
ex.:
git add aula1/divide.go
```

Como salvar o arquivo adicionado com o git add?

comando: ***git commit -s -m "mensagem do commit"***

```shell script
git commit -s -m "adiciona funcao para calcular raiz quadrada"
```

como enviar as mudanças salvas para os servidores do git-hub?

```shell script
git push origin master
```
## Funções e arquivos de teste

como criarum arquivo de teste?
```shell script
o nome da "bagagem" e test.go no final

dica: para cada teste fazer uma nova função
```

