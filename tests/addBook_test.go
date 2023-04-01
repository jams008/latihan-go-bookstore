package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jams008/latihan-bookstore/pkg/books"
	"github.com/jams008/latihan-bookstore/pkg/common/models"
)

func TestAddBook(t *testing.T) {
	viper.SetConfigFile("../pkg/common/envs/.env")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Inisialisasi database
	dsn := viper.GetString("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize database: %s", err.Error())
	}

	// Inisialisasi router
	r := gin.Default()
	books.RegisterRoutes(r, db)

	// Membuat request POST untuk menambahkan buku baru
	newBook := models.Book{
		Title:       "Book Title Unit Test",
		Author:      "Book Author Unit Test",
		Description: "Book Description Unit Test",
	}
	requestBody, _ := json.Marshal(newBook)
	req, err := http.NewRequest("POST", "/books/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err.Error())
	}

	// Menjalankan request POST
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Memeriksa hasil response
	assert.Equal(t, http.StatusCreated, w.Code)

	var responseBook models.Book
	err = json.Unmarshal(w.Body.Bytes(), &responseBook)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err.Error())
	}

	assert.Equal(t, newBook.Title, responseBook.Title)
	assert.Equal(t, newBook.Author, responseBook.Author)
}
