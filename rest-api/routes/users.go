package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

// login menangani permintaan login user dengan mengikat payload JSON yang masuk ke model User.
// Fungsi ini mencoba mem-parsing body request sebagai JSON dan mengisi struct user.
// Jika parsing gagal, maka akan merespons dengan 500 Internal Server Error dan pesan yang sesuai.
//
// Metode context.ShouldBindJSON digunakan untuk mengikat payload JSON ke struct.
// ShouldBindJSON akan mengembalikan error jika body request tidak dapat di-parse sebagai JSON
// atau jika struktur JSON tidak sesuai dengan model User yang diharapkan.

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}
