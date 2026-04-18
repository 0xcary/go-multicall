package multicall

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCall_BadABI(t *testing.T) {
	r := require.New(t)

	const oneValueABI = `[
		{
			"constant":true,
			"inputs": [
				{
					"name":"val1",
					"type":"bool"
				}
			],
			"name":"testFunc",
			"outputs": [
				{
					"name":"val1",
					"type":"bool"
				}
			],
			"payable":false,
			"stateMutability":"view",
			"type":"function"
		}
	` // missing closing ] at the end

	_, err := NewContract(oneValueABI, "0x")
	r.Error(err)
	r.ErrorContains(err, "unexpected EOF")
}

// TestUnpackResult_SafeTypeAssertion verifies UnpackResult never panics regardless of
// the concrete type stored in Outputs.
func TestUnpackResult_SafeTypeAssertion(t *testing.T) {
	r := require.New(t)

	// nil Outputs → nil
	r.Nil((&Call{}).UnpackResult())

	// struct pointer (common case) → nil, no panic
	r.Nil((&Call{Outputs: new(struct{ Val bool })}).UnpackResult())

	// []interface{} → returns the slice as-is
	out := []interface{}{true, "hello"}
	r.Equal(out, (&Call{Outputs: out}).UnpackResult())
}
