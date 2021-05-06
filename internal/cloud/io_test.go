package cloud

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tilt-dev/tilt/internal/store"
	"github.com/tilt-dev/tilt/internal/testutils"
)

func TestWriteSnapshotTo(t *testing.T) {
	ctx, _, _ := testutils.CtxAndAnalyticsForTest()
	buf := bytes.NewBuffer(nil)
	state := store.NewState()
	err := WriteSnapshotTo(ctx, *state, buf)
	assert.NoError(t, err)
	assert.Equal(t, `{
  "view": {
    "runningTiltBuild": {

    },
    "versionSettings": {
      "checkUpdates": true
    },
    "tiltCloudSchemeHost": "https:",
    "logList": {
      "fromCheckpoint": -1,
      "toCheckpoint": -1
    },
    "tiltStartTime": "0001-01-01T00:00:00Z",
    "uiResources": [
      {
        "metadata": {
          "name": "(Tiltfile)"
        },
        "status": {
          "buildHistory": [
            {

            }
          ],
          "runtimeStatus": "not_applicable",
          "updateStatus": "pending"
        }
      }
    ]
  }
}
`, buf.String())
}
