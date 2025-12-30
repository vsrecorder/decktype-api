package beta

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Card struct {
	Name      string `json:"name"`
	DetailURL string `json:"detail_url"`
	ImageURL  string `json:"image_url"`
	Count     int    `json:"count"`
}

type AcespecCard struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type DeckCard struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type DeckType struct {
	MainTitle   string       `json:"main_title"`
	SubTitle    string       `json:"sub_title"`
	MainCards   []*DeckCard  `json:"main_cards"`
	SubCards    []*DeckCard  `json:"sub_cards"`
	AcespecCard *AcespecCard `json:"acespec_card"`
}

func getAcespecCard(ctx *gin.Context, deckCode string) *AcespecCard {
	resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/" + deckCode + "/acespec")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusInternalServerError, "Failed to fetch deck data: "+resp.Status)
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return nil
	}

	var acespecCard *AcespecCard
	if err := json.Unmarshal(body, acespecCard); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return nil
	}

	return acespecCard
}

func analyzeJoltik(cardlist map[string]int, deck []*Card) *DeckType {
	if cardlist["バチュル"] >= 2 && cardlist["サーフゴーex"] >= 3 {
		var mainTitle string = "バチュル&サーフゴーex"
		var subTitle string
		var mainCards []*DeckCard
		var subCards []*DeckCard

		cards := []string{"バチュル", "サーフゴーex"}
		for _, cardname := range cards {
			for _, card := range deck {
				if card.Name == cardname {
					mainCards = append(
						mainCards,
						&DeckCard{
							Name:     card.Name,
							ImageURL: card.ImageURL,
						},
					)
					break
				}
			}
		}

		decktype := &DeckType{
			MainTitle: mainTitle,
			MainCards: mainCards,
			SubTitle:  subTitle,
			SubCards:  subCards,
		}

		return decktype
	} else if cardlist["バチュル"] >= 2 && cardlist["サーフゴーex"] == 0 {
		var mainTitle string = "バチュルバレット"
		var subTitle string
		var mainCards []*DeckCard
		var subCards []*DeckCard

		cards := []string{"バチュル"}
		for _, cardname := range cards {
			for _, card := range deck {
				if card.Name == cardname {
					mainCards = append(
						mainCards,
						&DeckCard{
							Name:     card.Name,
							ImageURL: card.ImageURL,
						},
					)
					break
				}
			}
		}

		decktype := &DeckType{
			MainTitle: mainTitle,
			MainCards: mainCards,
			SubTitle:  subTitle,
			SubCards:  subCards,
		}

		return decktype
	}

	return nil
}

func analyzeGholdengo_ex(cardlist map[string]int, deck []*Card) *DeckType {
	if cardlist["サーフゴーex"] >= 3 && cardlist["バチュル"] == 0 {
		var mainTitle string = "サーフゴーex"
		var subTitle string
		var mainCards []*DeckCard
		var subCards []*DeckCard

		cards := []string{"サーフゴーex"}
		for _, cardname := range cards {
			for _, card := range deck {
				if card.Name == cardname {
					mainCards = append(
						mainCards,
						&DeckCard{
							Name:     card.Name,
							ImageURL: card.ImageURL,
						},
					)
					break
				}
			}
		}

		if cardlist["ルナトーン"] >= 2 && cardlist["ソルロック"] >= 2 {
			cards = []string{"ルナトーン", "ソルロック"}
			for _, cardname := range cards {
				for _, card := range deck {
					if card.Name == cardname {
						subCards = append(
							subCards,
							&DeckCard{
								Name:     card.Name,
								ImageURL: card.ImageURL,
							},
						)
						break
					}
				}
			}
			subTitle = "ルナトーン/ソルロック"
		}

		decktype := &DeckType{
			MainTitle: mainTitle,
			MainCards: mainCards,
			SubTitle:  subTitle,
			SubCards:  subCards,
		}

		return decktype
	}

	return nil
}
