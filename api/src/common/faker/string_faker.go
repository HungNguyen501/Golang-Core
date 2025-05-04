package faker

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Faker struct {
	faker gofakeit.Faker
}

func NewFaker() Faker {
	return Faker{
		faker: *gofakeit.New(time.Now().UnixNano()),
	}
}

func (f *Faker) FakeName() string {
	return f.faker.Name()
}

func (f *Faker) FakeEmail() string {
	return f.faker.Email()
}

func (f *Faker) FakeUUID() string {
	return f.faker.UUID()
}

func (f *Faker) FakePhone() string {
	return f.faker.Phone()
}

func (f *Faker) FakePassword() string {
	return f.faker.Password(true, true, true, true, false, 10)
}
