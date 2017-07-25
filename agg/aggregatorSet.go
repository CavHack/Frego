package agg

import (

		"fmt"
		"sync"

		"github.com/cavhack/Frego/protos"
		"github.com/pkg/errors"


)



type AggregatorSet struct {

	mutex sync.RWMutex
	aggs map[string]Aggregator

}

func NewSet() *AggregatorSet {

	return &AggregatorSet {

		aggs:  make(map[string]Aggregator),
	}


}



func NewSetFromMessages(messages []*protos.Aggregator) (*AggregatorSet, error) {

	result := &AggregatorSet{aggs: make(map[string]Aggregator)}

	if messages == nil {

			return result, nil
	}

	for _, proto := range messages {

		id := proto.Id

		_, ok := result.aggs[id]
		if ok {
				return nil, fmt.Errorf("duplicate aggregator id: %s", id)
		}

		agg, err := newAggregator(proto.Type)
		if err != nil {

			return nil, err
		}

		err = agg.SetBytes(proto.value)
		if err != nil {
			return nil, err
		}

		result.aggs[id] = agg

	}


return result, nil


}

func (set *AggregatorSet) GetValue(id string) (interface{}, bool) {

	set.mutex.RLock()
	defer set.mutex.RUnlock()

	agg, ok := set.aggs[id]

}










