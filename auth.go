package googleplay

import (
   "bytes"
   "github.com/89z/parse/protobuf"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const (
   Sleep = 16 * time.Second
   agent = "Android-Finsky (sdk=99,versionCode=99999999)"
)

type Auth struct {
   url.Values
}

func (a Auth) Delivery(dev *Device, app string, ver int) (*Delivery, error) {
   req, err := http.NewRequest("GET", origin + "/fdfe/delivery", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Get("Auth")},
      "User-Agent": {agent},
      "X-DFE-Device-ID": {dev.String()},
   }
   req.URL.RawQuery = url.Values{
      "doc": {app},
      "vc": {strconv.Itoa(ver)},
   }.Encode()
   res, err := roundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   wrap := new(responseWrapper)
   if err := protobuf.NewDecoder(buf).Decode(wrap); err != nil {
      return nil, err
   }
   return &wrap.Payload.DeliveryResponse, nil
}

func (a Auth) Details(dev *Device, app string) (*Details, error) {
   req, err := http.NewRequest("GET", origin + "/fdfe/details", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Get("Auth")},
      "X-DFE-Device-ID": {dev.String()},
   }
   req.URL.RawQuery = url.Values{
      "doc": {app},
   }.Encode()
   res, err := roundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   wrap := new(responseWrapper)
   if err := protobuf.NewDecoder(buf).Decode(wrap); err != nil {
      return nil, err
   }
   return &wrap.Payload.DetailsResponse, nil
}

// Purchase app. Only needs to be done once per Google account.
func (a Auth) Purchase(dev *Device, app string) error {
   buf := url.Values{
      "doc": {app},
   }.Encode()
   req, err := http.NewRequest(
      "POST", origin + "/fdfe/purchase", strings.NewReader(buf),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Get("Auth")},
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {agent},
      "X-DFE-Device-ID": {dev.String()},
   }
   res, err := roundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

// This seems to return `StatusOK`, even with invalid requests, and the response
// body only contains a token, that doesnt seem to indicate success or failure.
// Only way I know to check, it to try the `deviceID` with a `details` request
// or similar. Also, after the POST, you need to wait at least 16 seconds
// before the `deviceID` can be used.
func (a Auth) Upload(dev *Device, con Config) error {
   enc, err := protobuf.NewEncoder(con)
   if err != nil {
      return err
   }
   buf, err := enc.Encode()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", origin + "/fdfe/uploadDeviceConfig", bytes.NewReader(buf),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Get("Auth")},
      "User-Agent": {agent},
      "X-DFE-Device-ID": {dev.String()},
   }
   res, err := roundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

type Delivery struct {
   Status int32 `json:"1"`
   AppDeliveryData struct {
      DownloadURL string `json:"3"`
   } `json:"2"`
}

type Details struct {
   DocV2 struct {
      Details struct {
         AppDetails struct {
            DeveloperName string `json:"1"`
            VersionCode int32 `json:"3"`
            Version string `json:"4"`
            InstallationSize int64 `json:"9"`
            Permission []string `json:"10"`
         } `json:"1"`
      } `json:"13"`
   } `json:"4"`
}

type responseWrapper struct {
   Payload struct {
      DetailsResponse Details `json:"2"`
      DeliveryResponse Delivery `json:"21"`
   } `json:"1"`
}
