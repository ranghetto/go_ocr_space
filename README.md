# OCR_Space
OCR library is based on [OCR.Space](https://ocr.space/) API, meant to read text from pdf files and images with Go language

## Index
1. [Get your API key](#getapi)
2. [Installation](#install)
3. [Update](#update)
4. [Basic Usage](#usage)
5. [Example Code](#example)

### <a name="getapi"></a>1. Get your API key
You can get your personal API key [here](https://ocr.space/ocrapi) (we will need this later)

### 2. <a name="install"></a>Installation
`go get -t github.com/ranghetto/go_ocr_space`

### 3. <a name="update"></a>Update
Delete the folder situated in $GO_PATH/src/github.com/ranghetto/go_ocr_space and then run this command:

`go get -t github.com/ranghetto/go_ocr_space`

### 4. <a name="usage"></a>Basic Usage
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
	
	config := ocr.InitConfig("yourApiKeyHere", "eng", OCREngine2)
	//More code here...
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

Now we can go ahead and start reading some text; there are three method that allow you to do it:

`config.ParseFromUrl("https://example.com/image.png")`

`config.ParseFromLocal("path/to/the/image.jpg")`

`config.ParseFromBase64("data:image/jpeg;base64,873hf9qehq98efwuehf...")`

Method names are self explanatory.

Remember:
.ParseFromBase64 need as parameter a valid Base64 format like `data:<file>/<extension>;base64,<image>` where:
* `<file>` is `application` in case of a pdf file or `image` in case of an image
* `<extension>` is the extension of the file you encode. Only valid are `pdf`, `jpg`, `png` and `gif`
* `<image>` is the actual encode of your file 

So basically these methods will give you back the whole struct complete of all parameters that [OCR.Space](https://ocr.space/ocrapi) provides to you.

If you are only interested in the output text call `.justText()` method at the end of one of the three methods mentioned above.

## 5. <a name="example"></a>Example Code
```go
package main

import (
	"fmt"
	ocr "github.com/ranghetto/go_ocr_space"
)

func main() {
	//this is a demo api key 
	apiKey:="helloworld"

    //setting up the configuration 
    config := ocr.InitConfig(apiKey , "eng", OCREngine2)

    //actual converting the image from the url (struct content) 
    result, err := config.ParseFromUrl("https://www.azquotes.com/picture-quotes/quote-maybe-we-should-all-just-listen-to-records-and-quit-our-jobs-jack-white-81-40-26.jpg")
    if err != nil {
    	fmt.Println(err)
    }
    //printing the just the parsed text
	fmt.Println(result.JustText())
}
```
### Example output
![](https://www.azquotes.com/picture-quotes/quote-maybe-we-should-all-just-listen-to-records-and-quit-our-jobs-jack-white-81-40-26.jpg)
```
Maybe we should all just listen to 
records and quit our jobs 
Jack White 
AZ QUOTES
```
