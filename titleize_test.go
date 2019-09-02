package strings

import "testing"

func TestTitleize(t *testing.T) {
	tests := []struct{
		text string
		want string
	}{
		{
			text: "kev-kev",
			want: "Kev Kev",
		},
		{
			text: "kev.kev",
			want: "Kev Kev",
		},
		{
			text: "kev_kev",
			want: "Kev Kev",
		},
		{
			text: "kev kev",
			want: "Kev Kev",
		},
		{
			text: "kevKev",
			want: "Kev Kev",
		},
		{
			text: "KevKev",
			want: "Kev Kev",
		},
		{
			text: "kevKEV",
			want: "Kev KEV",
		},
	}

	for _, test := range tests {
		t.Run(test.text, func(st *testing.T) {
			got := Titleize(test.text)
			if got != test.want {
				st.Errorf("wanted %s, got %s", test.want, got)
			}
		})
	}
}
