package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	IsInit    bool // check if translator initialized or not
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

// func (t *Translator) InitTranslator(initLang string) {
// 	// added by mohamed hassan to support multilanguage
// 	t.bundle = i18n.NewBundle(language.English)
// 	t.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
// 	t.bundle.MustLoadMessageFile("./lang/en.json")
// 	t.bundle.MustLoadMessageFile("./lang/zh.json")
// 	t.bundle.MustLoadMessageFile("./lang/ar.json")
// 	t.localizer = i18n.NewLocalizer(t.bundle, initLang) // should add additionl check here
// 	t.IsInit = true
// 	// end of adding
// }

// added by mohamed hassan to support multilanguage
func (t *Translator) InitTranslator(initLang string, langPath string) {

	langFiles, err := ioutil.ReadDir(langPath)
	if err != nil {
		fmt.Printf("InitTranslator() Error: %v", err)
	}

	t.bundle = i18n.NewBundle(language.English)
	t.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, langFile := range langFiles {
		if !langFile.IsDir() {
			langFilePath := langPath + langFile.Name()
			fmt.Printf("Language file: %s loaded.\r\n", langFilePath)
			t.bundle.MustLoadMessageFile(langFilePath)
		}
	}

	t.localizer = i18n.NewLocalizer(t.bundle, initLang) // should add additionl check here
	t.IsInit = true
	// end of adding
}

func (t *Translator) SetTranslatorLanguage(lang string) {
	t.localizer = i18n.NewLocalizer(t.bundle, lang)
}

func (t *Translator) TranslateMessage(messageID string) string {
	translatedMsg, err := t.localizer.LocalizeMessage(&i18n.Message{ID: messageID})
	if err != nil {
		return messageID
	}
	//return t.localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
	return translatedMsg
}
