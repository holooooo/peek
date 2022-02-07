package manager

import (
	"context"
	"fmt"
	"net"
	"peek/src/gen/peek/v1"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	matev1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var logger = logrus.New().WithField("component", "manager")

// errors
var (
	// AgentNotFound for agent
	AgentNotFound = fmt.Errorf("not found")
)

type Server struct {
	v1.UnimplementedManagerServiceServer
	config Config

	clientSet *kubernetes.Clientset
}

type Config struct {
	ServerAddress string

	AgentImage     string
	AgentNamespace string
}

func Serve(clientSet *kubernetes.Clientset, config Config) error {
	server := &Server{
		config:    config,
		clientSet: clientSet,
	}
	lis, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	var opts = []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
	}
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterManagerServiceServer(grpcServer, server)
	logrus.Infof("ready to serve")
	return grpcServer.Serve(lis)
}

func (s *Server) CreateAgent(ctx context.Context, target *v1.Target) (*v1.Agent, error) {
	agentObj, _ := s.GetAgent(ctx, target)
	if agentObj != nil {
		return agentObj, nil
	}
	return s.createAgent(ctx, target)
}

func (s *Server) createAgent(ctx context.Context, target *v1.Target) (*v1.Agent, error) {
	pod, err := s.clientSet.CoreV1().Pods(target.Namespace).Get(ctx, target.Pod, matev1.GetOptions{})
	if err != nil {
		return nil, err
	}
	var Privileged = true
	agentPod := corev1.Pod{
		ObjectMeta: matev1.ObjectMeta{
			Name:      "peek-" + hash(pod.ObjectMeta),
			Namespace: s.config.AgentNamespace,
			Labels: map[string]string{
				"peek-agent": "true",
				"target-pod": pod.Name,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "main",
					Image:   s.config.AgentImage,
					Command: []string{"/apps/agent"},
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceMemory: *resource.NewQuantity(2*1024, resource.BinarySI),
							corev1.ResourceCPU:    *resource.NewMilliQuantity(2000, resource.DecimalSI),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: *resource.NewQuantity(1*1024, resource.BinarySI),
							corev1.ResourceCPU:    *resource.NewMilliQuantity(1000, resource.DecimalSI),
						},
					},
					ImagePullPolicy: corev1.PullIfNotPresent,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &Privileged, // TODO make it reasonable
					},
				},
			},
			RestartPolicy:    corev1.RestartPolicyAlways,
			NodeName:         pod.Spec.NodeName,
			HostPID:          true,
			ImagePullSecrets: pod.Spec.ImagePullSecrets,
			Hostname:         pod.Spec.Hostname,
			Affinity:         pod.Spec.Affinity,
			SchedulerName:    pod.Spec.SchedulerName,
			Tolerations:      pod.Spec.Tolerations,
			HostAliases:      pod.Spec.HostAliases,
			DNSConfig:        pod.Spec.DNSConfig,
		},
	}
	_, err = s.clientSet.CoreV1().Pods(target.Namespace).Create(ctx, &agentPod, matev1.CreateOptions{})
	if err != nil {
		logger.Infof("create new agent %s/%s failed: %s", agentPod.Namespace, agentPod.Name, err)
		return nil, err
	}
	logger.Infof("create new agent %s/%s successed", agentPod.Namespace, agentPod.Name)
	return s.GetAgent(ctx, target)
}

func (s *Server) GetAgent(ctx context.Context, target *v1.Target) (*v1.Agent, error) {
	pods, err := s.clientSet.CoreV1().Pods(target.GetNamespace()).List(ctx, matev1.ListOptions{
		LabelSelector: "peek-agent=true,target-pod=" + target.Pod,
	})
	if err != nil {
		return nil, err
	}
	if len(pods.Items) == 0 {
		return nil, AgentNotFound
	}
	pod := pods.Items[0]

	// create agent client
	c, err := getAgentClient(pod.Status.PodIP)
	if err != nil {
		return nil, err
	}

	agent, err := c.GetInfo(ctx, &v1.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return agent, nil
}

func getAgentClient(ip string) (v1.AgentServiceClient, error) {
	conn, err := grpc.Dial(ip+":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Warnf("connect to agent %s failed: %s", ip, err)
		return nil, err
	}
	c := v1.NewAgentServiceClient(conn)
	return c, nil
}

func (s *Server) GetJob(ctx context.Context, req *v1.ManagerGetJobRequest) (*v1.Job, error) {
	agent, err := s.GetAgent(ctx, req.Target)
	if err != nil {
		logger.Warnf("failed to get agent %s: %s", req.Target.Pod, err)
		return nil, err
	}

	c, err := getAgentClient(agent.Ip)
	return c.GetJob(ctx, &v1.Job{Name: req.JobName})
}