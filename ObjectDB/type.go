package objectdb

import (
	object "Endor/Object"
	"context"
)

type ObjectDB interface {
	// Store will store the object in the data store. The object will have a // name and kind, and the Store method should create a unique ID.
	Store(ctx context.Context, object object.Object) error
	// GetObjectByID will retrieve the object with the provided ID.
	GetObjectByID(ctx context.Context, id string) (object.Object, error)
	// GetObjectByName will retrieve the object with the given name.
	GetObjectByName(ctx context.Context, name string) (object.Object, error)
	// ListObjects will return a list of all objects of the given kind.
	ListObjects(ctx context.Context, kind string) ([]object.Object, error)
	// DeleteObject will delete the object.
	DeleteObject(ctx context.Context, id string) error
}
