package styreplan

import (
	"fmt"
	"github.com/prometheus/common/promlog"
	"os"
	"testing"
)

var promlogConfig = &promlog.Config{}
var logger = promlog.New(promlogConfig)

func TestBeregnDagerTilTokenUtløper(t *testing.T) {
	azureNøkkel := "ett&eller=annet&se=2023-04-13T19:22:04Z&s=sdfsdfsdf&"
	os.Setenv(AZURE_TOKEN_ENV_NØKKEL, azureNøkkel)

	dagerTilUtløpt := beregnDagerTilTokenUtløper(logger)
	if dagerTilUtløpt != 212 {
		t.Errorf("\nforventet at dagerTilUtløpt matchet, var: %d", dagerTilUtløpt)
	}

	fmt.Println("dagerTilUtløpt:", dagerTilUtløpt)
}

func TestBeregnDagerTilTokenUtløperMedUtløptToken(t *testing.T) {
	azureNøkkel := "ett&eller=annet&se=2022-04-13T19:22:04Z&s=sdfsdfsdf&"
	os.Setenv(AZURE_TOKEN_ENV_NØKKEL, azureNøkkel)

	dagerTilUtløpt := beregnDagerTilTokenUtløper(logger)
	if dagerTilUtløpt != -152 {
		t.Errorf("\nforventet at dagerTilUtløpt matchet, var: %d", dagerTilUtløpt)
	}

	fmt.Println("dagerTilUtløpt:", dagerTilUtløpt)
}
