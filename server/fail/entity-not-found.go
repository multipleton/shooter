package fail

import (
	"errors"
	"fmt"
)

func EntityNotFound(id int) error {
	message := fmt.Sprintf("entity with id=%d not found", id)
	return errors.New(message)
}
