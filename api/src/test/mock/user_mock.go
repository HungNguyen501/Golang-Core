package mock

import (
	"golang-core/api/src/client/request"
	"golang-core/api/src/common/faker"
)

var fk faker.Faker = faker.NewFaker()

func MockUserInsertRequest() request.UserInsertRequest {
	return request.UserInsertRequest{
		Email:    fk.FakeEmail(),
		Name:     fk.FakeName(),
		Password: fk.FakePassword(),
	}
}

func MockUserUpdateRequest(id string) request.UserUpdateRequest {
	email := fk.FakeEmail()
	name := fk.FakeName()
	password := fk.FakePassword()
	return request.UserUpdateRequest{
		ID:       id,
		Email:    &email,
		Name:     &name,
		Password: &password,
	}
}
