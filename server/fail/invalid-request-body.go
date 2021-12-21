package fail

import (
	"errors"
)

func InvalidRequestBody() error {
	message := "invalid request body"
	return errors.New(message)
}
