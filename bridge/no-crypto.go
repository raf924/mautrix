// Copyright (c) 2022 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build !cgo || nocrypto

package bridge

import (
	"errors"
)

func NewCryptoHelper(bridge *Bridge) Crypto {
	if bridge.Config.Bridge.GetEncryptionConfig().Allow {
		bridge.Log.Warnln("Bridge built without end-to-bridge encryption, but encryption is enabled in config")
	} else {
		bridge.Log.Debugln("Bridge built without end-to-bridge encryption")
	}
	return nil
}

var NoSessionFound = errors.New("nil")
