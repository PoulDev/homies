package avatar

import (
	"fmt"
	"math/rand/v2"
	"math"

	"github.com/zibbadies/homies/internal/homies/models"
)

type RGBColor struct {
	R int
	G int
	B int
}

type HSLColor struct {
	H float64
	S float64
	L float64
}

func HSL2RGB(color HSLColor) RGBColor {
	var rp, gp, bp float64;
	C := (1.0 - math.Abs(2.0 * color.L - 1.0)) * color.S;
	X := C * (1.0 - math.Abs(math.Mod(color.H*6, 2) - 1.0));
	m := color.L - C/2;

    if 0 <= color.H && color.H < 1.0/6.0 {
        rp, gp, bp = C, X, 0
    } else if 1.0/6.0 <= color.H && color.H < 2.0/6.0 {
        rp, gp, bp = X, C, 0
    } else if 2.0/6.0 <= color.H && color.H < 3.0/6.0 {
        rp, gp, bp = 0, C, X
    } else if 3.0/6.0 <= color.H && color.H < 4.0/6.0 {
        rp, gp, bp = 0, X, C
    } else if 4.0/6.0 <= color.H && color.H < 5.0/6.0 {
        rp, gp, bp = X, 0, C
    } else {
        rp, gp, bp = C, 0, X
    }

	return RGBColor{
		int((rp + m) * 255),
		int((gp + m) * 255),
		int((bp + m) * 255),
	}
}

func isDark(color RGBColor) bool {
	luminance := float64(0.2126)*float64(color.R) + float64(0.7152)*float64(color.G) + float64(0.0722)*float64(color.B);
	return luminance <= 190;
}

func color2hex(color RGBColor) string {
	return fmt.Sprintf("%x%x%x", color.R, color.G, color.B);
}

func RandAvatar() models.Avatar {
	hslColor := HSLColor{
		rand.Float64(),
		0.70,
		0.80,
	}

	rgbColor := HSL2RGB(hslColor);
	color := color2hex(rgbColor);

	firstPoint := fmt.Sprintf("%d %d", rand.IntN(4), 1+rand.IntN(2));
	secondPoint := fmt.Sprintf("%d %d", rand.IntN(4)+3, 1+rand.IntN(2));

	eyesSpace := (rand.Float32()*8 + 4) / 2;

	var faceColor string;
	if isDark(rgbColor) {
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
