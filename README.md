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

##Comandos go

como inicializar um programa com go module

```shell script
go mod init github.com/mmfg/classes
``` 

##Comandos linux

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


