package cli

import (
	"io"

	"github.com/solbero/tempconv/scale"
)

type config struct {
	temp      float64
	fromScale scale.Scale
	toScale   scale.Scale
	decimal   int
	unit      bool
	version   bool
	help      bool
}

const usageMsg = "try 'tempconv -h' for more information"

func fprinte(w io.Writer, msg string) {
	w.Write([]byte(msg + "\n" + usageMsg))
}
