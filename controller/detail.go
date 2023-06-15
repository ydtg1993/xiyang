package controller

import (
	"github.com/gocolly/colly"
	"pow/global/orm"
	"pow/model"
	"pow/robot"
	"pow/tools"
	"pow/tools/config"
	"pow/tools/rd"
	"regexp"
	"time"
)

func DetailScan()  {
	id, err := rd.LPop("xy:task")
	id = "1"
	/*if err != nil || id == "" {
		return
	}*/
	Seed := new(model.SourceSeed)
	if orm.Eloquent.Where("id = ?", id).First(&Seed); Seed.ID == 0 {
		return
	}
	bot := robot.GetColly()
	bot.OnHTML("div.content-info>span:nth-child(2)", func(e *colly.HTMLElement) {
		t := tools.ConvertToUTF8(e.DOM.Text())
		re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)
		match := re.FindStringSubmatch(t)
		if len(match) < 2 {
			return
		}
		dateTimeString := match[1]
		Seed.PublishTime, _ = time.Parse("2006-01-02 15:04:05", dateTimeString)
	})

	bot.OnHTML("div.content", func(e *colly.HTMLElement) {
		html,_ := e.DOM.Html()
		Seed.RawContent = tools.ConvertToUTF8(html)
		Seed.Content = html
	})

	aimUrl := "http://" + config.Spe.SourceUrl + "/"
	err = bot.Visit(aimUrl+Seed.SourceURL)
	if err != nil {
		return
	}
}
