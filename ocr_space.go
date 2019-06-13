package ocr_space

import (
	"encoding/json"
	"github.com/polds/imgbase64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type OCRText struct {
	ParsedResults                []ParsedResults `json:"ParsedResults"`
	OCRExitCode                  int             `json:"OCRExitCode"`
	IsErroredOnProcessing        bool            `json:"IsErroredOnProcessing"`
	ProcessingTimeInMilliseconds string          `json:"ProcessingTimeInMilliseconds"`
	SearchablePDFURL             string          `json:"SearchablePDFURL"`
}

type ParsedResults struct {
	TextOverlay       TextOverlay `json:"TextOverlay"`
	TextOrientation   string      `json:"TextOrientation"`
	FileParseExitCode int         `json:"FileParseExitCode"`
	ParsedText        string      `json:"ParsedText"`
	ErrorMessage      string      `json:"ErrorMessage"`
	ErrorDetails      string      `json:"ErrorDetails"`
}

type TextOverlay struct {
	Lines      []Lines `json:"Lines"`
	HasOverlay bool    `json:"HasOverlay"`
	Message    string  `json:"Message"`
}

type Lines struct {
	Words     []Words `json:"Words"`
	MaxHeight float64 `json:"MaxHeight"`
	MinTop    float64 `json:"MinTop"`
}

type Words struct {
	WordText string  `json:"WordText"`
	Left     float64 `json:"Left"`
	Top      float64 `json:"Top"`
	Height   float64 `json:"Height"`
	Width    float64 `json:"Width"`
}

type Config struct {
	ApiKey   string
	Language string
}

func InitConfig(apiKey string, language string) Config {
	var config Config
	config.ApiKey = apiKey
	config.Language = language
	return config
}

func (c Config) ConvertPdfFromUrl(imageUrl string) OCRText {
	resp, err := http.PostForm("https://api.ocr.space/parse/image",
		url.Values{
			"url":                          {imageUrl},
			"language":                     {c.Language},
			"apikey":                       {c.ApiKey},
			"isOverlayRequired":            {"true"},
			"isSearchablePdfHideTextLayer": {"true"},
			"scale": {"true"},
		},
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	return results
}

func (c Config) ConvertImageFromUrl(imageUrl string) OCRText {
	resp, err := http.PostForm("https://api.ocr.space/parse/image",
		url.Values{
			"url":      {imageUrl},
			"language": {c.Language},
			"apikey":   {c.ApiKey},
		},
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	return results
}

func (c Config) ConvertPdfFromLocal(localPath string) OCRText {

	baseImage := encodeToBase64(localPath)

	resp, err := http.PostForm("https://api.ocr.space/parse/image",
		url.Values{
			"base64Image":                  {baseImage},
			"language":                     {c.Language},
			"apikey":                       {c.ApiKey},
			"isOverlayRequired":            {"true"},
			"isSearchablePdfHideTextLayer": {"true"},
			"scale": {"true"},
		},
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	return results
}

func (c Config) ConvertImageFromLocal(localPath string) OCRText {

	baseImage := encodeToBase64(localPath)

	resp, err := http.PostForm("https://api.ocr.space/parse/image",
		url.Values{
			"base64Image": {baseImage},
			"language":    {c.Language},
			"apikey":      {c.ApiKey},
		},
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	return results
}

func (ocr OCRText) JustText() string {
	text := ""
	for _, page := range ocr.ParsedResults {
		text += page.ParsedText
	}
	return text
}

func base64ImgToPdf(imageString string) string {
	s := strings.Split(imageString, ";")
	pdfString := "data:application/pdf"
	if len(s) != 2 {
		return "Error parsing the base64 image"
	} else {
		return pdfString + ";" + s[1]
	}
}

func encodeToBase64(path string) string {
	baseImage, err := imgbase64.FromLocal(path)
	if nil != err {
		return "Error encoding base64"
	}
	format := strings.Split(path, ".")
	if format[1] == "pdf" {
		return base64ImgToPdf(baseImage)
	}
	return baseImage
}
