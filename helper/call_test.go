package helper

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/maticnetwork/bor/common"
	authTypes "github.com/maticnetwork/heimdall/auth/types"
	"github.com/maticnetwork/heimdall/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testTendermintNode = "tcp://localhost:26657"
)

//  TestCheckpointSigs decodes signers from checkpoint sigs data
func TestCheckpointSigs(t *testing.T) {
	t.Parallel()

	viper.Set(TendermintNodeFlag, testTendermintNode)
	viper.Set("log_level", "info")
	InitHeimdallConfig(os.ExpandEnv("$HOME/.heimdalld"))

	contractCallerObj, err := NewContractCaller()
	if err != nil {
		fmt.Println("Error creating contract caller")
	}

	txHashStr := "0x9c2a9e20e1fecdae538f72b01dd0fd5008cc90176fd603b92b59274d754cbbd8"
	txHash := common.HexToHash(txHashStr)

	voteSignBytes, sigs, txData, err := contractCallerObj.GetCheckpointSign(txHash)
	if err != nil {
		fmt.Println("Error fetching checkpoint tx input args")
	}

	fmt.Println("checkpoint args", "vote", hex.EncodeToString(voteSignBytes), "sigs", hex.EncodeToString(sigs), "txData", hex.EncodeToString(txData))

	signerList, err := FetchSigners(voteSignBytes, sigs)
	if err != nil {
		fmt.Println("Error fetching signer list from tx input args")
	}

	fmt.Println("signers list", signerList)
}

// FetchSigners fetches the signers' list
func FetchSigners(voteBytes []byte, sigInput []byte) ([]string, error) {
	const sigLength = 65

	signersList := make([]string, len(sigInput))

	// Calculate total stake Power of all Signers.
	for i := 0; i < len(sigInput); i += sigLength {
		signature := sigInput[i : i+sigLength]

		pKey, err := authTypes.RecoverPubkey(voteBytes, signature)
		if err != nil {
			fmt.Println("Error Recovering PubKey", "Error", err)
			return nil, err
		}

		signersList[i] = types.NewPubKey(pKey).Address().String()
	}

	return signersList, nil
}

//  TestPopulateABIs tests that package level ABIs cache works as expected
//  by not invoking json methods after contracts ABIs' init
func TestPopulateABIs(t *testing.T) {
	t.Parallel()

	viper.Set(TendermintNodeFlag, testTendermintNode)
	viper.Set("log_level", "info")
	InitHeimdallConfig(os.ExpandEnv("$HOME/.heimdalld"))

	fmt.Println("ABIs map should be empty and all ABIs not found")
	assert.True(t, len(ContractsABIsMap) == 0)
	_, found := ContractsABIsMap[RootChainABI]
	assert.False(t, found)
	_, found = ContractsABIsMap[StakingInfoABI]
	assert.False(t, found)
	_, found = ContractsABIsMap[StateReceiverABI]
	assert.False(t, found)
	_, found = ContractsABIsMap[StateSenderABI]
	assert.False(t, found)
	_, found = ContractsABIsMap[StakeManagerABI]
	assert.False(t, found)
	_, found = ContractsABIsMap[SlashManagerABI]
	assert.False(t, found)
	_, found = ContractsABIsMap[MaticTokenABI]
	assert.False(t, found)

	fmt.Println("Should create a new contract caller and populate its ABIs by decoding json")

	contractCallerObjFirst, err := NewContractCaller()
	if err != nil {
		fmt.Println("Error creating contract caller")
	}

	assert.Equalf(t, ContractsABIsMap[RootChainABI], &contractCallerObjFirst.RootChainABI,
		"values for %s not equals", RootChainABI)
	assert.Equalf(t, ContractsABIsMap[StakingInfoABI], &contractCallerObjFirst.StakingInfoABI,
		"values for %s not equals", StakingInfoABI)
	assert.Equalf(t, ContractsABIsMap[StateReceiverABI], &contractCallerObjFirst.StateReceiverABI,
		"values for %s not equals", StateReceiverABI)
	assert.Equalf(t, ContractsABIsMap[StateSenderABI], &contractCallerObjFirst.StateSenderABI,
		"values for %s not equals", StateSenderABI)
	assert.Equalf(t, ContractsABIsMap[StakeManagerABI], &contractCallerObjFirst.StakeManagerABI,
		"values for %s not equals", StakeManagerABI)
	assert.Equalf(t, ContractsABIsMap[SlashManagerABI], &contractCallerObjFirst.SlashManagerABI,
		"values for %s not equals", SlashManagerABI)
	assert.Equalf(t, ContractsABIsMap[MaticTokenABI], &contractCallerObjFirst.MaticTokenABI,
		"values for %s not equals", MaticTokenABI)

	fmt.Println("ABIs map should not be empty and all ABIs found")
	assert.True(t, len(ContractsABIsMap) == 8)
	_, found = ContractsABIsMap[RootChainABI]
	assert.True(t, found)
	_, found = ContractsABIsMap[StakingInfoABI]
	assert.True(t, found)
	_, found = ContractsABIsMap[StateReceiverABI]
	assert.True(t, found)
	_, found = ContractsABIsMap[StateSenderABI]
	assert.True(t, found)
	_, found = ContractsABIsMap[StakeManagerABI]
	assert.True(t, found)
	_, found = ContractsABIsMap[SlashManagerABI]
	assert.True(t, found)
	_, found = ContractsABIsMap[MaticTokenABI]
	assert.True(t, found)

	fmt.Println("Should create a new contract caller and populate its ABIs by using cached map")

	contractCallerObjSecond, err := NewContractCaller()
	if err != nil {
		fmt.Println("Error creating contract caller")
	}

	assert.Equalf(t, ContractsABIsMap[RootChainABI], &contractCallerObjSecond.RootChainABI,
		"values for %s not equals", RootChainABI)
	assert.Equalf(t, ContractsABIsMap[StakingInfoABI], &contractCallerObjSecond.StakingInfoABI,
		"values for %s not equals", StakingInfoABI)
	assert.Equalf(t, ContractsABIsMap[StateReceiverABI], &contractCallerObjSecond.StateReceiverABI,
		"values for %s not equals", StateReceiverABI)
	assert.Equalf(t, ContractsABIsMap[StateSenderABI], &contractCallerObjSecond.StateSenderABI,
		"values for %s not equals", StateSenderABI)
	assert.Equalf(t, ContractsABIsMap[StakeManagerABI], &contractCallerObjSecond.StakeManagerABI,
		"values for %s not equals", StakeManagerABI)
	assert.Equalf(t, ContractsABIsMap[SlashManagerABI], &contractCallerObjSecond.SlashManagerABI,
		"values for %s not equals", SlashManagerABI)
	assert.Equalf(t, ContractsABIsMap[MaticTokenABI], &contractCallerObjSecond.MaticTokenABI,
		"values for %s not equals", MaticTokenABI)

}
