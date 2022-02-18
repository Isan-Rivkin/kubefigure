package inputs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/isan-rivkin/kubefigure/sources"
	log "github.com/sirupsen/logrus"
)

var (
	AppName             = "kubefigure"
	configPrefixPattern = fmt.Sprintf("com.%s.", AppName)
	jsonPathDelim       = "$"
	tokenToSource       = map[string]sources.SourceType{
		"!terraform": sources.TerraformSource,
		"!vault":     sources.VaultSource,
		"!consulkv":  sources.ConsulSource,
	}
)

func trimQuotes(out string) string {

	out = strings.TrimRight(strings.TrimLeft(out, "'"), "'")
	out = strings.TrimRight(strings.TrimLeft(out, "\""), "\"")
	return out
}
func getStringAfter(line, substrRegex string) (string, []int, bool) {
	var out string
	m := regexp.MustCompile(substrRegex)
	r := m.FindStringIndex(line)
	if len(r) < 2 {
		return "", nil, false
	}

	if r[0] == r[1] {
		splitted := strings.SplitAfter(line, substrRegex)
		if len(splitted) > 1 {
			out = splitted[1]
		}
	} else {
		//out = strings.TrimLeft(line, line[:r[1]])
		splitted := strings.SplitAfter(line, substrRegex)
		if len(splitted) > 1 {
			out = strings.TrimSpace(splitted[1])
		}
	}

	return out, r, true
}

func getValueAfterSource(source sources.SourceType, line string) (string, string, bool) {

	sourceVal, _, found := getStringAfter(line, string(source))

	if !found {
		return "", "", found
	}

	sourceVal = trimQuotes(sourceVal)

	valIndicator, _, delimFound := getStringAfter(sourceVal, jsonPathDelim)

	if delimFound {
		sourceVal = strings.SplitAfter(sourceVal, jsonPathDelim)[0]
		sourceVal = strings.TrimRight(sourceVal, jsonPathDelim)
	}

	return sourceVal, valIndicator, true
}

func getSource(line string) (sources.SourceType, bool) {

	for token, source := range tokenToSource {
		match, err := regexp.MatchString(token, line)
		if err != nil {
			log.WithError(err).Error("failed finding source in raw line, failed regex")
			continue
		}
		if match {
			return source, true
		}
	}
	return "", false
}

func getSourceConfig(line string) {
	for _, source := range tokenToSource {
		pat := fmt.Sprintf("%s%s", configPrefixPattern, string(source))
		match, err := regexp.MatchString(pat, line)
		if err != nil {
			log.WithError(err).Error("failed finding source in raw line, failed regex")
			continue
		}
		if match {
			fmt.Println("found config source!! ", line)
		}
	}
}

func ScanLine(line string) (*ValueToken, bool) {
	getSourceConfig(line)
	if source, found := getSource(line); found {
		if sourceVal, jsonPath, found := getValueAfterSource(source, line); found {
			return NewValueToken(sourceVal, jsonPath, source), true
		}
	}
	return nil, false
}

func ScanRawFromFile(path string) error {
	var valTokens []*ValueToken
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lineTxt := scanner.Text()
		if t, found := ScanLine(lineTxt); found {
			valTokens = append(valTokens, t)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range valTokens {
		fmt.Println(fmt.Sprintf("%s => %s : %s", v.Type, v.Value, v.ValueIndicatorPath))
	}
	return nil
}
