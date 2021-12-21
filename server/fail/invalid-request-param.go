package fail

import (
	"errors"
	"fmt"
)

func InvalidRequestParam(param string) error {
	message := fmt.Sprintf("invalid request param: %s", param)
	return errors.New(message)
}
