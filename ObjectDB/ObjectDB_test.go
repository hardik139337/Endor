package objectdb

import (
	object "Endor/Object"
	"context"
	"fmt"
	"testing"
)

func TestObjectDB(t *testing.T) {

	db := NewRedisObjectDB("")
	s, err := db.client.Keys(context.Background(), "*").Result()
	if err != nil {
		panic(err)
	}
	db.client.Del(context.Background(), s...)

	// Create some test data
	person1 := &object.Person{
		Name:     "John",
		ID:       "",
		LastName: "Doe",
		Birthday: "01/01/1990",
	}
	animal1 := &object.Animal{
		Name:    "Max",
		ID:      "",
		Type:    "Dog",
		OwnerID: "",
	}
	person2 := &object.Person{
		Name:     "Jane",
		ID:       "",
		LastName: "Doe",
		Birthday: "01/01/1995",
	}
	animal2 := &object.Animal{
		Name:    "Charlie",
		ID:      "",
		Type:    "Cat",
		OwnerID: "",
	}

	// Test object creation and retrieval
	err = db.Store(context.Background(), person1)
	if err != nil {
		t.Fatalf("failed to store object: %v", err)
	}
	err = db.Store(context.Background(), animal1)
	if err != nil {
		t.Fatalf("failed to store object: %v", err)
	}
	err = db.Store(context.Background(), person2)
	if err != nil {
		t.Fatalf("failed to store object: %v", err)
	}
	err = db.Store(context.Background(), animal2)
	if err != nil {
		t.Fatalf("failed to store object: %v", err)
	}

	// Test object retrieval by ID
	obj, err := db.GetObjectByID(context.Background(), person1.GetID())
	if err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}
	if obj.GetKind() != person1.GetKind() || obj.GetName() != person1.GetName() {
		t.Errorf("retrieved object does not match stored object")
	}

	// Test object retrieval by name
	obj, err = db.GetObjectByName(context.Background(), animal1.GetName())
	if err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}
	if obj.GetKind() != animal1.GetKind() || obj.GetName() != animal1.GetName() {
		t.Errorf("retrieved object does not match stored object")
	}

	results, err := db.ListObjects(context.Background(), person1.GetKind())
	fmt.Printf("objs: %v\n", results)
	if err != nil {
		t.Fatalf("failed to retrieve object list: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 objects, got %d", len(results))
	}

	// Test object deletion
	err = db.DeleteObject(context.Background(), animal2.GetID())
	if err != nil {
		t.Fatalf("failed to delete object: %v", err)
	}
	_, err = db.GetObjectByID(context.Background(), animal2.GetID())
	if err == nil {
		t.Errorf("object was not deleted")
	}
}
