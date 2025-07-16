package AiServices

import (
	"fmt"
)

func Reply(prompt string) (string, error) {
	return fmt.Sprintf("You said: %s. AI bot replied!", prompt), nil
}
