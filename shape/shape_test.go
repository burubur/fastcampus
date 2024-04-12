package shape

import (
	"reflect"
	"testing"
)

func TestRectangle_Scale(t *testing.T) {
	type fields struct {
		Multipoint []Point
	}
	type args struct {
		ratio float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Rectangle
	}{
		{
			name: "1. scaled shape should be doubled from original rectangle",
			fields: fields{
				Multipoint: []Point{
					{0, 0},
					{0, 5},
					{5, 0},
					{5, 5},
				},
			},
			args: args{
				ratio: 2,
			},
			want: &Rectangle{
				[]Point{
					{0, 0},
					{0, 10},
					{10, 0},
					{5, 10},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				Multipoint: tt.fields.Multipoint,
			}
			if got := r.Scale(tt.args.ratio); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rectangle.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}
