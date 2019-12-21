package quranize

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	quranizeTest Quranize
)

func TestMain(m *testing.M) {
	quranizeTest = NewDefaultQuranize()
	os.Exit(m.Run())
}

func TestEncodeEmptyString(t *testing.T) {
	input := ""
	expected := []string{}
	actual := quranizeTest.Encode(input)
	assert.Equal(t, expected, actual)
}

func TestEncodeNonAlquran(t *testing.T) {
	input := "alfan nur fauzan"
	expected := []string{}
	actual := quranizeTest.Encode(input)
	assert.Equal(t, expected, actual)
}

func TestEncodeAlquran(t *testing.T) {
	testCases := map[string][]string{
		"tajri":                 {"تجري"},
		"alhamdulillah":         {"الحمد لله"},
		"bismillah":             {"بسم الله", "بشماله"},
		"wa'tasimu":             {"واعتصموا"},
		"wa'tasimu bihablillah": {"واعتصموا بحبل الله"},
		"shummun bukmun":        {"صم وبكم", "صم بكم", "الصم البكم"},
		"kahfi":                 {"الكهف"},
		"wabasyiris sobirin":    {"وبشر الصابرين"},
		"bissobri wassolah":     {"بالصبر والصلاة"},
		"ya aiyuhalladzina":     {"يا أيها الذين"},
		"syai in 'alim":         {"شيء عليم"},
		"'alal qoumil kafirin":  {"على القوم الكافرين"},
		"subhanalladzi asro":    {"سبحان الذي أسرى"},
		"sabbihisma robbikal":   {"سبح اسم ربك"},
		"hal ataka hadisul":     {"هل أتاك حديث"},
		"kutiba 'alaikumus":     {"كتب عليكم"},

		"bismillah hirrohman nirrohim":    {"بسم الله الرحمن الرحيم"},
		"alhamdu lillahi robbil 'alamin":  {"الحمد لله رب العالمين"},
		"arrohma nirrohim":                {"الرحمن الرحيم"},
		"maaliki yau middin":              {"مالك يوم الدين"},
		"iyya kanakbudu waiyya kanastain": {"إياك نعبد وإياك نستعين"},
		"ihdinash shirothol mustaqim":     {"اهدنا الصراط المستقيم"},
		"shirotholladzina an'am ta'alaihim ghoiril maghdzu bi'alaihim waladh dhollin": {"صراط الذين أنعمت عليهم غير المغضوب عليهم ولا الضالين"},
	}
	for input, expected := range testCases {
		actual := quranizeTest.Encode(input)
		assert.Equalf(t, expected, actual, "input: %s", input)
	}
}

func ExampleQuranize_Encode() {
	quranize := NewDefaultQuranize()
	fmt.Println(quranize.Encode("alhamdulillah hirobbil 'alamin"))
	// Output: [الحمد لله رب العالمين]
}

func ExampleQuranize_Locate() {
	quranize := NewDefaultQuranize()
	fmt.Println(quranize.Locate("الحمد لله رب العالمين"))
	// Output: [{1 2 0} {10 10 10} {39 75 13} {40 65 10}]
}

func TestLocateEmptyString(t *testing.T) {
	input := ""
	expected := zeroLocs
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateNonAlquran(t *testing.T) {
	input := "alfan"
	expected := zeroLocs
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquran(t *testing.T) {
	input := "بسم الله الرحمن الرحيم"
	expected := []Location{NewLocation(1, 1, 0), NewLocation(27, 30, 4)}
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquranBeforeBuildIndex(t *testing.T) {
	root := quranizeTest.root
	defer func() { quranizeTest.root = root }()
	quranizeTest.root = nil
	input := "بسم الله الرحمن الرحيم"
	expected := zeroLocs
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}
