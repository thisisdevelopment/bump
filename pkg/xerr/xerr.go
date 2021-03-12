package xerr

import (
	"os"

	"github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

// Exitif err with printf format
func Exitif(err error, format string, ctx ...interface{}) {
	if err != nil {
		if len(ctx) > 0 {
			format = aurora.Sprintf(format, ctx...)
		}
		log.Error(aurora.Sprintf("%v %s", aurora.BrightRed(err), aurora.Yellow(format)))
		os.Exit(-1)
	}
}
