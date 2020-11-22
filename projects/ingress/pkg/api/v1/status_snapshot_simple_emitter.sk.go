// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/stats"
	"go.uber.org/zap"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type StatusSimpleEmitter interface {
	Snapshots(ctx context.Context) (<-chan *StatusSnapshot, <-chan error, error)
}

func NewStatusSimpleEmitter(aggregatedWatch clients.ResourceWatch) StatusSimpleEmitter {
	return NewStatusSimpleEmitterWithEmit(aggregatedWatch, make(chan struct{}))
}

func NewStatusSimpleEmitterWithEmit(aggregatedWatch clients.ResourceWatch, emit <-chan struct{}) StatusSimpleEmitter {
	return &statusSimpleEmitter{
		aggregatedWatch: aggregatedWatch,
		forceEmit:       emit,
	}
}

type statusSimpleEmitter struct {
	forceEmit       <-chan struct{}
	aggregatedWatch clients.ResourceWatch
}

func (c *statusSimpleEmitter) Snapshots(ctx context.Context) (<-chan *StatusSnapshot, <-chan error, error) {
	snapshots := make(chan *StatusSnapshot)
	errs := make(chan error)

	untyped, watchErrs, err := c.aggregatedWatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	go errutils.AggregateErrs(ctx, errs, watchErrs, "status-emitter")

	go func() {
		currentSnapshot := StatusSnapshot{}
		timer := time.NewTicker(time.Second * 1)
		var previousHash uint64
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			previousHash = currentHash

			stats.Record(ctx, mStatusSnapshotOut.M(1))
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		defer func() {
			close(snapshots)
			close(errs)
		}()

		for {
			record := func() { stats.Record(ctx, mStatusSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case untypedList := <-untyped:
				record()

				currentSnapshot = StatusSnapshot{}
				for _, res := range untypedList {
					switch typed := res.(type) {
					case *KubeService:
						currentSnapshot.Services = append(currentSnapshot.Services, typed)
					case *Ingress:
						currentSnapshot.Ingresses = append(currentSnapshot.Ingresses, typed)
					default:
						select {
						case errs <- fmt.Errorf("StatusSnapshotEmitter "+
							"cannot process resource %v of type %T", res.GetMetadata().Ref(), res):
						case <-ctx.Done():
							return
						}
					}
				}

			}
		}
	}()
	return snapshots, errs, nil
}
