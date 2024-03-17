# Instala Go en tu máquina

Instalar Go en tu máquina es bastante sencillo. En el sitio oficial de Go: [https://go.dev/](https://go.dev/) hay instaladores, simplemente debes descargarlos y ejecutarlos.

Pero hay una manera mucho más práctica, que es usando `asdf`. Esta herramienta nos permite gestionar múltiples versiones de los lenguajes operativos y es por eso que recomiendo usarla.

Para instalar `asdf` necesitas tener instalado `git` y `curl`.
Voy a explicar como instalar esto en ambientes tipo Unix (como Linux o Mac). Para Windows recomiendo usar el instalador, o instalar usando Chocolatey, tal como describo más adelante (pero pierdes la habilidad de usar múltiples versiones del lenguaje).

## Soporte para múltiples versiones con ASDF

Como los lenguajes evolucionan a veces es bueno poder usar ASDF, que es una herramienta que nos permite mantener diversas versiones con ambientes virtuales.

Para instalar `asdf` debes ejecutar este comando:

    git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.

Luego asegurarte de que `asdf` esté accesible en el path. Si usas Bash agrega esta línen el archivo `.bashrc`:

    . "$HOME/.asdf/asdf.sh"

Si tienes Mac puedes ahorrarte todo la anterior usando brew:

    brew install asdf

Si tienes `asdf` instalado debes agregar el plugin para golang del siguiente modo:

        asdf plugin add golang

Luego puedes instalar la última versión de Go del siguiente modo:

        asdf install golang latest  

En los tutoriales que ejecutaremos usaremos Go 1.22.1, así que si quieres asegurarte de usar esa versión puedes hacer

        asdf install golang 1.22.1

Cuando quieras usar go, en el directorio de trabajo de tu proyecto debes hacer:

        asdf local golang 1.22.1

A continuación muestro otros modos de instalar Go en Windows y Mac.

## Instalando en Windows con Chocolatey

[Chocolatey](https://chocolatey.org/) es un gestor de paquetes para Windows bastante popular. Para instalar Go con este gestor ejecutas el siguiente comando en el terminal (en este caso estamos usando PowerShell)

        choco install golang -y

Tienes que abrir una nueva terminal para poder aplicar los cambios, y así podrás probar que todo esté instalado con el siguiente comando:

        go version

![](choco-install.gif)

## Instalando en Mac con Brew

Debes tener instalado [brew](https://brew.sh/). Simplemente ejecutas:

    brew install go
    go version


![](brew-install.gif)


[Volver al índice](../README.md)
