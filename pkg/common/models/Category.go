package models

type Category struct {
	Name        string
	Path        string
	CssClass    string
	Description string
	Icon        string
}

var Categories = map[string]*Category{
	"event": {
		Name:        "Events",
		Path:        "event",
		CssClass:    "tmbEvents",
		Description: "Latest events which may attract your attention.",
		Icon:        "event",
	},
	"game": {
		Name:     "Games",
		Path:     "game",
		CssClass: "tmbGames",
		Description: "Computer games of all kinds and for all platforms, " +
			"including consoles, portable gaming devices and smartphones.",
		Icon: "game",
	},
	"hard": {
		Name:     "Hardware",
		Path:     "hard",
		CssClass: "tmbHardware",
		Description: "Various computer hardware, from cases and power supply units " +
			"to microprocessors and GPUs.",
		Icon: "hard",
	},
	"life": {
		Name:     "Life",
		Path:     "life",
		CssClass: "tmbLife",
		Description: "Everything about life in the universe. [br]" +
			"From flowers and bees to animals and humanoids.",
		Icon: "life",
	},
	"media": {
		Name:     "Multimedia",
		Path:     "media",
		CssClass: "tmbMultimedia",
		Description: "Audio, video, animation, images, photography, literature. " +
			"Anything from a short ringtone sound to a full-length feature film.",
		Icon: "media",
	},
	"motor": {
		Name:     "Motorsport",
		Path:     "motor",
		CssClass: "tmbMotorsport",
		Description: "All kinds of automobile sports. From kart and buggy " +
			"to IndyCar and Formula One.",
		Icon: "motor",
	},
	"news": {
		Name:        "News",
		Path:        "news",
		CssClass:    "tmbNews",
		Description: "All kinds of news all around the world of planet Earth.",
		Icon:        "news",
	},
	"review": {
		Name:        "Reviews",
		Path:        "review",
		CssClass:    "tmbReviews",
		Description: "Reviews of all kinds. From bread toasters to supercomputers.",
		Icon:        "review",
	},
	"soft": {
		Name:        "Software",
		Path:        "soft",
		CssClass:    "tmbSoftware",
		Description: "Software of all kinds.",
		Icon:        "soft",
	},
	"tech": {
		Name:        "Technology",
		Path:        "tech",
		CssClass:    "tmbTechnology",
		Description: "Everything about past, modern and future technologies.",
		Icon:        "tech",
	},
}
