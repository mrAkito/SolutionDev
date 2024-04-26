package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"html/template"

	"io"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

type GPSDatas struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
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
	e.GET("/gpsdata", SendGPSdata)
	//   e.StartTLS(":443","server.crt","server.key")
	e.Logger.Fatal(e.Start(":8080"))
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

func SendGPSdata(c echo.Context) error {
	// ファイルを開く
	file, err := os.OpenFile("public/gps.csv", os.O_RDONLY, 0655)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var gpsData []GPSDatas
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		la, err := strconv.ParseFloat(record[0], 64)
		ln, err := strconv.ParseFloat(record[1], 64)

		data := GPSDatas{
			Latitude:  la,
			Longitude: ln,
		}
		gpsData = append(gpsData, data)
	}

	return c.JSON(http.StatusOK, gpsData)
}
