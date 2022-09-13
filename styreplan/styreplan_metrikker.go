package styreplan

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"regexp"
	"time"
)

const AZURE_TOKEN_ENV_NØKKEL = "AZURE_SAS_TOKEN"
const layout = "2006-01-02T15:04:05Z"

var /* const */ r = regexp.MustCompile("se=([0-9\\-]+T[0-9:]+Z)")

func Registrer_styreplan_metrikker(logger log.Logger) {
	err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pgbackrest_azuretoken_dagertilutloeper",
			Help: "Antall dager til azure sas token lest fra env.var " + AZURE_TOKEN_ENV_NØKKEL + " utløper",
		},
		func() float64 {
			level.Debug(logger).Log("msg", "Samler data for azure nøkkel")
			dager := beregnDagerTilTokenUtløper(logger)
			return float64(dager)
		},
	))

	if err != nil {
		level.Error(logger).Log("msg", "Feil ved registrering av pgbackrest_azure_nokkelutloeper", "err", err)
	}
}

func beregnDagerTilTokenUtløper(logger log.Logger) int {
	azureNøkkel := lesNøkkelFraMiljøvar()
	if !r.MatchString(azureNøkkel) {
		level.Debug(logger).Log("msg", "regexp matchet ikke nøkkel")
		return -1
	}
	match := r.FindStringSubmatch(azureNøkkel)
	if len(match) != 2 {
		level.Debug(logger).Log("msg", "forventet at regexp ga array med to innslag")
		return -1
	}
	dato, err := parseDate(match[1])
	if err != nil {
		level.Debug(logger).Log("msg", "feil under parsing av dato", "dato", match[1], "err", err)
		return -1
	}
	return dagerTil(dato)
}

func parseDate(value string) (time.Time, error) {
	valueReturn, err := time.Parse(layout, value)
	return valueReturn, err
}

func dagerTil(dato time.Time) int {
	dager := dato.Sub(time.Now()).Hours() / 24
	return int(dager)
}

func lesNøkkelFraMiljøvar() string {
	nøkkel := os.Getenv(AZURE_TOKEN_ENV_NØKKEL)
	return nøkkel
}
