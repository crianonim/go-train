package main

import "fmt"

func main() {
	fmt.Println("Welcome to Star Trader")
	player := Entity{name: "Jan", money: 12}
	baron := Entity{name: "Baron"}
	station := Place{name: "Station One"}
	station.put(&baron)
	station.put(&player)
	dock := Place{name: "Dock"}
	market := Place{name: "Market"}
	station.places = append(station.places, &dock)
	station.places = append(station.places, &market)
	fmt.Print(placeString(&station))
	fmt.Println(player, baron, station)
}
