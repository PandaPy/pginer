package config

import (
	"fmt"
)

var Version = "1.0.5"

var BannerText = fmt.Sprintf(`
    ____  ___________   ____________
   / __ \/ ____/  _/ | / / ____/ __ \
  / /_/ / / __ / //  |/ / __/ / /_/ /
 / ____/ /_/ // // /|  / /___/ _, _/
/_/    \____/___/_/ |_/_____/_/ |_|  	v%s

`, Version)
