package handlers

type Card struct {
	Name      string `json:"name"`
	DetailURL string `json:"detail_url"`
	ImageURL  string `json:"image_url"`
	Count     int    `json:"count"`
}

type MainCard struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type DeckType struct {
	Title     string      `json:"title"`
	MainCards []*MainCard `json:"main_cards"`
}

func analyze(title string, deck []*Card, cards []string) *DeckType {
	var mainCards []*MainCard

	for _, cardname := range cards {
		for _, card := range deck {
			if card.Name == cardname {
				mainCards = append(
					mainCards,
					&MainCard{
						Name:     card.Name,
						ImageURL: card.ImageURL,
					},
				)
				break
			}
		}
	}

	deckType := &DeckType{
		Title:     title,
		MainCards: mainCards,
	}

	return deckType
}
