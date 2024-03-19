package pwhelper

import (
	"encoding/base64"
	"fmt"
)

func DecodePw(encodedPw string) (string, error) {
    decodedText, err := base64.StdEncoding.DecodeString(encodedPw)
    if err != nil {
        return "", fmt.Errorf("error while decoding password: %s", err)
    }
    return string(decodedText), nil
}

func EncodePw(plainPw string) string {
    encodedText := base64.StdEncoding.EncodeToString([]byte(plainPw))
    return encodedText
}

