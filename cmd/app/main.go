package main

import (
	"hey-notes-api/database"

	server "hey-notes-api/internal/api/http"
	"hey-notes-api/internal/api/http/controllers"
	noteRepository "hey-notes-api/pkg/repositories/note"
	userRepository "hey-notes-api/pkg/repositories/user"
	authService "hey-notes-api/pkg/services/auth"
	noteService "hey-notes-api/pkg/services/note"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := database.NewDB()
	dbImpl := database.NewDbImpl(db)
	validator := validator.New()

	userRepo := userRepository.NewUserRepositoryImpl()
	noteRepo := noteRepository.NewNoteRepositoryImpl()
	authService := authService.NewAuthServiceImpl(userRepo, dbImpl, validator)
	noteService := noteService.NewNoteServiceImpl(noteRepo, dbImpl, validator)
	httpHandler := controllers.NewRoute(authService, noteService)
	server := server.NewHttpImpl(httpHandler)

	server.Listen()
}