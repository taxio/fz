package fz

import "testing"

func TestCalcLevenshteinDistance(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "apple vs play",
			args: args{
				s1: "apple",
				s2: "play",
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcLevenshteinDistance(tt.args.s1, tt.args.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcLevenshteinDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcLevenshteinDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
