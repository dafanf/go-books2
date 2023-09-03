package services

import (
	"database/sql"
	"go-structure-project/models"
)

// GetBooksFromDB retrieves a list of books from the database.
func GetBooksFromDB(db *sql.DB) ([]models.Book, error) {
	var books []models.Book
	query := "SELECT id, title, author, quantity FROM books"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

// GetBookByID retrieves a book by its ID from the database.
func GetBookByID(db *sql.DB, id string) (*models.Book, error) {
	var book models.Book
	query := "SELECT id, title, author, quantity FROM books WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// UpdateBookQuantity updates a book's quantity in the database.
func UpdateBookQuantity(db *sql.DB, b *models.Book) error {
	query := "UPDATE books SET quantity = ? WHERE id = ?"
	_, err := db.Exec(query, b.Quantity, b.ID)
	return err
}

// CreateBookInDB creates a new book in the database.
func CreateBookInDB(db *sql.DB, newBook models.Book) error {
	query := "INSERT INTO books (id, title, author, quantity) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
	return err
}
