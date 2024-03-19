# Múltiplos

El problema que vamos a resolver a continuación está tomado de acá: https://projecteuler.net/problem=1 y dice:

> Si listamos los números naturales bajo 10 que son múltiplos de 3 ó 5, tenemos 3,5,6 y 9. La suma de estos múltiplos es 23.
>
> Encuentra la suma de todos los múltiplos de 3 ó 5 bajo 1000 (mil).

# Una primera solución

La primera solución que implementaremos está en el archivo [solu1.go](/parte-1/multiplos/solu1.go), vamos a revisarlo linea a linea.

Primero declaramos nuestro package como `main`, porque queremos crear un ejecutable, y vamos a incluir el package `fmt` de la biblioteca estándar, para poder imprimir nuestros resultados en la consola:

```go
package main

include "fmt"
```

Tenemos el caso para los múltiplos de 3 ó 5 menores de 10, y nos piden el caso para los múltiplos menores a 1000. Así que vamos a declarar dos constantes con esos valores que las usaremos para llamar a la función que implementa nuestro algoritmo.


```go
const test = 10
const challenge = 1000

func main() {
    fmt.Println("Version 1")
    fmt.Println("Múltiplos", test, sumMults(test))
    fmt.Println("Múltiplos", challenge, sumMults(challenge))
}
```

La solución la implementamos en la función `sumMults`:

```go
func sumMults(n int) int {
    var sum int = 0
    var i int = 1
    for i < n {
        if i % 3 == 0 || i % 5 == 0 {
            sum = sum + i
        }
        i = i + 1
    }
    return sum
}
```

Nuestra función tiene un parámetro `n` de tipo int, y retorna un valor del mismo tipo. Más adelante veremos otras formas de declarar funciones, pero noten que el tipo se coloca después, al contrario de lo que pasa en los otros lenguajes derivados de C, en que el tipo va primero.

Luego en el cuerpo de la función nota cómo declaramos las variables que vamos a usar: `sum` e `i`. La forma para declarar es con la sentencia `var` seguida de la variable, luego el tipo y un valor inicial después del signo `=`.

En otros lenguajes el ciclo lo implementaríamo usando la sentencia `while`. En Go existe sólo una palabra clave para implementar ciclos, `for`. Después del for va la condición que se evalúa en cada iteración, cuando esta condición ya no se cumpla entonces el ciclo termina. Si omites la condición, se entiende que estás implementado un ciclo infinito (cuando escribes algo así: `for { ... }`). Después del for va la condición que se evalúa en cada iteración, cuando esta condición ya no se cumpla entonces el ciclo termina. Si omites la condición, se entiende que estás implementado un ciclo infinito (cuando escribes algo así: `for { ... }`).

El `if` no debería sorprender, es idéntico a lo que escribiríamos en C o Java, por ejemplo. Nota la diferencia con estos lenguajes, los paréntesis que son requeridos en estos, acá se omiten.


Otra cosa importante, las llaves `{` y `}` son obligatorias y la llave de apertura ('{') debe ir en la misma línea que las sentencias `if` o `for` (al igual que al inicio de una función).

Así que en Go esta sentencia produce un error:

```go 
if condicion 
{
    ...
}
else 
{
    ....
}
```

Otra cosa que hay que notar es que no agregamos punto y coma '`;`' al final de cada linea, en Go son opcionales. De hecho el compilador de Go los agrega por nosotros al final de cada línea siguiendo las reglas descritas acá: https://go.dev/doc/effective_go#semicolons.

Puedes probar esta solución ejecutando:

    go run solu1.go

## Una solución abreviada

Go permite declarar de forma más breve las variables y usar inferencia de tipos. Esto lo demostramos en la versión 2 de nuestra solución, disponible en el archivo [solu2.go](/parte-1/multiplos/solu2.go).

```go
// sumMults calcula la suma de los múltiplos de 3 y 5 menores a n
func sumMults(n int) int {
  sum := 0 
  for i := 1; i < n; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  return sum
}
```

Al igual que en C++ y otros lenguajes, los comentarios epiezan con `//`.

Fíjate que al hacer `sum := 0` declaramos una variable de tipo `int`, esto lo infiere el compilador a partir de la expresión y el contexto (como la función retorna `int` y sum es el valor retornado, debe ser de tipo `int`).

Noten la otra forma del `for`. Esto es igual que un for en un lenguaje derivado de C, con la diferencia que no van los paréntesis.

Con respecto al operador `++`, este sólo puede usar como sufijo, la expresión `++i` es inválida en Go. Nota que también existe el operador `+=`.

Prueba esta solución usando este comando:

        go run solu2.go


## FizzBuzz

Hay un problema parecido a este, que dice así:

> Imprime los números desde 0 a 100. Pero cuando sea múltiplo de 3 agrega la palabra Fizz. Si es mútiplo de 5 entonces agrega la palabra Buzz. Y si es múltiplo de ambos el texto FizzBuzz.

Intenta resolver este problema con lo que has visto hasta ahora. Si no, puedes encontrar la solución en el archivo [fizzbuzz.go](/parte-1/multiplos/fizzbuzz.go).

Fíjate que en vez de usar `if` he usado la sentencia `switch`. Esta es una de las formas de usarla, esto nos permite evitar muchos `if` `else if` que a veces son difíciles de leer.

```go
  for i := 1; i < 100; i++ {
    switch {
    case i % 3 == 0 && i % 5 == 0:
      fmt.Println(i, "FizzBuzz")
    case i % 3 == 0:
      fmt.Println(i, "Fizz")
    case i % 5 == 0:
      fmt.Println(i, "Buzz")
    default:
      fmt.Println(i)
    }
  }
```

Prueba esta solución ejecutando:

        go run fizzbuzz.go

## Ejercicios

1. ¿Qué pasa si ejecutas `go build` en la carpeta [multiplos](/parte-1/multiplos)?¿Cómo lo harías para compilar el ejecutable `fizzbuzz`?


2. Un problema con la solución presentada para múltiplos, es que es muy ineficiente. Si nos pidieran encontrar la suma para n igual a mil millones este programa va a tomar mucho tiempo iterando. ¿Se te ocurre un medio de implementar este algoritmo de un modo más eficiente?
