package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WilayahService interface {
	GetProvinces() ([]map[string]interface{}, error)
	GetRegencies(provinceID string) ([]map[string]interface{}, error)
	GetDistricts(regencyID string) ([]map[string]interface{}, error)
	GetVillages(districtID string) ([]map[string]interface{}, error)
}

type wilayahService struct{}

func NewWilayahService() WilayahService {
	return &wilayahService{}
}

func (s *wilayahService) fetchWilayah(url string) ([]map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("gagal fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca response body: %v", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("gagal decode json: %v", err)
	}

	return data, nil
}

func (s *wilayahService) GetProvinces() ([]map[string]interface{}, error) {
	return s.fetchWilayah("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
}

func (s *wilayahService) GetRegencies(provinceID string) ([]map[string]interface{}, error) {
	return s.fetchWilayah(fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regencies/%s.json", provinceID))
}

func (s *wilayahService) GetDistricts(regencyID string) ([]map[string]interface{}, error) {
	return s.fetchWilayah(fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/districts/%s.json", regencyID))
}

func (s *wilayahService) GetVillages(districtID string) ([]map[string]interface{}, error) {
	return s.fetchWilayah(fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/villages/%s.json", districtID))
}
