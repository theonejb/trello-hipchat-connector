package hipchat

import "encoding/json"
import "fmt"
import "log"
import "net/http"
import "net/url"

type Api struct {
	apiKey string
	apiUrl string
	logger *log.Logger
}

type Room struct {
	Api
	roomId   string
	RoomName string
}

func NewRoom(apiKey string, roomId string, roomName string, logger *log.Logger) *Room {
	urlWithAuthToken := fmt.Sprintf("%s?auth_token=%s", roomPostUrl, apiKey)
	return &Room{Api: Api{apiKey: apiKey,
		apiUrl: urlWithAuthToken, logger: logger}, roomId: roomId, RoomName: roomName}
}

func (room *Room) PostMessage(message, from, msgFormat, notify, color string) error {
	messageData := &url.Values{"room_id": {room.roomId}, "from": {from},
		"message": {message}, "message_format": {msgFormat},
		"notify": {notify}, "color": {color}, "format": {"json"}}

	resp, err := http.PostForm(room.apiUrl, *messageData)
	defer resp.Body.Close()
	if err != nil {
		room.logger.Println("Error encountered while posting message")
		return err
	}

	body := make([]byte, resp.ContentLength)
	if _, err = resp.Body.Read(body); err != nil {
		room.logger.Println("Error encountered while reading response message body")
		return err
	}

	decodedResponse, err := convertJsonResponseToStruct(body)
	if err != nil {
		room.logger.Println("Error encountered while converting response body to Json")
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
	apiErrorData `json:"error"`
	Status       string
}

func (ds *apiDecodedResponse) Error() string {
	return fmt.Sprintf("[Hipchat Error] Code: '%d' Type: '%s' Message: '%s'", ds.apiErrorData.Code,
		ds.apiErrorData.Type, ds.apiErrorData.Message)
}

func convertJsonResponseToStruct(responseBody []byte) (*apiDecodedResponse, error) {
	ds := &apiDecodedResponse{}
	parsedJson := &apiDecodedResponse{} //make(map[string]interface{})
	if err := json.Unmarshal(responseBody, &parsedJson); err != nil {
		return nil, err
	}
	return ds, nil
}
