# OCR_Space
OCR library is based on [OCR.Space](https://ocr.space/) API, meant to read text from pdf files and images with Go language

## Index
1. [Get your API key](#getapi)
2. [Installation](#install)
3. [Basic Usage](#usage)
4. [Example Code](#example)

### <a name="getapi"></a>1. Get your API key
You can get your personal API key [here](https://ocr.space/ocrapi) (we will need this later)

### 2. <a name="install"></a>Installation
`go get -t github.com/ranghetto/go_ocr_space`

## 3. <a name="usage"></a>Baisc Usage
You need at first to create a configuration:
```go
package main

import(
	/*
	Remember to run your program from your $GO_PATH/src/name_of_your_folder
	or to provide the right path to this library that is situated
	in $GO_PATH/src/github.com/ranghetto/go_ocr_space
	*/
	
	ocr "github.com/ranghetto/go_ocr_space"
	//Other libraries...
)

func main(){
	
	config := ocr.InitConfig("yourApiKeyHere", "eng")
	
}
```
The first parameter is your API key as a string and the second one is the code of the language you want read from file or image.
Here a list of all available languages and their code*:
* Arabic = `ara`
* Bulgarian = `bul`
* Chinese(Simplified) = `chs`
* Chinese(Traditional) = `cht`
* Croatian = `hrv`
* Czech = `cze`
* Danish = `dan`
* Dutch = `dut`
* English = `eng`
* Finnish = `fin`
* French = `fre`
* German = `ger`
* Greek = `gre`
* Hungarian = `hun`
* Korean = `kor`
* Italian = `ita`
* Japanese = `jpn`
* Polish = `pol`
* Portuguese = `por`
* Russian = `rus`
* Slovenian = `slv`
* Spanish = `spa`
* Swedish = `swe`
* Turkish = `tur`

Now we can go ahead and start reading some text; there are four method that allow you to do it:

`config.ConvertImageFromUrl("https://example.com/image.png")`

`config.ConvertPdfFromUrl("https://example.com/myfile.pdf")`

`config.ConvertImageFromLocal("path/to/the/image.jpg")`

`config.ConvertPdfFromLocal("path/to/the/file.pdf")`

Method names are self explanatory.

So basically this will give you back the whole struct complete of all parameters that [OCR API](https://ocr.space/ocrapi) provide to you.

If you are only interested in the output text call `.justText()` method at thew end of the four once mentioned above.

## 4. <a name="example"></a>Example Code
```go
package main

import (
	"fmt"
	ocr "go_ocr_space"
)

func main() {
	//this is a demo api key 
	apiKey:="helloworld"

    //setting up the configuration 
    config := ocr.InitConfig(apiKey , "eng")

    //actual converting the image from the url (struct content) 
    result := config.ConvertImageFromUrl("https://i.pinimg.com/originals/74/1a/37/741a37fe653930e27e4b5e9c61f30ca0.jpg")

    //printing the just the parsed text
	fmt.Println(result.JustText())
}
```
### Example output
![](https://i.pinimg.com/originals/74/1a/37/741a37fe653930e27e4b5e9c61f30ca0.jpg)
```
"There are many people out
there who will tell you that
you can't. What you've got
to do is turn around and say,
'Watch Me'."
-Jack White
```
