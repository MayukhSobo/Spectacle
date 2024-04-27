package app

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"strconv"
)

var _banner = `
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

func Banner() string {
	return _banner
}

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

func makeRampStyles(bannerProp *BannerStyleProperties) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(bannerProp.BannerGradientStartColor)
	cB, _ := colorful.Hex(bannerProp.BannerGradientEndColor)
	steps := float64(len(_banner))
	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}

func BannerRendered() string {
	var bannerRendered string
	for i, each := range homeScreenStyle.BannerStyle {
		bannerRendered += each.Render(string(_banner[i]))
	}
	bannerRendered += "\n"
	return bannerRendered
}
