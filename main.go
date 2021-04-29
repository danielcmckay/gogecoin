package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/joho/godotenv"
	models "gogecoin.com/models"
)

var priceData []float64

func main() {
	doge := models.BuildDogeStruct(connect())
	priceData = append(priceData, doge.Data.Id.Quote.USD.Price)
	renderUi(0, 0, 25, 50, doge, priceData)

}

func connect() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("id", "74")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", goDotEnvVariable("API_KEY"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("ping")

	return string(respBody)
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func renderUi(height int, width int, offsetX int, offsetY int, doge models.Doge, data []float64) {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize: $v", err)
	}
	defer ui.Close()
                                                                   
	p0 := widgets.NewPlot()
	p0.Title = "DOGE Tracker"
	p0.SetRect(0, 5, 60, 35)
	p0.Marker = widgets.MarkerDot
	p0.PlotType = widgets.ScatterPlot
	p0.Data = make([][]float64, 2)
	p0.Data[0] = data
	p0.BorderStyle.Fg = ui.ColorGreen

		p4 := widgets.NewParagraph()
		p4.Title = "Text Box with Wrapping"
		p4.Text =
		`[xxkkkOOOOOOOOO000O0kdk00KKKXXXXK00OddOKKXXXXXXXXXX
		xxkkOOOOOOOOOO00OO0kldOKKKKKKKK0Oxdc:d0KKXXXXXXXXX
		xkkOOOOOOOOOkkkOkkOkxxkkOOOOOkkdlc;;:oOKKXXXXXXXXX
	kkkkkkxxxxxxxxkkOO0K0kddxxxkkxocc;.'cd0KKXXXXXXXXX
	xxxxxxxddxxxkkO00kOKOddkOOkxdxkkd;;cok00KXXXXXXXXX
	ddxxkkkkkkOOOKNKl'lOkxkOkxdodxxkOOxl:d00KXXXXXXXXX
	dxxkOOOOO000KNWKkxO00Od:;;''okOOOOOxldO0KXXXXXXXXX
	xxkkOOOOOO00XN0llld0K0kkxddk0KXK000Okxk0KKXXXXXXXX
	kkkkkkkkkOO0XXc   'kK000KKKKKKKK000OkxkOKKKXXXXXXX
	00OOOOOkOOOKX0:. .:xOOOO000000OOOOOxddxk0KKXXXXXXX
	KKKK0000000KXXd'..';cccdkkOOOOO000OxdddxO0KXXXXXXX
	KKKKKKKKKKKKKXKOdoolooxkkkkOO00OOkxdolodk0KXXXXXXX
	KKKKKKKKKKKKKKK000OkkkkxkkO00OOOkxdolodxk0KXXXXXXX
	000000KKKKKKKK00OOOkkkkxxxkkkkkkkxdddkOOO0KKXXXXXX
	00000000000KKXK0OkxxddddxkkkkOOOOOO00000O0KKKKKKKK
	OOOOOO000000KXK0OkkxdxxkOOOOOOkO00000K00OO0KKKKKKK](fg:green)`
		p4.SetRect(60, 5, 114, 22)
		p4.BorderStyle.Fg = ui.ColorGreen

	p5 := widgets.NewParagraph()
	p5.Title = "Last updated: " + doge.Data.Id.Quote.USD.Last_updated
	p5.SetRect(60, 22, 114, 35)
	p5.Text =
		`[Current Price: $` + strconv.FormatFloat(doge.Data.Id.Quote.USD.Price, 'f', -1, 32) + "\n\n"+ `Percent change in last hour: ` + strconv.FormatFloat(doge.Data.Id.Quote.USD.Percent_change_1h, 'f', 2, 32) + "%%\n\n" + `Percent change in last 24 hours: ` + strconv.FormatFloat(doge.Data.Id.Quote.USD.Percent_change_24h, 'f', 2, 32) + "%%\n\n" + `Percent change in last week: ` + strconv.FormatFloat(doge.Data.Id.Quote.USD.Percent_change_7d, 'f', 2, 32) + "%%\n](fg:green)"
	p5.BorderStyle.Fg = ui.ColorGreen
	ui.Render(p0, p4, p5)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(5 * time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID { // event string/identifier
			case "q", "<C-c>": // press 'q' or 'C-c' to quit
				return
			}
		// use Go's built-in tickers for updating and drawing data

		case <- ticker:
			doge = models.BuildDogeStruct(connect())
			priceData = append(priceData, doge.Data.Id.Quote.USD.Price)
			ui.Render(p0, p5)
		}
	}
}
