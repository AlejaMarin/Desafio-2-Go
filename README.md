# Desafío Final Go

## Especialización en BackEnd III
## Certified Tech Developer
## Digital House

---

#### [Enunciado](https://drive.google.com/file/d/1EnCe3Y_7OTaSxv_aDEqArK5fQhUIybEQ/view?usp=sharing)

## Integrantes
> Alejandra Marin  
> Vanesa Vilte

---

## Base de Datos
La base de datos fue realizada en MySQL, compuesta por tres tablas: Paciente, Dentista y Turnos. 
El modelo EER se planteó de la siguiente manera:
> Turno: Posee una relación de muchos a uno con Dentista.
>
> Turno: Posee una relación de muchos a uno con Paciente.

!["Modelo EER"](/EER.png)

Las tablas Paciente y Dentista tienen una columna llamada Activo, que cambia de estado ante una request de tipo Delete, lo cual evita eliminar información de la base de datos. El objetivo de esta implementación es manejar el error de la base de datos al querer eliminar datos que estan sujetos a la constraint de las foreing key (IdPaciente, IdDentista).
En el caso de la tabla Turnos, se puede eliminar un turno de forma directa.

El script que crea la base de datos con algunos datos de prueba, esta en:  
[Script DB](/database.sql)

## Observación
Al realizar una request de tipo POST en los endpoint de Paciente, Dentista y Turno, omitir el campo "id". Ya que este valor es generado automáticamente en la base de datos, ya que su propiedad es ser un valor auto incremental.

## Swagger
La API fue documentada en Swagger. Para ingresar a la interfaz de la misma, la ruta es la siguiente:

    http://localhost:8080/swagger/index.html

## Postman
En Postman se realizó las pruebas a los diferentes endpoints de la API.  
[Colección](/Sistema%20de%20Reserva%20de%20Turnos.postman_collection.json)

---

### 2023