Base Structure in Repository

FindAll()
↓
QueryContext()
↓
\*sql.Rows
↓
for rows.Next()

FindByID()
↓
QueryRowContext()
↓
\*sql.Row
↓
Scan()


```golang

    package main

import (
	// _ "github.com/go-sql-driver/mysql"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)

	}

	// Load .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Failed to load .env")
	// }

	//  Configure DSN
	// dsn := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/%s?parseTime=true",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	// Connect DB
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Database connected!")

	// test cli
	// if len(os.Args) < 3 {
	// 	fmt.Println("Usage: go run ./cmd/todo <name>")
	// 	return
	// }

	// var greetings string
	// getGreetings := os.Args[1]

	// switch getGreetings {
	// case "goodmorning":
	// 	greetings = "Good Morning"
	// case "goodafternoon":
	// 	greetings = "Good Afternoon"
	// case "goodevening":
	// 	greetings = "Good Evening"
	// default:
	// 	log.Fatal("Invalid Input")
	// }

	// name := os.Args[2]

	// fmt.Printf("Hello %s, %s!\n", greetings, name)

	// Belajar Struct
	// newTodo := todo.Todo{
	// 	ID:          1,
	// 	Title:       "Belajar Go",
	// 	Description: "Memahami struct dan method",
	// 	Completed:   false,
	// }

	// fmt.Println(newTodo)

	// Array Slice
	// todos := []todo.Todo{
	// 	{
	// 		ID:        1,
	// 		Title:     "Belajar Go",
	// 		Completed: false,
	// 	},
	// 	{
	// 		ID:        2,
	// 		Title:     "Belajar MySQL",
	// 		Completed: false,
	// 	},
	// 	{
	// 		ID:        3,
	// 		Title:     "Membuat CLI",
	// 		Completed: true,
	// 	},
	// }

	// printTodos(todos)

	// item := todo.Todo{
	// 	ID:    1,
	// 	Title: "Belajar Go",
	// }

	// item.Complete()

	// fmt.Println(item.IsCompleted())

	// item.Reopen()

	// fmt.Println(item.IsCompleted())

	// Error Handling
	// result, err := divide(10, 0)

	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Result: %v\n", result)

	// Repository
	// repo := todo.NewMemoryRepository()

	// repo.Create()
	// repo.FindAll()
	// repo.FindByID()

	// Init Context
	// ctx := context.Background()
	// repository := todo.NewMemoryRepository()
	// service := todo.NewService(repository)

	// Test Create
	// item, err := service.Create(
	// 	ctx,
	// 	"Belajar Go",
	// 	"Memahami Interface dan repository",
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf(
	// 	"Created Todo: #%d %s\n",
	// 	item.ID,
	// 	item.Title,
	// )

	// data, err := service.List(ctx)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// printTodos(data)

}

// Function
// func printTodos(todos []todo.Todo) {
// 	for _, item := range todos {
// 		fmt.Printf("[%d] %s\n", item.ID, item.Title)
// 	}
// }

// Error Handling
// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("Cannot divide by zero")
// 	}

// 	return a / b, nil
// }


```
