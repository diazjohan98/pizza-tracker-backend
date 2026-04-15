package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	Username string
}

func (h *Handler) HandleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", LoginData{})
}

func (h *Handler) HandleLoginPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{Error: "Invalid input" + err.Error()})
		return
	}

	user, err := h.users.AuthenticateUser(form.Username, form.Password)
	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "Invalid Credentials",
		})
		return
	}

	// setSessionValue(c, "userID", fmt.Sprintf("%v", user.ID))
	// setSessionValue(c, "username", user.Username)

	errSession := setSessionValue(c, user.ID, user.Username)
	if errSession != nil {
		// Si SQLite falla, ahora lo veremos en rojo en la pantalla de Login
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "Error de base de datos al crear sesión: " + errSession.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) HandleLogout(c *gin.Context) {
	if err := ClearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) HandleAdminDashboard(c *gin.Context) {
	username := GetSessionString(c, "username")

	c.HTML(http.StatusOK, "admin.tmpl", AdminDashboardData{
		Username: username,
	})
}
