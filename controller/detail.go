package controller

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/ydtg1993/ant"
	"log"
	"pow/global/orm"
	"pow/model"
	"pow/robot"
	"pow/tools/config"
	"pow/tools/rd"
	"regexp"
	"strings"
	"time"
)

func DetailScan() {
	id, err := rd.LPop("xy:task")
	id = "1"
	/*if err != nil || id == "" {
		return
	}*/
	Seed := new(model.SourceSeed)
	if orm.Eloquent.Where("id = ?", id).First(&Seed); Seed.ID == 0 {
		return
	}
	bot := ant.Get([]int{})
	if bot == nil {
		return
	}
	for try := 0; try < 10; try++ {
		aimUrl := "http://" + config.Spe.SourceUrl + "/"
		bot.WebDriver.Get(aimUrl + Seed.SourceURL)
		Seed.RawContent, err = bot.WebDriver.PageSource()
		if err != nil {
			if try > 9 {
				return
			}
			if try > 3 {
				bot.Proxy(robot.GetProxy())
			}
			bot.WebDriver.Refresh()
			continue
		}
		break
	}
	dateDom, err := bot.WebDriver.FindElement(selenium.ByCSSSelector, "div.content-info>span:nth-child(2)")
	if err == nil {
		t, _ := dateDom.Text()
		re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
		match := re.FindStringSubmatch(t)
		if len(match) < 2 {
			return
		}
		dateTimeString := match[1]
		Seed.PublishTime, _ = time.Parse("2006-01-02 15:04:05", dateTimeString)
	}

	contentDom, err := bot.WebDriver.FindElement(selenium.ByCSSSelector, "div.content")
	if err != nil {
		return
	}

	Seed.RawContent, _ = contentDom.GetAttribute("innerHTML")
	imgDom, err := contentDom.FindElement(selenium.ByTagName, "img")
	if err == nil {
		Seed.BigCover, _ = imgDom.GetAttribute("src")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader("<div id='body'>"+Seed.RawContent+"</div>"))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("p[style='text-align: center;']").Each(func(i int, selection *goquery.Selection) {
		selection.Remove()
	})
	lastP := doc.Find("p:last-child")
	if lastP.Length() > 0 {
		Seed.Links,_ = lastP.Html()
	}
	doc.Find("p:last-child").Remove()
	result, err := doc.Find("div#body").Html()
	if err == nil {
		Seed.Content = result
	}
	orm.Eloquent.Save(Seed)
}
