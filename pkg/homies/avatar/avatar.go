package avatar

import (
	"fmt"
	"math/rand/v2"

	"github.com/zibbadies/homies/internal/homies/models"
)

type Color struct {
	R int
	G int
	B int
}

var colors = []Color{
	{255, 212, 120},
	{255, 125, 122},
	{122, 164, 255},
	{122, 209, 255},
	{122, 255, 213},
	{213, 122, 255},
	{164, 255, 122},
	{122, 164, 255},
	{213, 122, 255},
	{122, 255, 147},
	{122, 164, 255},
	{255, 122, 231},
}

func isDark(color Color) bool {
	luminance := float64(0.2126)*float64(color.R) + float64(0.7152)*float64(color.G) + float64(0.0722)*float64(color.B);
	return luminance <= 190;
}

func changeColor(v int) int {
	val := float64(v) + rand.Float64()*30.0 - 15
	if val > 255 {
		val = 255;
	} else if v < 0 {
		val = 0;
	}
	return int(val);
}

func color2hex(color Color) string {
	return fmt.Sprintf("%x%x%x", color.R, color.G, color.B);
}

func RandAvatar() models.Avatar {
	oColor := colors[rand.IntN(len(colors))];
	RGBcolor := Color{
		changeColor(oColor.R),
		changeColor(oColor.G),
		changeColor(oColor.B),
	}

	color := color2hex(RGBcolor);

	firstPoint := fmt.Sprintf("%d %d", rand.IntN(4), 1+rand.IntN(2));
	secondPoint := fmt.Sprintf("%d %d", rand.IntN(4)+3, 1+rand.IntN(2));

	eyesSpace := (rand.Float32()*8 + 4) / 2;

	var faceColor string;
	if isDark(RGBcolor) {
		faceColor = "EFEFEF";
	} else {
		faceColor = "010101";
	}

	return models.Avatar{
		BgColor:   color,
		FaceColor: faceColor,
		FaceX:     rand.Float32()*39 + 7,
		FaceY:     rand.Float32()*40 + 7,
		LeX:       6 - eyesSpace,
		LeY:       2 + rand.Float32()*4,
		ReX:       6 + eyesSpace,
		ReY:       2 + rand.Float32()*4,
		Bezier:    fmt.Sprintf("%s %s 6 0", firstPoint, secondPoint),
	}
}
