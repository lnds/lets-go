# Hola Mundo

Es un ejemplo clásico, introducido por Khernighan y Ritchie en su famoso libro sobre el lenguaje C. La idea es escribir un programa que imprima en la consola la frase "Hola Mundo".

El siguiente video muestra como lo implementamos en el programa que puedes encontrar en el subdirectorio (o carpeta) [hola](/parte-1/hola).

![](hola.gif)

Primero creamos la carpeta [hola](/parte-1/hola), instalamos Go usando  dentro de esta el módulo del mismo nombre. Esto se hace usando `asdf`. Luego lo más interesante es este comando:

        go mod init hola

Esto inicializa nuestro proyecto creando un módulo llamado `hola`. Con esto se crea el archivo `go.mod`, que contiene toda la descripción de nuestro proyecto, por ahora lo más destacado es que declara la versión de Go que estamos usando.

Luego creamos nuestro programa editando el archivo `main.go`. De este lo más destacado es que declaramos en la primera linea el nombre del paquete.
 

 ```go
 package main
 ```

En Go el módulo es la unidad que se usa para distribuir el software, ya sea en la forma de aplicaciones, o de bibliotecas. Un módulo puede contener uno o más paquetes (`package`). Y si se trata de una aplicación debe incluir el package `main`. Además el punto de entrada de nuestra aplicación es una función (que se declaran usando la  sentencia `func`) también llamada `main`.

Para poder escribir el mensaje en la consola usaremos la función `Println()` que pertenece a la bibloteca estándar de Go y se encuentra en el package `fmt`. Es por eso que importamos ese package con la sentencia:

```go
import "fmt"
```

Fíjate que el nombre del paquete se escribe entre comillas, porque es un string. Esto es un tanto diferente a otros lenguajes de programación, más adelante hablaremos en detalle de módulos y paquetes y en ese momento explicaremos todo esto.
 

Finalmente nuestro programa queda así:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola Mundo")
}
```

Para ejecutarlo simplemente haces:

        go run main.go


Y puedes compilarlo, generando un ejecutable, con esta instrucción:

        go build

Que crea un binario que puedes ejecutar directamente:

        ./hola


Todo esto lo puedes ver en el video.

Ahora es tu turno, intenta reproducir todo esto en tu entorno. Puedes usar otro editor de texto, yo uso Neovim, pero puedes usar el que más te acomode, una buena recomendación es Visual Studio Code. 


[Volver al índice](../README.md)
