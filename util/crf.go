package util

func GetCrf(code string, width, height int) (crf string) {
	switch code {
	case "h265":
		if width == 1080 || height == 1080 {
			crf = "22"
		} else if width == 720 || height == 720 {
			crf = "23"
		} else if width == 480 || height == 480 {
			crf = "24"
		}
	case "vp9":
		if width == 1080 || height == 1080 {
			crf = "31"
		} else if width == 720 || height == 720 {
			crf = "32"
		} else if width == 480 || height == 480 {
			crf = "33"
		}
	case "av1":
		if width == 1080 || height == 1080 {
			crf = "24"
		} else if width == 720 || height == 720 {
			crf = "24"
		} else if width == 480 || height == 480 {
			crf = "24"
		}
	default:
		panic("不是合法的编码")
	}
	return crf
}
