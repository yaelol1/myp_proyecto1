# myp_proyecto1

Proyecto de chat, cliente y servidor para la clase de modelado y programación.

## Lenguaje

Go
---
El proyecto está hecho en Go, principalmente porque es un lenguaje con orientación a objetos, rápido y con
una comunidad bastante grande, que al momento de buscar bibliotecas para por ejemplo Json que se ocupará
en este proyecto; será fácil encontrarlas.

### Ventajas

- Sintaxis fácil de leer
- Tipos implícitos de variables
- Interfaces y Objetos
- Gorutines
- Godoc para generar la documentación
- Go test para las pruebas unitarias
- Go build para compilar

### Desventajas
- Bindings para interfaces gráficas no son comunes

## Objetos (Estructuras en go)

### Paquete cliente

- Cliente
  - Nombre
  - Conexión
  - Cuartos

### Paquete Servidor

- Cuarto
  - Nombre
  - Integrantes
- Servidor
  - Cuartos

## Compilar
Para usar el cliente, el servidor tiene que estar corriendo.

### Servidor
```
$ cd servidor
$ go build .
$ ./servidor
```

### Cliente

```
$ cd cliente
$ go build .
$ ./cliente
```
