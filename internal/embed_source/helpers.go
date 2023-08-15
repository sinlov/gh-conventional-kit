package embed_source

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/aymerick/raymond"
	"reflect"
)

// RegisterSettings
// most of this can use embed_source.RegisterSettings(DefaultFunctions)
func RegisterSettings(funcSettings map[string]interface{}) {
	for name, function := range sprig.GenericFuncMap() {
		if invalidHelper(name) {
			continue
		}
		val := reflect.ValueOf(function)
		funcType := val.Type()
		if funcType.NumOut() != 1 {
			continue
		}

		funcSettings[name] = function
	}

	raymond.RegisterHelpers(funcSettings)
}

func invalidHelper(name string) bool {
	invalids := []string{
		"buildCustomCert",
		"decryptAES",
		"derivePassword",
		"encryptAES",
		"fail",
		"genCA",
		"genPrivateKey",
		"genSelfSignedCert",
		"genSignedCert",
		"hello",
		"mustAppend",
		"mustCompact",
		"mustDateModify",
		"mustDeepCopy",
		"mustFirst",
		"mustHas",
		"mustInitial",
		"mustLast",
		"mustMerge",
		"mustMergeOverwrite",
		"mustPrepend",
		"mustPush",
		"mustRegexFind",
		"mustRegexFindAll",
		"mustRegexMatch",
		"mustRegexReplaceAll",
		"mustRegexReplaceAllLiteral",
		"mustRegexSplit",
		"mustRest",
		"mustReverse",
		"mustSlice",
		"mustToDate",
		"mustToJson",
		"mustToPrettyJson",
		"mustToRawJson",
		"mustUniq",
		"mustWithout",
		"must_date_modify",
		"semver",
		"semverCompare",
		"trimall",
	}

	for _, invalid := range invalids {
		if name == invalid {
			return true
		}
	}

	return false
}
