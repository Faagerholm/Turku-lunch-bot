package config

const BotWelcome = "Welcome to Turku lunch bot!" +
	" I will help you select which restaurant you should visit today." +
	" I can show you todays menu, just search for your desired restaurant."

type Restaurant struct {
	Url string
	Idx int
}

func RestaurantList() map[string]Restaurant {
	return map[string]Restaurant{
		"Arken": {"https://www.karkafeerna.fi/en/lunch", 0},
		"Gado":  {"https://www.karkafeerna.fi/en/lunch", 1},
		"Kåren": {"https://www.karkafeerna.fi/en/lunch", 2},
		"ASA":   {"https://www.karkafeerna.fi/en/lunch", 3},
	}
}
