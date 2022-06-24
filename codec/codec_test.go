package codec

import (
	"testing"
)

type ColorGroup struct {
	Id     int      `json:"id" xml:"id,attr" msg:"id"`
	Name   string   `json:"name" xml:"name" msg:"name"`
	Colors []string `json:"colors" xml:"colors" msg:"colors"`
}

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var pbGroup = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

func BenchmarkJSONCodec_Encode(b *testing.B) {
	var raw = make([]byte, 0, 1024)
	serializer := JSONCodec{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		raw, _ = serializer.Encode(group)
	}
	b.ReportMetric(float64(len(raw)), "bytes")
}

func BenchmarkPBCodec_Encode(b *testing.B) {
	var raw = make([]byte, 0, 1024)
	serializer := PBCodec{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		raw, _ = serializer.Encode(&pbGroup)
	}
	b.ReportMetric(float64(len(raw)), "bytes")
}
