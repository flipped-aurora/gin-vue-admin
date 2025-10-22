package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	IsInit           bool // check if translator initialized or not
	bundle           *i18n.Bundle
	localizer        *i18n.Localizer
	defaultLocalizer *i18n.Localizer
	mu               sync.RWMutex // 保护localizer和defaultLocalizer的并发访问
}

// func (t *Translator) InitTranslator(initLang string) {
// 	// added by mohamed hassan to support multilanguage
// 	t.bundle = i18n.NewBundle(language.English)
// 	t.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
// 	t.bundle.MustLoadMessageFile("./lang/en.json")
// 	t.bundle.MustLoadMessageFile("./lang/zh.json")
// 	t.bundle.MustLoadMessageFile("./lang/ar.json")
// 	t.bundle.MustLoadMessageFile("./lang/ru.json")
// 	t.localizer = i18n.NewLocalizer(t.bundle, initLang) // should add additionl check here
// 	t.IsInit = true
// 	// end of adding
// }

// added by mohamed hassan to support multilanguage
func (t *Translator) InitTranslator(initLang string, langPath string) {

	langFiles, err := ioutil.ReadDir(langPath)
	if err != nil {
		fmt.Printf("InitTranslator() Error reading language directory '%s': %v\n", langPath, err)
		return
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
	fmt.Printf("Translator initialized successfully with language: %s\n", initLang)
	// end of adding
}

func (t *Translator) InitTranslatorEx(initLang string, defaultLang string, langPath string) {

	langFiles, err := ioutil.ReadDir(langPath)
	if err != nil {
		fmt.Printf("InitTranslatorEx() Error reading language directory '%s': %v\n", langPath, err)
		return
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

	if defaultLang != "" {
		t.defaultLocalizer = i18n.NewLocalizer(t.bundle, defaultLang) // should add additionl check here
	} else {
		t.defaultLocalizer = i18n.NewLocalizer(t.bundle, "zh") // set default langauge as chinese in case of there is no transalation found
	}

	t.IsInit = true
	fmt.Printf("Translator initialized successfully with language: %s, default: %s\n", initLang, defaultLang)
	// end of adding
}

func (t *Translator) SetTranslatorLanguage(lang string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.localizer = i18n.NewLocalizer(t.bundle, lang)
}

func (t *Translator) SetTranslatorDefaultLanguage(lang string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.defaultLocalizer = i18n.NewLocalizer(t.bundle, lang)
}

func (t *Translator) TranslateMessage(messageID string) string {
	t.mu.RLock()
	localizer := t.localizer
	defaultLocalizer := t.defaultLocalizer
	t.mu.RUnlock()

	// 检查localizer是否已初始化
	if localizer == nil {
		fmt.Printf("Warning: localizer is not initialized, returning original message: %s\n", messageID)
		return messageID
	}

	translatedMsg, err := localizer.LocalizeMessage(&i18n.Message{ID: messageID})
	if err != nil || translatedMsg == "" { // if translation fail use default language transalator
		if err != nil {
			fmt.Printf("Translation error for message '%s': %v\n", messageID, err)
		}

		// 检查defaultLocalizer是否已初始化
		if defaultLocalizer == nil {
			fmt.Printf("Warning: defaultLocalizer is not initialized, returning original message: %s\n", messageID)
			return messageID
		}

		translatedMsg, err = defaultLocalizer.LocalizeMessage(&i18n.Message{ID: messageID})
		if err != nil {
			fmt.Printf("Default translation error for message '%s': %v\n", messageID, err)
			return messageID
		}
	}
	//return t.localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
	return translatedMsg
}
