package beta

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetM2a(ctx *gin.Context) {
	deckCode := ctx.Param("id")
	resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/" + deckCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusInternalServerError, "Failed to fetch deck data: "+resp.Status)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var deck []*Card
	if err := json.Unmarshal(body, &deck); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	cardlist := make(map[string]int)
	for _, card := range deck {
		_, ok := cardlist[card.Name]
		if !ok {
			cardlist[card.Name] = card.Count
		} else {
			cardlist[card.Name] += card.Count
		}
	}

	if deckType := analyzeGholdengo_ex(cardlist, deck); deckType != nil {
		ctx.JSON(http.StatusOK, deckType)
		return
	}

	ctx.JSON(http.StatusNoContent, DeckType{})
}
