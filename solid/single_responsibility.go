package solid

// The Single Responsibility Principle (SRP) is a software design principle that
// states that every module or class in a program should have only one reason
// to change. This means that a class or module should only be responsible for one thing or behavior.

// Here is an example of implementing the SRP in Go:

// Suppose we have a User struct that represents a user in our application.
// The User struct has several methods such as Create, Update, and Delete, as well as some helper methods like Validate.
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func (u *User) Create() error {
	// Code to create user in database
	return nil
}

func (u *User) Update() error {
	// Code to update user in database
	return nil
}

func (u *User) Delete() error {
	// Code to delete user from database
	return nil
}

func (u *User) Validate() error {
	// Code to validate user data
	return nil
}

// In this implementation, the User struct is responsible for several things:
// creating, updating, and deleting users in the database, as well as validating user data.

// To apply SRP, we should separate the responsibilities of the User struct into separate classes or modules.
// For example, we could create a UserRepository struct that is responsible for creating, updating, and
// deleting users in the database:
type UserRepository struct {
	// Database connection
}

func (ur *UserRepository) Create(user *User) error {
	// Code to create user in database
	return nil
}

func (ur *UserRepository) Update(user *User) error {
	// Code to update user in database
	return nil
}

func (ur *UserRepository) Delete(user *User) error {
	// Code to delete user from database
	return nil
}

// And we could create a separate UserValidator struct that is responsible for validating user data:
type UserValidator struct{}

func (uv *UserValidator) Validate(user *User) error {
	// Code to validate user data
	return nil
}

// By separating the responsibilities of the User struct into separate classes or modules,
// we have made our code more modular, easier to maintain, and more testable.
// Each class or module is responsible for only one thing, which makes it easier to reason about and modify without affecting other parts of the code.
