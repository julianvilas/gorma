//************************************************************************//
// API "congo": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// User Model Description
type User struct {
	ID        int `gorm:"primary_key"` // This is the User Model PK field
	Bio       *string
	City      *string
	Country   *string
	Email     string
	Firstname string
	Lastname  string
	Proposals []Proposal // has many Proposals
	Reviews   []Review   // has many Reviews
	State     *string
	CreatedAt time.Time  // timestamp
	UpdatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	Get(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error

	ListAppUser(ctx context.Context) []*app.User
	OneUser(ctx context.Context, id int) (*app.User, error)

	ListAppUserLink(ctx context.Context) []*app.UserLink
	OneUserLink(ctx context.Context, id int) (*app.UserLink, error)

	UpdateFromCreateUserPayload(ctx context.Context, payload *app.CreateUserPayload, id int) error

	UpdateFromUpdateUserPayload(ctx context.Context, payload *app.UpdateUserPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id int) (User, error) {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "user", "get"}, now)
	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.RecordNotFound {
		return User{}, nil
	}

	return native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) []User {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "user", "list"}, now)
	var objs []User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error listing User", goa.KV{"error", err.Error()})
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *UserDB) Add(ctx context.Context, model *User) (*User, error) {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "user", "add"}, now)
	err := m.Db.Create(model).Error
	if err != nil {
		goa.Error(ctx, "error updating User", goa.KV{"error", err.Error()})
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "user", "update"}, now)
	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id int) error {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "user", "delete"}, now)
	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.Error(ctx, "error retrieving User", goa.KV{"error", err.Error()})
		return err
	}

	return nil
}

// UserFromCreateUserPayload Converts source CreateUserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromCreateUserPayload(payload *app.CreateUserPayload) *User {
	user := &User{}
	user.Email = payload.Email
	if payload.Country != nil {
		user.Country = payload.Country
	}
	user.Firstname = payload.Firstname
	user.Lastname = payload.Lastname
	if payload.State != nil {
		user.State = payload.State
	}
	if payload.Bio != nil {
		user.Bio = payload.Bio
	}
	if payload.City != nil {
		user.City = payload.City
	}

	return user
}

// UpdateFromCreateUserPayload applies non-nil changes from CreateUserPayload to the model
// and saves it
func (m *UserDB) UpdateFromCreateUserPayload(ctx context.Context, payload *app.CreateUserPayload, id int) error {
	now := time.Now()

	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromcreateUserPayload"}, now)
	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving User", goa.KV{"error", err.Error()})
		return err
	}
	obj.Lastname = payload.Lastname
	if payload.State != nil {
		obj.State = payload.State
	}
	if payload.Bio != nil {
		obj.Bio = payload.Bio
	}
	if payload.City != nil {
		obj.City = payload.City
	}
	if payload.Country != nil {
		obj.Country = payload.Country
	}
	obj.Firstname = payload.Firstname
	obj.Email = payload.Email

	err = m.Db.Save(&obj).Error
	return err
}

// UserFromUpdateUserPayload Converts source UpdateUserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromUpdateUserPayload(payload *app.UpdateUserPayload) *User {
	user := &User{}
	if payload.State != nil {
		user.State = payload.State
	}
	if payload.Bio != nil {
		user.Bio = payload.Bio
	}
	if payload.City != nil {
		user.City = payload.City
	}
	if payload.Country != nil {
		user.Country = payload.Country
	}
	if payload.Firstname != nil {
		user.Firstname = *payload.Firstname
	}
	if payload.Lastname != nil {
		user.Lastname = *payload.Lastname
	}
	user.Email = payload.Email

	return user
}

// UpdateFromUpdateUserPayload applies non-nil changes from UpdateUserPayload to the model
// and saves it
func (m *UserDB) UpdateFromUpdateUserPayload(ctx context.Context, payload *app.UpdateUserPayload, id int) error {
	now := time.Now()

	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromupdateUserPayload"}, now)
	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving User", goa.KV{"error", err.Error()})
		return err
	}
	if payload.State != nil {
		obj.State = payload.State
	}
	if payload.Bio != nil {
		obj.Bio = payload.Bio
	}
	if payload.City != nil {
		obj.City = payload.City
	}
	if payload.Country != nil {
		obj.Country = payload.Country
	}
	if payload.Firstname != nil {
		obj.Firstname = *payload.Firstname
	}
	if payload.Lastname != nil {
		obj.Lastname = *payload.Lastname
	}
	obj.Email = payload.Email

	err = m.Db.Save(&obj).Error
	return err
}
