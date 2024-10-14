package cli

import (
	"fmt"

	"github.com/PandaPy/pginer/component"
)

var asciiArtTemplate = `
    ____  ___________   ____________
   / __ \/ ____/  _/ | / / ____/ __ \
  / /_/ / / __ / //  |/ / __/ / /_/ /
 / ____/ /_/ // // /|  / /___/ _, _/
/_/    \____/___/_/ |_/_____/_/ |_|     v%s
`

func NewVersionText(version string) string {
	versionText := fmt.Sprintf(asciiArtTemplate, version)
	return component.FgBlue.Width(0).Render(versionText)
}
