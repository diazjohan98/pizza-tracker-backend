# 🍕 Pizza Tracker - Real-Time Order Management System

Un sistema completo para la gestión y seguimiento de órdenes de una pizzería, desarrollado con **Go (Golang)**. Este proyecto incluye un panel de administración interactivo y una vista de cliente con actualizaciones en tiempo real, demostrando el uso de concurrencia, Server-Sent Events (SSE) y renderizado dinámico de plantillas.

## 🚀 Características Principales

### 📡 Motor en Tiempo Real (Real-Time)

- **Server-Sent Events (SSE):** Flujo de datos unidireccional y eficiente para actualizar interfaces sin recargar la página.
- **Gestor de Notificaciones Concurrente:** Implementación de un patrón Pub/Sub utilizando canales de Go (`chan`) y `sync.RWMutex` para manejar múltiples clientes simultáneos de forma segura y sin fugas de memoria.

### 👨‍🍳 Panel de Administración (Admin Dashboard)

- **Gestión de Órdenes (CRUD):** Visualización dinámica de órdenes activas con actualización de estados en tiempo real mediante menús desplegables auto-enviables (`onchange`).
- **Seguridad y Sesiones:** Autenticación protegida mediante middleware y manejo de sesiones con `gin-contrib/sessions` respaldado por la base de datos.
- **Borrado en Cascada:** Eliminación segura de órdenes y sus ítems asociados para mantener la integridad de la base de datos usando GORM.

### 🍕 Vista del Cliente (Customer Tracker)

- **Seguimiento Visual:** Barra de progreso interactiva construida con Tailwind CSS y animaciones JavaScript que reacciona a los eventos del servidor.
- **Actualizaciones Automáticas:** El cliente recibe notificaciones SSE cuando el administrador cambia el estado de su pizza (Ej: De "Baking" a "Ready").

---

## 🛠️ Tecnologías Utilizadas

**Backend:**

- [Go (Golang)](https://go.dev/)
- [Gin Web Framework](https://gin-gonic.com/) - Enrutamiento y middleware.
- [GORM](https://gorm.io/) - ORM para la interacción ágil con la base de datos.
- **SQLite** - Base de datos ligera para almacenamiento persistente y sesiones.

**Frontend:**

- **Go HTML Templates** (`html/template`) - Renderizado del lado del servidor con inyección de funciones personalizadas (FuncMaps).
- [Tailwind CSS](https://tailwindcss.com/) - Estilos de utilidad para un diseño web responsivo, moderno y con efecto _Glassmorphism_.
- **JavaScript (Vanilla)** - Lógica del cliente para el manejo del `EventSource` y la manipulación del DOM.

---

## 📂 Arquitectura del Proyecto

El proyecto sigue una estructura limpia y modularizada:

```text
📦 pizza-shop-order
 ┣ 📂 cmd                  # Punto de entrada y controladores principales
 ┃ ┣ 📜 main.go            # Configuración del servidor Gin y rutas
 ┃ ┣ 📜 admin.go           # Controladores del panel de administración
 ┃ ┣ 📜 customer.go        # Controladores de la vista del cliente
 ┃ ┣ 📜 router.go          # Definición de endpoints y middlewares
 ┃ ┣ 📜 event.go           # Endpoints de streaming para SSE
 ┃ ┣ 📜 notification.go    # Motor Pub/Sub concurrente
 ┃ ┗ 📜 utils.go           # Funciones de ayuda (carga de plantillas, env vars)
 ┣ 📂 internal/models      # Lógica de dominio e interacciones con la BD
 ┃ ┣ 📜 models.go          # Configuración de GORM
 ┃ ┣ 📜 order.go           # Estructuras de la Orden e Ítems y operaciones CRUD
 ┃ ┗ 📜 user.go            # Estructura de Usuarios y autenticación
 ┣ 📂 template             # Vistas del Frontend
 ┃ ┣ 📜 admin.tmpl         # Dashboard UI
 ┃ ┣ 📜 customer.tmpl      # Order Tracker UI
 ┃ ┣ 📜 login.tmpl         # Interfaz de acceso
 ┃ ┗ 📂 static             # Assets estáticos (imágenes)
 ┗ 📂 data                 # Almacenamiento de SQLite (auto-generado)
⚙️ Instalación y Ejecución
Prerrequisitos
Go 1.20+ instalado en tu sistema.

Git.

Pasos para iniciar el servidor local
Clonar el repositorio:

Bash
git clone [https://github.com/tu-usuario/pizza-tracker-backend.git](https://github.com/tu-usuario/pizza-tracker-backend.git)
cd pizza-tracker-backend
Instalar dependencias:

Bash
go mod tidy
Ejecutar el proyecto:

Bash
go run ./cmd
Acceder a la aplicación:

Crear una orden de prueba (Cliente): http://localhost:8080/

Acceder al panel de control (Admin): http://localhost:8080/login

👨‍💻 Autor
Johan Sebastian Vasquez Diaz

Software Developer| Frontend explorando el mundo Full-Stack con Go.
```
