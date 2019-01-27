package workers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type workerMock struct {
}

func (workerMock) Process(context.Context, []byte) error {
	panic("implement me")
}

func TestRegisterWorker(t *testing.T) {

	w := getWorkers("random_queue")
	require.Nil(t, w)
	require.NoError(t, RegisterWorker("random_queue", &workerMock{}))
	require.Error(t, RegisterWorker("random_queue", &workerMock{}))
	w = getWorkers("random_queue")
	require.NotNil(t, w)
}
