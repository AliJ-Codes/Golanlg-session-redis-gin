package session

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func CreateSessionID() (string, error){
	session_id := make([]byte, 32)
	if _,err := io.ReadFull(rand.Reader, session_id); err != nil{
		return "", err
	}
	return hex.EncodeToString(session_id), nil
}