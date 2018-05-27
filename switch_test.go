package switch_bytes_string

import "testing"

var strShortV = "a"

func BenchmarkStr_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Str(strShortV)
	}
}

var strMedV = "abcdefgabcdefgabcdefgabcdefgabcde"

func BenchmarkStr_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Str(strMedV)
	}
}

var strLongV = "abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga"

func BenchmarkStr_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Str(strLongV)
	}
}

func BenchmarkStrConst_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = StrConst(strShortV)
	}
}

func BenchmarkStrConst_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = StrConst(strMedV)
	}
}

func BenchmarkStrConst_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = StrConst(strLongV)
	}
}

var btsShort = []byte("abcdefgijabcdefgijabcdefg")

func BenchmarkBts_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Bts(btsShort)
	}
}

func BenchmarkBtsConst_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = BtsConst(btsShort)
	}
}

var btsMed = []byte("abcdefgabcdefgabcdefgabcdefgabcde")

func BenchmarkBts_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Bts(btsMed)
	}
}

func BenchmarkBtsConst_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = BtsConst(btsMed)
	}
}

var btsLong = []byte("abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga")

func BenchmarkBts_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Bts(btsLong)
	}
}

func BenchmarkBtsConst_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = BtsConst(btsLong)
	}
}
