package internal

import (
	"awesomeProject/crypto"
	"awesomeProject/utils"
	"bytes"
	"encoding/json"
	"fmt"
)

func CreateInvitation() (string, error) {
	cert, err := crypto.CertFingerprint("cert.pem")
	if err != nil {
		return "", err
	}

	inv := map[string]interface{}{"Cert": cert, "Addr": utils.Settings.Ip}

	invByrs, err := json.Marshal(inv)
	if err != nil {
		return "", err
	}

	return crypto.ToBase64(invByrs), nil
}

func ShowInvitation(invCode string) {
	var inv map[string]interface{}
	json.Unmarshal([]byte(crypto.FromBase64(invCode)), &inv)
	utils.Print(inv, 1)
	fmt.Println(inv)
}

func FingerprintFormat(fingerprint string) {
	var buf bytes.Buffer

	for i, f := range fingerprint {
		if i > 0 {
			fmt.Fprintf(&buf, ":")
		}
		fmt.Fprintf(&buf, "%02X", f)
	}
}
