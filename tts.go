package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
)

var client = initClient()

func initClient() *tts.Client {
	config, err := LoadConfig()
	ep(err)

	profile := profile.NewClientProfile()
	c, err := tts.NewClient(common.NewCredential(config.SecretId, config.SecretKey), config.Region, profile)
	ep(err)
	return c
}

func hash(text string) string {
	i := md5.New()
	i.Write([]byte(text))
	bytes := i.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}

func convert(text string, chatName string, speed string) []byte {
	sessionId := hash(text)
	voiceType := parseCharName(chatName)
	speedV := parseSpeed(speed)

	req := tts.NewTextToVoiceRequest()
	req.Text = &text
	req.SessionId = &sessionId
	req.VoiceType = &voiceType
	req.Speed = &speedV

	rsp, err := client.TextToVoice(req)
	ep(err)

	bytes, err := base64.StdEncoding.DecodeString(*rsp.Response.Audio)
	ep(err)
	return bytes
}

func parseCharName(name string) int64 {
	i, err := strconv.Atoi(name)
	if err != nil {
		return int64(601008)
	}

	return int64(i)
}

func parseSpeed(speed string) float64 {
	i, err := strconv.Atoi(speed)
	if err != nil {
		return 0
	}

	if i < 10 {
		return float64(i)*2/5 - 4
	} else {
		return float64(i)*3/20 - 1.5
	}

}
