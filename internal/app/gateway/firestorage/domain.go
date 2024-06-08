package firestorage

import "time"

type StudentDb struct {
	Id             string    `firestore:"id"`
	FullName       string    `firestore:"full_name"`
	DocumentNumber string    `firestore:"document_number"`
	BirthDate      time.Time `firestore:"birth_date"`
	MotherName     string    `firestore:"mother_name"`
	FatherName     string    `firestore:"father_name"`
	Errors         []string  `firestore:"errors"`
	CreatedAt      time.Time `firestore:"created_at"`
}
