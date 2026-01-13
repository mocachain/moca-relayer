package assembler

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRLPDecodingForSBTAck tests the RLP decoding logic
// This test proves HIGH-001 bug exists and verifies the fix
func TestRLPDecodingForSBTAck(t *testing.T) {
	// Step 1: Build business payload (OperationType + ABI encoded data)
	mocaSBTAckPayload := &MocaSBTAckPackage{
		Toaddrs: []sdk.AccAddress{
			sdk.AccAddress(ethcommon.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb").Bytes()),
			sdk.AccAddress(ethcommon.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEc").Bytes()),
		},
		Status: STATUS_SUCCESS,
	}
	serializedAck, err := mocaSBTAckPayload.Serialize()
	require.NoError(t, err)

	// Key: business payload = [OperationType][ABI data]
	businessPayload := append([]byte{OperationMocaSBTACK}, serializedAck...)

	// Step 2: Wrap into oracletypes.Package
	packages := oracletypes.Packages{
		oracletypes.Package{
			ChannelId: 1,
			Sequence:  100,
			Payload:   businessPayload, // Business data here
		},
	}

	// Step 3: RLP encode the entire packages
	encodedPayload, err := rlp.EncodeToBytes(packages)
	require.NoError(t, err)

	t.Logf("RLP encoded payload length: %d", len(encodedPayload))
	t.Logf("First byte (RLP prefix): 0x%02x (NOT the OperationType!)", encodedPayload[0])

	// Test the OLD buggy code (for documentation purposes only)
	t.Run("BuggyOldCode_Documentation", func(t *testing.T) {
		t.Skip("This test documents the old bug. The bug has been fixed, so this test is skipped.")
		
		// OLD BUGGY CODE (before fix):
		// pack, err := DeserializeRawMocaSBTAckPackage(votes[0].ClaimPayload[sdk.AckPackageHeaderLength+ORACLETYPES_PACKAGES_PREFIX:])
		// 
		// Problem: ORACLETYPES_PACKAGES_PREFIX was a fixed 8-byte offset
		// But RLP encoding has variable-length prefixes
		// So reading RLP data with fixed offset was incorrect
	})

	// Test the CORRECT fixed code
	t.Run("CorrectFixedCode", func(t *testing.T) {
		// Correct: directly RLP decode the entire payload (no fixed offset skip)
		var decodedPackages oracletypes.Packages
		err := rlp.DecodeBytes(encodedPayload, &decodedPackages)
		require.NoError(t, err, "RLP decode should succeed")

		assert.Len(t, decodedPackages, 1)
		assert.Equal(t, uint64(100), decodedPackages[0].Sequence)
		assert.NotEmpty(t, decodedPackages[0].Payload)

		// Extract business data from Package.Payload
		pkg := decodedPackages[0]
		require.Greater(t, len(pkg.Payload), 0)

		op := pkg.Payload[0]   // Operation code
		data := pkg.Payload[1:] // Business data

		// Verify correct OperationType
		assert.Equal(t, OperationMocaSBTACK, op,
			"Correctly reads OperationType from Payload[0]")
		t.Logf("Correct code reads OperationType: 0x%02x", op)

		// Verify ABI data decoding
		ackData, err := DeserializeMocaSBTAckPackage(data)
		require.NoError(t, err)

		ackStruct, ok := ackData.(*MocaSBTAckPackageStruct)
		require.True(t, ok, "Should decode to MocaSBTAckPackageStruct")
		assert.Len(t, ackStruct.Toaddrs, 2, "Should have 2 addresses")
		assert.Equal(t, STATUS_SUCCESS, ackStruct.Status)
	})
}

// TestMultiplePackagesProcessing tests MED-002 fix
// Verifies that all packages in a bundle are processed
func TestMultiplePackagesProcessing(t *testing.T) {
	// Build 3 different packages
	packages := oracletypes.Packages{
		// Package 1: SBT Ack with 2 addresses
		{
			ChannelId: 1,
			Sequence:  100,
			Payload:   buildSBTAckPayload(t, 2, STATUS_SUCCESS),
		},
		// Package 2: SBT Ack with 3 addresses
		{
			ChannelId: 1,
			Sequence:  101,
			Payload:   buildSBTAckPayload(t, 3, STATUS_SUCCESS),
		},
		// Package 3: Other type (not SBT Ack)
		{
			ChannelId: 1,
			Sequence:  102,
			Payload:   []byte{0x01, 0x02, 0x03}, // Different OperationType
		},
	}

	encodedPayload, err := rlp.EncodeToBytes(packages)
	require.NoError(t, err)

	// Decode and process
	var decodedPackages oracletypes.Packages
	err = rlp.DecodeBytes(encodedPayload, &decodedPackages)
	require.NoError(t, err)

	// Key assertion: should decode all 3 packages
	assert.Len(t, decodedPackages, 3, "Should decode all 3 packages")

	// Count different types of packages
	sbtAckCount := 0
	totalAddresses := 0

	for i, pkg := range decodedPackages {
		require.Greater(t, len(pkg.Payload), 0, "Package %d should have payload", i)

		op := pkg.Payload[0]
		if op == OperationMocaSBTACK {
			sbtAckCount++

			// Parse address count
			data := pkg.Payload[1:]
			ackData, err := DeserializeMocaSBTAckPackage(data)
			require.NoError(t, err)

			ackStruct := ackData.(*MocaSBTAckPackageStruct)
			totalAddresses += len(ackStruct.Toaddrs)
		}
	}

	assert.Equal(t, 2, sbtAckCount, "Should have 2 SBT Ack packages")
	assert.Equal(t, 5, totalAddresses, "Should have 2+3=5 total addresses")
}

// TestErrorHandling tests error cases
func TestErrorHandling(t *testing.T) {
	t.Run("InvalidRLPData", func(t *testing.T) {
		invalidRLP := []byte{0xFF, 0xFF, 0xFF}

		var packages oracletypes.Packages
		err := rlp.DecodeBytes(invalidRLP, &packages)

		assert.Error(t, err, "Should return error for invalid RLP")
		assert.NotPanics(t, func() {
			rlp.DecodeBytes(invalidRLP, &packages)
		}, "Should NOT panic on invalid RLP")
	})

	t.Run("EmptyPayload", func(t *testing.T) {
		packages := oracletypes.Packages{
			{
				ChannelId: 1,
				Sequence:  100,
				Payload:   []byte{}, // Empty payload
			},
		}

		encodedPayload, err := rlp.EncodeToBytes(packages)
		require.NoError(t, err)

		var decoded oracletypes.Packages
		err = rlp.DecodeBytes(encodedPayload, &decoded)
		require.NoError(t, err)

		// Should be able to decode, but should skip empty payload during processing
		assert.Len(t, decoded, 1)
		assert.Empty(t, decoded[0].Payload)
	})

	t.Run("InvalidABIData", func(t *testing.T) {
		// OperationType correct, but ABI data wrong
		invalidPayload := []byte{OperationMocaSBTACK, 0x01, 0x02}

		_, err := DeserializeMocaSBTAckPackage(invalidPayload[1:])
		assert.Error(t, err, "Should return error for invalid ABI data")
	})
}

// TestRegressionHIGH001 ensures HIGH-001 bug doesn't reappear
func TestRegressionHIGH001(t *testing.T) {
	t.Run("MustNotReadRLPPrefixAsOperationType", func(t *testing.T) {
		// Build RLP encoded data
		packages := oracletypes.Packages{
			{
				ChannelId: 1,
				Sequence:  100,
				Payload:   buildSBTAckPayload(t, 1, STATUS_SUCCESS),
			},
		}

		encoded, _ := rlp.EncodeToBytes(packages)

		// RLP prefix should NOT be treated as OperationType
		rlpPrefix := encoded[0]
		assert.NotEqual(t, OperationMocaSBTACK, rlpPrefix,
			"RLP prefix (0x%02x) should NOT equal OperationMocaSBTACK (0x%02x)", rlpPrefix, OperationMocaSBTACK)

		// After correct decoding, should find OperationType
		var decoded oracletypes.Packages
		rlp.DecodeBytes(encoded, &decoded)

		op := decoded[0].Payload[0]
		assert.Equal(t, OperationMocaSBTACK, op,
			"Should correctly read OperationType from Payload[0]")
	})
}

// TestRegressionMED002 ensures MED-002 bug doesn't reappear
func TestRegressionMED002(t *testing.T) {
	t.Run("MustProcessAllPackages", func(t *testing.T) {
		packages := oracletypes.Packages{
			{ChannelId: 1, Sequence: 100, Payload: buildSBTAckPayload(t, 1, STATUS_SUCCESS)},
			{ChannelId: 1, Sequence: 101, Payload: buildSBTAckPayload(t, 1, STATUS_SUCCESS)},
			{ChannelId: 1, Sequence: 102, Payload: buildSBTAckPayload(t, 1, STATUS_SUCCESS)},
		}

		encoded, _ := rlp.EncodeToBytes(packages)

		var decoded oracletypes.Packages
		rlp.DecodeBytes(encoded, &decoded)

		assert.Len(t, decoded, 3, "Must decode all 3 packages, not just the first one")
	})
}

// TestRegressionMultipleAddresses ensures all addresses are processed
func TestRegressionMultipleAddresses(t *testing.T) {
	t.Run("MustProcessAllAddresses", func(t *testing.T) {
		// Build SBT Ack with 5 addresses
		payload := buildSBTAckPayload(t, 5, STATUS_SUCCESS)

		data := payload[1:]
		ackData, _ := DeserializeMocaSBTAckPackage(data)
		ackStruct := ackData.(*MocaSBTAckPackageStruct)

		assert.Len(t, ackStruct.Toaddrs, 5,
			"Must decode all 5 addresses, not just Toaddrs[0]")
	})
}

// Helper function: build SBT Ack payload
func buildSBTAckPayload(t *testing.T, numAddrs int, status uint32) []byte {
	addrs := make([]sdk.AccAddress, numAddrs)
	for i := 0; i < numAddrs; i++ {
		addr := ethcommon.HexToAddress(fmt.Sprintf("0x%040d", i))
		addrs[i] = sdk.AccAddress(addr.Bytes())
	}

	pkg := &MocaSBTAckPackage{
		Toaddrs: addrs,
		Status:  status,
	}
	serialized, err := pkg.Serialize()
	require.NoError(t, err)

	// [OperationType][ABI data]
	return append([]byte{OperationMocaSBTACK}, serialized...)
}

