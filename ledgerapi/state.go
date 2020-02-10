// Copyright the Hyperledger Fabric contributors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ledgerapi

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateCompositeKey creates a composite key
func CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return shim.CreateCompositeKey(objectType, attributes)
}

// StateInterface placeholder
type StateInterface interface {
	GetValue() ([]byte, error)
	GetKey() string
	GetSplitKey() (string, []string, error)
}

// StateIteratorInterface placeholder
type StateIteratorInterface interface{}

// State implements StateInterface
type State struct {
	collection string
	ctx        contractapi.TransactionContextInterface
	key        string
	value      []byte
}

// GetValue returns the byte value of the state
func (s *State) GetValue() ([]byte, error) {
	if s.value != nil {
		return s.value, nil
	}

	bytes, err := s.ctx.GetStub().GetState(s.key)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s in collection %s. %s", s.key, s.collection, err.Error())
	}

	return bytes, nil
}

// GetKey returns the key of the state
func (s *State) GetKey() string {
	return s.key
}

// GetSplitKey returns the component parts that formed the composite key of the
// state
func (s *State) GetSplitKey() (string, []string, error) {
	stub := shim.ChaincodeStub{}
	objectType, attrs, _ := stub.SplitCompositeKey(s.key)

	return objectType, attrs, nil
}
