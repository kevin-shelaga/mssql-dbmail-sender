package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strings"
	"time"
)

//Config interface
var Config map[string]interface{}

//GetDateTime return date time now as string
func GetDateTime() string {
	return (time.Now().Format("2006-01-02 15:04:05.000"))
}

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Address: str}
	return strings.Trim(addr.String(), " <>")
}

//ComposeMimeMail return formatted email for smtp
func ComposeMimeMail(to string, from string, subject string, body string, importance string, format string) []byte {
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = encodeRFC2047(subject)
	header["X-Priority"] = importance
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = format
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(message)
}

//GetConfig return config as interface
func GetConfig(file string) map[string]interface{} {

	jsonFile, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &Config)

	return Config
}
