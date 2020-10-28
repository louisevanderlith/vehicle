package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/vehicle/core"
	"net/http"
)

func FetchVehicleInfo(web *http.Client, host string, k hsk.Key) (core.Vehicle, error) {
	url := fmt.Sprintf("%s/info/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Vehicle{}, err
	}

	defer resp.Body.Close()

	result := core.Vehicle{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
