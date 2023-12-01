package lang

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/gin-contrib/sessions"
	"strconv"
	"strings"
)

type Service struct {
	bundle    *i18n.Bundle
	ctx       *gin.Context
	localizer *i18n.Localizer
}

const LangID = "LangID"

func New(c *gin.Context, bundle *i18n.Bundle) (Service, string) {
    session := sessions.Default(c)
    lang := c.Query("lang")
    if (lang != "") {
        session.Set(LangID, lang)
        session.Save()
    } else {
        langSession := session.Get(LangID)
        if (langSession != nil) {
            lang = langSession.(string)
        } else {
            accept := c.GetHeader("Accept-Language")
            langQ := ParseAcceptLanguage(accept)
            lang = langQ[0].Lang
        }
    }

	localizer := i18n.NewLocalizer(bundle, lang, "pl")
	return Service{
		bundle:    bundle,
		ctx:       c,
		localizer: localizer,
	}, lang
}

func (s *Service) Trans(str string) string {
	for _, m := range translationMessages {
		if m.ID == str {
			localizedString, _ := s.localizer.Localize(&i18n.LocalizeConfig{
				DefaultMessage: &m,
			})
			return localizedString
		} else if m.Other == str {
			localizedString, _ := s.localizer.Localize(&i18n.LocalizeConfig{
				DefaultMessage: &m,
			})
			return localizedString
		}
	}
	return str
}

// https://siongui.github.io/2015/02/22/go-parse-accept-language/
type LangQ struct {
	Lang string
	Q    float64
}

func ParseAcceptLanguage(acptLang string) []LangQ {
	var lqs []LangQ

	langQStrs := strings.Split(acptLang, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			lq := LangQ{langQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				panic(err)
			}
			lq := LangQ{langQ[0], q}
			lqs = append(lqs, lq)
		}
	}
	return lqs
}
