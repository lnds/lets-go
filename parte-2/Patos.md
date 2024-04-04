# Si camina como pato...

Seguro has escuchado del concepto Duck Typing. Es una característica de los lenguajes dinámicos, en que los tipos se determinan en tiempo de ejecución.

Por ejemplo, considera este código en Python:

```python
class Duck:
    def swim(self):
        print("Duck swimming")

    def fly(self):
        print("Duck flying")

class Whale:
    def swim(self):
        print("Whale swimming")

for animal in [Duck(), Whale()]:
    animal.swim()
    animal.fly()
```

El método `swim` al estar definido en ambas clases, es invocado por el intérprete en tiempo de ejecución y se ejecutará sin problemas. El caso del método `fly` sólo está disponible en la clase `Duck`, así que eso provocará una falla.

En los lenguajes de tipos estático, como Java, esto no pasa, porque esto se verifica en tiempo de compilación.

Go implementa algo distinto. La definición de qué es un pato se puede declarar a priori, mediante una interfaz, de este modo:

```go
type Duck interface {
    Fly()
    Swim()
}
```

Nota que una `interface` se declara como un tipo. Entonces podemos usarla en una función del siguiente modo:

```go
func LakeSimulation(ducks []Duck) {
    for _, duck := range ducks {
        duck.Fly()
        duck.Swim()
    }
}
```

Acá `LakeSimulation` es una función que nos permite implementar la simulación de un lago en que hay patos. Por ahora es bastante sencilla, simplemente invocamos los métodos `Fly()` y  `Swim()` para cada elemento en el arreglo que definimos.

Ahora podemos crear una struct que implemente los métodos de la interfaz de este modo:

```go 
type BlackDuck struct {
    name string
}

func (duck *BlackDuck) Fly() {
    fmt.Printf("%s duck is flying\n", duck.name)
}

func (duck *BlackDuck) Swim() {
    fmt.Printf("%s duck is swimming\n", duck.name)
}

func NewBlackDuck(name string) Duck {
    return &BlackDuck{
        name: name,
    }
}
```


Y podemos probar esto del siguiente modo:

```go
func main() {
    ducks := []ducks.Duck{
        ducks.NewBlackDuck("Daffy"),
        ducks.NewBlackDuck("Donald"),
    }
    LakeSimulation(ducks)
}
```

## Cisnes

Los cisnes se parecen a los patos, así que podemos aprovechar este hecho para agregar cisnes a nuestra simulación.

Agregamos el archivo `swans.go` a nuestro package.


```go
type Swan int 

func (swan Swan) Fly() {
    fmt.Println("Swan", swan, "is flying")
}

func (swan Swan) Swim() {
    fmt.Println("Swan", swan, "is swimming")
}

```

Fíjate que el tipo `Swan` es simplemente un `int`. Ojo, ya no es un tipo entero, es un nuevo tipo que hereda todo el comportamiento de un `int`, pero que puede tener métodos, como `Fly()` y `Swan()`.

Nuestros cisnes son bastante anónimos, así que su identidad será el valor numérico que se le de al declararlo.

Ahora podemos modificar nuestra simulación agregando un par de cisnes:

```go
func main() {
    ducks := []Duck{
        ducks.NewBlackDuck("Daffy"),
        ducks.NewBlackDuck("Donald"),
        ducks.Swan(100),
        ducks.Swan(42),
    }
    LakeSimulation(ducks)
}
```

Esto es muy interesante, podemos asociar interfaces a cualquier tipo, basta con agregar los métodos que la componen en ese nuevo tipo.

## Ejercicios

1. Crea un nuevo package llamado `swans` y mueve el archivo `swans.go` a este. ¿Qué otras modificaciones debes hacer al código para que siga funcionando la simulación?
2. Modifica la simulación para que el 25% del tiempo los patos naden y el 35% del tiempo vuelen.
3. Agrega el tipo `Goose` que implemente el comportamiento de un Ganso. Una diferencia es que el ganso cada vez que emprende el vuelo grazna, implementa el método `Quack` que se encarga de graznar y sólo debe estar disponible en la struct `Goose`.

[Volver al índice](../README.md)
