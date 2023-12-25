package service

import (
	"30.80.1/pkg/storage/postgresql"
	"fmt"
	"log"
)

func Run() {
	var err error
	dbURL := "postgres://postgres:postgres@localhost:5432/tasks"

	db, err := postgresql.Init(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	//task := model.Task{
	//	Opened:     time.Now().Unix(),
	//	AuthorID:   5,
	//	AssignedID: 3,
	//	Title:      "Task N",
	//	Content:    "Description for Task N",
	//}

	//id, err := db.NewTask(task)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(id)

	tasks, err := db.Tasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)

	//task, err := db.TaskById(1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(task)

	//tasks, err := db.TasksByAuthor(2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(tasks)

	//db.UpdateTask(2, task)
	//fmt.Println(db.TaskById(2))

	//err = db.DeleteTask(20)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
