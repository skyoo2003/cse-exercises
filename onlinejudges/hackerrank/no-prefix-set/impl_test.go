package noprefixset

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestNoPrefixSet1(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`7
		aab
		defgab
		abcde
		aabcde
		cedaaa
		bbbbbbbbbb
		jabjjjad`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	NoPrefixSet(in)
}

func TestNoPrefixSet2(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`4
		aab
		aac
		aacghgh
		aabghgh`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	NoPrefixSet(in)
}

func TestNoPrefixSet3(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`100
		hgiiccfchbeadgebc
		biiga
		edchgb
		ccfdbeajaeid
		ijgbeecjbj
		bcfbbacfbfcfbhcbfjafibfhffac
		ebechbfhfcijcjbcehbgbdgbh
		ijbfifdbfifaidje
		acgffegiihcddcdfjhhgadfjb
		ggbdfdhaffhghbdh
		dcjaichjejgheiaie
		d
		jeedfch
		ahabicdffbedcbdeceed
		fehgdfhdiffhegafaaaiijceijdgbb
		beieebbjdgdhfjhj
		ehg
		fdhiibhcbecddgijdb
		jgcgafgjjbg
		c
		fiedahb
		bhfhjgcdbjdcjjhaebejaecdheh
		gbfbbhdaecdjaebadcggbhbchfjg
		jdjejjg
		dbeedfdjaghbhgdhcedcj
		decjacchhaciafafdgha
		a
		hcfibighgfggefghjh
		ccgcgjgaghj
		bfhjgehecgjchcgj
		bjbhcjcbbhf
		daheaggjgfdcjehidfaedjfccdafg
		efejicdecgfieef
		ciidfbibegfca
		jfhcdhbagchjdadcfahdii
		i
		abjfjgaghbc
		bddeejeeedjdcfcjcieceieaei
		cijdgbddbceheaeececeeiebafgi
		haejgebjfcfgjfifhihdbddbacefd
		bhhjbhchdiffb
		jbbdhcbgdefifhafhf
		ajhdeahcjjfie
		idjajdjaebfhhaacecb
		bhiehhcggjai
		bjjfjhiice
		aif
		gbbfjedbhhhjfegeeieig
		fefdhdaiadefifjhedaieaefc
		hgaejbhdebaacbgbgfbbcad
		heghcb
		eggadagajjgjgaihjdigihfhfbijbh
		jadeehcciedcbjhdeca
		ghgbhhjjgghgie
		ibhihfaeeihdffjgddcj
		hiedaegjcdai
		bjcdcafgfjdejgiafdhfed
		fgdgjaihdjaeefejbbijdbfabeie
		aeefgiehgjbfgidcedjhfdaaeigj
		bhbiaeihhdafgaciecb
		igicjdajjdegbceibgebedghihihh
		baeeeehcbffd
		ajfbfhhecgaghgfdadbfbb
		ahgaccehbgajcdfjihicihhc
		bbjhih
		a
		cdfcdejacaicgibghgddd
		afeffehfcfiefhcagg
		ajhebffeh
		e
		hhahehjfgcj
		ageaccdcbbcfidjfc
		gfcjahbbhcbggadcaebae
		gi
		edheggceegiedghhdfgabgcd
		hejdjjbfacggdccgahiai
		ffgeiadgjfgecdbaebagij
		dgaiahge
		hdbaifh
		gbhccajcdebcig
		ejdcbbeiiebjcagfhjfdahbif
		g
		ededbjaaigdhb
		ahhhcibdjhidbgefggdjebfcf
		bdigjaehfchebiedajcjdh
		fjehjgbdbaiifi
		fbgigbdcbcgffdicfcidfdafghajc
		ccajeeijhhb
		gaaagfacgiddfahejhbgdfcfbfeedh
		gdajaigfbjcdegeidgaccjfi
		fghechfchjbaebcghfcfbdicgaic
		cfhigaciaehacdjhfcgajgbhhgj
		edhjdbdjccbfihiaddij
		cbbhagjbcadegicgifgghai
		hgdcdhieji
		fbifgbhdhagch
		cbgcdjea
		dggjafcajhbbbaja
		bejihed
		eeahhcggaaidifdigcfjbficcfhjj`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	NoPrefixSet(in)
	/*
		BAD SET
		d
	*/
}
