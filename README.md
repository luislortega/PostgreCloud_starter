# 🙋🏽‍♂️ Hey there! 
Este es un repositorio basico y rapido desde lo tecnico hasta lo teorico, de como subir una base de datos a la nube... usaremos codigo en GoLang para hacer los test de conexion pero sirve para todo, espero les sirva

# 🐱‍🏍¿Como instalar todo lo necesario?  
### GoLang
1. [Descargar GoLang](https://golang.org/)
2. Crear un *workspace* de GoLang (solo desde ahi podras compilarlo)
3. Crear una variable de enterno **GoPath** que tenga como direccion tu *workspace* 
4. Escribe el comando `$ go version` en la terminal para verificar que toda la instalacion haya sido correcta

Para otras personas que usan **Linux**:
[Como instalar GoLang desde Linux](https://www.youtube.com/watch?v=0e_C1B8fDvg)

### PostgreSQL
Tambien demasiado sencillo en su [sitio oficial](https://www.postgresql.org/download/) hay un tutorial demasiado completo para instalarlo en diferentes sistemas operativos

### Heroku
Esto es demasiado sencillo solo necesitas [crear tu cuenta de heroku](https://dashboard.heroku.com/apps) crear un proyecto nuevo y poner que tenga el add-ons de PostgreSQL que rapidamente nos dara nuestras llaves de acceso a ella

# 🐱‍👤Hora de construir nuestra primera db en cloud!

1. ir a tu workspace/src
2. crear una carpeta para tu proyecto nuevo
3. crear main.go adentro de tu carpeta
4. ejecutar `$ go get -u github.com/lib/pq`
5. ir a `C:\Program Files\PostgreSQL\11\bin` desde la CLI
6. ejecutar`$ psql -U postgres -h localhost` (la contraseña por default es null entonces dale enter)
7. Ponemos las credenciales en nuestro ejemplo para conectarse a la database

``` 
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = ""
	port = 5432  //default port for PostgreSQL...
	user = ""
	password = ""
	dbname = ""
)

func  main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
```

9. Habilita el SSL
10. Insert data

``` 
package main

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const (
	host = ""
	port = 5432  //default port for PostgreSQL...
	user = ""
	password = ""
	dbname = ""
)

func  main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec("CREATE TABLE IF NOT EXISTS example (id integer, data varchar(32))")
	if err != nil {
		panic(err)
	}
	insForm, err := db.Prepare("INSERT INTO example (id, data) VALUES(1,'test')")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec()
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
``` 
11. Select data from Heroku server
``` 
package main

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const (
	host = ""
	port = 5432
	user = ""
	password = ""
	dbname = ""
)

type  Example  struct {
	id int
	data string
}

func  main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
				host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec("CREATE TABLE IF NOT EXISTS example (id integer, data varchar(32))")
	if err != nil {
	panic(err)
	}
	insForm, err := db.Prepare("INSERT INTO example (id, data) VALUES(1,'test')")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec()
	var  myExample Example
	query := "SELECT id, data FROM example WHERE id = $1"
	err = db.QueryRow(query, 1).Scan(&myExample.id, &myExample.data)
	if err != nil {
		fmt.Printf("Failed to execute query: ", err)
	}
	fmt.Printf("usuario:{ id: %d, data: %s} \n", myExample.id, myExample.data)
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
``` 

# 🤓 Resolucion de existenciales....

**¿Por que escoger Heroku como mi proveedor cloud?**
* Permite al desarrollador centrarse en el código en lugar de en la infraestructura.
* Mejore la productividad del equipo de desarrollo de aplicaciones en la nube
* Ofrece facturación única para todos los proyectos desglosados ​​por equipo
* Supervise y mejore el rendimiento a través de una rica supervisión de aplicaciones
* Ayuda a los interesados ​​en el desarrollo, el control de calidad y el negocio a crear un panel unificado.
* Soporte de lenguajes modernos de código abierto
* Integración de Salesforce de alto rendimiento
* Escalabilidad simple
* El equipo de operaciones y seguridad de Heroku está listo al instante para ayudarlo las 24 horas del día, los 7 días de la semana
* Ecosistema de servicios y herramientas de plataforma líder
* Le ayuda a concentrarse en la innovación, no en las operaciones
* La arquitectura de Heroku Enterprise ofrece un tiempo de inactividad mínimo o nulo durante las actualizaciones del sistema.
* Gestión y permisos rápidos del ciclo de vida de las aplicaciones
* Le permite eliminar la fricción del desarrollo.
* Ofrece un potente panel de control y CLI
* Se integra con flujos de trabajo de desarrolladores familiares
* Previsibilidad y conocimiento del costo de desarrollo y mantenimiento de aplicaciones
* Un montón de herramientas de apoyo
* Para principiantes y para startups
* Le permite crear un nuevo servidor en solo 10 segundos usando la interfaz de Heroku Command Line.
* Esta plataforma de computación en la nube se encarga de parchear los sistemas y mantener todo saludable.
* Una gama de funcionalidades automatizadas que incluyen escalado, configuración, instalación y otras
* Fácil integración con otros productos de AWS
* Curva de aprendizaje media
* Ofrece la mejor experiencia para desarrolladores de su clase
* Conectividad segura a los datos de Salesforce para crear una vista única del cliente
* Flexibilidad para personalizar y respaldar las necesidades únicas del flujo de trabajo de DevOps

**¿Por que escoger servicios de cloud mas avanzados como AWS?**
* Debe optar por AWS cuando tenga equipos de DevOps que puedan configurar y administrar la infraestructura
* Tiene muy poco tiempo para invertir en la implementación de una nueva versión de su aplicación web o móvil.
* AWS ofrece un proceso de implementación sencillo para una aplicación
* El servicio web de AWS es una opción ideal cuando su proyecto necesita una alta potencia informática
* Le ayuda a mejorar la productividad del equipo de desarrollo de aplicaciones
* Una gama de funcionalidades automatizadas que incluyen la configuración, el escalado, la instalación y otras
* Utilice AWS para tener control total sobre los recursos.
* AWS permite a las organizaciones utilizar modelos de programación, sistemas operativos, bases de datos y arquitecturas ya familiares.
* Es un servicio rentable que le permite pagar solo por lo que usa, sin ningún compromiso inicial o a largo plazo.
* No necesitará gastar dinero en ejecutar y mantener centros de datos.
* Ofrece implementaciones rápidas
* Puede agregar o eliminar capacidad fácilmente.
* Se le permite el acceso a la nube rápidamente con capacidad ilimitada.
* El costo total de propiedad es muy bajo en comparación con cualquier servidor privado / dedicado.
* Ofrece facturación y gestión centralizadas
* Le permite implementar su aplicación en varias regiones del mundo con solo unos pocos clics

