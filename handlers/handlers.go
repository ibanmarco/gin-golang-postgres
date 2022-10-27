package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ibanmarco/gin-golang-postgres/initializers"
	"github.com/ibanmarco/gin-golang-postgres/models"
	"net/http"
)

var Book struct {
	Author    *string `gorm:"type:varchar(255);not null" json:"author"`
	Title     *string `json:"title"`
	Content   *string `gorm:"type:text;not null"`
	Publisher *string `json:"publisher"`
	Year      uint16  `json:"year"`
}

var ErrBookNotFound = fmt.Errorf("Book not found")

func WelcomeHandler(ctx *gin.Context) {
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome " + name})
}

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Are you right?"})
}

func ListBooksHandler(ctx *gin.Context) {
	var books []models.Books
	initializers.DB.Find(&books)
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}

func CreateBookHandler(ctx *gin.Context) {
	ctx.Bind(&Book)
	book := models.Books{
		Author:    Book.Author,
		Title:     Book.Title,
		Content:   Book.Content,
		Publisher: Book.Publisher,
		Year:      Book.Year,
	}

	payload := initializers.DB.Create(&book)
	if payload.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request 400"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}
func findBook(id string, ctx *gin.Context) (models.Books, error) {
	var book models.Books
	initializers.DB.First(&book, id)

	if book.ID == 0 {
		return book, ErrBookNotFound
	}

	return book, nil
}

func GetBookHandler(ctx *gin.Context) {
	//id, err := strconv.Atoi(ctx.Param("id"))
	//if err != nil {
	//	log.Fatal("Failed to convert to integer")
	//}

	book, err := findBook(ctx.Param("id"), ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func UpdateBookHandler(ctx *gin.Context) {
	var book models.Books
	id := ctx.Param("id")

	ctx.Bind(&Book)
	initializers.DB.First(&book, id)
	initializers.DB.Model(&book).Updates(models.Books{
		Author:    Book.Author,
		Title:     Book.Title,
		Content:   Book.Content,
		Publisher: Book.Publisher,
		Year:      Book.Year,
	})

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := findBook(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	initializers.DB.Delete(&models.Books{}, id)

	ctx.JSON(http.StatusOK, gin.H{"message": "Book was deleted"})
}
