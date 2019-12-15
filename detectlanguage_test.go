package detectlanguage_test

import (
	"os"

	"github.com/detectlanguage/detectlanguage-go"
)

var client = detectlanguage.New(os.Getenv("DETECTLANGUAGE_API_KEY"))
