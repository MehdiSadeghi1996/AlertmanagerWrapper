package alerts_controller

//func CreateUser(c *gin.Context) {
//	var user users.User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		restErr := errors.NewBadRequestError("invalid json")
//		c.JSON(restErr.Status, restErr)
//		return
//	}
//	result, saveErr := services.CreateUser(user)
//	if saveErr != nil {
//		c.JSON(saveErr.Status, saveErr)
//		return
//	}
//	c.JSON(http.StatusCreated, result)
//}
