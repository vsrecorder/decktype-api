package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSV10(ctx *gin.Context) {
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

	deckTypes := []*DeckType{}

	if cardlist["タケルライコex"] >= 2 && cardlist["オーガポン みどりのめんex"] >= 2 {
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

	if cardlist["ダイオウドウex"] >= 2 {
		deckType := analyze(
			"ダイオウドウex",
			deck,
			[]string{
				"ダイオウドウex",
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

	if cardlist["ハピナスex"] >= 2 {
		deckType := analyze(
			"ハピナスex",
			deck,
			[]string{
				"ハピナスex",
				"マシマシラ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ガオガエンex"] >= 2 {
		deckType := analyze(
			"ガオガエンex",
			deck,
			[]string{
				"ガオガエンex",
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

	if cardlist["スコヴィランex"] >= 3 {
		deckType := analyze(
			"スコヴィランex",
			deck,
			[]string{
				"スコヴィランex",
				"オーガポン みどりのめんex",
				"ユキメノコ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if (cardlist["イーブイex"] >= 1 || cardlist["イーブイ"] >= 1) && (cardlist["ブースターex"] >= 1 || cardlist["シャワーズex"] >= 1 || cardlist["サンダースex"] >= 1 ||
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

	if cardlist["イイネイヌ"] >= 3 {
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

	if cardlist["ロケット団のバンギラス"] >= 2 {
		deckType := analyze(
			"ロケット団のバンギラス",
			deck,
			[]string{
				"ロケット団のバンギラス",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のデンリュウ"] >= 2 {
		deckType := analyze(
			"ロケット団のデンリュウ",
			deck,
			[]string{
				"ロケット団のデンリュウ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のペルシアンex"] >= 2 {
		deckType := analyze(
			"ロケット団のペルシアンex",
			deck,
			[]string{
				"ロケット団のペルシアンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のニドキングex"] >= 2 {
		deckType := analyze(
			"ロケット団のニドキングex",
			deck,
			[]string{
				"ロケット団のニドキングex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のニドクイン"] >= 2 {
		deckType := analyze(
			"ロケット団のニドクイン",
			deck,
			[]string{
				"ロケット団のニドクイン",
				"ニドキング",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のアーボック"] >= 2 {
		deckType := analyze(
			"ロケット団のアーボック",
			deck,
			[]string{
				"ロケット団のアーボック",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ロケット団のファイヤーex"] >= 2 {
		deckType := analyze(
			"ロケット団のファイヤーex",
			deck,
			[]string{
				"ロケット団のファイヤーex",
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

	if cardlist["テラパゴスex"] >= 3 {
		deckType := analyze(
			"テラパゴスex",
			deck,
			[]string{
				"テラパゴスex",
				"ヨルノズク",
				"バッフロン",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヨノワール"] >= 3 && cardlist["サマヨール"] >= 3 && cardlist["ヨマワル"] >= 3 {
		deckType := analyze(
			"カースドボム",
			deck,
			[]string{
				"ヨノワール",
				"サマヨール",
				"ヨマワル",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["トドロクツキex"] >= 2 && cardlist["トドロクツキ"] <= 2 && (cardlist["モモワロウ"] == 0 && cardlist["アラブルタケ"] == 0) {
		deckType := analyze(
			"トドロクツキex",
			deck,
			[]string{
				"トドロクツキex",
				"トドロクツキ",
				"モモワロウ",
				"アラブルタケ",
				"危険な密林",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["トドロクツキ"] == 4 && (cardlist["イダイナキバ"] >= 1 || cardlist["コライドン"] >= 1) && cardlist["オーリム博士の気迫"] == 4 && cardlist["探検家の先導"] >= 3 {
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

	if (cardlist["トドロクツキex"] >= 2 || cardlist["トドロクツキ"] >= 2) && cardlist["モモワロウ"] >= 2 && cardlist["アラブルタケ"] >= 2 && cardlist["オーリム博士の気迫"] == 4 && cardlist["危険な密林"] >= 3 {
		deckType := analyze(
			"毒トドロクツキ",
			deck,
			[]string{
				"トドロクツキex",
				"トドロクツキ",
				"モモワロウ",
				"アラブルタケ",
				"危険な密林",
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

	if cardlist["マンムーex"] >= 2 {
		deckType := analyze(
			"マンムーex",
			deck,
			[]string{
				"マンムーex",
				"ピジョットex",
				"キョジオーン",
				"バシャーモex",
				"ガブリアスex",
				"レントラーex",
				"レントラー",
				"ヨノワール",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ウガツホムラex"] >= 2 {
		deckType := analyze(
			"ウガツホムラex",
			deck,
			[]string{
				"ウガツホムラex",
				"トドロクツキex",
				"モモワロウ",
				"アラブルタケ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヤバソチャex"] >= 1 || cardlist["ヤバソチャ"] >= 2 {
		deckType := analyze(
			"ヤバソチャex",
			deck,
			[]string{
				"ヤバソチャex",
				"ヤバソチャ",
				"オーガポン みどりのめんex",
				"テツノイサハex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["デスカーンex"] >= 2 {
		deckType := analyze(
			"デスカーンex",
			deck,
			[]string{
				"デスカーンex",
				"ノココッチ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["フーディンex"] >= 2 {
		deckType := analyze(
			"フーディンex",
			deck,
			[]string{
				"フーディンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ペンドラー"] >= 2 {
		deckType := analyze(
			"ペンドラー",
			deck,
			[]string{
				"ペンドラー",
				"モモワロウ",
				"アラブルタケ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["レントラーex"] >= 3 {
		deckType := analyze(
			"レントラーex",
			deck,
			[]string{
				"レントラーex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["エースバーンex"] >= 2 {
		deckType := analyze(
			"エースバーンex",
			deck,
			[]string{
				"エースバーンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["エレキブルex"] >= 2 {
		deckType := analyze(
			"エレキブルex",
			deck,
			[]string{
				"エレキブルex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ビークインex"] >= 2 {
		deckType := analyze(
			"ビークインex",
			deck,
			[]string{
				"ビークインex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["キョジオーン"] >= 2 {
		deckType := analyze(
			"キョジオーン",
			deck,
			[]string{
				"キョジオーン",
				"ピジョットex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["デカヌチャンex"] >= 2 {
		deckType := analyze(
			"デカヌチャンex",
			deck,
			[]string{
				"デカヌチャンex",
				"ノココッチ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ブーバーン"] >= 3 && cardlist["ボルケニオンex"] >= 2 {
		deckType := analyze(
			"ブーバーン & ボルケニオンex",
			deck,
			[]string{
				"ブーバーン",
				"ボルケニオンex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ルガルガン"] >= 3 && cardlist["スパイクエネルギー"] >= 3 {
		deckType := analyze(
			"ルガルガン",
			deck,
			[]string{
				"ルガルガン",
				"スパイクエネルギー",
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

	if cardlist["ローブシン"] >= 3 {
		deckType := analyze(
			"ローブシン",
			deck,
			[]string{
				"ローブシン",
				"アラブルタケ",
				"モモワロウ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イダイナキバ"] >= 3 && cardlist["ニュートラルセンター(ACE SPEC)"] == 1 {
		deckType := analyze(
			"イダイナキバLO",
			deck,
			[]string{
				"イダイナキバ",
				"ヒビキのウソッキー",
				"クラッシュハンマー",
				"ニュートラルセンター(ACE SPEC)",
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

	if cardlist["ガチグマ アカツキ"] >= 3 {
		deckType := analyze(
			"ガチグマ アカツキ",
			deck,
			[]string{
				"ガチグマ アカツキ",
				"マラカッチ",
				"マシマシラ",
				"ラティアスex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ヒードラン"] >= 3 {
		deckType := analyze(
			"ヒードラン",
			deck,
			[]string{
				"ヒードラン",
				"メタング",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ワナイダーex"] >= 3 {
		deckType := analyze(
			"ワナイダーex",
			deck,
			[]string{
				"ワナイダーex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イルカマンex"] >= 3 {
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

	if cardlist["リーリエのピッピex"] >= 3 && cardlist["リーリエのしんじゅ"] >= 3 {
		deckType := analyze(
			"リーリエのピッピex",
			deck,
			[]string{
				"リーリエのピッピex",
				"リーリエのしんじゅ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["テツノイバラex"] >= 3 {
		deckType := analyze(
			"テツノイバラex",
			deck,
			[]string{
				"テツノイバラex",
				"クラッシュハンマー",
				"ポケモンキャッチャー",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ホエルオー"] >= 3 {
		deckType := analyze(
			"ホエルオー",
			deck,
			[]string{
				"ホエルオー",
				"セグレイブ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["イワパレス"] >= 2 {
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

	if cardlist["バチュル"] >= 2 && (cardlist["テツノカイナex"] >= 1 || cardlist["ピカチュウex"] >= 1 || cardlist["テツノイサハex"] >= 1) {
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

	if (cardlist["オーガポン みどりのめんex"] >= 1 || cardlist["オーガポン いどのめんex"] >= 1 || cardlist["オーガポン いしずえのめんex"] >= 1) && (cardlist["テラパゴスex"] >= 1 || cardlist["ピカチュウex"] >= 1 || cardlist["テツノイサハex"] >= 1 || cardlist["リーリエのピッピex"] >= 1) && cardlist["ゼロの大空洞"] >= 2 {
		deckType := analyze(
			"テラスタルバレット",
			deck,
			[]string{
				"オーガポン みどりのめんex",
				"オーガポン いどのめんex",
				"オーガポン いしずえのめんex",
				"テラパゴスex",
				"ピカチュウex",
				"テツノイサハex",
				"リーリエのピッピex",
				"ゼロの大空洞",
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

	if cardlist["ミロカロスex"] >= 2 {
		deckType := analyze(
			"ミロカロスex",
			deck,
			[]string{
				"ミロカロスex",
				"オンバーンex",
				"オーガポン いしずえのめんex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["リキキリンex"] >= 2 {
		deckType := analyze(
			"リキキリンex",
			deck,
			[]string{
				"リキキリンex",
				"オンバーンex",
				"オーガポン いしずえのめんex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["ユキメノコ"] >= 2 && cardlist["マシマシラ"] >= 3 && cardlist["マリィのオーロンゲex"] == 0 {
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

	if (cardlist["カミッチュ"] >= 3 || cardlist["アズマオウ"] >= 3) && cardlist["バチンキー"] >= 3 && cardlist["お祭り会場"] >= 3 {
		deckType := analyze(
			"おまつりおんど",
			deck,
			[]string{
				"カミッチュ",
				"アズマオウ",
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

	if cardlist["オンバーンex"] >= 2 && cardlist["モモワロウ"] >= 2 && cardlist["アラブルタケ"] >= 2 && cardlist["危険な密林"] >= 3 {
		deckType := analyze(
			"オンバーンex",
			deck,
			[]string{
				"オンバーンex",
				"モモワロウ",
				"アラブルタケ",
				"危険な密林",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["おはやし笛"] >= 2 && cardlist["クセロシキのたくらみ"] >= 1 && cardlist["ビワ"] >= 1 {
		deckType := analyze(
			"コントロール",
			deck,
			[]string{
				"ロケット団のリーシャン",
				"ヒビキのウソッキー",
				"ミロカロス",
				"ゲノセクト",
				"イーユイex",
				"ディンルーex",
				"おはやし笛",
				"クセロシキのたくらみ",
				"ビワ",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	if cardlist["シャリタツex"] >= 2 {
		deckType := analyze(
			"シャリタツex",
			deck,
			[]string{
				"シャリタツex",
				"リザードンex",
				"バシャーモex",
				"ゲッコウガex",
				"ドラパルトex",
				"ピジョットex",
			},
		)
		deckTypes = append(deckTypes, deckType)
	}

	ctx.JSON(http.StatusOK, deckTypes)
}
