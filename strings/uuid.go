package strings

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func ParseID() (uuid.UUID, error) {
	u1 := uuid.UUID{}
	nid := uuid.NullUUID{}
	fmt.Println(u1 == nid.UUID)
	id := uuid.NewV4()
	fmt.Println(id.String())
	err := u1.Scan(id.String())
	fmt.Println(u1.String())
	fmt.Println(u1 == id)
	return u1, err
}
