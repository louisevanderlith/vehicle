package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/vehicle/core"
	"io/ioutil"
	"net/http"
)

func FetchVehicleInfo(web *http.Client, host string, k hsk.Key) (core.Vehicle, error) {
	url := fmt.Sprintf("%s/info/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Vehicle{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Vehicle{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Vehicle{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllVehicles(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/info/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Vehicle{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

