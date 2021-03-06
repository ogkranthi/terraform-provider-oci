// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// EncryptDataDetails The representation of EncryptDataDetails
type EncryptDataDetails struct {

	// The OCID of the key to encrypt with.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The plaintext data to encrypt.
	Plaintext *string `mandatory:"true" json:"plaintext"`

	// Information that can be used to provide an encryption context for the
	// encrypted data. The length of the string representation of the associated data
	// must be fewer than 4096 characters.
	AssociatedData map[string]string `mandatory:"false" json:"associatedData"`

	// Information that provides context for audit logging. You can provide this additional
	// data as key-value pairs to include in the audit logs when audit logging is enabled.
	LoggingContext map[string]string `mandatory:"false" json:"loggingContext"`
}

func (m EncryptDataDetails) String() string {
	return common.PointerString(m)
}
