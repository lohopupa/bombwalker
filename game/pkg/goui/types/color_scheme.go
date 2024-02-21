package types


type ColorScheme struct {
	PrimaryColor Color
	AccentColor Color
	PrimaryColorHighlight Color
	AccentColorHighlight Color	
}

func DefaultColorScheme() ColorScheme {
	return ColorScheme{
		PrimaryColor: FromHexString("#303030"),
		AccentColor: Color{139, 139, 139, 255},
		PrimaryColorHighlight: FromHexString("#404040"),
		AccentColorHighlight: Color{150, 150, 150, 255},
	}
}