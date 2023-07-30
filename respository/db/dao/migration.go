package dao

import (
	"fmt"
	"test-todolist/respository/model"
)

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.UserList{}, &model.TaskService{})
	fmt.Println("------------------------------")

	if err != nil {
		fmt.Println("%+v", err)
		return
	}
}
