package types


type ColorScheme struct {
	PrimaryColor Color
	AccentColor Color
	PrimaryColorHighlight Color
	AccentColorHighlight Color	
}

func DefaultColorScheme() ColorScheme {
	return ColorScheme{
		PrimaryColor: FromHexString("#303030E0"),
		AccentColor: Color{181, 165, 166, 255},
		// AccentColor: Color{139, 139, 139, 255},
		PrimaryColorHighlight: FromHexString("#404040"),
		AccentColorHighlight: Color{181, 165, 166, 255},
	}
}