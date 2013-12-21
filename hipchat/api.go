package hipchat

import "encoding/json"
import "fmt"
import "net/http"
import "net/url"

type HipchatApi struct {
	apiKey string
	apiUrl string
}

type HipchatRoom struct {
	HipchatApi
	roomId string
}

func NewHipchatRoom(apiKey string, roomId string) *HipchatRoom {
	urlWithAuthToken := fmt.Sprintf("%s?auth_token=%s", roomPostUrl, apiKey)
	return &HipchatRoom{HipchatApi: HipchatApi{apiKey: apiKey, apiUrl: urlWithAuthToken}, roomId: roomId}
}

func (room *HipchatRoom) PostMessage(message, from, msgFormat, notify, color string) error {
	messageData := &url.Values{"room_id": {room.roomId}, "from": {from},
		"message": {message}, "message_format": {msgFormat},
		"notify": {notify}, "color": {color}, "format": {"json"}}

	resp, err := http.PostForm(room.apiUrl, *messageData)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body := make([]byte, resp.ContentLength)
	if _, err = resp.Body.Read(body); err != nil {
		return err
	}

	decodedResponse, err := convertJsonResponseToStruct(body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return decodedResponse
	}

	return nil
}

type apiErrorData struct {
	Code    int
	Type    string
	Message string
}
type apiDecodedResponse struct {
	apiErrorData
	Status string
}

func (ds *apiDecodedResponse) Error() string {
	return fmt.Sprintf("[Hipchat Error] Code: '%d' Type: '%s' Message: '%s'", ds.apiErrorData.Code,
		ds.apiErrorData.Type, ds.apiErrorData.Message)
}

func convertJsonResponseToStruct(responseBody []byte) (*apiDecodedResponse, error) {
	ds := &apiDecodedResponse{}
	parsedJson := make(map[string]interface{})
	if err := json.Unmarshal(responseBody, &parsedJson); err != nil {
		return nil, err
	}
	if val, ok := parsedJson["status"]; ok {
		ds.Status = val.(string)
	}
	if val, ok := parsedJson["error"]; ok {
		errData := val.(map[string]interface{})
		ds.apiErrorData.Code = int(errData["code"].(float64))
		ds.apiErrorData.Type = errData["type"].(string)
		ds.apiErrorData.Message = errData["message"].(string)
	}

	return ds, nil
}
