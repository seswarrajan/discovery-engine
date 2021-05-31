package systempolicy

import (
	"testing"

	"github.com/accuknox/knoxAutoPolicy/src/libs"
	"github.com/accuknox/knoxAutoPolicy/src/types"
	"github.com/stretchr/testify/assert"
)

// ==================== //
// == Exact Matching == //
// ==================== //

func TestIsExistingPolicy(t *testing.T) {
	exist := types.KubeArmorSystemPolicy{
		Metadata: map[string]string{
			"name": "test",
		},

		Spec: types.KubeArmorSpec{
			Selector: types.Selector{
				MatchLabels: map[string]string{
					"app": "test1",
				},
			},
		},
	}

	newOne := types.KubeArmorSystemPolicy{}
	libs.DeepCopy(&newOne, &exist)

	assert.True(t, IsExistingPolicy([]types.KubeArmorSystemPolicy{exist}, newOne))
}

// ======================= //
// == Policy Name Check == //
// ======================= //

func TestGeneratePolicyName(t *testing.T) {
	exist := types.KubeArmorSystemPolicy{
		Metadata: map[string]string{},

		Spec: types.KubeArmorSpec{
			Selector: types.Selector{
				MatchLabels: map[string]string{
					"app": "test1",
				},
			},
		},
	}

	updated := GeneratePolicyName(map[string]bool{}, exist, "testcluster")

	assert.Equal(t, updated.Metadata["clusterName"], "testcluster")
}
