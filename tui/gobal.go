package tui

import (
	"fmt"
)

var version = "1.0.1"

var asciiArt = fmt.Sprintf(`
    ____  ___________   ____________
   / __ \/ ____/  _/ | / / ____/ __ \
  / /_/ / / __ / //  |/ / __/ / /_/ /
 / ____/ /_/ // // /|  / /___/ _, _/
/_/    \____/___/_/ |_/_____/_/ |_|  	v%s

`, version)

// var titleColor = color.New(color.FgBlue).Add(color.Bold)

// var questionMarkColor = color.New(color.FgGreen)
// var promptTextColor = color.New(color.FgWhite).Add(color.Bold)
// var optionTextColor = color.New(color.FgWhite)
// var optionTextLightColor = color.New(color.FgWhite)
// var helpTextColor = color.New(color.FgHiBlack)
