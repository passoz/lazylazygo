package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		cmd  string
		name string
		arg  string
	)

	argCount := len(os.Args)

	if argCount < 4 && os.Args[1] != "new" {
		fmt.Println("Not enough arguments")
		os.Exit(0)
	}

	for i := 1; i < argCount; i++ {
		if os.Args[i] == "new" {
			cmd = "new"
			name = os.Args[i+1]
			break
		} else if os.Args[i] == "generate" {
			cmd = "generate"
			name = os.Args[i+1]
			arg = os.Args[i+2]
			break
		}
	}

	fmt.Println(cmd, name, arg)

	flag.Parse()

	switch cmd {
	case "new":
		fmt.Println("Creating new project")
		createProject(name)
		fmt.Println("Project name: ", name, " created")
	case "generate":
		fmt.Println("Generating code.")
		generateCode()
		fmt.Println("Code generated")
	default:
		fmt.Println("Unknown command")
	}

}

func createProject(name string) {

	directories := []string{
		"internal",
		"internal/entity",
		"internal/repository",
		"internal/service",
		"internal/controller",
		"internal/middleware",
		"internal/database",
		"internal/database/schema",
		"internal/handler",
		"internal/schema",
		"internal/tools",
		"internal/types",
		"cmd",
		"cmd/api",
	}

	files := []string{
		"cmd/api/main.go",
		"go.mod",
		"internal/database/database.go",
		"internal/entity/user.go",
		"internal/repository/user_repository.go",
		"internal/service/user_service.go",
		"internal/controller/user_controller.go",
		"internal/handler/user_handler.go",
		"internal/service/auth_service.go",
		"internal/controller/auth_controller.go",
		"internal/handler/auth_handler.go",
		"internal/middleware/auth_middleware.go",
		"internal/database/schema/schema.sql",
		"internal/database/schema/query.sql",
	}

	for _, dir := range directories {
		os.Mkdir(name+"/"+dir, 0755)
	}

	for _, file := range files {
		os.Create(name + "/" + file)
	}

	mainCode := "package main\n\nfunc main () {}\n"
	writeCode(mainCode, name+"/cmd/api/main.go")

	modCode := "module " + name + "\n\ngo 1.21.6\n"
	writeCode(modCode, name+"/go.mod")

	dbCode := "package database\n\nfunc Connect() {}\n"
	writeCode(dbCode, name+"/internal/database/database.go")

	userCode := generateEntity("User", map[string]string{
		"id":         "string",
		"email":      "string",
		"created_at": "int64",
		"updated_at": "int64",
	})
	writeCode(userCode, name+"/internal/entity/user.go")

	userRepoCode := generateRepository("User",
		map[string]string{
			"db": "sql.DB",
		},
		map[string]string{
			"CreateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"GetUser":        "id string",
			"GetUsers":       "",
			"GetUserByEmail": "email string",
			"UpdateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"DeleteUser":     "id string",
		})
	writeCode(userRepoCode, name+"/internal/repository/user_repository.go")

	userServiceCode := generateService("User",
		map[string]string{
			"userRepo": "UserRepository",
		},
		map[string]string{
			"CreateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"GetUser":        "id string",
			"GetUsers":       "",
			"GetUserByEmail": "email string",
			"UpdateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"DeleteUser":     "id string",
		})
	writeCode(userServiceCode, name+"/internal/service/user_service.go")

	userControllerCode := generateController("User",
		map[string]string{
			"userService": "UserService",
		},
		map[string]string{
			"CreateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"GetUser":        "id string",
			"GetUsers":       "",
			"GetUserByEmail": "email string",
			"UpdateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"DeleteUser":     "id string",
		})
	writeCode(userControllerCode, name+"/internal/controller/user_controller.go")

	userHandlerCode := generateHandler("User",
		map[string]string{
			"userController": "UserController",
		},
		map[string]string{
			"CreateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"GetUser":        "id string",
			"GetUsers":       "",
			"GetUserByEmail": "email string",
			"UpdateUser":     "id string, email string, password string, created_at int64, updated_at int64",
			"DeleteUser":     "id string",
		})
	writeCode(userHandlerCode, name+"/internal/handler/user_handler.go")

	identityServiceCode := generateInterface("Identity", map[string]string{
		"Login":         "email string, password string",
		"Signup":        "email string, password string",
		"Logout":        "token string",
		"Refresh":       "token string",
		"Delete":        "token string",
		"OTP":           "email string, otp string",
		"ConfirmEmail":  "email string, token string",
		"ResetPassword": "email string, token string, password string",
	})
	writeCode(identityServiceCode, name+"/internal/service/identity_service.go")

	authServiceCode := generateService("Auth",
		map[string]string{
			"userRepo": "UserRepository",
		},
		map[string]string{
			"Login":         "email string, password string",
			"Signup":        "email string, password string",
			"Logout":        "token string",
			"Refresh":       "token string",
			"Delete":        "token string",
			"OTP":           "email string, otp string",
			"ConfirmEmail":  "email string, token string",
			"ResetPassword": "email string, token string, password string",
		})
	writeCode(authServiceCode, name+"/internal/service/auth_service.go")

	authControllerCode := generateController("Auth",
		map[string]string{
			"authService": "AuthService",
		},
		map[string]string{
			"Login":         "email string, password string",
			"Signup":        "email string, password string",
			"Logout":        "token string",
			"Refresh":       "token string",
			"Delete":        "token string",
			"OTP":           "email string, otp string",
			"ConfirmEmail":  "email string, token string",
			"ResetPassword": "email string, token string, password string",
		})
	writeCode(authControllerCode, name+"/internal/controller/auth_controller.go")

	authHandlerCode := generateHandler("Auth",
		map[string]string{
			"authController": "AuthController",
		},
		map[string]string{
			"Login":         "email string, password string",
			"Signup":        "email string, password string",
			"Logout":        "token string",
			"Refresh":       "token string",
			"Delete":        "token string",
			"OTP":           "email string, otp string",
			"ConfirmEmail":  "email string, token string",
			"ResetPassword": "email string, token string, password string",
		})
	writeCode(authHandlerCode, name+"/internal/handler/auth_handler.go")

	authMiddlewareCode := "package middlewares\n\nfunc AuthMiddleware() {}\n"
	writeCode(authMiddlewareCode, name+"/internal/middleware/auth_middleware.go")

	schemaCode := "drop database if exists " + name + ";\ncreate database " + name + ";\n\n"
	schemaCode += "create table users (\n\tid text not null,\n\temail varchar(255) not null,\n\tcreated_at integer not null,\n\tupdated_at integer not null,\n\tprimary key (id)\n);\n"
	writeCode(schemaCode, name+"/internal/database/schema/schema.sql")

	queryCode := "select * from users;\n"
	writeCode(queryCode, name+"/internal/database/schema/query.sql")

}

func generateCode() {
	fmt.Println("Generating code")
}

func writeCode(code string, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	_, err = f.Write([]byte(code))

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
}

func generateStruct(name string, args map[string]string) string {
	srt := "type " + name + " struct {\n"
	for k, v := range args {
		srt += "\t" + k + " " + v + "\n"
	}
	srt += "}\n\n"
	return srt
}

func generateInterface(name string, args map[string]string) string {
	itf := "type " + name + " interface {\n"
	for k, v := range args {
		itf += "\t" + k + " " + v + "\n"
	}
	itf += "}\n\n"
	return itf
}

func generateMethods(name string, args map[string]string) string {
	mtd := ""
	for k, v := range args {
		mtd += "func (u *" + name + ") " + k + " (" + v + ") {}\n\n"
	}
	return mtd
}

func generateGenericNew(name string) string {
	return "func New" + name + "() {\n\t\treturn &" + name + "{\n\t}\n}\n\n"
}

func generateHandler(name string, fields map[string]string, args map[string]string) string {
	str := "package handler\n\n"
	str += "import (\n\t\"../controller\"\n)\n\n"
	str += generateGenericNew(name + "Handler")
	str += generateStruct(name+"Handler", fields)
	str += generateMethods(name+"Handler", args)

	return str
}

func generateController(name string, fields map[string]string, args map[string]string) string {
	str := "package controller\n\n"
	str += "import (\n\t\"../service\"\n)\n\n"
	str += generateGenericNew(name + "Controller")
	str += generateStruct(name+"Controller", fields)
	str += generateMethods(name+"Controller", args)

	return str
}

func generateService(name string, fields map[string]string, args map[string]string) string {
	str := "package service\n\n"
	str += "import (\n\t\"../repository\"\n)\n\n"
	str += generateGenericNew(name + "Service")
	str += generateStruct(name+"Service", fields)
	str += generateMethods(name+"Service", args)

	return str
}

func generateRepository(name string, fields map[string]string, args map[string]string) string {
	str := "package repository\n\n"
	str += "import (\n\t\"../repository\"\n)\n\n"
	str += generateGenericNew(name + "Service")
	str += generateStruct(name+"Service", fields)
	str += generateMethods(name+"Service", args)

	return str
}

func generateEntity(name string, fields map[string]string) string {
	str := "package service\n\n"
	str += generateGenericNew(name + "Service")
	str += generateStruct(name+"Service", fields)

	return str
}
