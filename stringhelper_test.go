package stringhelper

import "testing"

func TestCamelize(t *testing.T) {
	tests := []string{
		"kev-kev",
		"kev.kev",
		"kev_kev",
		"kev kev",
		"kevKev",
		"KevKev",
	}

	for _, test := range tests {
		t.Run(test, func(st *testing.T) {
			want := "kevKev"
			got := Camelize(test)
			if got != want {
				st.Errorf("wanted %s, got %s", want, got)
			}
		})
	}
}

func TestPascalize(t *testing.T) {
	tests := []string{
		"kev-kev",
		"kev.kev",
		"kev_kev",
		"kev kev",
		"kevKev",
		"KevKev",
	}

	for _, test := range tests {
		t.Run(test, func(st *testing.T) {
			want := "KevKev"
			got := Pascalize(test)
			if got != want {
				st.Errorf("wanted %s, got %s", want, got)
			}
		})
	}
}

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
