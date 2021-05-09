package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TranslateData struct {
	Data        interface{}
	SourceLang  string
	TargetLangs interface{}
	Records     []interface{}
}

func Translate(OriginalText, SourceLang, TargetLang string) (string, error) {
	var Text []string
	var Result []interface{}

	SourceLang = strings.ToLower(SourceLang)
	TargetLang = strings.ToLower(TargetLang)

	URL := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" +
		SourceLang + "&tl=" + TargetLang + "&dt=t&q=" + url.QueryEscape(OriginalText)

	Response, Error := http.Get(URL)
	if Error != nil {
		return "", errors.New("Error getting translate.googleapis.com")
	}
	defer Response.Body.Close()

	Body, Error := ioutil.ReadAll(Response.Body)
	if Error != nil {
		return "", errors.New("Error reading response Body")
	}

	if strings.Contains(string(Body), `<title>Error 400 (Bad Request)`) {
		return "", errors.New("Error 400 (Bad Request)")
	}

	Error = json.Unmarshal(Body, &Result)
	if Error != nil {
		return "", errors.New("Error unmarshaling data")
	}
	if len(Result) > 0 {
		Outter := Result[0]
		for _, Inner := range Outter.([]interface{}) {
			for _, Line := range Inner.([]interface{}) {
				Text = append(Text, fmt.Sprintf("%v", Line))
				break
			}
		}
		return strings.Join(Text, ""), nil
	} else {
		return "", errors.New("No translated data in response")
	}
}

func TranslateArrayOfStrings(TranslateArrayMap map[string]interface{}, languges []string) (map[string]interface{}, error) {
	CoulmObj := make(map[string]interface{})
	Bool := true
	for Bool {
		if TranslateArrayMap["EN"] != nil {
			TranslateMap := fmt.Sprintf("%v", TranslateArrayMap["EN"])
			TranslateMap = TranslateMap[1 : len(TranslateMap)-1]
			CoulmObj["EN"] = strings.Split(TranslateMap, " ")
			for _, Lang := range languges {
				if Lang != "EN" {
					TranslatedMap, Error := Translate(TranslateMap, "EN", Lang)
					if Error != nil {
						return nil, Error
					}
					CoulmObj[Lang] = strings.Split(TranslatedMap, " ")
				}
			}
			Bool = false
		} else {
			for Key, Obj := range TranslateArrayMap {
				TranslateMap := fmt.Sprintf("%v", Obj.([]interface{}))
				TranslateMap = TranslateMap[1 : len(TranslateMap)-1]
				TranslatedMap, Error := Translate(TranslateMap, Key, "EN")
				if Error != nil {
					return nil, Error
				}
				TranslateArrayMap["EN"] = strings.Split(TranslatedMap, " ")
				break
			}
		}
	}
	return CoulmObj, nil
}

func TranslateArrayOfMap(DescMap map[string]interface{}, languges []string) (map[string]interface{}, error) {
	Bool := true

	for Bool {
		if DescMap["EN"] != nil {
			Desc := fmt.Sprintf("%v", DescMap["EN"])
			Desc = Desc[1 : len(Desc)-1]
			for _, Lang := range languges {
				if Lang != "EN" {
					TranslateDesc, Error := Translate(Desc, "EN", Lang)
					if Error != nil {
						return nil, Error
					}
					DescMap[Lang] = TranslateDesc
				}
			}
			Bool = false
		} else {
			for Key, Obj := range DescMap {
				Desc := fmt.Sprintf("%v", Obj.(string))
				Desc = Desc[1 : len(Desc)-1]
				TranslateDesc, Error := Translate(Desc, Key, "EN")
				if Error != nil {
					return nil, Error
				}
				DescMap["EN"] = TranslateDesc
				break
			}
		}
	}
	return DescMap, nil
}

func (TranslateObj *TranslateData) NewTranslate() (interface{}, error) {

	DataMap, ok := TranslateObj.Data.(map[string]interface{})
	if ok {
		for _, Obj := range DataMap {
			_, ok := Obj.([]interface{})
			if ok {
				TranslatedMap, Error := TranslateArrayOfStrings(DataMap, TranslateObj.TargetLangs.([]string))
				if Error != nil {
					return nil, Error
				}
				return TranslatedMap, nil
			}
			_, ok = Obj.(string)
			if ok {
				TranslatedMap, Error := TranslateArrayOfMap(DataMap, TranslateObj.TargetLangs.([]string))
				if Error != nil {
					return nil, Error
				}
				return TranslatedMap, nil
			}
		}
	}

	DataString, ok := TranslateObj.Data.(string)
	if ok {
		TranslatedText, Error := Translate(DataString, TranslateObj.SourceLang, TranslateObj.TargetLangs.(string))
		if Error != nil {
			return nil, Error
		}
		return TranslatedText, nil
	}
	return nil, errors.New("Error in Transalting")
}

func main() {
	// Case 1
	var NewTranslateArray TranslateData
	transalteArray := make(map[string]interface{})
	transalteArray["EN"] = []interface{}{"Try", "Translate", "Words"}
	NewTranslateArray.Data = transalteArray
	NewTranslateArray.SourceLang = "EN"
	NewTranslateArray.TargetLangs = []string{"AR", "FR", "SP"}
	TransaltedData, Error := NewTranslateArray.NewTranslate()
	if Error != nil {
		fmt.Println("Error", Error.Error())
	}
	fmt.Println("Case 1", TransaltedData)
	/////////////////////////////
	// Case 2
	var NewTranslateMap TranslateData
	transalteMap := make(map[string]interface{})
	transalteMap["EN"] = "Try Translate Words"
	NewTranslateMap.Data = transalteMap
	NewTranslateMap.SourceLang = "EN"
	NewTranslateMap.TargetLangs = []string{"AR", "FR", "SP"}
	TransaltedData, Error = NewTranslateMap.NewTranslate()
	if Error != nil {
		fmt.Println("Error", Error.Error())
	}
	fmt.Println("Case 2", TransaltedData)
	/////////////////////////////
	// Case 3
	var NewTranslateString TranslateData
	NewTranslateString.Data = "Try Translate Words"
	NewTranslateString.SourceLang = "EN"
	NewTranslateString.TargetLangs = "FR"
	TransaltedData, Error = NewTranslateString.NewTranslate()
	if Error != nil {
		fmt.Println("Error", Error.Error())
	}
	fmt.Println("Case 3", TransaltedData)
	///////////////////////////////
}
