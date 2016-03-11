package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateUUID is used to generate a random UUID
func GenerateUUID() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %v", err)
	}
	buf[6] = (buf[6] & 0x0F) | 0x40 // Version: https://tools.ietf.org/html/rfc4122#section-4.1.3
	buf[8] = (buf[8] & 0x3F) | 0x80 // Variant: https://tools.ietf.org/html/rfc4122#section-4.1.1

	return FormatUUID(buf)
}

func FormatUUID(buf []byte) (string, error) {
	if len(buf) != 16 {
		return "", fmt.Errorf("wrong length byte slice (%d)", len(buf))
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16]), nil
}

func ParseUUID(uuid string) ([]byte, error) {
	if len(uuid) != 36 {
		return nil, fmt.Errorf("uuid string is wrong length")
	}

	hyph := []byte("-")

	if uuid[8] != hyph[0] ||
		uuid[13] != hyph[0] ||
		uuid[18] != hyph[0] ||
		uuid[23] != hyph[0] {
		return nil, fmt.Errorf("uuid is improperly formatted")
	}

	hexStr := uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:36]

	ret, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	if len(ret) != 16 {
		return nil, fmt.Errorf("decoded hex is the wrong length")
	}

	return ret, nil
}
