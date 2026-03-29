//
// Copyright (C) 2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

// mergeExtensions merges the extensions map into the already-marshaled data bytes at top level.
// The extensions keys will overwrite existing keys if there is a conflict.
func mergeExtensions(data []byte, extensions map[string]any, unmarshalFn func([]byte, any) error,
	marshalFn func(any) ([]byte, error)) ([]byte, error) {
	var m map[string]any
	if err := unmarshalFn(data, &m); err != nil {
		return nil, err
	}
	for k, v := range extensions {
		m[k] = v
	}
	return marshalFn(m)
}
