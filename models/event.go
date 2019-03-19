/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"

	"github.com/ugorji/go/codec"
)

/*
 * This file is for the Event model in EdgeX
 *
 *
 * Event struct to hold event data
 */
type Event struct {
	ID       string    `json:"id" codec:"omitempty"`
	Pushed   int64     `json:"pushed" codec:"omitempty"`
	Device   string    `json:"device" codec:"omitempty"` // Device identifier (name or id)
	Created  int64     `json:"created" codec:"omitempty"`
	Modified int64     `json:"modified" codec:"omitempty"`
	Origin   int64     `json:"origin" codec:"omitempty"`
	Readings []Reading `json:"readings" codec:"omitempty"` // List of readings
}

func encodeAsCBOR(e Event) ([]byte, error) {
	var handle codec.CborHandle
	var byteBuffer = make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&byteBuffer, &handle)

	err := enc.Encode(e)
	if err != nil {
		return []byte{}, err
	}

	return byteBuffer, nil
}

// Custom marshaling to make empty strings null
func (e Event) MarshalJSON() ([]byte, error) {
	test := struct {
		ID       *string   `json:"id,omitempty"`
		Pushed   int64     `json:"pushed,omitempty"`
		Device   *string   `json:"device,omitempty"` // Device identifier (name or id)
		Created  int64     `json:"created,omitempty"`
		Modified int64     `json:"modified,omitempty"`
		Origin   int64     `json:"origin,omitempty"`
		Readings []Reading `json:"readings,omitempty"` // List of readings
	}{
		Pushed:   e.Pushed,
		Created:  e.Created,
		Modified: e.Modified,
		Origin:   e.Origin,
	}

	// Empty strings are null
	if e.ID != "" {
		test.ID = &e.ID
	}
	if e.Device != "" {
		test.Device = &e.Device
	}

	// Empty arrays are null
	if len(e.Readings) > 0 {
		test.Readings = e.Readings
	}

	return json.Marshal(test)
}

func (e Event) String() string {
	out, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(out)
}

func (e Event) CBOR() []byte {
	cbor, err := encodeAsCBOR(e)
	if err != nil {
		return []byte{}
	}

	return cbor
}
