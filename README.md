# mda

Para ejecutar este proyecto ejecute:
1.  Clonar el repositorio 
```$ git clone https://github.com/udistrital/mda```
2. Moverse al reposotorio 
```$ cd mda```
3. Docker compose de mda
```$ docker-compose up```
(```$ docker-compose up &``` si quiere seguir utilizando la misma terminal)

4. Cargar el fuente del binario **boostrap** en la terminal 
```$ source bootstrap```


Con esto tendrá disponible el generador el cual podrá ejecutar como:

```$ oasgen init  <proyecto>```
```$ oasgen generate <<entidad>>```

## Cómo crear un nuevo proyecto:
* Ejecutar ```$ oasgen init <<nombre_proyecto>>```
* Moverse al proyecto recién creado ```$ cd <nombre_proyecto>```
* Probar el proyecto generado ```$ docker-compose up```

## Cómo agregar entidades al proyecto:
* Moverse a la carpeta entities del proyecto ```$ cd <nombre_proyecto>/entities```
* Ejecutar ```$ oassgen generate <entidad>``` (por defecto para pruebas ```oasgen generate example.ent```)
* Esto generará una carpeta con el nombre **srcgen** donde se encuentra el código fuente listo para su aplicación tanto
backend como frontend

## Errores conocidos:
*Error:* `ERROR: for beego  Cannot create container for service beego: Conflict. The container name "/beego" is already in use by container "CONTAINER_ID". You have to remove (or rename) that container to be able to reuse that name.`

*Solución:* ```$ docker rm -f CONTAINER_ID```
