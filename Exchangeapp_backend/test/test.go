package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	chromeCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(chromeCtx,
		chromedp.Navigate("https://www.baidu.com"),
		chromedp.WaitVisible(`[name="wd"]`, chromedp.ByQuery),
		chromedp.SendKeys(`[name="wd"]`, "人民币 泰铢 汇率"),
		chromedp.Click(`[id="su"]`, chromedp.NodeVisible),
		chromedp.WaitVisible(`#content_left`, chromedp.ByID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var result string
			if err := chromedp.Run(ctx,
				chromedp.WaitVisible(`.money-style_59F57`, chromedp.ByQuery),
				chromedp.Text(`.money-style_59F57`, &result, chromedp.ByQueryAll),
			); err != nil {
				return err
			}

			lines := strings.Split(result, "\n")
			for _, line := range lines {
				if strings.Contains(line, "泰铢") {
					// 使用正则表达式匹配数字部分
					re := regexp.MustCompile(`(\d+\.\d+)`)
					match := re.FindStringSubmatch(line)
					if len(match) > 1 {
						rate, err := strconv.ParseFloat(match[1], 64)
						if err != nil {
							return err
						}
						fmt.Println(rate)
						return nil
					}
				}
			}
			return fmt.Errorf("exchange rate not found")
		}),
	)

	if err != nil {
		fmt.Println("Failed to retrieve exchange rate from web: " + err.Error())
		return
	}
}
