package vatansms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendToService(url string, body interface{}, response interface{}) error {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fmt.Println(string(requestBody))
	rawResponse, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer rawResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBody))

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return err
	}

	return nil
}

func (settings Settings) GetSettings() (resp SettingsResponse, err error) {
	var response SettingsResponse
	err = sendToService(SettingsUrl, settings, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (oneToNSms OneToN) SendOneToN() (resp SmsResponse, err error) {
	var response SmsResponse
	err = sendToService(OneToNUrl, oneToNSms, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (nToNSms NToNSms) SendNToNSms() (resp SmsResponse, err error) {
	var response SmsResponse
	err = sendToService(NToNUrl, nToNSms, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (reportStats ReportStats) GetReportStats() (resp ReportStatsResponse, err error) {
	var response ReportStatsResponse
	err = sendToService(ReportStatsUrl, reportStats, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (reportDetails ReportDetails) GetReportDetails() (resp ReportDetailsResponse, err error) {
	var response ReportDetailsResponse
	err = sendToService(ReportDetailsUrl, reportDetails, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
