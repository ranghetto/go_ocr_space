package ocr_space

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type OCRText struct {
	ParsedResults []struct {
		TextOverlay struct {
			Lines []struct {
				Words []struct {
					WordText string  `json:"WordText"`
					Left     float64 `json:"Left"`
					Top      float64 `json:"Top"`
					Height   float64 `json:"Height"`
					Width    float64 `json:"Width"`
				} `json:"Words"`

				MaxHeight float64 `json:"MaxHeight"`
				MinTop    float64 `json:"MinTop"`
			} `json:"Lines"`

			HasOverlay bool   `json:"HasOverlay"`
			Message    string `json:"Message"`
		} `json:"TextOverlay"`

		TextOrientation   string `json:"TextOrientation"`
		FileParseExitCode int    `json:"FileParseExitCode"`
		ParsedText        string `json:"ParsedText"`
		ErrorMessage      string `json:"ErrorMessage"`
		ErrorDetails      string `json:"ErrorDetails"`
	} `json:"ParsedResults"`

	OCRExitCode                  int    `json:"OCRExitCode"`
	IsErroredOnProcessing        bool   `json:"IsErroredOnProcessing"`
	ProcessingTimeInMilliseconds string `json:"ProcessingTimeInMilliseconds"`
	SearchablePDFURL             string `json:"SearchablePDFURL"`
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

func (c Config) ConvertPdfFromUrl(pdfUrl string) OCRText {
	resp, err := http.PostForm("https://api.ocr.space/parse/image",
		url.Values{
			"url":                          {pdfUrl},
			"language":                     {c.Language},
			"apikey":                       {c.ApiKey},
			"isOverlayRequired":            {"true"},
			"isSearchablePdfHideTextLayer": {"true"},
			"scale": {"true"},
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	if results.IsErroredOnProcessing == true {
		err = fmt.Errorf("OCR.Api: Error on processing file %s", pdfUrl)
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
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	if results.IsErroredOnProcessing == true {
		err = fmt.Errorf("OCR.Api: Error on processing file %s", imageUrl)
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
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var results OCRText

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	if results.IsErroredOnProcessing == true {
		err = fmt.Errorf("OCR.Api: Error on processing file %s", localPath)
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
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
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

func base64Format(encoded string, localPath string) string {
	s := strings.Split(localPath, ".")
	lastElement := s[len(s)-1]
	pdfString := "data:application/pdf;base64,"
	imageString := "data:image/" + lastElement + ";base64,"

	if lastElement == "pdf" {
		return pdfString + encoded
	} else if lastElement == "png" || lastElement == "jpg" || lastElement == "gif" {
		return imageString + encoded
	} else {
		return "File type not valid. PDF, JPG, PNG or GIF supported."
	}

}

func encodeToBase64(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)

	return base64Format(encoded, path)
}
