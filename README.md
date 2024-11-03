
# Instrucciones para la Configuración del Proyecto en Go

> **Repositorio del Proyecto**: [GitHub - Estructura del proyecto](https://github.com/EladioRocha/go-structure)  

### 1. Crear el Proyecto y Configurar Módulos

1.  **Crea una carpeta** para el proyecto y abre la terminal en esa carpeta.
    
2.  Ejecuta el siguiente comando para inicializar el módulo Go:
        
    `go mod init <nombre-del-proyecto>` 
    
3.  Luego, ejecuta:
    
    `go mod tidy` 
    
4.  Es importante instalar los siguientes módulos necesarios para este proyecto:
    
    ```bash
    go get github.com/gin-gonic/gin
    go get gorm.io/gorm
    go get gorm.io/driver/postgres
    go get github.com/joho/godotenv
    ```
    

### 2. Estructura del Proyecto

La estructura del proyecto puede organizarse de la siguiente manera:

```go
<nombre-del-proyecto>/
├── db/
│   └── db.go
├── handlers/
│   └── users.go
├── models/
│   └── user.go
├── routes/
│   └── routes.go
└── main.go
```

-   **db/db.go**: Configuración de la base de datos.
-   **handlers/users.go**: Controladores para manejar solicitudes relacionadas con usuarios.
-   **models/user.go**: Definición del modelo `User`.
-   **routes/routes.go**: Definición de las rutas de la API.
-   **main.go**: Punto de entrada para inicializar el servidor.

### 3. Definición de Campos en Estructuras con Tipos Específicos

A continuación, se explica cómo definir tipos de campos en una estructura utilizando las etiquetas `gorm:"type"` para almacenar correctamente diferentes tipos de datos en la base de datos.

#### Tipos de Datos en GORM (`type`)

| Tipo de Dato (`type`)        | Descripción                                                                                              | Ejemplo de Uso                        |
|------------------------------|----------------------------------------------------------------------------------------------------------|---------------------------------------|
| `int`                        | Entero de tamaño estándar en SQL.                                                                        | `gorm:"type:int"`                     |
| `smallint`                   | Entero pequeño (2 bytes).                                                                               | `gorm:"type:smallint"`                |
| `bigint`                     | Entero grande (8 bytes).                                                                                | `gorm:"type:bigint"`                  |
| `serial`                     | Entero auto-incremental.                                                                                 | `gorm:"type:serial"`                  |
| `varchar(n)`                 | Cadena de texto de longitud variable con un límite máximo de `n` caracteres.                            | `gorm:"type:varchar(255)"`            |
| `char(n)`                    | Cadena de texto de longitud fija con `n` caracteres.                                                     | `gorm:"type:char(10)"`                |
| `text`                       | Texto de longitud ilimitada.                                                                            | `gorm:"type:text"`                    |
| `boolean`                    | Valor booleano (`true` o `false`).                                                                       | `gorm:"type:boolean"`                 |
| `decimal(p, s)`              | Número decimal con `p` dígitos totales y `s` decimales.                                                  | `gorm:"type:decimal(10,2)"`           |
| `date`                       | Fecha sin tiempo.                                                                                       | `gorm:"type:date"`                    |
| `timestamp`                  | Marca de tiempo (fecha y hora).                                                                         | `gorm:"type:timestamp"`               |
| `json`                       | Almacena datos JSON en formato de texto (PostgreSQL y MySQL).                                           | `gorm:"type:json"`                    |
| `uuid`                       | Identificador único universal.                                                                          | `gorm:"type:uuid"`                    |


#### Ejemplo de Estructura de Datos Usando GORM

Aquí tienes un ejemplo de cómo definir una estructura `User` en `models/user.go`, utilizando diferentes tipos de datos con precisión para caracteres fijos, dinámicos y decimales:

```go
package models

import (
    "time"
    "github.com/shopspring/decimal"
)

// User estructura que representa a un usuario en la base de datos
type User struct {
    ID          uint            `gorm:"primaryKey;type:serial"`
    Username    string          `gorm:"type:varchar(50);not null"`       // Carácter dinámico de hasta 50 caracteres
    Email       string          `gorm:"type:varchar(100);unique;not null"` // Carácter dinámico, único y no nulo
    Password    string          `gorm:"type:char(60);not null"`          // Carácter fijo, 60 caracteres (para hash)
    Balance     decimal.Decimal `gorm:"type:decimal(10,2);not null"`     // Decimal con precisión de 10 dígitos y 2 decimales
    Active      bool            `gorm:"type:boolean;default:true"`       // Booleano con valor predeterminado
    CreatedAt   time.Time       `gorm:"autoCreateTime"`                  // Fecha de creación automática
    UpdatedAt   time.Time       `gorm:"autoUpdateTime"`                  // Fecha de actualización automática
}
```

### Explicación de los Campos

-   **ID**: Identificador auto-incremental (`serial`).
-   **Username**: Carácter dinámico con un límite de 50 caracteres.
-   **Email**: Carácter dinámico, único y no nulo, con un límite de 100 caracteres.
-   **Password**: Carácter fijo de 60 caracteres (útil para almacenar hashes de contraseñas).
-   **Balance**: Campo decimal con precisión, permite manejar dinero con exactitud (`decimal(10,2)`).
-   **Active**: Campo booleano con valor predeterminado en `true`.
-   **CreatedAt**: Fecha y hora de creación, generada automáticamente por GORM.
-   **UpdatedAt**: Fecha y hora de última actualización, actualizada automáticamente.