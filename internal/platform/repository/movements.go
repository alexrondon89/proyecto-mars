package repository

type Movement map[string]string

func GetOrientationRepository() []Movement {
	return []Movement{
		{
			"orientation": "N",
			"turn":        "L",
			"result":      "W",
		},
		{
			"orientation": "N",
			"turn":        "R",
			"result":      "E",
		},
		{
			"orientation": "W",
			"turn":        "L",
			"result":      "S",
		},
		{
			"orientation": "W",
			"turn":        "R",
			"result":      "N",
		},
		{
			"orientation": "E",
			"turn":        "L",
			"result":      "N",
		},
		{
			"orientation": "E",
			"turn":        "R",
			"result":      "S",
		},
		{
			"orientation": "S",
			"turn":        "R",
			"result":      "W",
		},
		{
			"orientation": "S",
			"turn":        "L",
			"result":      "E",
		},
	}
}
