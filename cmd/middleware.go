package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetSessionString(c, "userID")

		if userID == "" {
			// Si la petición es para la API de React, devolvemos JSON
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
				c.Abort()
				return
			}
			// Si es para las vistas de Go, redirigimos normal
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		_, err := h.users.GetUserByID(userID)
		if err != nil {
			ClearSession(c)
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
				c.Abort()
				return
			}
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
