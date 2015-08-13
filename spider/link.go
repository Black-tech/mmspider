package spider

import (
	"github.com/PuerkitoBio/goquery"
)

// box-inner imgString
func GetImg(url string) (urls []string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" {
			urls = append(urls, testString(s.Text())...)
		}
	})
	return
}
func GetImgLinks(url string) (count int64, err error) {
	count = 0
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		img_href, exists1 := s.Parent().Attr("href")
		img_title, exists2 := s.Parent().Attr("title")
		img_src, exists3 := s.Attr("src")

		if exists1 && exists2 && exists3 {
			img_count++
			count++
			imgs[img_count] = &Image{
				Id:   img_count,
				Name: img_title,
				Src:  img_src,
				Href: img_href,
			}
		}
	})
	return
}
