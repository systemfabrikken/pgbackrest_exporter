package styreplan

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func Registrer_styreplan_metrikker(logger log.Logger) {
	err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pgbackrest_azure_nokkelutloeper",
			Help: "Antall dager til azzure nøkkel utløper",
		},
		func() float64 {
			level.Debug(logger).Log("msg", "Samler data for azure nøkkel")
			return float64(time.Now().Unix())
		},
	))

	if err != nil {
		level.Error(logger).Log("msg", "Feil ved registrering av pgbackrest_azure_nokkelutloeper", "err", err)
	}

}
