package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claims, err := utils.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		fmt.Println(err)
	}
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claims.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	//claim,err:=utils.ParseToken(c.GetHeader("Authorization"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		//fmt.Println(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, err := utils.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		fmt.Println(err)
	}
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		//fmt.Println(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	//claim,err:=utils.ParseToken(c.GetHeader("Authorization"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		//fmt.Println(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func SearchTasks(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		//fmt.Println(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	//claim,_:=utils.ParseToken(c.GetHeader("Authorization"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		//fmt.Println(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
