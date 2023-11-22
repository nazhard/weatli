package program

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "os"
  
  "github.com/joho/godotenv"
)

type Weather struct {
  Location struct {
    Name    string `json:"name"`
    Country string `json:"country"`
  } `json:"location"`
  Current struct {
    TempC   float64 `json:"temp_c"`
    Condition struct {
      Text string `json:"text"`
    } `json:"condition"`
  } `json:"current"`
}

func Run() {
  q := "bandung"
  
  if len(os.Args) >= 2 {
    q = os.Args[1]
  }
  
  godotenv.Load(".env")
  key := os.Getenv("API_KEY")
  res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + q + "&days=1&aqi=no&alerts=no*")
  if err != nil {
    panic(err)
  }
  
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    panic("Weather API not avaiable")
  }
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    panic(err)
  }
//  fmt.Println(string(body))
  
  var weather Weather
  err = json.Unmarshal(body, &weather)
  if err != nil {
    panic(err)
  }
  
  location, current := weather.Location, weather.Current

  fmt.Printf(
    "%s, %s: %.0fC %s\n",
    location.Name,
    location.Country,
    current.TempC,
    current.Condition.Text,
  )
}