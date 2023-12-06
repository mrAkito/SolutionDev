package main

import (
	"net/http"
	"html/template"

	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}

type CommonData struct {
	Title string
}


func main() {
  // インスタンスを作成
  e := echo.New()
  t := &Template{
	templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
  e.Renderer = t

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Static("/static", "public/static")
  // ルートを設定
  e.GET("/", viewMainPage) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する
  e.GET("/savegps", SaveGPS)
  e.GET("/getgoogleAPI", GetGoogleAPI)
  e.StartTLS(":443","server.crt","server.key")
}


func viewMainPage(c echo.Context) error {
	var common = CommonData{Title: "Main Page"}

	data := struct {
		CommonData
		Constent string
	}{
		CommonData: common,
		Constent:   "Hello, Test!",
	}
	return c.Render(http.StatusOK, "mainPage", data)
}

func GetGoogleAPI(c echo.Context) error {
	endpoint := "https://maps.googleapis.com/maps/api/staticmap"

	params := "?center=Tokyo&zoom=13&size=600x300&key=AIzaSyAFICfLJzeKZA0ll7VKKSPWMEVLsSeaER8"
	resp, err := http.Get(endpoint + params)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	outFile, err := os.Create("map.png")
	io.Copy(outFile, resp.Body)
	return c.String(http.StatusOK, "OK")
}