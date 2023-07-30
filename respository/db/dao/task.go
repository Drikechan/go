package dao

import (
	"context"
	"gorm.io/gorm"
	"test-todolist/respository/model"
	"test-todolist/types"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}

	return &TaskDao{NewClient(ctx)}
}

func (t *TaskDao) CreateTask(task *model.TaskService) error {
	return t.Model(&model.TaskService{}).Create(task).Error
}

func (t *TaskDao) ListTask(currentPage, pageSize int, userId uint) (r []*model.TaskService, total int64, err error) {
	err = t.Model(&model.TaskService{}).Preload("User").Where("uid=?", userId).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&r).Error
	return
}

func (t *TaskDao) UpdateTask(uid uint, req *types.UpdateListTaskReq) error {
	s := new(model.TaskService)
	err := t.Model(&model.TaskService{}).Where("id=? AND uid=?", req.ID, uid).First(&s).Error

	if err != nil {
		return err
	}

	if req.Status != 0 {
		s.Status = req.Status
	}

	if req.Content != "" {
		s.Content = req.Content
	}

	if req.Title != "" {
		s.Title = req.Title
	}
	return t.Save(s).Error

}

func (t *TaskDao) FindTaskIdByUserId(uId, id uint) (r *model.TaskService, err error) {
	err = t.Model(&model.TaskService{}).Where("id=? AND uid=?", id, uId).First(&r).Error
	return
}

func (t *TaskDao) DeleteTask(uId, id uint) error {
	r, err := t.FindTaskIdByUserId(uId, id)
	if err != nil {
		return err
	}

	return t.Delete(&r).Error
}
