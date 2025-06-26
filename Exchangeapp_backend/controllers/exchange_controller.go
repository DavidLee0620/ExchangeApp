package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/model"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate model.ExchangeRate
	if err := ctx.ShouldBindBodyWithJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	exchangeRate.Data = time.Now()
	var rate float64
	chromeCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	a := exchangeRate.FromCurrency
	b := exchangeRate.ToCurrency

	err := chromedp.Run(chromeCtx,
		chromedp.Navigate("https://www.baidu.com"),
		chromedp.WaitVisible(`[name="wd"]`, chromedp.ByQuery), //等待搜索框
		chromedp.SendKeys(`[name="wd"]`, a+" "+b+" 汇率"),       //输入
		chromedp.Click(`[id="su"]`, chromedp.NodeVisible),     //点击搜索一下
		chromedp.WaitVisible(`#content_left`, chromedp.ByID),  //等待展示
		chromedp.ActionFunc(func(ctx context.Context) error {
			var result string
			if err := chromedp.Run(ctx,
				chromedp.WaitVisible(`.money-style_59F57`, chromedp.ByQuery),      //等待汇率展示
				chromedp.Text(`.money-style_59F57`, &result, chromedp.ByQueryAll), //F12下查询到元素的控件，抓取元素的内容 1 人民币 ≈ 4.5286 泰铢
				//                           1 泰铢 ≈ 0.2208 人民币
			); err != nil {
				return err
			}

			lines := strings.Split(result, "\n")
			for _, line := range lines {
				pattern := `(\d+\.\d+)\s*` + regexp.QuoteMeta(b)
				re := regexp.MustCompile(pattern)
				match := re.FindStringSubmatch(line)
				// 使用正则表达式匹配数字部分

				var err error
				if len(match) > 1 {
					rate, err = strconv.ParseFloat(match[1], 64)
					if err != nil {
						return err
					}
					fmt.Println(rate)
					return nil
				}

			}
			return fmt.Errorf("exchange rate not found")
		}),
	)

	if err != nil {
		fmt.Println("Failed to retrieve exchange rate from web: " + err.Error())
		return
	}

	exchangeRate.Rate = rate
	if err := global.DB.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	if err := global.DB.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, exchangeRate)

}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRate []model.ExchangeRate

	if err := global.DB.Find(&exchangeRate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}

		return
	}
	ctx.JSON(http.StatusOK, exchangeRate)
}
