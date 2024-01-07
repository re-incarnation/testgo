package stalcraftgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	stalcraftgo "github.com/re-incarnation/stalcraftgo/object"
)

type EXBO struct {
	ClientId   string
	Url        string
	EXBOApiKey string
}

type Params map[string]interface{}

func Init(clientId string, url string, apiKey string) *EXBO {
	//exp url: http://exbo.net (without last /)
	var exbo EXBO
	exbo.ClientId = clientId
	exbo.Url = url
	exbo.EXBOApiKey = apiKey

	return &exbo
}

func (exbo *EXBO) RequestAuth(params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprint("https://eapi.stalcraft.net/")

	req, err := http.NewRequest("POST", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestRegions() (response []byte, err error) {
	req, err := http.NewRequest("GET", "", &bytes.Buffer{})
	if err != nil {
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		var e stalcraftgo.Error

		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}

	return body, nil
}

func (exbo *EXBO) RequestEmission(Region string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf("https://eapi.stalcraft.net/", Region, "/emission")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestAuctionItemPriceHistory(Region string, Item string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf(exbo.Url, "%s/%s", Region, "/auction", "%s/%s", Item, "/history")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestAuctionActiveItemLots(Region string, Item string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf(exbo.Url, "%s/%s", Region, "/auction", "%s/%s", Item, "/lots")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestCharacterProfile(Region string, Character string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf(exbo.Url, "%s/%s", Region, "/character/by-name/", Character, "/profile")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestListOfCharacter(VPadlu string) {}

func (exbo *EXBO) RequestClanInformation(Region string, ClanId string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf(exbo.Url, "%s/%s", Region, "/clan/", ClanId, "/info")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestClanMembers(Region string, ClanId string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf(exbo.Url, "%s/%s", Region, "/clan/", ClanId, "/members")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestClanList(Region string, params Params) (response []byte, err error) {
	querry := url.Values{}
	FormatParams(&querry, params)
	reqRaw := bytes.NewBufferString(querry.Encode())

	Url := fmt.Sprintf(exbo.Url, "%s/%s", Region, "/clans")
	req, err := http.NewRequest("GET", Url, reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("Client-Id", exbo.ClientId)
	req.Header.Set("Client-Secret", exbo.EXBOApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e stalcraftgo.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}
	return body, nil
}

func (exbo *EXBO) RequestUnmarshalRegions(params Params, obj interface{}) error {
	resp, err := exbo.RequestRegions()
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalEmission(Region string, params Params, obj interface{}) error {
	resp, err := exbo.RequestEmission(Region, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalAuctionItemPriceHistory(Region string, Item string, params Params, obj interface{}) error {
	resp, err := exbo.RequestAuctionItemPriceHistory(Region, Item, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalAuctionActiveItemLots(Region string, Item string, params Params, obj interface{}) error {
	resp, err := exbo.RequestAuctionActiveItemLots(Region, Item, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalCharacterProfile(Region string, Character string, params Params, obj interface{}) error {
	resp, err := exbo.RequestCharacterProfile(Region, Character, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

/*
func (exbo *EXBO) RequestUnmarshalListOfCharacters(VPadlu, params Params, obj interface{}) error {

		resp, err := exbo.RequestListOfCharacter(VPadlu)
		if err != nil {
			return err
		}

		return json.Unmarshal(resp, &obj)
}
*/

func (exbo *EXBO) RequestUnmarshalClanInformation(Region string, ClanId string, params Params, obj interface{}) error {
	resp, err := exbo.RequestClanInformation(Region, ClanId, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalClanMembers(Region string, ClanId string, params Params, obj interface{}) error {
	resp, err := exbo.RequestClanMembers(Region, ClanId, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalClanList(Region string, params Params, obj interface{}) error {
	resp, err := exbo.RequestClanList(Region, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func (exbo *EXBO) RequestUnmarshalAuth(params Params, obj interface{}) error {
	resp, err := exbo.RequestAuth(params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

func FormatParams(q *url.Values, params Params) {
	for k, v := range params {
		key, val := FormatValue(v, k, q)
		if key != "" && val != "" {
			q.Set(key, val)
		}
	}
}

func FormatValue(value interface{}, key string, q *url.Values) (k, v string) {
	if value == nil || key == "" {
		return
	}

	switch iface := value.(type) {
	case bool:
		return key, fmt.Sprintf("%t", iface)
	case Params:
		for kI, vI := range iface {
			keyAns, valAns := FormatValue(vI, kI, q)
			q.Set(fmt.Sprintf("%s[%s]", key, keyAns), valAns)
		}
	case []int:
		for _, val := range iface {
			if q.Get(fmt.Sprintf("%s[]", key)) != "" {
				q.Add(fmt.Sprintf("%s[]", key), strconv.Itoa(val))
			} else {
				q.Set(fmt.Sprintf("%s[]", key), strconv.Itoa(val))
			}
		}
	default:
		return key, fmt.Sprintf("%v", iface)
	}

	return
}
