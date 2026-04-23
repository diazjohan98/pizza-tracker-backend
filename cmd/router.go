package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupRouter(router *gin.Engine, h *Handler, store sessions.Store) {
	router.Use(sessions.Sessions("pizza-tracker", store))
	router.GET("/", h.ServeNewOrderForm)
	router.POST("/new-order", h.HandleNewOrderPost)
	router.GET("/customer/:id", h.serveCustomer)
	router.GET("/notifications", h.notificationHandler)

	router.GET("/login", h.HandleLoginGet)
	router.POST("/login", h.HandleLoginPost)
	router.POST("/logout", h.HandleLogout)

	admin := router.Group("/admin")
	admin.Use(h.AuthMiddleware())
	{
		admin.GET("", h.ServerAdminDashboard)
		admin.POST("/order/:id/update", h.HandlerOrderPut)
		admin.POST("/order/:id/delete", h.HandlerOrderDelete)
		admin.GET("/notifications", h.adminNotificationHandler)
	}

	router.Static("/static", "./template/static")

	//* ---------------------------------
	//* FUNCIONES PARA LA API REACT
	//*----------------------------------

	api := router.Group("/api")
	{
		api.GET("/form-data", h.ApiGetOrderFormData)

		api.POST("/orders", h.ApiHandleNewOrderPost)

		api.GET("/orders/:id", h.ApiServeCustomer)
	}

}
