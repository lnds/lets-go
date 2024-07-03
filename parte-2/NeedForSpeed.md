# Need For Speed

Vamos a aprender a usar structs en Go resolviendo dos ejercicios de la plataforma [Exercism](https://exercism.org/).


## Resolviendo "Need For Speed"

El primer ejercicio se llama "Need For Speed" y lo encuentras en esta dirección: https://exercism.org/tracks/go/exercises/need-for-speed. Te sugiero que lo leas, porque además tiene una introducción al concepto de estructuras en Go.


El problema nos pide crear un par de estructuras de datos para representar autos y pistas. Esto se puede usar por ejemplo en un simulador de carreras.
 
Para esto vamos a crear un módulo que llamaremos `speed` como siempre:

    mkdir speed
    go mod init speed

En el package `speed` vamos a declarar nuestra estructura `Car`:


```go
package speed

// Car implements a simulated remote controller car.
type Car struct {
    speed        int
    batteryDrain int
    battery      int
    distance     int
}
```

Tal como nos dicen en el ejercicio es una buena práctica crear funciones que nos permitan crear nuestras estructuras. Normalmente estas empiezan con el prefijo `New`. Vamos a declarar una para crear nuestros carros:

```go
// NewCar creates a new Car with full battery and given specs
func NewCar(speed, batteryDrain int) Car {
    return Car{
        speed:        speed,
        batteryDrain: batteryDrain,
        battery:      100,
    }
}
```


Fíjate cómo definimos el valor de una estructura usando el nombre de cada atributo seguido de dos puntos ':' y el valor.

Junto con resolver este ejercicio vamos a aprender a crear tests en Go.

Para esto vamos a crear el archivo `speed_test.go` con el siguiente contenido:

```go
// Archivo speed_test.go
package speed

import "testing"

func TestNewCar(t *testing.T) {

    tests := []struct{
                name string
                car Car
                expected Car
        }{
                {
                        name: "Create a new car.",
                        car: Car{
                                speed: 5,
                                batteryDrain: 2,
                                battery: 100,
                        },
                        expected: Car{
                                speed: 5,
                                batteryDrain: 2,
                                battery: 100,
                        },

                },
        }

        for _,tt :=range tests{
                t.Run(tt.name, func(t *testing.T){
                        got := NewCar(tt.car.speed, tt.car.batteryDrain)

                        if got!=tt.expected{
                                t.Errorf("NewCar(%v,%v) = %v; expected %v", tt.car.speed, tt.car.batteryDrain,got, tt.expected)
                        }
                })
        }
}

```

Este código muestra una técnica muy usada para escribir tests en Go.

Usualmente declaramos un arreglo con estructuras declaradas inline, por ejemplo:

```go
tests := []struct{
    name string
    car Car
    expected Car
}
```


En este caso no es necesario declarar el nombre de la estructura porque no será usado.

Usamos el package `testing` que está en la biblioteca estándar. Lo que hace go es que al ejecutar el comando `go test` busca todos los archivos con sufijo `_test`, y dentro de estos las funciones que empiezan con `Test`. A cada una de estas funciones les entrega como parámetro un contexto de testing que es de tipo `*testint.T`. El `*` significa que nos pasa una referencia a esa estructura, pero esto lo vamos a ver en un momento. Notarás que en el for se hace una llamada a `t.Run`, ese es un método de la estructura `testing.T`, concepto que también vamos a ver más abajo.

Lo importante por ahora es que ejecutamos el test haciendo:

        go test

Y obtenemos:

        PASS
        ok      speed   0.448s


Vamos a implementar la estructura que representa la pista de carreras, `Track`:

```go
// Track implements a race track.
type Track struct {
        distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
        return Track{
                distance: distance,
        }
}
```


Y el problema nos pide crear una función llamada `Drive` que implemente el movimiento del carro:

```go
// Drive drives the car one time. If there is not enough battry to drive on more time,
// the car will not move but use the leftover battery.
func Drive(car Car) Car {
        if car.battery >= car.batteryDrain {
                car.distance += car.speed
                car.battery -= car.batteryDrain
        }
        return car
}
```


Nota como nuestra función toma el auto como parámetro y aplica las transformaciones a sus valores internos, retornando el valor recibido como argumento, pero modificado.


La otra función que nos piden implementar es `CanFinish` que nos debe retornar verdadero si el auto es capaz de finalizar el recorrido de la pista dada.


```go
// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
        for car.battery > 0 {
                car = Drive(car)
        }
        return car.distance >= track.distance
}
```


No voy a incluir el test completo, pero se encuentra en la carpeta [speed](/parte-2/speed/).

Estudia las funciones `TestDrive` y en especial `TestCanFinish`.

## Ejercicios

1. ¿Qué problema le vea a las funciones que hemos definido?
2. Implementa una función llamada `Charge()` que restaure la batería al 100. Luego implementa una función que se llame `Finish()` que reciba un `Car` y un `Track` y que retorne el valor de batería que quedó disponible en el carro después de recorrer la pista entera (usa la función `Charge` para implementar esto).


## Elon's Toys

Nuestra forma de modelar el auto es un tanto limitada. Nos gustaría usar un enfoque más "orientado al objeto", en que el carro tenga métodos con el que podamos modificar su estado interno.

Go nos permite hacer eso. Y ese es el objetivo del problema "Elon's Toys" que se encuentra en esta dirección: https://exercism.org/tracks/go/exercises/elons-toys


Vamos a crear nuestro módulo del siguiente modo:


        mkdir elons-toys
        cd elons-toys
        go mod init elon


Una primera implementación de `Car` sería:

```go
package elon

// Car implements a remote controlled Car
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}


// NewCar creates a new car with given specs
func NewCar(speed, batteryDrain int) *Car {
  return &Car{
    speed: speed,
    batteryDrain: batteryDrain,
    battery: 100,
  }
}


func (car *Car) Drive() {
  if car.battery >= car.batteryDrain {
    car.distance += car.speed
    car.battery -= car.batteryDrain
  }
}

```

Fíjate cómo ahora `NewCar` retorna una referencia a `Car`. La notación `*Car` indica que el tipo es una referencia a `Car`. La notación `&Car` nos permite obtener la referencia.

Si has usado otros lenguajes, como C++ o Java estarás esperando un operador como `new`. Este existe también en Go pero la expresión usada en el return es una forma más corta de crear una nueva instancia de Car y retornar su referencia, y es la forma preferida en Go.

Por último, puedes ver como hemos implementado el método `Drive()`. En Go el método se declara usando `func` pero anteponiendo al nombre del método una referencia a la estructura, como en este caso `func (car *Car)`.

Como `Drive()` altera el estado interno de `car` no es necesario retornarlo como lo hicimos en el caso anterior.

## Ejercicios

1. Completa el ejercicio "Elon's Toys" tal como esta descrito en Exercism. Copia los tests disponibles en ese sitio.
2. Agrega los métodos `Charge()` y `Finish()` que se comortan tal como los describimos en la sección anterior.

[Volver al índice](../README.md)
