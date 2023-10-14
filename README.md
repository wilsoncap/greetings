#Saludos en Go

Este paquete proporciona una forma simple de pptener saludos personalizados en Go


##Isntalacion
Ejecute el siguiente comando para instalar el paquete:

```bash
go get -u github.com/wilsoncap/greetings
```

## Uso

Aqui tienes un ejemplo de como utilizar el apquete en tu codigo

```go
package main

import (
	"fmt"
	"github.com/wilsoncap/greetings"
)

func main() {
	messages, err := greetings.Hello("Wilson")

	if err != nil{
		fmt.Println("Ocurrio un error: ", err)
	}

	fmt.Println(message)
}
```