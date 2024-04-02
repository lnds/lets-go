# Lo público y lo privado

Partiremos con un ejemplo sencillo. Consideremos el problema de transformar entre distintos tipos de unidades.

Para empezar a explorar esto primero vamos a crear un módulo que llamaremos `units`:

        mkdir units
        cd units
        go mod init units

Vamos a crear el archivo `main.go`:

```go
package main

import (
	"fmt"
)

func toCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

func toFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}

func main() {
	celsius := toCelsius(95)
	fahrenheit := toFahrenheit(35)

	fmt.Println(35, "grados celsius corresponden a", fahrenheit, "grados fahrenheit")
	fmt.Println(95, "grados fahrenheit corresponden a", celsius, "grados celsius")
}
```

Como nuestro módulo puede crecer con otras funciones para transforma entre distintos tipos de unidades, nos conviene "empaquetar" estas dos funciones que hemos creado en un paquete que llamaremos "temperature".

Para esto vamos a crear la carpeta `temperature` con un archivo llamado `temperature.go`:

        mkdir temperature
        cd temperature
        touch temperature.go

Vamos a copiar las funciones `toCelsius` y `toFahrenheit`
en este archivo, nota que definimos el nombre usando la cláusula `package`:


```go
package temperature


func toCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

func toFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}
```


Y modificaremos el archivo main.go del siguiente modo:

```go

package main

import (
	"fmt"
    "units/temperature"
)


func main() {
	celsius := toCelsius(95)
	fahrenheit := toFahrenheit(35)

	fmt.Println(35, "grados celsius corresponden a", fahrenheit, "grados fahrenheit")
	fmt.Println(95, "grados fahrenheit corresponden a", celsius, "grados celsius")
}
```

Para incluir un paquete debemos indicar el nombre del módulo seguido del paquete. Como nuestro módulo se llama `units` debemos incluir `units/temperature`.

Sin embargo al ejecutar `go run main.go` obtenemos un error que dice:

        go run main.go
        # command-line-arguments
        ./main.go:6:2: "units/temperature" imported and not used


La razón es que las funciones `toCelsius` y `toFahrenheit` son privadas. En Go toda función cuyo identificador comienza con minúscula es privada y sólo visible dentro del paquete. Para volverla pública el identificador de la función debe comenzar con mayúsculas. Corrijamos esto.
Primero el archivo `temperature/temperature.go`: 


```go
package temperature 

func ToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0 
}

func toFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}
```

Luego el archivo main.go:

```go
package main

import (
	"fmt"

	"units/temperature"
)

func main() {
	celsius := temperature.ToCelsius(95)
	fahrenheit := temperature.ToFahrenheit(35)

	fmt.Println(35, "grados celsius corresponden a", fahrenheit, "grados fahrenheit")
	fmt.Println(95, "grados fahrenheit corresponden a", celsius, "grados celsius")
}
```

A veces el nombre del package puede ser muy largo. En esos casos se puede crear un alias de este modo:


```go
package main

import (
	"fmt"

    tmp	"units/temperature"
)

func main() {
	celsius := tmp.ToCelsius(95)
	fahrenheit := tmp.ToFahrenheit(35)

	fmt.Println(35, "grados celsius corresponden a", fahrenheit, "grados fahrenheit")
	fmt.Println(95, "grados fahrenheit corresponden a", celsius, "grados celsius")
}
``` 

## Go Lint

Si bien nuestro paquete está correcto, se recomienda documentar los packages.

Hay un utilitario que permite verificar esta situación. Para instalarlo hacemos:

        go install golang.org/x/lint/golint@latest

Una vez instalado hacemos:

        golint temperature

Obteniendo:

        temperature/temperature.go:3:1: exported function ToCelsius should have comment or be unexported
        temperature/temperature.go:7:1: exported function ToFahrenheit should have comment or be unexported


Esto se corrige agregando comentarios adecuados:

```go
package temperature

// ToCelsius convert the given temperature (in celsius) to fahrenheit
func ToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}
 
// ToFahrenheit convert the given temperature (in fahrenheit) to celsius
func ToFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}
```


## Ejercicio

Ahora tu turno. 

1. Agrega la función `ToKelvin`. ¿Qué problema presenta agregar esta función? ¿Cómo tendríamos que modificar las interfaces de nuestras funciones para adaptarla a esta nueva situación?
2. Agrega otro paquete para transformar unidades de peso (libras, kilos, onzas, etcétera). Llámalo `weights` y crea las respectivas funciones.
