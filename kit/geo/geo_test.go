package geo_test

import (
	"testing"

	"github.com/burubur/fastcampus/kit/geo"
	"github.com/stretchr/testify/assert"
)

func TestHaversine(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	testCases := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "calculate distance between monas and stasiun kota",
			args: args{
				lat1: -6.176421464109725,
				lon1: 106.8230166265814,
				lat2: -6.136538249584232,
				lon2: 106.81373546121458,
			},
			want: 4.3,
		},
		{
			name: "calculate distance stasiun kota and somewhere else",
			args: args{
				lat1: -6.136538249584232,
				lon1: 106.81373546121458,
				lat2: -6.134866061903904,
				lon2: 106.83123886704062,
			},
			want: 1.5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := geo.Haversine(tc.args.lat1, tc.args.lon1, tc.args.lat2, tc.args.lon2)
			assert.GreaterOrEqual(t, got, tc.want, "the distance should be greater or equal the expected number")
		})
	}
}

func BenchmarkHaversine(b *testing.B) {
	for i := 0; i < 1000; i++ {
		geo.Haversine(-6.176421464109725, 106.8230166265814, -6.136538249584232, 106.81373546121458)
	}
}

func BenchmarkSphericalLawOfCosines(b *testing.B) {
	for i := 0; i < 1000; i++ {
		geo.SphericalLawOfCosines(-6.176421464109725, 106.8230166265814, -6.136538249584232, 106.81373546121458)
	}
}
