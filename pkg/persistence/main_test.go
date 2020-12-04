package persistence

import (
	"github.com/covid-tracing-mongolia/backend-server/pkg/config"
	"github.com/covid-tracing-mongolia/backend-server/pkg/keyclaim"
	"os"
	"strings"
	"testing"
)

var (
	token1 = strings.Repeat("a", 20)
	token2 = strings.Repeat("b", 20)
	onApi = "ONApi"
)


// TestMain this gets called instead of the regular testing main method and allows us to run setup code
func TestMain(m *testing.M)  {

	// Initialise Authenticator object
	os.Setenv("KEY_CLAIM_TOKEN", token1 +"=" + onApi + ":"+ token2 +"=428")

	config.InitConfig()
	SetupLookup(keyclaim.NewAuthenticator())

	os.Exit(m.Run())
}