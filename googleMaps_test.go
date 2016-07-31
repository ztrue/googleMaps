package googleMaps

import "testing"

func TestLink(t *testing.T) {
  var tests = []struct {
    lat, lon, zoom float64
    expected string
  }{
    {39.1943744, -92.5944788, 6.61, "https://www.google.com/maps/@39.1943744,-92.5944788,6.61z"},
    {44.7841814, 20.4479806, 14.12, "https://www.google.com/maps/@44.7841814,20.4479806,14.12z"},
    {-35.3065367, 141.2022485, 4.21, "https://www.google.com/maps/@-35.3065367,141.2022485,4.21z"},
  }
  for i, c := range tests {
    actual := Link(c.lat, c.lon, c.zoom)
    if actual != c.expected {
      t.Errorf("Link(%v, %v, %v) == %v, expected %v in case %v", c.lat, c.lon, c.zoom, actual, c.expected, i)
    }
  }
}
