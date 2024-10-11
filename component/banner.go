package component

import (
	"github.com/PandaPy/pginer/config"
	"github.com/fatih/color"
)

func Banner() string {
	return Text(config.BannerText).Color(color.FgBlue).Bold().String()
}
