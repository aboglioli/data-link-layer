# Data Link simulation

Este proyecto intenta simular el funcionamiento de la **Capa de Enlace** en
redes. Se realizó para la cátedra **Redes de Información** de FRM-UTN.

El proyecto puede encontrarse en
[github.com/aboglioli/data-link-layer](https://github.com/aboglioli/data-link-layer)

Está desarrollado en [Go](https://golang.org). Se eligió este lenguaje por su
facilidad para manejar múltiples hilos, ya que *multi-threading* es en lo que
está enfocado.

El principal funcionamiento de la capa de red, es decir, sus servicios, son:

- Detección y/o corrección de errores
- Control de flujo

Hay una implementación básica de un receptor y emisor *utópicos simples*. No
hay detección de errores, sin embargo, si se controla el flujo.

Leer comentarios del código para entender el funcionamiento básico.

## Uso

Ejecutar receptor primero:

```
go run cmd/utopian_simplex_receiver.go
```

Luego, ejecutar emisor:

```
go run cmd/utopian_simplex_sender.go
```

## Información, referencias, ejemplos e implementaciones

- [Go encoding bytes](https://medium.com/learning-the-go-programming-language/encoding-data-with-the-go-binary-package-42c7c0eb3e73)
- https://github.com/stefano-lupo/Data-Link-Protocol-Simulation
- https://github.com/topics/data-link-layer
- https://github.com/polltery/Data-layer-protocol-example-in-Java
- https://github.com/jonaias/data-link-byte-framing
- https://github.com/harryawk/flow-control-simulation
- https://github.com/azmainadel/data-link-layer-implementation
- https://github.com/sharanyakamath/Flow-Control-in-Data-Link-Layer
- https://github.com/briancain/Layers-Protocol
- https://stackoverflow.com/questions/14667615/how-do-i-port-this-code-to-java
- http://www.dsi.unive.it/~franz/reti/dll/Protocolli.html