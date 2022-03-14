package tests

import (
	"crypto/ed25519"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"testing"

	"github.com/cheqd/cheqd-node/x/cheqd/types"
	"github.com/multiformats/go-multibase"

	"github.com/stretchr/testify/require"
)

func TestCreateDID(t *testing.T) {
	var err error
	keys := GenerateTestKeys()
	cases := []struct {
		valid   bool
		name    string
		keys    map[string]KeyPair
		signers []string
		msg     *types.MsgCreateDidPayload
		errMsg  string
	}{
		{
			valid: true,
			name:  "Valid: Works",
			keys: map[string]KeyPair{
				ImposterKey1: GenerateKeyPair(),
			},
			signers: []string{ImposterKey1},
			msg: &types.MsgCreateDidPayload{
				Id:             ImposterDID,
				Authentication: []string{ImposterKey1},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         ImposterKey1,
						Type:       Ed25519VerificationKey2020,
						Controller: ImposterDID,
					},
				},
			},
		},
		{
			valid: true,
			name:  "Valid: Works with Key Agreement",
			keys: map[string]KeyPair{
				ImposterKey1: GenerateKeyPair(),
				AliceKey1:    keys[AliceKey1],
			},
			signers: []string{ImposterKey1, AliceKey1},
			msg: &types.MsgCreateDidPayload{
				Id:           ImposterDID,
				KeyAgreement: []string{ImposterKey1},
				Controller:   []string{AliceDID},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         ImposterKey1,
						Type:       Ed25519VerificationKey2020,
						Controller: ImposterDID,
					},
				},
			},
		},
		{
			valid: true,
			name:  "Valid: Works with Assertion Method",
			keys: map[string]KeyPair{
				ImposterKey1: GenerateKeyPair(),
				AliceKey1:    keys[AliceKey1],
			},
			signers: []string{AliceKey1, ImposterKey1},
			msg: &types.MsgCreateDidPayload{
				Id:              ImposterDID,
				AssertionMethod: []string{ImposterKey1},
				Controller:      []string{AliceDID},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         ImposterKey1,
						Type:       Ed25519VerificationKey2020,
						Controller: ImposterDID,
					},
				},
			},
		},
		{
			valid: true,
			name:  "Valid: Works with Capability Delegation",
			keys: map[string]KeyPair{
				ImposterKey1: GenerateKeyPair(),
				AliceKey1:    keys[AliceKey1],
			},
			signers: []string{AliceKey1, ImposterKey1},
			msg: &types.MsgCreateDidPayload{
				Id:                   ImposterDID,
				CapabilityDelegation: []string{ImposterKey1},
				Controller:           []string{AliceDID},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         ImposterKey1,
						Type:       Ed25519VerificationKey2020,
						Controller: ImposterDID,
					},
				},
			},
		},
		{
			valid: true,
			name:  "Valid: Works with Capability Invocation",
			keys: map[string]KeyPair{
				ImposterKey1: GenerateKeyPair(),
				AliceKey1:    keys[AliceKey1],
			},
			signers: []string{AliceKey1, ImposterKey1},
			msg: &types.MsgCreateDidPayload{
				Id:                   ImposterDID,
				CapabilityInvocation: []string{ImposterKey1},
				Controller:           []string{AliceDID},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         ImposterKey1,
						Type:       Ed25519VerificationKey2020,
						Controller: ImposterDID,
					},
				},
			},
		},
		{
			valid: true,
			name:  "Valid: With controller works",
			msg: &types.MsgCreateDidPayload{
				Id:         ImposterDID,
				Controller: []string{AliceDID, BobDID},
			},
			signers: []string{AliceKey1, BobKey3},
			keys: map[string]KeyPair{
				AliceKey1: keys[AliceKey1],
				BobKey3:   keys[BobKey3],
			},
		},
		{
			valid: true,
			name:  "Valid: Full message works",
			keys: map[string]KeyPair{
				"did:cheqd:test:1111111111111111#key-1": GenerateKeyPair(),
				"did:cheqd:test:1111111111111111#key-2": GenerateKeyPair(),
				"did:cheqd:test:1111111111111111#key-3": GenerateKeyPair(),
				"did:cheqd:test:1111111111111111#key-4": GenerateKeyPair(),
				"did:cheqd:test:1111111111111111#key-5": GenerateKeyPair(),
				AliceKey1:                             keys[AliceKey1],
				BobKey1:                               keys[BobKey1],
				BobKey2:                               keys[BobKey2],
				BobKey3:                               keys[BobKey3],
				CharlieKey1:                           keys[CharlieKey1],
				CharlieKey2:                           keys[CharlieKey2],
				CharlieKey3:                           keys[CharlieKey3],
			},
			signers: []string{
				"did:cheqd:test:1111111111111111#key-1",
				"did:cheqd:test:1111111111111111#key-5",
				AliceKey1,
				BobKey1,
				BobKey2,
				BobKey3,
				CharlieKey1,
				CharlieKey2,
				CharlieKey3,
			},
			msg: &types.MsgCreateDidPayload{
				Id: "did:cheqd:test:1111111111111111",
				Authentication: []string{
					"did:cheqd:test:1111111111111111#key-1",
					"did:cheqd:test:1111111111111111#key-5",
				},
				Context:              []string{"abc", "de"},
				CapabilityInvocation: []string{"did:cheqd:test:1111111111111111#key-2"},
				CapabilityDelegation: []string{"did:cheqd:test:1111111111111111#key-3"},
				KeyAgreement:         []string{"did:cheqd:test:1111111111111111#key-4"},
				AlsoKnownAs:          []string{"SomeUri"},
				Service: []*types.Service{
					{
						Id:              "did:cheqd:test:1111111111111111#service-1",
						Type:            "DIDCommMessaging",
						ServiceEndpoint: "ServiceEndpoint",
					},
				},
				Controller: []string{"did:cheqd:test:1111111111111111", AliceDID, BobDID, CharlieDID},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         "did:cheqd:test:1111111111111111#key-1",
						Type:       Ed25519VerificationKey2020,
						Controller: "did:cheqd:test:1111111111111111",
					},
					{
						Id:         "did:cheqd:test:1111111111111111#key-2",
						Type:       Ed25519VerificationKey2020,
						Controller: "did:cheqd:test:1111111111111111",
					},
					{
						Id:         "did:cheqd:test:1111111111111111#key-3",
						Type:       Ed25519VerificationKey2020,
						Controller: "did:cheqd:test:1111111111111111",
					},
					{
						Id:         "did:cheqd:test:1111111111111111#key-4",
						Type:       "Ed25519VerificationKey2020",
						Controller: "did:cheqd:test:1111111111111111",
					},
					{
						Id:         "did:cheqd:test:1111111111111111#key-5",
						Type:       "Ed25519VerificationKey2020",
						Controller: "did:cheqd:test:1111111111111111",
					},
				},
			},
		},
		{
			valid: false,
			name:  "Not Valid: Second controller did not sign request",
			msg: &types.MsgCreateDidPayload{
				Id:         ImposterDID,
				Controller: []string{AliceDID, BobDID},
			},
			signers: []string{AliceKey1},
			keys: map[string]KeyPair{
				AliceKey1: keys[AliceKey1],
			},
			errMsg: fmt.Sprintf("signer: %s: signature is required but not found", BobDID),
		},
		{
			valid: false,
			name:  "Not Valid: No signature",
			msg: &types.MsgCreateDidPayload{
				Id:         ImposterDID,
				Controller: []string{AliceDID, BobDID},
			},
			errMsg: fmt.Sprintf("signer: %s: signature is required but not found", AliceDID),
		},
		{
			valid: false,
			name:  "Not Valid: Controller not found",
			msg: &types.MsgCreateDidPayload{
				Id:         ImposterDID,
				Controller: []string{AliceDID, NotFounDID},
			},
			signers: []string{AliceKey1, ImposterKey1},
			keys: map[string]KeyPair{
				AliceKey1:    keys[AliceKey1],
				ImposterKey1: GenerateKeyPair(),
			},
			errMsg: fmt.Sprintf("%s: DID Doc not found", NotFounDID),
		},
		{
			valid: false,
			name:  "Not Valid: Wrong signature",
			msg: &types.MsgCreateDidPayload{
				Id:         ImposterDID,
				Controller: []string{AliceDID},
			},
			signers: []string{AliceKey1},
			keys: map[string]KeyPair{
				AliceKey1: keys[BobKey1],
			},
			errMsg: fmt.Sprintf("method id: %s : invalid signature detected", AliceKey1),
		},
		{
			valid: false,
			name:  "Not Valid: DID signed by wrong controller",
			msg: &types.MsgCreateDidPayload{
				Id:             ImposterDID,
				Authentication: []string{ImposterKey1},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:                 ImposterKey1,
						Type:               Ed25519VerificationKey2020,
						Controller:         ImposterDID,
						PublicKeyMultibase: "z" + base58.Encode(keys[ImposterKey1].PublicKey),
					},
				},
			},
			signers: []string{AliceKey1},
			keys: map[string]KeyPair{
				AliceKey1: keys[AliceKey1],
			},
			errMsg: fmt.Sprintf("signer: %s: signature is required but not found", ImposterDID),
		},
		{
			valid: false,
			name:  "Not Valid: DID self-signed by not existing verification method",
			msg: &types.MsgCreateDidPayload{
				Id:             ImposterDID,
				Authentication: []string{ImposterKey1},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:                 ImposterKey1,
						Type:               Ed25519VerificationKey2020,
						Controller:         ImposterDID,
						PublicKeyMultibase: "z" + base58.Encode(keys[ImposterKey1].PublicKey),
					},
				},
			},
			signers: []string{ImposterKey2},
			keys: map[string]KeyPair{
				ImposterKey2: GenerateKeyPair(),
			},
			errMsg: fmt.Sprintf("%s: verification method not found", ImposterKey2),
		},
		{
			valid: false,
			name:  "Not Valid: Self-signature not found",
			msg: &types.MsgCreateDidPayload{
				Id:             ImposterDID,
				Controller:     []string{AliceDID, ImposterDID},
				Authentication: []string{ImposterKey1},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:                 ImposterKey1,
						Type:               Ed25519VerificationKey2020,
						Controller:         ImposterDID,
						PublicKeyMultibase: "z" + base58.Encode(keys[ImposterKey1].PublicKey),
					},
				},
			},
			signers: []string{AliceKey1, ImposterKey2},
			keys: map[string]KeyPair{
				AliceKey1:    keys[AliceKey1],
				ImposterKey2: GenerateKeyPair(),
			},
			errMsg: fmt.Sprintf("%s: verification method not found", ImposterKey2),
		},
		{
			valid: false,
			name:  "Not Valid: DID Doc already exists",
			keys: map[string]KeyPair{
				CharlieKey1: GenerateKeyPair(),
			},
			signers: []string{CharlieKey1},
			msg: &types.MsgCreateDidPayload{
				Id:             CharlieDID,
				Authentication: []string{CharlieKey1},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         CharlieKey1,
						Type:       Ed25519VerificationKey2020,
						Controller: CharlieDID,
					},
				},
			},
			errMsg: fmt.Sprintf("%s: DID Doc exists", CharlieDID),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			msg := tc.msg
			setup := InitEnv(t, keys)

			for _, vm := range msg.VerificationMethod {
				if vm.PublicKeyMultibase == "" {
					vm.PublicKeyMultibase, err = multibase.Encode(multibase.Base58BTC, tc.keys[vm.Id].PublicKey)
				}
				require.NoError(t, err)
			}

			signerKeys := map[string]ed25519.PrivateKey{}
			for _, signer := range tc.signers {
				signerKeys[signer] = tc.keys[signer].PrivateKey
			}

			did, err := setup.SendCreateDid(msg, signerKeys)

			if tc.valid {
				require.Nil(t, err)
				require.Equal(t, tc.msg.Id, did.Id)
				require.Equal(t, tc.msg.Controller, did.Controller)
				require.Equal(t, tc.msg.VerificationMethod, did.VerificationMethod)
				require.Equal(t, tc.msg.Authentication, did.Authentication)
				require.Equal(t, tc.msg.AssertionMethod, did.AssertionMethod)
				require.Equal(t, tc.msg.CapabilityInvocation, did.CapabilityInvocation)
				require.Equal(t, tc.msg.CapabilityDelegation, did.CapabilityDelegation)
				require.Equal(t, tc.msg.KeyAgreement, did.KeyAgreement)
				require.Equal(t, tc.msg.AlsoKnownAs, did.AlsoKnownAs)
				require.Equal(t, tc.msg.Service, did.Service)
				require.Equal(t, tc.msg.Context, did.Context)
			} else {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			}
		})
	}
}

//func TestUpdateDid(t *testing.T) {
//	keys := map[string]KeyPair{
//		AliceKey1:   GenerateKeyPair(),
//		AliceKey2:   GenerateKeyPair(),
//		BobKey1:     GenerateKeyPair(),
//		BobKey2:     GenerateKeyPair(),
//		BobKey3:     GenerateKeyPair(),
//		BobKey4:     GenerateKeyPair(),
//		CharlieKey1: GenerateKeyPair(),
//		CharlieKey2: GenerateKeyPair(),
//		CharlieKey3: GenerateKeyPair(),
//		ImposterKey1: GenerateKeyPair(),
//	}
//	setup := Setup()
//	keys, err := setup.CreateTestDIDs(nil)
//	require.NoError(t, err)
//
//	cases := []struct {
//		valid   bool
//		name    string
//		signers []string
//		msg     *types.MsgUpdateDidPayload
//		errMsg  string
//	}{
//		{
//			valid:   true,
//			name:    "Key rotation works",
//			signers: []string{AliceKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   false,
//			name:    "Try to add controller without self-signature",
//			signers: []string{BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID},
//				Authentication: []string{AliceKey1},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey1,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:alice not found: invalid signature detected",
//		},
//		{
//			valid:   false,
//			name:    "Add controller and replace authentication without old signature do not work",
//			signers: []string{BobKey1, AliceKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID},
//				Authentication: []string{AliceKey1},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey1,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "did:cheqd:test:alice#key-1: verification method not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			name:    "Add controller work",
//			signers: []string{BobKey1, AliceKey2},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   true,
//			name:    "Add controller without signature work (signatures of old controllers are present)",
//			signers: []string{BobKey1, AliceKey2},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:charlie not found: invalid signature detected",
//		},
//		{
//			valid:   false,
//			name:    "Replace controller work without new signature do not work",
//			signers: []string{BobKey1, AliceKey2},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:charlie not found: invalid signature detected",
//		},
//		{
//			valid:   false,
//			name:    "Replace controller without old signature do not work",
//			signers: []string{AliceKey2, CharlieKey3},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:bob not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			name:    "Replace controller work",
//			signers: []string{AliceKey2, CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   true,
//			name:    "Add second controller works",
//			signers: []string{AliceKey2, CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   true,
//			name:    "Add verification method without signature controller work",
//			signers: []string{CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				KeyAgreement:   []string{AliceKey1},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//					{
//						Id:         AliceKey1,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   false,
//			name:    "Remove verification method without signature controller do not work",
//			signers: []string{CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:alice not found: invalid signature detected",
//		},
//		{
//			valid:   false,
//			name:    "Remove verification method wrong authentication detected",
//			signers: []string{AliceKey1, CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "did:cheqd:test:alice#key-1: verification method not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			name:    "Add second authentication works",
//			signers: []string{AliceKey2, CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey1, AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey1,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: BobDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   false,
//			name:    "Remove self authentication without signature do not work",
//			signers: []string{CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: BobDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:alice not found: invalid signature detected",
//		},
//		{
//			valid:   false,
//			name:    "Change self controller verification without signature do not work",
//			signers: []string{CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey1, AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey1,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: CharlieDID,
//					},
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: BobDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:alice not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			signers: []string{AliceKey2, CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Controller:     []string{BobDID, CharlieDID},
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: BobDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   false,
//			name:    "Change controller to self without old controllers signatures does not work",
//			signers: []string{AliceKey2},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: BobDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:bob not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			name:    "Change controller to self works",
//			signers: []string{AliceKey2, CharlieKey3, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: BobDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   false,
//			name:    "Change verification method controller without old signature",
//			signers: []string{AliceKey2, CharlieKey3},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: CharlieDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:bob not found: invalid signature detected",
//		},
//		{
//			valid:   false,
//			name:    "Change verification method controller without new signature",
//			signers: []string{AliceKey2, BobKey1},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: CharlieDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:charlie not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			name:    "Change verification method controller",
//			signers: []string{AliceKey2, BobKey1, CharlieKey3},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: CharlieDID,
//					},
//				},
//			},
//		},
//		{
//			valid:   false,
//			name:    "Change to self verification method without controller signature",
//			signers: []string{AliceKey2},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//			errMsg: "signature did:cheqd:test:charlie not found: invalid signature detected",
//		},
//		{
//			valid:   true,
//			name:    "Change to self verification method without controller signature",
//			signers: []string{AliceKey2, CharlieKey3},
//			msg: &types.MsgUpdateDidPayload{
//				Id:             AliceDID,
//				Authentication: []string{AliceKey2},
//				VerificationMethod: []*types.VerificationMethod{
//					{
//						Id:         AliceKey2,
//						Type:       "Ed25519VerificationKey2020",
//						Controller: AliceDID,
//					},
//				},
//			},
//		},
//	}
//
//	for _, tc := range cases {
//		t.Run(tc.name, func(t *testing.T) {
//			msg := tc.msg
//
//			for _, vm := range msg.VerificationMethod {
//				encoded, err := multibase.Encode(multibase.Base58BTC, keys[vm.Id].PublicKey)
//				require.NoError(t, err)
//				vm.PublicKeyMultibase = encoded
//			}
//
//			signerKeys := map[string]ed25519.PrivateKey{}
//			for _, signer := range tc.signers {
//				signerKeys[signer] = keys[signer].PrivateKey
//			}
//
//			did, err := setup.SendUpdateDid(msg, signerKeys)
//
//			if tc.valid {
//				require.Nil(t, err)
//				require.Equal(t, tc.msg.Id, did.Id)
//				require.Equal(t, tc.msg.Controller, did.Controller)
//				require.Equal(t, tc.msg.VerificationMethod, did.VerificationMethod)
//				require.Equal(t, tc.msg.Authentication, did.Authentication)
//				require.Equal(t, tc.msg.AssertionMethod, did.AssertionMethod)
//				require.Equal(t, tc.msg.CapabilityInvocation, did.CapabilityInvocation)
//				require.Equal(t, tc.msg.CapabilityDelegation, did.CapabilityDelegation)
//				require.Equal(t, tc.msg.KeyAgreement, did.KeyAgreement)
//				require.Equal(t, tc.msg.AlsoKnownAs, did.AlsoKnownAs)
//				require.Equal(t, tc.msg.Service, did.Service)
//				require.Equal(t, tc.msg.Context, did.Context)
//			} else {
//				require.Error(t, err)
//				require.Equal(t, tc.errMsg, err.Error())
//			}
//			// Reset by recreating databases and reinit for TestSetup
//			setup, _ = InitEnv(t, keys)
//		})
//	}
//}

func TestHandler_DidDocAlreadyExists(t *testing.T) {
	setup := Setup()

	_, _, _ = setup.InitDid("did:cheqd:test:alice")
	_, _, err := setup.InitDid("did:cheqd:test:alice")

	require.Error(t, err)
	require.Equal(t, "DID is already used by DIDDoc did:cheqd:test:alice: DID Doc exists", err.Error())
}
