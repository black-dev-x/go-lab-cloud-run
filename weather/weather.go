package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/black-dev-x/go-lab-cloud-run/config"
)

type Weather struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func Get(location string) (*Weather, error) {

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", config.WEATHER_API_KEY, location)
	httpResp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != 200 {
		println("Error making request to weather API:", err)
		return nil, fmt.Errorf("Error getting weather data")
	}

	var response Weather
	err = json.NewDecoder(httpResp.Body).Decode(&response)
	return &response, err
}
