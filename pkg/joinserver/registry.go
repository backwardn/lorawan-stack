// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
)

type DeviceRegistry interface {
	GetByEUI(ctx context.Context, joinEUI types.EUI64, devEUI types.EUI64) (*ttnpb.EndDevice, error)
	SetByEUI(ctx context.Context, joinEUI types.EUI64, devEUI types.EUI64, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, error)) (*ttnpb.EndDevice, error)
}

func DeleteDevice(ctx context.Context, r DeviceRegistry, joinEUI types.EUI64, devEUI types.EUI64) error {
	_, err := r.SetByEUI(ctx, joinEUI, devEUI, func(*ttnpb.EndDevice) (*ttnpb.EndDevice, error) { return nil, nil })
	return err
}

func CreateDevice(ctx context.Context, r DeviceRegistry, dev *ttnpb.EndDevice) (*ttnpb.EndDevice, error) {
	if dev.EndDeviceIdentifiers.JoinEUI == nil || dev.EndDeviceIdentifiers.DevEUI == nil {
		return nil, errInvalidIdentifiers
	}
	dev, err := r.SetByEUI(ctx, *dev.EndDeviceIdentifiers.JoinEUI, *dev.EndDeviceIdentifiers.DevEUI, func(stored *ttnpb.EndDevice) (*ttnpb.EndDevice, error) {
		if stored != nil {
			return nil, errDuplicateIdentifiers
		}
		return dev, nil
	})
	if err != nil {
		return nil, err
	}
	return dev, nil
}

type KeyRegistry interface {
	GetByID(ctx context.Context, devEUI types.EUI64, id string) (*ttnpb.SessionKeys, error)
	SetByID(ctx context.Context, devEUI types.EUI64, id string, f func(*ttnpb.SessionKeys) (*ttnpb.SessionKeys, error)) (*ttnpb.SessionKeys, error)
}

func DeleteKeys(ctx context.Context, r KeyRegistry, devEUI types.EUI64, id string) error {
	_, err := r.SetByID(ctx, devEUI, id, func(*ttnpb.SessionKeys) (*ttnpb.SessionKeys, error) { return nil, nil })
	return err
}

func CreateKeys(ctx context.Context, r KeyRegistry, devEUI types.EUI64, ks *ttnpb.SessionKeys) (*ttnpb.SessionKeys, error) {
	if devEUI.IsZero() || ks.SessionKeyID == "" {
		return nil, errInvalidIdentifiers
	}
	ks, err := r.SetByID(ctx, devEUI, ks.SessionKeyID, func(stored *ttnpb.SessionKeys) (*ttnpb.SessionKeys, error) {
		if stored != nil {
			return nil, errDuplicateIdentifiers
		}
		return ks, nil
	})
	if err != nil {
		return nil, err
	}
	return ks, nil
}