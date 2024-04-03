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
    ducks := []Duck{
        NewBlackDuck("Daffy"),
        NewBlackDuck("Donald"),
    }
    LakeSimulation(ducks)
}
```

## Cisnes

Los cisnes se parecen a los patos, así que podemos aprovechar este hecho para agregar cisnes a nuestra simulación.

Agregamos el archivo `swans.go` a nuestro package.


```go
type Swan struct {
    id int
}

func (swan *Swan) Fly() {
    fmt.Println("Swan", swan.id, "is flying")
}

func (swan *Swan) Swim() {
    fmt.Println("Swan", swan.id, "is swimming")
}

```

El '"constructor' lo vamos a declarar así:

```go
func NewSwan() Duck {
    return &Swan{
        id: rand.IntN(1000)
    }
}
```

Nuestros cisnes son bastante anónimos, así que su identidad será un número aleatorio.

Ahora podemos modificar nuestra simulación agregando un par de cisnes:

```go
func main() {
    ducks := []Duck{
        NewBlackDuck("Daffy"),
        NewBlackDuck("Donald"),
        NewSwan(),
        NewSwan(),
    }
    LakeSimulation(ducks)
}
```


Nota que las funciones que crean las estructuras tienen como tipo de retorno `Duck`. Incluso la función `NewSwan` se ve extraña retornando el tipo `Swan`, pero esto se debe a que si queremos tener este grado de polimorfismo, nuestros "constructores" tienen que retornar una interfaz.

Otra cosa que se debe notar es que nuestras funciones constructoras retornan referencias a las estructuras. Te dejo como ejercicio averiguar que pasa si retornamos la estructura en vez de una referencia a la misma.

Todo el ejercicio se encuentra en la carpeta [ducks](/parte-2/ducks)
