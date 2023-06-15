package controller

import (
	"fmt"
	"github.com/gocolly/colly"
	"pow/global/orm"
	"pow/model"
	"pow/robot"
	"pow/tools"
	"pow/tools/config"
	"strings"
	"time"
)

func MenuScan() {
	menu := map[string]string{
		"电影":  "film",
		"国产剧": "neidi",
		"美剧":  "meiju",
		"韩剧":  "hanjv",
		"港剧":  "gangjv",
		"日剧":  "rijv",
		"动画":  "comic",
		"明星":  "mingxing",
		"音乐":  "music",
		"纪录片": "jilupian",
		"精选":  "nice",
	}

	for m, uri := range menu {
		go Catch(m, uri)
	}
}

func Catch(menu, uri string) {
	bot := robot.GetColly()
	tryLimit := 10
	aimUrl := "http://" + config.Spe.SourceUrl + "/"

	end := 1
	switch uri {
	case "film":
		uri += "/15_%d.html"
		end = 1308
		break
	case "neidi":
		uri += "37_%d.html"
		end = 160
		break
	case "meiju":
		uri += "19_%d.html"
		end = 175
		break
	case "hanjv":
		uri += "list_34_%d.html"
		end = 55
		break
	case "gangjv":
		uri += "list_32_%d.html"
		end = 34
		break
	case "rijv":
		uri += "list_33_%d.html"
		end = 73
		break
	case "comic":
		uri += "17_%d.html"
		end = 182
		break
	case "mingxing":
		uri += "35_%d.html"
		end = 18
		break
	case "music":
		uri += "4_%d.html"
		end = 262
		break
	case "jilupian":
		uri += "39_%d.html"
		end = 71
		break
	case "nice":
		uri += "21_%d.html"
		end = 127
		break
	}
	for start := 1; start <= end; start++ {
		for try := 1; try < tryLimit; try++ {
			url := fmt.Sprintf(aimUrl+uri, start)
			bot.OnHTML("ul.imdb-list>li", func(e *colly.HTMLElement) {
				T := e.DOM.Find("p.s-title")
				url, _ := T.Find("a").Attr("href")
				url = strings.Replace(url, aimUrl, "", 1)
				Seed := new(model.SourceSeed)
				if Seed.Exists(url) == true {
					return
				}
				Seed.Type = menu
				Seed.Title, _ = tools.ConvertToUTF8(T.Text())
				Seed.SourceURL = url
				Seed.Description, _ = tools.ConvertToUTF8(e.DOM.Find("p.s-desc").Text())
				Seed.Cover, _ = e.DOM.Find("img").Attr("src")
				Seed.Tag = model.Tags{}
				err := orm.Eloquent.Create(&Seed).Error
				if err != nil {
					msg := fmt.Sprintf("资源入库失败 url = %s err = %s", url, err.Error())
					model.RecordFail(url, msg, "资源入库", 1)
				} else {
					//rd.RPush()
				}
			})

			err := bot.Visit(url)
			if err != nil {
				bot = robot.GetColly()
				if try == 30 {
					model.RecordFail(url, "无法抓取分类列表页信息 :"+url, "列表错误", 0)
				}
			} else {
				break
			}
			t := time.NewTicker(time.Second * 3)
			<-t.C
		}
	}
}
