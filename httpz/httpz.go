package httpz

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gzorm/commons/core/mapping"
	"github.com/gzorm/commons/encoding"
	"github.com/gzorm/commons/header"
	"github.com/gzorm/commons/rest/pathvar"

	"io"
	"net/http"
	"strings"

	"golang.org/x/text/language"

	"github.com/gzorm/commons/core/logx"

	enLang "github.com/go-playground/locales/en"
	brLang "github.com/go-playground/locales/pt_BR"
	zhLang "github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	brTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

const (
	formKey           = "form"
	pathKey           = "path"
	maxMemory         = 32 << 20 // 32MB
	maxBodyLen        = 8 << 20  // 8MB
	separator         = ";"
	tokensInAttribute = 2
	xForwardedFor     = "X-Forwarded-For"
)

var supportLang map[string]string

var (
	formUnmarshaler = mapping.NewUnmarshaler(formKey, mapping.WithStringValues())
	pathUnmarshaler = mapping.NewUnmarshaler(pathKey, mapping.WithStringValues())
	xValidator      = NewValidator()
)

// Parse parses the request.
func Parse(r *http.Request, v any, isValidate bool) error {
	if r.Header.Get(header.ContentType) == "application/xml" || r.Header.Get(header.ContentType) == "text/xml" ||
		strings.Contains(r.Header.Get(header.ContentType), "application/xml") ||
		strings.Contains(r.Header.Get(header.ContentType), "text/xml") {
		var buf strings.Builder
		reader := io.LimitReader(r.Body, 8<<20)
		teeReader := io.TeeReader(reader, &buf)
		decoder := xml.NewDecoder(teeReader)
		// 解析 XML 数据到 Person 结构体
		err := decoder.Decode(v)
		if err != nil {
			return err
		}
	}
	if err := ParseJsonBody(r, v); err != nil {
		return err
	}

	if err := ParsePath(r, v); err != nil {
		return err
	}

	if err := ParseForm(r, v); err != nil {
		return err
	}

	if err := ParseHeaders(r, v); err != nil {
		return err
	}

	if isValidate {
		if errMsg := xValidator.Validate(v, r.Header.Get("Accept-Language")); errMsg != "" {
			return fmt.Errorf("%v,%v", xValidator.ErrorCode, errMsg)
		}
	}
	return nil
}

// ParseHeaders parses the headers request.
func ParseHeaders(r *http.Request, v any) error {
	return encoding.ParseHeaders(r.Header, v)
}

// ParseForm parses the form request.
func ParseForm(r *http.Request, v any) error {
	params, err := GetFormValues(r)
	if err != nil {
		return err
	}

	return formUnmarshaler.Unmarshal(params, v)
}

// ParseHeader parses the request header and returns a map.
func ParseHeader(headerValue string) map[string]string {
	ret := make(map[string]string)
	fields := strings.Split(headerValue, separator)

	for _, field := range fields {
		field = strings.TrimSpace(field)
		if len(field) == 0 {
			continue
		}

		kv := strings.SplitN(field, "=", tokensInAttribute)
		if len(kv) != tokensInAttribute {
			continue
		}

		ret[kv[0]] = kv[1]
	}

	return ret
}

// ParseJsonBody parses the post request which contains json in body.
func ParseJsonBody(r *http.Request, v any) error {
	if withJsonBody(r) {
		reader := io.LimitReader(r.Body, maxBodyLen)
		return mapping.UnmarshalJsonReader(reader, v)
	}

	return mapping.UnmarshalJsonMap(nil, v)
}

// ParsePath parses the symbols reside in url path.
// Like http://localhost/bag/:name
func ParsePath(r *http.Request, v any) error {
	vars := pathvar.Vars(r)
	m := make(map[string]any, len(vars))
	for k, v := range vars {
		m[k] = v
	}

	return pathUnmarshaler.Unmarshal(m, v)
}

func withJsonBody(r *http.Request) bool {
	return r.ContentLength > 0 && strings.Contains(r.Header.Get(header.ContentType), header.ApplicationJson)
}

// GetFormValues returns the form values.
func GetFormValues(r *http.Request) (map[string]any, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	if err := r.ParseMultipartForm(maxMemory); err != nil {
		if err != http.ErrNotMultipart {
			return nil, err
		}
	}

	params := make(map[string]any, len(r.Form))
	for name := range r.Form {
		formValue := r.Form.Get(name)
		if len(formValue) > 0 {
			params[name] = formValue
		}
	}

	return params, nil
}

// GetRemoteAddr returns the peer address, supports X-Forward-For.
func GetRemoteAddr(r *http.Request) string {
	v := r.Header.Get(xForwardedFor)
	if len(v) > 0 {
		return v
	}

	return r.RemoteAddr
}

type Validator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     map[string]ut.Translator
	ErrorCode int
}

func NewValidator() *Validator {
	v := Validator{}
	// set default error code to 3
	v.ErrorCode = 3

	en := enLang.New()
	zh := zhLang.New()
	br := brLang.New()
	v.Uni = ut.New(zh, en, br)
	v.Validator = validator.New()
	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")
	brTrans, _ := v.Uni.GetTranslator("pt_BR")

	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans
	v.Trans["pt"] = brTrans
	// add support languages
	initSupportLanguages()

	err := enTranslations.RegisterDefaultTranslations(v.Validator, enTrans)
	if err != nil {
		logx.Errorw("register English translation failed", logx.Field("detail", err.Error()))
		return nil
	}
	err = zhTranslations.RegisterDefaultTranslations(v.Validator, zhTrans)
	if err != nil {
		logx.Errorw("register Chinese translation failed", logx.Field("detail", err.Error()))
		return nil
	}
	err = brTranslations.RegisterDefaultTranslations(v.Validator, brTrans)
	if err != nil {
		logx.Errorw("register pt-BR translation failed", logx.Field("detail", err.Error()))
		return nil
	}

	return &v
}

func (v *Validator) Validate(data any, lang string) string {
	err := v.Validator.Struct(data)
	if err == nil {
		return ""
	}

	targetLang, parseErr := ParseAcceptLanguage(lang)
	if parseErr != nil {
		return parseErr.Error()
	}

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		transData := errs.Translate(v.Trans[targetLang])
		s := strings.Builder{}
		for _, v := range transData {
			s.WriteString(v)
			s.WriteString(" ")
		}
		return s.String()
	}

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return invalid.Error()
	}

	return ""
}

func ParseAcceptLanguage(lang string) (string, error) {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		return "", errors.New("fail to parse accept language")
	}

	for _, v := range tags {
		if val, ok := supportLang[v.String()]; ok {
			return val, nil
		}
	}

	return "zh", nil
}

func initSupportLanguages() {
	supportLang = make(map[string]string)
	supportLang["zh"] = "zh"
	supportLang["zh-CN"] = "zh"
	supportLang["en"] = "en"
	supportLang["en-US"] = "en"
	supportLang["pt-BR"] = "pt"
	supportLang["pt"] = "pt"

}

// RegisterValidation registers the validation function to validator
func RegisterValidation(tag string, fn validator.Func) {
	if err := xValidator.Validator.RegisterValidation(tag, fn); err != nil {
		logx.Must(errors.Join(err, errors.New("failed to register the validation function, tag is "+tag)))
	}
}

// RegisterValidationTranslation registers the validation translation for validator
func RegisterValidationTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc,
	translationFn validator.TranslationFunc) {
	if err := xValidator.Validator.RegisterTranslation(tag, trans, registerFn, translationFn); err != nil {
		logx.Must(errors.Join(err, errors.New("failed to register the validation translation, tag is "+tag)))
	}
}

// SetValidatorErrorCode sets the error code for validator when errors occurs
func SetValidatorErrorCode(code int) {
	xValidator.ErrorCode = code
}
