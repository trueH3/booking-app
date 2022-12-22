package helper

type IUserRepository interface {
	SaveBooking(user *User)
	GetAllBookings() []*User
}

type UserRepositoryStruct struct{}

func (repository UserRepositoryStruct) SaveBooking(user *User) {
	transaction := GetDatabase().Txn(true)
	if err := transaction.Insert("user", user); err != nil {
		panic(err)
	}
	transaction.Commit()
}

func (repository UserRepositoryStruct) GetAllBookings() []*User {
	transaction := GetDatabase().Txn(false)
	usersAsResultIterator, err := transaction.Get("user", "id")
	if err != nil {
		panic(err)
	}

	bookings := []*User{}

	for obj := usersAsResultIterator.Next(); obj != nil; obj = usersAsResultIterator.Next() {
		bookings = append(bookings, obj.(*User))
	}
	transaction.Commit()
	return bookings
}
