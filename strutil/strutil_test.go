package strutil

import "testing"

func TestEqualNoStrict(t *testing.T) {
	table := []struct {
		cityName1 string
		cityName2 string
		expected  bool
	}{
		{
			cityName1: "Buenos Aires",
			cityName2: "buenos aires",
			expected:  true,
		},
		{
			cityName1: "Buenos Aires",
			cityName2: "Buenos Aires",
			expected:  true,
		},
		{
			cityName1: "Buenos Aires",
			cityName2: " Buenos Aires ",
			expected:  true,
		},
		{
			cityName1: "Buenos Aires",
			cityName2: "\n\t\r Büénòs Áïrés	\n\t\r",
			expected:  true,
		},
		{
			cityName1: "Buenos Aires",
			cityName2: "Aires Buenos",
			expected:  false,
		},
	}

	for _, test := range table {
		if EqualNoStrict(test.cityName1, test.cityName2) != test.expected {
			t.Errorf(
				`Expected "%s=%s->%v", got %v`,
				test.cityName1,
				test.cityName2,
				test.expected,
				!test.expected,
			)
		}
	}
}
