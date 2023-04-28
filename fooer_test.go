package main

import "testing"

func Test_fooer(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"9 should be Foo", args{9}, "Foo"},
        {"3 should be Foo", args{3}, "Foo"},
        {"1 is not Foo", args{1}, "1"},
        {"0 should be Foo", args{0}, "Foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fooer(tt.args.input); got != tt.want {
				t.Errorf("fooer() = %v, want %v", got, tt.want)
			}
		})
	}
}
