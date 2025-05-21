package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

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

func main() {
	r := gin.Default()

	r.GET(
		"/decktypes/:id",
		Get,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    ":8930",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 3 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

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

func Get(ctx *gin.Context) {
	deckCode := ctx.Param("id")
	resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/" + deckCode)

	// タケルライコex
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/ngnLQQ-N9tdRR-ignLgn")

	// ドラパルトex
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/HLQgin-PvQExD-gQLgQg")

	// 悪リザードンex & ドラパルトex
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/vVkfVF-lxMyv8-VvwvFk")

	// バチュルバレット
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/SMpSpM-REy1lZ-EypSyX")

	// ロケット団のミューツーex
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/QLLngg-o6YtFY-LQQNgH")

	// 古代バレット
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/SypMyy-rw6242-XyU2pM")

	// オーダイル
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/VkwFFv-RlGSbH-5vk5vd")

	// ユキメノコ & マシマシラ
	//resp, err := http.Get("https://vsrecorder.mobi/api/v1/deckcards/6nQQgN-QhjfmY-nNgQnH")

	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	var deck []*Card
	if err := json.Unmarshal(body, &deck); err != nil {
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

	deckTypes := []*DeckType{}

	if cardlist["タケルライコex"] >= 3 {
		deckType := analyze(
			"タケルライコex",
			deck,
			[]string{
				"タケルライコex",
				"オーガポン みどりのめんex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["リザードンex"] >= 2 {
		deckType := analyze(
			"リザードンex",
			deck,
			[]string{
				"リザードンex",
				"ドラパルトex",
				"ピジョットex",
				"ヨルノズク",
				"ヨノワール",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ドラパルトex"] >= 2 && cardlist["ドロンチ"] >= 3 && cardlist["ドラメシヤ"] >= 3 && cardlist["リザードンex"] == 0 {
		deckType := analyze(
			"ドラパルトex",
			deck,
			[]string{
				"ドラパルトex",
				"ヨノワール",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["マリィのオーロンゲex"] >= 2 {
		deckType := analyze(
			"マリィのオーロンゲex",
			deck,
			[]string{
				"マリィのオーロンゲex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["サーナイトex"] >= 2 {
		deckType := analyze(
			"サーナイトex",
			deck,
			[]string{
				"サーナイトex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ブリジュラスex"] >= 3 {
		deckType := analyze(
			"ブリジュラスex",
			deck,
			[]string{
				"ブリジュラスex",
				"ホップのバイウールー",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ソウブレイズex"] >= 3 {
		deckType := analyze(
			"ソウブレイズex",
			deck,
			[]string{
				"ソウブレイズex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["サーフゴーex"] >= 3 {
		deckType := analyze(
			"サーフゴーex",
			deck,
			[]string{
				"サーフゴーex",
				"ドラパルトex",
				"ノココッチ",
				"ヒビキのバクフーン",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["バシャーモex"] >= 3 {
		deckType := analyze(
			"バシャーモex",
			deck,
			[]string{
				"バシャーモex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ダイゴのメタグロスex"] >= 3 {
		deckType := analyze(
			"ダイゴのメタグロスex",
			deck,
			[]string{
				"ダイゴのメタグロスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["パピナスex"] >= 3 {
		deckType := analyze(
			"パピナスex",
			deck,
			[]string{
				"パピナスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["サザンドラex"] >= 3 {
		deckType := analyze(
			"サザンドラex",
			deck,
			[]string{
				"サザンドラex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ナンジャモのハラバリーex"] >= 2 {
		deckType := analyze(
			"ナンジャモのハラバリーex",
			deck,
			[]string{
				"ナンジャモのハラバリーex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヒビキのバクフーン"] >= 3 && cardlist["ヒビキの冒険"] == 4 {
		deckType := analyze(
			"ヒビキのバクフーン",
			deck,
			[]string{
				"ヒビキのバクフーン",
				"ヒビキの冒険",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["カミツオロチex"] >= 2 {
		deckType := analyze(
			"カミツオロチex",
			deck,
			[]string{
				"カミツオロチex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ブースターex"] >= 2 {
		deckType := analyze(
			"ブースターex",
			deck,
			[]string{
				"ブースターex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["シロナのガブリアスex"] >= 2 {
		deckType := analyze(
			"シロナのガブリアスex",
			deck,
			[]string{
				"シロナのガブリアスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["オーダイル"] >= 2 {
		deckType := analyze(
			"オーダイル",
			deck,
			[]string{
				"オーダイル",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["クエスパトラex"] >= 2 {
		deckType := analyze(
			"クエスパトラex",
			deck,
			[]string{
				"クエスパトラex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イイネイヌ"] == 4 {
		deckType := analyze(
			"イイネイヌ",
			deck,
			[]string{
				"イイネイヌ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のミュウツーex"] >= 2 && cardlist["ロケット団のワナイダー"] >= 3 {
		deckType := analyze(
			"ロケット団のミュウツーex",
			deck,
			[]string{
				"ロケット団のミュウツーex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["パオジアンex"] >= 2 && cardlist["セグレイブ"] >= 2 {
		deckType := analyze(
			"パオジアンex",
			deck,
			[]string{
				"パオジアンex",
				"セグレイブ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["トドロクツキex"] >= 2 && cardlist["イダイナキバ"] == 0 && cardlist["コライドン"] == 0 {
		deckType := analyze(
			"トドロクツキex",
			deck,
			[]string{
				"トドロクツキex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["Nのゾロアークex"] >= 3 && cardlist["Nのヒヒダルマ"] >= 2 && cardlist["Nのレシラム"] >= 1 {
		deckType := analyze(
			"Nのゾロアークex",
			deck,
			[]string{
				"Nのゾロアークex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ミライドンex"] >= 2 && cardlist["バチュル"] == 0 {
		deckType := analyze(
			"ミライドンex",
			deck,
			[]string{
				"ミライドンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["バチュル"] >= 2 && cardlist["テツノカイナex"] >= 1 && cardlist["ピカチュウex"] >= 1 {
		deckType := analyze(
			"バチュル",
			deck,
			[]string{
				"バチュル",
				"ミライドンex",
				"テツノカイナex",
				"ピカチュウex",
				"テツノイサハex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ホップのザシアンex"] >= 2 {
		deckType := analyze(
			"ホップのザシアンex",
			deck,
			[]string{
				"ホップのザシアンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["オリーヴァex"] >= 2 {
		deckType := analyze(
			"オリーヴァex",
			deck,
			[]string{
				"オリーヴァex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のポリゴンZ"] >= 3 {
		deckType := analyze(
			"ロケット団のポリゴンZ",
			deck,
			[]string{
				"ロケット団のポリゴンZ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ミロカロスex"] >= 2 {
		deckType := analyze(
			"ミロカロスex",
			deck,
			[]string{
				"ミロカロスex",
				"リキキリンex",
				"オンバーンex",
				"オーガポン いしずえのめんex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ユキメノコ"] >= 3 && cardlist["マシマシラ"] >= 3 && cardlist["マリィのオーロンゲex"] == 0 {
		deckType := analyze(
			"ユキメノコ & マシマシラ",
			deck,
			[]string{
				"ユキメノコ",
				"マシマシラ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["トドロクツキ"] == 4 && (cardlist["ハバタクカミ"] >= 1 || cardlist["イダイナキバ"] >= 1 || cardlist["コライドン"] >= 1) && cardlist["オーリム博士の気迫"] == 4 && cardlist["探検家の先導"] == 4 {
		deckType := analyze(
			"古代バレット",
			deck,
			[]string{
				"トドロクツキ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["カミッチュ"] >= 3 && cardlist["バチンキー"] >= 3 && cardlist["お祭り会場"] >= 3 {
		deckType := analyze(
			"おまつりおんど",
			deck,
			[]string{
				"カミッチュ",
				"バチンキー",
				"お祭り会場",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	ctx.JSON(http.StatusOK, deckTypes)
}
