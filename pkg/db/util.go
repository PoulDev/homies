package db

import (
	"github.com/google/uuid"
)

func PreattyError(err error) error { // !! TODO
	return err
}

func UUID2Bytes(uid uuid.UUID) []byte {
	return uid[:]
}

func UUIDString2Bytes(s_uuid string) ([]byte, error) {
	new_uuid, err := uuid.Parse(s_uuid)
	if (err != nil) {
		return []byte{}, err
	}

	return UUID2Bytes(new_uuid), nil
}

func UUIDBytes2String(b_uuid []byte) (string, error) {
	new_uuid, err := uuid.FromBytes(b_uuid)
	if (err != nil) { return "", err; }

	return new_uuid.String(), nil
}
