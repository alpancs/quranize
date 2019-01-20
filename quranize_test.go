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
		"tajri":                 []string{"تجري"},
		"alhamdulillah":         []string{"الحمد لله"},
		"bismillah":             []string{"بسم الله", "بشماله"},
		"wa'tasimu":             []string{"واعتصموا"},
		"wa'tasimu bihablillah": []string{"واعتصموا بحبل الله"},
		"shummun bukmun":        []string{"صم وبكم", "صم بكم", "الصم البكم"},
		"kahfi":                 []string{"الكهف"},
		"wabasyiris sobirin":    []string{"وبشر الصابرين"},
		"bissobri wassolah":     []string{"بالصبر والصلاة"},
		"ya aiyuhalladzina":     []string{"يا أيها الذين"},
		"syai in 'alim":         []string{"شيء عليم"},
		"'alal qoumil kafirin":  []string{"على القوم الكافرين"},
		"subhanalladzi asro":    []string{"سبحان الذي أسرى"},
		"sabbihisma robbikal":   []string{"سبح اسم ربك"},
		"hal ataka hadisul":     []string{"هل أتاك حديث"},
		"kutiba 'alaikumus":     []string{"كتب عليكم"},

		"bismillah hirrohman nirrohim":    []string{"بسم الله الرحمن الرحيم"},
		"alhamdu lillahi robbil 'alamin":  []string{"الحمد لله رب العالمين"},
		"arrohma nirrohim":                []string{"الرحمن الرحيم"},
		"maaliki yau middin":              []string{"مالك يوم الدين"},
		"iyya kanakbudu waiyya kanastain": []string{"إياك نعبد وإياك نستعين"},
		"ihdinash shirothol mustaqim":     []string{"اهدنا الصراط المستقيم"},
		"shirotholladzina an'am ta'alaihim ghoiril maghdzu bi'alaihim waladh dhollin": []string{"صراط الذين أنعمت عليهم غير المغضوب عليهم ولا الضالين"},
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
