# 🍕 Pizza Tracker - Full-Stack Real-Time Order Management System

Un sistema completo para la gestión y seguimiento de órdenes de una pizzería, desarrollado con una arquitectura moderna que separa el **Backend en Go (Golang)** y el **Frontend en React**.

Este proyecto demuestra la construcción de una API REST robusta, el manejo de concurrencia, Server-Sent Events (SSE) para actualizaciones en tiempo real y el despliegue mediante contenedores **Docker**.

> 🔗 **Repositorio del Frontend:** La interfaz visual de este proyecto se encuentra en un repositorio separado. Puedes ver el código aquí: [diazjohan98/pizza-tracker-frontend](https://github.com/diazjohan98/pizza-tracker-frontend).

---

## 🚀 Características Principales

### 📡 Motor en Tiempo Real (Backend)

- **Server-Sent Events (SSE):** Flujo de datos unidireccional y eficiente para actualizar el frontend sin recargar la página.
- **Gestor de Notificaciones Concurrente:** Implementación de un patrón Pub/Sub utilizando canales de Go (`chan`) y `sync.RWMutex` para manejar múltiples clientes simultáneos de forma segura y sin fugas de memoria.
- **API RESTful:** Rutas protegidas y estructuradas bajo un enrutador limpio usando el framework Gin.

### 👨‍🍳 Panel de Administración (Frontend React)

- **Gestión de Órdenes (CRUD):** Visualización dinámica de órdenes activas con actualización de estados en tiempo real consumiendo la API de Go.
- **Seguridad y Sesiones:** Autenticación protegida mediante middleware en Go y manejo de cookies seguras (`credentials: "include"`) entre orígenes distintos (CORS).
- **Borrado en Cascada:** Eliminación segura de órdenes y sus ítems asociados para mantener la integridad de la base de datos usando GORM.

### 🍕 Vista del Cliente (Customer Tracker)

- **Seguimiento Visual:** Barra de progreso interactiva construida con React, GSAP y Tailwind CSS v4.
- **Actualizaciones Automáticas:** El cliente recibe notificaciones SSE cuando el administrador cambia el estado de su pizza (Ej: De "Baking" a "Ready").

---

## 🛠️ Tecnologías Utilizadas

**Backend (Este Repositorio):**

- **Go (Golang)** - Lógica del servidor y concurrencia.
- **Gin Web Framework** - Enrutamiento, CORS y middlewares.
- **GORM** - ORM para la interacción ágil con la base de datos.
- **SQLite** - Base de datos ligera para almacenamiento persistente y sesiones.
- **Docker & Docker Compose** - Contenerización de la arquitectura completa.

**Frontend ([Ver Repositorio](https://github.com/diazjohan98/pizza-tracker-frontend)):**

- **React & Vite** - Renderizado del lado del cliente y empaquetado ultrarrápido.
- **Tailwind CSS v4** - Estilos de utilidad y diseño responsivo.
- **JavaScript / API Fetch** - Consumo de la API REST y manejo del objeto `EventSource`.

---

## ⚙️ Instalación y Ejecución con Docker

Para facilitar la evaluación del proyecto, ambos repositorios están configurados para levantarse simultáneamente usando Docker Compose.

**Prerrequisitos:**

- [Docker](https://www.docker.com/) y Docker Compose instalados.
- Git.

### Pasos para iniciar el ecosistema local:

1. **Crear una carpeta principal** para alojar ambos proyectos:
   ```bash
   mkdir mi-pizza-order
   cd mi-pizza-order
   ```
2. Clonar el Backend (Este repositorio):

```
Bash
git clone [https://github.com/diazjohan98/pizza-tracker-backend.git](https://github.com/diazjohan98/pizza-tracker-backend.git) Pizza-shop-order-backend
```

3. Clonar el Frontend (En la misma carpeta principal):

```
Bash
git clone [https://github.com/diazjohan98/pizza-tracker-frontend.git](https://github.com/diazjohan98/pizza-tracker-frontend.git) Pizza-tracker-order-frontend
```

4. Levantar los contenedores:
   Entra a la carpeta del backend (donde se encuentra el archivo docker-compose.yml) y ejecuta:

```
Bash
cd Pizza-shop-order-backend
docker compose up --build -d
```

🎯 Acceder a la aplicación:
Vista del Cliente (Crear orden): http://localhost:3000/

Panel de Control (Admin): http://localhost:3000/login

(El backend se ejecutará silenciosamente en el puerto :8080)

```
👨‍💻 Autor
Johan Sebastian Vasquez Diaz
Ingeniero en Sistemas | Frontend (React, Vue, TS) | Go enthusiast
```
