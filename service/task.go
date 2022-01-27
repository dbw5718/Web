package service

import (
	"time"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}
type ShowTaskService struct{}
type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}
type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}
type DeleteTaskService struct{}

func (service *CreateTaskService) Create(id uint) serializer.Response {

	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetMsg(e.ERROR),
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}

func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	err := model.DB.First(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetMsg(e.ERROR) + "查询失败",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(e.SUCCESS),
	}
}

func (service *ListTaskService) List(uid uint) (response serializer.Response) {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count)),
	}
}

func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)

	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status
	err := model.DB.Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetMsg(e.ERROR) + "更新失败",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Data:   serializer.BuildTask(task),
		Msg:    "更新成功",
	}
}

func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetMsg(e.ERROR),
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    "删除成功",
	}
}
