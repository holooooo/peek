package agent

import (
	"context"
	"net"
	peekv1 "peek/src/gen/peek/v1"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/client-go/kubernetes"
)

var (
	logger = logrus.New().WithField("component", "agent")
)

type Server struct {
	peekv1.UnimplementedAgentServiceServer
	config    Config
	agent     *peekv1.Agent
	clientSet *kubernetes.Clientset
	worker    *Worker
}

type Config struct {
	ServerAddress string
}

func Serve(clientSet *kubernetes.Clientset, config Config) error {
	server := &Server{
		config:    config,
		clientSet: clientSet,
		worker:    NewWorker(make(chan struct{})), // TODO Graceful exit
	}
	lis, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	var opts = []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
	}
	grpcServer := grpc.NewServer(opts...)
	peekv1.RegisterAgentServiceServer(grpcServer, server)
	logrus.Infof("ready to serve")
	return grpcServer.Serve(lis)
}

func (s *Server) CreateJob(_ context.Context, in *peekv1.Job) (*peekv1.Job, error) {
	s.worker.AddJob(in)
	return in, nil
}

func (s *Server) GetJobs(_ context.Context, _ *peekv1.EmptyRequest) (*peekv1.Jobs, error) {
	return &peekv1.Jobs{
		Jobs: s.worker.GetJobs(),
	}, nil
}

func (s *Server) GetJob(_ context.Context, job *peekv1.Job) (*peekv1.Job, error) {
	return s.worker.GetJob(job.Name), nil
}

func (s *Server) GetInfo(_ context.Context, _ *peekv1.EmptyRequest) (*peekv1.Agent, error) {
	return s.agent, nil
}