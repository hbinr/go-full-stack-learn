package model

type User struct {
	UserID   int64
	Age      uint8
	UserName string
	Password string
	Email    string
}

// TableName sets the insert table name for this struct type
func (b *User) TableName() string {
	return "user"
}
