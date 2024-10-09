package password

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Hash(password string) (string, error)
	Compare(hashed string, plain string) error
}

func Init(cost int) {
	service = &bcryptPasswordService{cost: cost}
}

var service Service

func Hash(password string) (string, error) {
	return service.Hash(password)
}

func Compare(hashed, plain string) error {
	return service.Compare(hashed, plain)
}

type bcryptPasswordService struct {
	cost int // Cost factor for bcrypt (default is 10)
}

// HashPassword hashes the provided password using bcrypt.
func (b *bcryptPasswordService) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *bcryptPasswordService) Compare(hashedPassword, textPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(textPassword))
}
