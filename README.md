**Concurrencia**: Ejecuta varias cosas al mismo tiempo, de forma desordenada, bajo un mismo hilo.
Por ejemplo:

**Servicio A**: Ejecuta el 10%

**Servicio B**: Ejecuta el 25%

**Servicio A**: Ejecuta el 15%

**Servicio B**: Ejecuta el 5%

Así sucesivamente hasta que algún programa termine.


**Paralelismo**: Ejecuta varias funcionalidades al mismo tiempo, bajo diferentes hilos o núcleos. Estas
ejecuciones se hacen de forma independiente y no están relacionadas entre si. Vease worker threads de nodejs. Escalamiento horizontal de la aplicación. Esto va a depender de las capacidades del servidor, de la cantidad de núcleos que este posea.