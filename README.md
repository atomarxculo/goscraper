# JKAnime API No oficial

Esta aplicación pretende ser una API la cual, al pasarle el nombre de un anime, busque todos los capitulos disponibles y la opción de descargartelos, ahorrandote tener que ver miles de anuncios hasta poder ver un capitulo.

**ToDo:**

- [x] Al buscar un anime aparece dos veces la URL raíz con una doble barra al final (Ej: "https://jkanime.net/anime/" y "https://jkanime.net/anime//")
- [ ] Frontend en React.
- [ ] Poder descargar los capitulos directamente del visualizador de video, no de los enlaces que tenga en la sección "Descargas". Esto es debido a problemas con Zippyshare en España.
- [ ] Que te muestre todos los enlaces para descargar el anime, en vez de tener que buscar capitulo por capitulo.
- [ ] Mejorar tiempos de busqueda. Integrar con Zippkin para medir tiempos.
- [ ] Testing con animes que tengan muchos capitulos. Las pruebas realizadas son con animes que tienen 24 capitulos máximo.
- [ ] Sugerencias en vez de busquedas absolutas.
- [ ] Manejo de errores.
- [ ] Genera un fichero de log, además de mostrar más información por consola. 
- [ ] Integración con JDownloader o similar para que mande las URLs de descarga automaticamente al gestor y descargue los capitulos.
- [ ] Documentación API con Swagger.
- [ ] Integración con Redis/Elasticsearch para guardar datos en caché.
- [ ] Continuará...
