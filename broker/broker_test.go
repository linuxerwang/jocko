package broker

import (
	"context"
	"reflect"
	"testing"

	"github.com/travisjeffery/jocko"
	"github.com/travisjeffery/jocko/protocol"
	"github.com/travisjeffery/simplelog"
)

func TestNew(t *testing.T) {
	type args struct {
		id   int32
		opts []BrokerFn
	}
	tests := []struct {
		name    string
		args    args
		want    *Broker
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_Run(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		ctx       context.Context
		requestc  <-chan jocko.Request
		responsec chan<- jocko.Response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			b.Run(tt.args.ctx, tt.args.requestc, tt.args.responsec)
		})
	}
}

func TestBroker_Join(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		addrs []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.Join(tt.args.addrs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleAPIVersions(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		req    *protocol.APIVersionsRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.APIVersionsResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleAPIVersions(tt.args.header, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleAPIVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleCreateTopic(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		reqs   *protocol.CreateTopicRequests
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.CreateTopicsResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleCreateTopic(tt.args.header, tt.args.reqs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleCreateTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleDeleteTopics(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		reqs   *protocol.DeleteTopicsRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.DeleteTopicsResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleDeleteTopics(tt.args.header, tt.args.reqs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleDeleteTopics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleLeaderAndISR(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		req    *protocol.LeaderAndISRRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.LeaderAndISRResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleLeaderAndISR(tt.args.header, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleLeaderAndISR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleOffsets(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		req    *protocol.OffsetsRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.OffsetsResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleOffsets(tt.args.header, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleOffsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleProduce(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		req    *protocol.ProduceRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.ProduceResponses
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleProduce(tt.args.header, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleProduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleMetadata(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		req    *protocol.MetadataRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.MetadataResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleMetadata(tt.args.header, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_handleFetch(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		header *protocol.RequestHeader
		r      *protocol.FetchRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protocol.FetchResponses
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.handleFetch(tt.args.header, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.handleFetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_clusterMembers(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []*jocko.ClusterMember
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.clusterMembers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.clusterMembers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_isController(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.isController(); got != tt.want {
				t.Errorf("Broker.isController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_topicPartitions(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFound []*jocko.Partition
		wantErr   protocol.Error
	}{
		{
			name: "partitions found",
			fields: fields{
				topicMap: map[string][]*jocko.Partition{"topic": []*jocko.Partition{{ID: 1}}},
			},
			args:      args{topic: "topic"},
			wantFound: []*jocko.Partition{{ID: 1}},
			wantErr:   protocol.ErrNone,
		},
		{
			name: "partitions not found",
			fields: fields{
				topicMap: map[string][]*jocko.Partition{"topic": []*jocko.Partition{{ID: 1}}},
			},
			args:      args{topic: "not_topic"},
			wantFound: nil,
			wantErr:   protocol.ErrUnknownTopicOrPartition,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			gotFound, gotErr := b.topicPartitions(tt.args.topic)
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("Broker.topicPartitions() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Broker.topicPartitions() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestBroker_topics(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	topicMap := map[string][]*jocko.Partition{
		"topic": []*jocko.Partition{{ID: 1}},
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string][]*jocko.Partition
	}{
		{
			name: "topic map returned",
			fields: fields{
				topicMap: topicMap},
			want: topicMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.topics(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.topics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_partition(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		topic     string
		partition int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *jocko.Partition
		want1  protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			got, got1 := b.partition(tt.args.topic, tt.args.partition)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.partition() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Broker.partition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBroker_createPartition(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		partition *jocko.Partition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if err := b.createPartition(tt.args.partition); (err != nil) != tt.wantErr {
				t.Errorf("Broker.createPartition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBroker_clusterMember(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		id int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *jocko.ClusterMember
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.clusterMember(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.clusterMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_startReplica(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		partition *jocko.Partition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.startReplica(tt.args.partition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.startReplica() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_createTopic(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		topic             string
		partitions        int32
		replicationFactor int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.createTopic(tt.args.topic, tt.args.partitions, tt.args.replicationFactor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.createTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_deleteTopic(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.deleteTopic(tt.args.topic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.deleteTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_deletePartitions(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		tp *jocko.Partition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if err := b.deletePartitions(tt.args.tp); (err != nil) != tt.wantErr {
				t.Errorf("Broker.deletePartitions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBroker_Shutdown(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if err := b.Shutdown(); (err != nil) != tt.wantErr {
				t.Errorf("Broker.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBroker_becomeFollower(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		topic          string
		partitionID    int32
		partitionState *protocol.PartitionState
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.becomeFollower(tt.args.topic, tt.args.partitionID, tt.args.partitionState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.becomeFollower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroker_becomeLeader(t *testing.T) {
	type fields struct {
		logger      *simplelog.Logger
		id          int32
		topicMap    map[string][]*jocko.Partition
		replicators map[*jocko.Partition]*Replicator
		brokerAddr  string
		logDir      string
		raft        jocko.Raft
		serf        jocko.Serf
		shutdownCh  chan struct{}
		shutdown    bool
	}
	type args struct {
		topic          string
		partitionID    int32
		partitionState *protocol.PartitionState
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				logger:      tt.fields.logger,
				id:          tt.fields.id,
				topicMap:    tt.fields.topicMap,
				replicators: tt.fields.replicators,
				brokerAddr:  tt.fields.brokerAddr,
				logDir:      tt.fields.logDir,
				raft:        tt.fields.raft,
				serf:        tt.fields.serf,
				shutdownCh:  tt.fields.shutdownCh,
				shutdown:    tt.fields.shutdown,
			}
			if got := b.becomeLeader(tt.args.topic, tt.args.partitionID, tt.args.partitionState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Broker.becomeLeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		rs []int32
		r  int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.rs, tt.args.r); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
