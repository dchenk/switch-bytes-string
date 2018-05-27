package switch_bytes_string

import "testing"

func BenchmarkStr_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Str("a")
	}
}

func BenchmarkStr_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Str("abcdefgabcdefgabcdefgabcdefgabcde")
	}
}

func BenchmarkStr_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Str("abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga")
	}
}

func BenchmarkStrConst_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = StrConst("a")
	}
}

func BenchmarkStrConst_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = StrConst("abcdefgabcdefgabcdefgabcdefgabcde")
	}
}

func BenchmarkStrConst_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = StrConst("abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga")
	}
}

var btsShort = []byte("abcdefgijabcdefgijabcdefg")

func BenchmarkBts_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Bts(btsShort)
	}
}

var btsMed = []byte("abcdefgabcdefgabcdefgabcdefgabcde")

func BenchmarkBts_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Bts(btsMed)
	}
}

var btsLong = []byte("abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga")

func BenchmarkBts_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Output = Bts(btsLong)
	}
}
