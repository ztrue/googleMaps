package googleMaps

import "fmt"

const UrlPattern = "https://www.google.com/maps/@%v,%v,%vz"

func Link(lat, lon, zoom float64) string {
  return fmt.Sprintf(UrlPattern, lat, lon, zoom)
}
