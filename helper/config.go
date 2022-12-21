package helper

type FontConfig struct {
	FontSize   float64
	BaseX      float64
	BaseY      float64
	IncrementY float64
}

func (f *FontConfig) ToTitleConfig(titleLength int) *FontConfig {
	if titleLength == 1 {
		f.FontSize = 72
		f.BaseX = 300
		f.BaseY = 116
	} else {
		f.FontSize = 46
		f.BaseX = 300
		f.BaseY = 84
		f.IncrementY = 48
	}
	return f
}

func (f *FontConfig) ToTextConfig() *FontConfig {
	f.FontSize = 56
	f.BaseX = 110
	f.BaseY = 268
	f.IncrementY = 80
	return f
}
