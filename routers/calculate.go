package routers

func (r Router) AdminRouters() {
	adminGroup := r.router.Group("/api/v1")

	adminGroup.POST("/calculate", r.handler.CalculateCredit)
}
