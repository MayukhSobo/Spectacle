package home

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

var banner = `
.d8888b.                             888                      888
d88P  Y88b                            888                      888
Y88b.                                 888                      888
 "Y888b.   88888b.   .d88b.   .d8888b 888888  8888b.   .d8888b 888  .d88b.
    "Y88b. 888 "88b d8P  Y8b d88P"    888        "88b d88P"    888 d8P  Y8b
      "888 888  888 88888888 888      888    .d888888 888      888 88888888
Y88b  d88P 888 d88P Y8b.     Y88b.    Y88b.  888  888 Y88b.    888 Y8b.
 "Y8888P"  88888P"   "Y8888   "Y8888P  "Y888 "Y888888  "Y8888P 888  "Y8888
           888
           888
           888
`

func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}

func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}

func makeRampStyles(colorStart, colorEnd string, banner string) (s []lipgloss.Style) {
	cA, err := colorful.Hex(colorStart)
	if err != nil {
		panic(err)
	}
	cB, err := colorful.Hex(colorEnd)
	if err != nil {
		panic(err)
	}
	steps := float64(len(banner))
	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}

func GradientBanner(banner *Banner) string {
	var bannerRendered string
	for i, each := range makeRampStyles(
		banner.BannerStatingColor,
		banner.BannerEndingColor,
		banner.BannerText,
	) {
		bannerRendered += each.Render(string(banner.BannerText[i]))
	}
	bannerRendered += "\n"
	return bannerRendered
}
