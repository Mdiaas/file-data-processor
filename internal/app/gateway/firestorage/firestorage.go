package firestorage

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
	"github.com/sirupsen/logrus"
)

type fstore struct {
	client     *firestore.Client
	collection string
}

func (f *fstore) AddStudent(ctx context.Context, student entity.Student) error {
	collection := f.client.Collection(f.collection)
	studentData := StudentToStudentDb(student)
	_, _, err := collection.Add(ctx, studentData)
	if err != nil {
		logrus.New().WithContext(ctx).WithError(err).Error("failed to add student to firestore")
		return err
	}
	return nil
}

func New(client *firestore.Client, studentCollection string) dataprovider.Firestorage {
	return &fstore{
		client:     client,
		collection: studentCollection,
	}
}
