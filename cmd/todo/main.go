package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/razenry/cli-todo-go/internal/database"
	"github.com/razenry/cli-todo-go/internal/todo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := database.NewMySQL()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repository := todo.NewMySQLRepository(db)
	service := todo.NewService(repository)

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command>")
	}

	command := os.Args[1]

	ctx := context.Background()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <title>")
		}

		title := os.Args[2]

		var description string

		if len(os.Args) < 4 {
			description = ""
		}

		if len(os.Args) == 4 {
			description = os.Args[3]
		}

		item, err := service.Create(
			ctx,
			title,
			description,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			// "Todo created: ID=%d, Title=%s\n",
			"Todo created: ID=%d, Title=%s, Description=%s, \n",
			item.ID,
			item.Title,
			item.Description,
		)
	case "list":
		items, err := service.List(ctx)

		if err != nil {
			log.Fatal(err)
		}

		for _, item := range items {

			var isCompleted string

			if item.Completed {
				isCompleted = "Done"
			} else {
				isCompleted = "Pending"
			}

			fmt.Printf(
				"%d. %s [%s]\n",
				item.ID,
				item.Title,
				isCompleted,
			)
		}
	case "show":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo show <id>")
			return
		}

		id, err := strconv.ParseUint(os.Args[2], 10, 64)

		if err != nil {
			fmt.Println("ID must be a number")
			return
		}

		item, err := service.Get(ctx, id)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("status", item.Completed)

		var isCompleted string

		if item.Completed {
			isCompleted = "Done"
		} else {
			isCompleted = "Pending"
		}

		fmt.Printf(
			"ID:%d\n Title:%s\n Description:%s\n CreatedAt:%s\n UpdateAt:%s\n Status:[%s]\n",
			item.ID,
			item.Title,
			item.Description,
			item.CreatedAt,
			item.UpdatedAt,
			isCompleted,
		)
	case "update":

		if len(os.Args) < 3 {
			fmt.Println("Usage: todo update <id>")
			return
		}

		id, err := strconv.ParseUint(os.Args[2], 10, 64)

		if err != nil {
			fmt.Println("ID must be a number")
			return
		}

		if len(os.Args) < 4 {
			fmt.Println("Usage: todo update <id> <title>")
			return
		}

		newTitle := os.Args[3]

		var newDesc string

		if len(os.Args) >= 5 {
			newDesc = os.Args[4]
		}

		err = service.Update(
			ctx,
			newTitle,
			newDesc,
			id,
		)

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf(
			"Todo %d updated successfully\n",
			id,
		)
	case "delete":
		// cek argument
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}

		// ParseUint
		id, err := strconv.ParseUint(os.Args[2], 10, 64)

		if err != nil {
			fmt.Println("ID must be a number")
			return
		}

		//  service.Get()
		item, err := service.Get(ctx, id)
		if err != nil {
			fmt.Println(err)
			return
		}

		// service.Delete()
		err = service.Delete(ctx, id)
		if err != nil {
			fmt.Println(err)
			return
		}

		// tampilkan success
		fmt.Printf(`%v deleted successfully`, item.Title)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo complete <id>")
			return
		}

		// ParseUint
		id, err := strconv.ParseUint(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("ID must be a number")
			return
		}

		//  service.Complete()
		err = service.Complete(ctx, id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(
			"Todo %d completed successfully\n",
			id,
		)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

}
