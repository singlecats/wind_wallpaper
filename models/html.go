package models

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "net/http"
)

func List(response *http.Response) ([]map[string]string, string) {
    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        panic(err)
    }
    content := make([]map[string]string, 24)
    doc.Find(".thumb-listing-page li").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        code := getCode(s)
        typ := getType(s)
        content[i] = make(map[string]string, 5)
        content[i]["url"] = getUrl(s)
        content[i]["src"] = getImagesSrc(s)
        content[i]["size"] = getSize(s)
        content[i]["type"] = typ
        content[i]["code"] = code
        content[i]["link"] = fmt.Sprintf(FullImgUrl, code[0:2], code, typ)
    })
    page , _:= doc.Find(".thumb-listing-page-header").Html()
    return content, page
}

func getUrl(section *goquery.Selection) string {
    url, err := section.Find("a").Attr("href")
    if err == false {
        panic(err)
    }
    return url
}

func getImagesSrc(section *goquery.Selection) string {
    src, err := section.Find(".lazyload").Attr("data-src")
    if err == false {
        panic(err)
    }
    return src
}

func getSize(section *goquery.Selection) string {
    size := section.Find(".thumb-info .wall-res").Text()
    return size
}

func getType(section *goquery.Selection) string {
    typ := section.Find(".thumb-info .png").Eq(0).Text()
    if typ != "" {
        return "png"
    }
    return "jpg"
}

func getCode(section *goquery.Selection) string {
    code, _ := section.Find(".thumb").Eq(0).Attr("data-wallpaper-id")
    return code
}

