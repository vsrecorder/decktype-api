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

	"github.com/gin-contrib/cors"
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
	r.SetTrustedProxies(nil)
	r.Use(cors.New(cors.Config{
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Access-Control-Request-Method",
			"Authorization",
			"Content-Type",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowOrigins: []string{
			//"*",
			"http://localhost:3000",
			"https://local.vsrecorder.mobi",
			"https://decktype.vsrecorder.mobi",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

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

	if cardlist["タケルライコex"] >= 2 && cardlist["オーガポン みどりのめんex"] >= 3 {
		deckType := analyze(
			"タケルライコex",
			deck,
			[]string{
				"タケルライコex",
				"オーガポン みどりのめんex",
				"タケルライコ",
				"コライドン",
				"チヲハウハネ",
				"テツノイサハex",
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
				"ピジョットex",
				"ヨルノズク",
				"ヨノワール",
				"テラパゴスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ドラパルトex"] >= 2 && cardlist["ドロンチ"] >= 3 && cardlist["ドラメシヤ"] >= 3 {
		deckType := analyze(
			"ドラパルトex",
			deck,
			[]string{
				"ドラパルトex",
				"ヨノワール",
				"ピジョットex",
				"ロケット団のクロバットex",
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
				"ユキメノコ",
				"マシマシラ",
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

	if cardlist["ブリジュラスex"] >= 2 {
		deckType := analyze(
			"ブリジュラスex",
			deck,
			[]string{
				"ブリジュラスex",
				"ホップのバイウールー",
				"ノココッチ",
				"モモワロウ",
				"アラブルタケ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ソウブレイズex"] >= 2 {
		deckType := analyze(
			"ソウブレイズex",
			deck,
			[]string{
				"ソウブレイズex",
				"ノココッチ",
				"ブロロローム",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["サーフゴーex"] >= 2 {
		deckType := analyze(
			"サーフゴーex",
			deck,
			[]string{
				"サーフゴーex",
				"ドラパルトex",
				"ノココッチ",
				"ヒビキのバクフーン",
				"ハッサム",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["バシャーモex"] >= 2 {
		deckType := analyze(
			"バシャーモex",
			deck,
			[]string{
				"バシャーモex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ゲッコウガex"] >= 2 {
		deckType := analyze(
			"ゲッコウガex",
			deck,
			[]string{
				"ゲッコウガex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ダイゴのメタグロスex"] >= 2 {
		deckType := analyze(
			"ダイゴのメタグロスex",
			deck,
			[]string{
				"ダイゴのメタグロスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["パピナスex"] >= 2 {
		deckType := analyze(
			"パピナスex",
			deck,
			[]string{
				"パピナスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["サザンドラex"] >= 2 {
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
				"ナンジャモのタイカイデン",
				"ナンジャモのビリリダマ",
				"ミライドンex",
				"タケルライコex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヒビキのバクフーン"] >= 2 && cardlist["ヒビキの冒険"] == 4 {
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

	if cardlist["ロケット団のクロバットex"] >= 2 {
		deckType := analyze(
			"ロケット団のクロバットex",
			deck,
			[]string{
				"ロケット団のクロバットex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イーブイex"] >= 1 && cardlist["イーブイ"] >= 1 && (cardlist["ブースターex"] >= 1 || cardlist["シャワーズex"] >= 1 || cardlist["サンダースex"] >= 1 ||
		cardlist["エーフィex"] >= 1 || cardlist["ブラッキーex"] >= 1 ||
		cardlist["リーフィアex"] >= 1 || cardlist["グレイシアex"] >= 1 || cardlist["ニンフィアex"] >= 1) {
		deckType := analyze(
			"ブイズバレット",
			deck,
			[]string{
				"イーブイex",
				"ブースターex",
				"シャワーズex",
				"サンダースex",
				"エーフィex",
				"ブラッキーex",
				"リーフィアex",
				"グレイシアex",
				"ニンフィアex",
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
				"シロナのロズレイド",
				"シロナのミカルゲ",
				"ユキメノコ",
				"マシマシラ",
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
				"ロケット団のワナイダー",
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

	if cardlist["Nのゾロアークex"] >= 3 && (cardlist["Nのヒヒダルマ"] >= 2 || cardlist["Nのレシラム"] >= 1 || cardlist["Nのシンボラー"] >= 1) {
		deckType := analyze(
			"Nのゾロアークex",
			deck,
			[]string{
				"Nのゾロアークex",
				"Nのヒヒダルマ",
				"Nのレシラム",
				"Nのシンボラー",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヒビキのホウオウex"] >= 2 && cardlist["グレンアルマ"] == 0 {
		deckType := analyze(
			"ヒビキのホウオウex",
			deck,
			[]string{
				"ヒビキのホウオウex",
				"ヒビキのマグカルゴ",
				"ヒビキのカイロス",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヒビキのホウオウex"] >= 2 && cardlist["グレンアルマ"] >= 2 {
		deckType := analyze(
			"ひおくりバレット",
			deck,
			[]string{
				"ヒビキのホウオウex",
				"グレンアルマ",
				"オーガポン いどのめんex",
				"テツノカイナex",
				"リーリエのピッピex",
				"レジギガス",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ハルクジラex"] >= 2 {
		deckType := analyze(
			"ハルクジラex",
			deck,
			[]string{
				"ハルクジラex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["メガヤンマex"] >= 2 {
		deckType := analyze(
			"メガヤンマex",
			deck,
			[]string{
				"メガヤンマex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["マスカーニャex"] >= 2 {
		deckType := analyze(
			"マスカーニャex",
			deck,
			[]string{
				"マスカーニャex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヤドキング"] >= 3 && cardlist["夜のアカデミー"] >= 3 {
		deckType := analyze(
			"ヤドキング",
			deck,
			[]string{
				"ヤドキング",
				"キュレム",
				"ローブシン",
				"レジギガス",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イダイナキバ"] >= 3 && cardlist["ニュートラルセンター"] == 1 {
		deckType := analyze(
			"イダイナキバLO",
			deck,
			[]string{
				"イダイナキバ",
				"ヒビキのウソッキー",
				"クラッシュハンマー",
				"ニュートラルセンター",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["バンギラス"] >= 3 {
		deckType := analyze(
			"バンギラス",
			deck,
			[]string{
				"バンギラス",
				"ノココッチ",
				"ドロンチ",
				"ピジョットex",
				"シャンデラ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イルカマンex"] >= 2 {
		deckType := analyze(
			"イルカマンex",
			deck,
			[]string{
				"イルカマンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["アマージョex"] >= 2 {
		deckType := analyze(
			"アマージョex",
			deck,
			[]string{
				"アマージョex",
				"ユキメノコ",
				"マシマシラ",
				"ピジョットex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イワパレス"] >= 3 {
		deckType := analyze(
			"イワパレス",
			deck,
			[]string{
				"イワパレス",
				"テツノイバラex",
				"オーガポン いしずえのめんex",
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
				"レアコイル",
				"テツノカイナex",
				"ピカチュウex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["バチュル"] >= 2 && cardlist["テツノカイナex"] >= 1 && cardlist["ピカチュウex"] >= 1 {
		deckType := analyze(
			"バチュルバレット",
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
				"ホップのカビゴン",
				"ホップのウッウ",
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
				"ハバタクカミ",
				"イダイナキバ",
				"コライドン",
				"トドロクツキex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["カットロトム"] >= 1 && cardlist["ヒートロトム"] >= 1 && cardlist["ウォッシュロトム"] >= 1 && cardlist["ロトム"] >= 1 {
		deckType := analyze(
			"ロトムバレット",
			deck,
			[]string{
				"カットロトム",
				"ヒートロトム",
				"ウォッシュロトム",
				"ロトム",
				"スピンロトム",
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

	if cardlist["ミライドン"] >= 2 && cardlist["テツノカシラex"] >= 2 && cardlist["テクノレーダー"] >= 2 {
		deckType := analyze(
			"未来バレット",
			deck,
			[]string{
				"ミライドン",
				"テツノカシラex",
				"テツノカイナex",
				"テツノブジンex",
				"テツノイサハex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	ctx.JSON(http.StatusOK, deckTypes)
}
