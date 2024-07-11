package hashpalte

import "testing"

func TestHashplate(t *testing.T) {
	type args struct {
		value    string
		hasEmoji []bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hello World!",
			args: args{
				value:    "Hello World!",
				hasEmoji: []bool{},
			},
			want: "渝F·WGVA2",
		},
		{
			name: "Hello World! emoji",
			args: args{
				value:    "Hello World!",
				hasEmoji: []bool{true},
			},
			want: "🍢 渝F·WGVA2 🪣",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hashplate(tt.args.value, tt.args.hasEmoji...); got != tt.want {
				t.Errorf("Hashplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
