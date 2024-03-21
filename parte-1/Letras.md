# Letras

Vamos a implementar un programa que cuenta la frecuencia de las letras. Pero si le pasamos ciertos parámetros podemos contar sólo las vocales, las consonantes o números.

En este ejemplo vamos a mostrar varias prácticas comunes usadas en el lenguaje. Aprenderemos a usar arreglos y mapas (diccionarios), y veremos otras funciones disponibles en la biblioteca estándar.

## La interfaz de nuestro programa

Nuestro programa usará la linea de comandos (cli). Si llamamos `letras` a este programa, entonces la forma de usar nuestro programa será:


        letras [-c] [-d] [-v] archivos...


Por ejemplo, para ver la frecuencia de cada letra, junto con la cantidad de vocales de este archivo que estás leyendo, tendríamos que ejecutar este comando:

        letras -v Letras.md

## Procesando los parámetros de nuestro programa

Los parámetros a un programa en Go se reciben desde el entorno en el sistema operativo. Estos están disponibles a través de la variable `Args` que se encuentra en el paquete `os`. La variable Args es un arreglo de strings[^1] con los parámetros que se le entregan al programa. El parámetro inicial `os.Args[0]` contiene el nombre del programa (`letras` en nuestro caso), como no nos interesa ese argumento, podemos obtener los parámetros descartando el primero mediante la expresisón `os.Args[1:]`. Esa expresión retorna lo que se conoce como un `slice`, que corresponde a los elementos del arreglo a partir de la posición 1 (como en muchos lenguajes, en Go los arreglos se indexan desde 0).

Con esto, podemos entender la primera parte de nuestra función `main`:

```go
func main() {
	args := os.Args[1:]
	var vowels, consonants, digits bool
	var files []string
	for i, arg := range args {
		switch arg {
		case "-v":
			vowels = true
		case "-c":
			consonants = true
		case "-d":
			digits = true
		default:
			if strings.HasPrefix(arg, "-v") {
				log.Fatal("argumento %s inválido", args[i])
			} else {
				files = append(files, arg)
			}
		}
	}
```

La variable `args` tendrá la lista de argumentos.
Lo otro que llama la atención en este código es la linea que empieza así:

```go
for i, arg := range args {
```

Lo que hace `range` es crear una suerte de iterador. Si su argumento (en este caso el slice `args`) es un arreglo, lo que hace es que retorna el índice y el valor de cada elemento del slice (o de un arreglo) en cada iteración. Así si invocamos el programa del siguiente modo:

        letras -v -c Letras.md

Entonces obtendremos estos valores en cada iteración:

i | arg
--|----
0 | -v
1 | -c
2 | Letras.md
 
El `switch` debería ser fácil de entender, pues funciona de forma similar a lo que hay en otros lenguajes. A diferencia de C no es necesario colocar la sentencia `break`.

El `if` que aparece en la rama `default` del `switch` arroja un error y termina la ejecución del programa si se pasa un argumento que no soportamos. Esto se hace usando la función `log.Fatal`. En realidad hay mejores formas de manejar esto, pero quise incluir esta función para poder incluir más funciones de la biblioteca estándar. En los ejercicios te desafiaré para que modifiques este código para manejar este error de una forma más "elegante".

Esta línea es interesante:

```go
				files = append(files, arg)
```

Si se dan cuenta la variable `files` fue declarada así:

```go
var files []string
```

Eso declara un slice vacío (el valor de files es `nil`). Lo que hace `append` es crear un nuevo slice que contendrá el valor de `arg` al final. Esto es funcionalmente idéntico a agregar un elemento al final del arreglo, pero internamente se maneja usando slices y hay que tener esto claro, pues es fuente de algunos errores, como veremos más adelante.

La lógica para procesar los archivos se encuentra en la función `processFiles` que comienza con este código:

```go
func processFiles(files []string, vowels, consonants, digits bool) {
	var countVowels, countConsonants, countDigits int

  letters := make(map[rune]int)
```

Recibimos la lista de archivos en la variable `files`, y recibimos los parámetros que nos indican si debemos mostrar los contadores de ciertos tipos de caracteres.

Luego declaramos las variables que usaremos para contar los distintos tipos de caracteres y luego viene la declaración de la variable `letters`, esta corresponde a lo que se conoce en Go como un map (mapa), que es el equivalente a los diccionarios o hashmap de otros lenguajes. El operador `make` se usa en Go para crear mapas, arreglos y channels (que veremos más adelante). 

Luego tenemos un loop para leer cada archivo que hemos recibido en la variable `files`

```go
	for _, file := range files {
		buf, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
```

Fíjate como usamos range para iterar sobre todos los archivos en el loop, de modo que en cada iteración la variable `file` contiene el nombre del archivo que vamos a procesar.

Usamos la función `os.ReadFile()` para leer todo el archivo. Esta función retorna un arreglo de bytes con el contenido del archivo y un error. Esta es la forma de manejar errores en Go. Más adelante lo veremos en detalle, pero por ahora basta con decir que si err es `nil` no hubo error, y al contrario, si algo falló el detalle del problema lo tenemos en la variable `err`. En nuestro caso la forma de manejar el error es imprimir el mensaje del error y continuar con el loop.

A continuación, si no hay error procesamos el contenido del archivo:

```go
		runes := string(buf)
		for _, r := range runes {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				continue
			}
			letters[r]++
			c := unicode.ToLower(r)
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				countDigits++
			case 'a', 'e', 'i', 'o', 'u':
				countVowels++
			default:
				countConsonants++
			}
		}
	}
```

Como `buf` es un arreglo de bytes, lo transformamos en un string y lo asignamos a la variable `runes`. En Go los strings se asume que representan el contenido en formato  unicode. Así que en vez de caracteres, cada elemento de un string es conocido como `rune`. por eso nuestro mapa `letters` tiene como llave el tipo `rune`, que es la representación de un caracter en unicode.

Fíjate el uso de `range runes` para obtener cada caracter del string `runes`.

Filtramos los caracteres que no sean letras o dígitos usando el package `unicode` y luego con la expresión `letters[r]++` incrementamos el contador para el caracter contenido en `r`. Esta es una de las características más agradables de Go, y es que el manejo de mapas es muy sencillo y práctico (a mi me recuerda a como lo hace AWK).

A estas alturas el `switch` debería ser fácil de entender, y nos sirve para llevar la cuenta de los tipos de caracteres que hemos seleccionado.


Finalmente imprimimos los resultados:

```go
	fmt.Println("Letra|Cantidad")

	for key, val := range letters {
		fmt.Printf("%c | %d\n", key, val)
	}
	if vowels {
		fmt.Println("Vocales:", countVowels)
	}
	if consonants {
		fmt.Println("Consonantes:", countConsonants)
	}
	if digits {
		fmt.Println("Dígitos:", countDigits)
	}
```

De este código lo más interesante es el `for`. Fíjate que al ser `letters` un map, este ciclo inicia la variable `key` con el valor de la llave del map, en este caso la rune, y en val el valor del contador de cada letra.

Finalmente los `if` que vienen se encargan de imprimir los contadores que nos han pedido.

Y como puedes ver con este ejemplo hemos aprendido como se manipulan arreglos de tamaño variable y el manejo básico de los mapas en Go.

## Ejercicios

1. Prueba cambiar la llamada a `log.Fatal()` por un `fmt.Println()` seguido de `return` y ve que pasa.
2. Cuando el usuario no ingrese ningún archivo deberíamos mostrar un mensaje indicando que no hay nada para procesar, implementa este cambio.
3. Muestra la frecuencia de letras expresada como un porcentaje.
4. Muestra la frecuencia ordenada alfabeticamente.
5. Muestra la frecuencia de los valores más altos a los valores más bajos.


[Volver al índice](../README.md)

[^1]: Técnicamente `os.Args` es un slice, los arreglos tienen un tamaño fijo.
