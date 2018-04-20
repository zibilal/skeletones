package uuid

import (
	"github.com/satori/go.uuid"
	"fmt"
)

type ID string

func (u ID) String() string {
	return string(u)
}

func (u ID) Valid() bool {
	str := string(u)
	_, err := uuid.FromString(str)

	return err == nil
}

func GenerateID() ID {
	id := fmt.Sprintf("%s", uuid.NewV4())
	return ID(id)
}
