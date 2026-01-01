package utils

import "github.com/google/uuid"

func GetBytesFromUUID(uuid uuid.UUID) ([]byte, error) {
	bytes, err := uuid.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}