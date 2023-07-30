package service

import (
	"context"
	"fmt"
	"sync"
	"test-todolist/pkg/ctl"
	"test-todolist/respository/db/dao"
	"test-todolist/respository/model"
	"test-todolist/types"
	"time"
)

var TaskSrvIns *TaskSrv

var TaskSrvOnes sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnes.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (s *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return
	}

	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)

	if err != nil {
		return
	}

	task := &model.TaskService{
		User:      *user,
		Uid:       user.ID,
		Content:   req.Content,
		Status:    0,
		Title:     req.Title,
		StartTime: time.Now().Unix(),
	}

	err = dao.NewTaskDao(ctx).CreateTask(task)

	if err != nil {
		fmt.Println("新增任务错误")
		return
	}
	return ctl.RespSuccess(), err
}

func (s *TaskSrv) GetTaskList(ctx context.Context, req *types.ListTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return
	}
	taskList, total, err := dao.NewTaskDao(ctx).ListTask(req.CurrentPage, req.PageSize, u.Id)
	if err != nil {
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range taskList {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			Status:    v.Status,
			StartTime: v.StartTime,
			CreateAt:  v.CreatedAt.Unix(),
			EndTime:   v.EndTime,
			//View: v.View(),
		})
	}

	return ctl.RespList(taskRespList, total), nil

}

func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateListTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return
	}

	err = dao.NewTaskDao(ctx).UpdateTask(u.Id, req)

	if err != nil {
		return
	}

	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return
	}

	err = dao.NewTaskDao(ctx).DeleteTask(u.Id, req.ID)
	if err != nil {
		return
	}
	return ctl.RespSuccess(), nil
}
