# Proyecto Mars

## Contexto

La NASA envía un archivo al satélite "Mars" en donde se define una nueva ruta exploratoria a cada uno de
los robots que se encuentran en el planeta marte. Para enviar la nueva ruta exploratoria a los robots, el satélite debe
conectarse con cada uno de ellos de manera secuencial y enviarle las instrucciones, de manera que si un satélite llega
a salirse de la zona de movimiento, este punto fronterizo sea mapeado y sirva de información al siguiente satélite.

## Flujo

El componente inicial "handler" está encargado de procesar y validar el archivo de entrada.
Luego el componente "service" contiene la lógica necesaria para procesar la información recibida desde el "handler" y enviarla a cada cliente.
Finalmente, el componente "cliente" que será el robot, recibe la instrucción del componente "service" y se encarga de ejecutar la acción solicitada.

## Caso de uso

Se deje un archivo de texto "file.txt" el cual contiene instrucciones para siete robots. Por alguna razon, 2 de los 7
robots contienen información errónea; uno tiene coordenadas de posición inválida (4 argumentos en lugar de 3) mientras el
otro tiene más de 100 instrucciones (el máximo permitido es 100 instrucciones).

El resultado esperado para el primer robot es perderse en la coordenada (50,51), ya que siempre va para adelante y sale
de la zona de movimiento.

El segundo robot aprenderá sobre el punto de perdida del primer robot, pero luego de girar a la derecha y avanzar se
perderá en la coordenada (51,50).

El tercer robot habrá aprendido de los dos primeros robots y terminará su recorrido en el punto de partida, habiendo
cambiado solo su orientación al sur.

El cuarto y quinto robot habrán completado exitosamente su ruta exploratoria.

## Trazabilidad

Para comprender un poco el flujo de información se deja instrumentado en el proyecto el sistema de trazabilidad distribuido
JAEGER.

Para poder visualizar el jaeger UI es necesario correr el siguiente comando docker

```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.32
```

Luego, ya se podrá acceder a la dirección http://localhost:16686

Fuente: https://www.jaegertracing.io/docs/1.32/getting-started/#all-in-one