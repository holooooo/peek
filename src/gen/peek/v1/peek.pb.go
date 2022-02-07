// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: peek.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobCommand int32

const (
	JobCommand_binary         JobCommand = 0
	JobCommand_go_profile_cpu JobCommand = 1
)

// Enum value maps for JobCommand.
var (
	JobCommand_name = map[int32]string{
		0: "binary",
		1: "go_profile_cpu",
	}
	JobCommand_value = map[string]int32{
		"binary":         0,
		"go_profile_cpu": 1,
	}
)

func (x JobCommand) Enum() *JobCommand {
	p := new(JobCommand)
	*p = x
	return p
}

func (x JobCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_peek_proto_enumTypes[0].Descriptor()
}

func (JobCommand) Type() protoreflect.EnumType {
	return &file_peek_proto_enumTypes[0]
}

func (x JobCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobCommand.Descriptor instead.
func (JobCommand) EnumDescriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{0}
}

type JobStatus int32

const (
	JobStatus_queued    JobStatus = 0
	JobStatus_running   JobStatus = 1
	JobStatus_failed    JobStatus = 2
	JobStatus_completed JobStatus = 3
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0: "queued",
		1: "running",
		2: "failed",
		3: "completed",
	}
	JobStatus_value = map[string]int32{
		"queued":    0,
		"running":   1,
		"failed":    2,
		"completed": 3,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_peek_proto_enumTypes[1].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_peek_proto_enumTypes[1]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{1}
}

type AgentType int32

const (
	Agent_pod AgentType = 0
	Agent_vm  AgentType = 1
)

// Enum value maps for AgentType.
var (
	AgentType_name = map[int32]string{
		0: "pod",
		1: "vm",
	}
	AgentType_value = map[string]int32{
		"pod": 0,
		"vm":  1,
	}
)

func (x AgentType) Enum() *AgentType {
	p := new(AgentType)
	*p = x
	return p
}

func (x AgentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentType) Descriptor() protoreflect.EnumDescriptor {
	return file_peek_proto_enumTypes[2].Descriptor()
}

func (AgentType) Type() protoreflect.EnumType {
	return &file_peek_proto_enumTypes[2]
}

func (x AgentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentType.Descriptor instead.
func (AgentType) EnumDescriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{1, 0}
}

type AgentAgentStatus int32

const (
	Agent_initializing AgentAgentStatus = 0
	Agent_running      AgentAgentStatus = 1
	Agent_terminating  AgentAgentStatus = 2
	Agent_terminated   AgentAgentStatus = 3
	Agent_released     AgentAgentStatus = 4
)

// Enum value maps for AgentAgentStatus.
var (
	AgentAgentStatus_name = map[int32]string{
		0: "initializing",
		1: "running",
		2: "terminating",
		3: "terminated",
		4: "released",
	}
	AgentAgentStatus_value = map[string]int32{
		"initializing": 0,
		"running":      1,
		"terminating":  2,
		"terminated":   3,
		"released":     4,
	}
)

func (x AgentAgentStatus) Enum() *AgentAgentStatus {
	p := new(AgentAgentStatus)
	*p = x
	return p
}

func (x AgentAgentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentAgentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_peek_proto_enumTypes[3].Descriptor()
}

func (AgentAgentStatus) Type() protoreflect.EnumType {
	return &file_peek_proto_enumTypes[3]
}

func (x AgentAgentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentAgentStatus.Descriptor instead.
func (AgentAgentStatus) EnumDescriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{1, 1}
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Pod       string `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{0}
}

func (x *Target) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Target) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Target      *Target          `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Ip          string           `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Status      AgentAgentStatus `protobuf:"varint,4,opt,name=status,proto3,enum=peek.v1.AgentAgentStatus" json:"status,omitempty"`
	Jobs        []*Job           `protobuf:"bytes,5,rep,name=jobs,proto3" json:"jobs,omitempty"`
	CurrentJobs []*Job           `protobuf:"bytes,6,rep,name=current_jobs,json=currentJobs,proto3" json:"current_jobs,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{1}
}

func (x *Agent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Agent) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Agent) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Agent) GetStatus() AgentAgentStatus {
	if x != nil {
		return x.Status
	}
	return Agent_initializing
}

func (x *Agent) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *Agent) GetCurrentJobs() []*Job {
	if x != nil {
		return x.CurrentJobs
	}
	return nil
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command JobCommand `protobuf:"varint,2,opt,name=command,proto3,enum=peek.v1.JobCommand" json:"command,omitempty"`
	Args    []string   `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	Config  *JobConfig `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	Result  *JobResult `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	Status  JobStatus  `protobuf:"varint,6,opt,name=status,proto3,enum=peek.v1.JobStatus" json:"status,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{2}
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetCommand() JobCommand {
	if x != nil {
		return x.Command
	}
	return JobCommand_binary
}

func (x *Job) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Job) GetConfig() *JobConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Job) GetResult() *JobResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Job) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_queued
}

type JobConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAsync bool  `protobuf:"varint,1,opt,name=isAsync,proto3" json:"isAsync,omitempty"`
	Timeout int32 `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *JobConfig) Reset() {
	*x = JobConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobConfig) ProtoMessage() {}

func (x *JobConfig) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobConfig.ProtoReflect.Descriptor instead.
func (*JobConfig) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{3}
}

func (x *JobConfig) GetIsAsync() bool {
	if x != nil {
		return x.IsAsync
	}
	return false
}

func (x *JobConfig) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type JobResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*JobResult_Stdout
	//	*JobResult_DownloadUrl
	Result   isJobResult_Result `protobuf_oneof:"result"`
	Error    string             `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ExitCode int32              `protobuf:"varint,4,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *JobResult) Reset() {
	*x = JobResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResult) ProtoMessage() {}

func (x *JobResult) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResult.ProtoReflect.Descriptor instead.
func (*JobResult) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{4}
}

func (m *JobResult) GetResult() isJobResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *JobResult) GetStdout() string {
	if x, ok := x.GetResult().(*JobResult_Stdout); ok {
		return x.Stdout
	}
	return ""
}

func (x *JobResult) GetDownloadUrl() string {
	if x, ok := x.GetResult().(*JobResult_DownloadUrl); ok {
		return x.DownloadUrl
	}
	return ""
}

func (x *JobResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *JobResult) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type isJobResult_Result interface {
	isJobResult_Result()
}

type JobResult_Stdout struct {
	Stdout string `protobuf:"bytes,1,opt,name=stdout,proto3,oneof"`
}

type JobResult_DownloadUrl struct {
	DownloadUrl string `protobuf:"bytes,2,opt,name=download_url,json=downloadUrl,proto3,oneof"`
}

func (*JobResult_Stdout) isJobResult_Result() {}

func (*JobResult_DownloadUrl) isJobResult_Result() {}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{5}
}

type AgentAliveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alive bool `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *AgentAliveResponse) Reset() {
	*x = AgentAliveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentAliveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentAliveResponse) ProtoMessage() {}

func (x *AgentAliveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentAliveResponse.ProtoReflect.Descriptor instead.
func (*AgentAliveResponse) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{6}
}

func (x *AgentAliveResponse) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

type Jobs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *Jobs) Reset() {
	*x = Jobs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jobs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jobs) ProtoMessage() {}

func (x *Jobs) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jobs.ProtoReflect.Descriptor instead.
func (*Jobs) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{7}
}

func (x *Jobs) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type ManagerGetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	JobName string  `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
}

func (x *ManagerGetJobRequest) Reset() {
	*x = ManagerGetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peek_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerGetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerGetJobRequest) ProtoMessage() {}

func (x *ManagerGetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peek_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerGetJobRequest.ProtoReflect.Descriptor instead.
func (*ManagerGetJobRequest) Descriptor() ([]byte, []int) {
	return file_peek_proto_rawDescGZIP(), []int{8}
}

func (x *ManagerGetJobRequest) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *ManagerGetJobRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

var File_peek_proto protoreflect.FileDescriptor

var file_peek_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x65,
	0x65, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x38, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x22,
	0xd3, 0x02, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x6a,
	0x6f, 0x62, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x65, 0x65, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x2f, 0x0a,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x22, 0x17,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x76, 0x6d, 0x10, 0x01, 0x22, 0x5c, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x10, 0x04, 0x22, 0xe4, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x0a,
	0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41,
	0x73, 0x79, 0x6e, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x88,
	0x01, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a,
	0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x28, 0x0a, 0x04, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x20, 0x0a,
	0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x65,
	0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22,
	0x5a, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x2d, 0x0a, 0x0b, 0x6a,
	0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x70, 0x75, 0x10, 0x01, 0x2a, 0x40, 0x0a, 0x0a, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03, 0x32, 0xc6, 0x01, 0x0a,
	0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x0c, 0x2e, 0x70,
	0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x0c, 0x2e, 0x70, 0x65, 0x65,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70,
	0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x0c, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x1a, 0x0c, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x6f, 0x62, 0x22, 0x00, 0x32, 0xaa, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x65, 0x65, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peek_proto_rawDescOnce sync.Once
	file_peek_proto_rawDescData = file_peek_proto_rawDesc
)

func file_peek_proto_rawDescGZIP() []byte {
	file_peek_proto_rawDescOnce.Do(func() {
		file_peek_proto_rawDescData = protoimpl.X.CompressGZIP(file_peek_proto_rawDescData)
	})
	return file_peek_proto_rawDescData
}

var file_peek_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_peek_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_peek_proto_goTypes = []interface{}{
	(JobCommand)(0),              // 0: peek.v1.job_command
	(JobStatus)(0),               // 1: peek.v1.job_status
	(AgentType)(0),               // 2: peek.v1.Agent.type
	(AgentAgentStatus)(0),        // 3: peek.v1.Agent.agent_status
	(*Target)(nil),               // 4: peek.v1.Target
	(*Agent)(nil),                // 5: peek.v1.Agent
	(*Job)(nil),                  // 6: peek.v1.Job
	(*JobConfig)(nil),            // 7: peek.v1.job_config
	(*JobResult)(nil),            // 8: peek.v1.job_result
	(*EmptyRequest)(nil),         // 9: peek.v1.EmptyRequest
	(*AgentAliveResponse)(nil),   // 10: peek.v1.AgentAliveResponse
	(*Jobs)(nil),                 // 11: peek.v1.Jobs
	(*ManagerGetJobRequest)(nil), // 12: peek.v1.ManagerGetJobRequest
}
var file_peek_proto_depIdxs = []int32{
	4,  // 0: peek.v1.Agent.target:type_name -> peek.v1.Target
	3,  // 1: peek.v1.Agent.status:type_name -> peek.v1.Agent.agent_status
	6,  // 2: peek.v1.Agent.jobs:type_name -> peek.v1.Job
	6,  // 3: peek.v1.Agent.current_jobs:type_name -> peek.v1.Job
	0,  // 4: peek.v1.Job.command:type_name -> peek.v1.job_command
	7,  // 5: peek.v1.Job.config:type_name -> peek.v1.job_config
	8,  // 6: peek.v1.Job.result:type_name -> peek.v1.job_result
	1,  // 7: peek.v1.Job.status:type_name -> peek.v1.job_status
	6,  // 8: peek.v1.Jobs.jobs:type_name -> peek.v1.Job
	4,  // 9: peek.v1.ManagerGetJobRequest.target:type_name -> peek.v1.Target
	9,  // 10: peek.v1.AgentService.GetInfo:input_type -> peek.v1.EmptyRequest
	6,  // 11: peek.v1.AgentService.CreateJob:input_type -> peek.v1.Job
	9,  // 12: peek.v1.AgentService.GetJobs:input_type -> peek.v1.EmptyRequest
	6,  // 13: peek.v1.AgentService.GetJob:input_type -> peek.v1.Job
	4,  // 14: peek.v1.ManagerService.CreateAgent:input_type -> peek.v1.Target
	4,  // 15: peek.v1.ManagerService.GetAgent:input_type -> peek.v1.Target
	12, // 16: peek.v1.ManagerService.GetJob:input_type -> peek.v1.ManagerGetJobRequest
	5,  // 17: peek.v1.AgentService.GetInfo:output_type -> peek.v1.Agent
	6,  // 18: peek.v1.AgentService.CreateJob:output_type -> peek.v1.Job
	11, // 19: peek.v1.AgentService.GetJobs:output_type -> peek.v1.Jobs
	6,  // 20: peek.v1.AgentService.GetJob:output_type -> peek.v1.Job
	5,  // 21: peek.v1.ManagerService.CreateAgent:output_type -> peek.v1.Agent
	5,  // 22: peek.v1.ManagerService.GetAgent:output_type -> peek.v1.Agent
	6,  // 23: peek.v1.ManagerService.GetJob:output_type -> peek.v1.Job
	17, // [17:24] is the sub-list for method output_type
	10, // [10:17] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_peek_proto_init() }
func file_peek_proto_init() {
	if File_peek_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peek_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentAliveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jobs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peek_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerGetJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_peek_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*JobResult_Stdout)(nil),
		(*JobResult_DownloadUrl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_peek_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_peek_proto_goTypes,
		DependencyIndexes: file_peek_proto_depIdxs,
		EnumInfos:         file_peek_proto_enumTypes,
		MessageInfos:      file_peek_proto_msgTypes,
	}.Build()
	File_peek_proto = out.File
	file_peek_proto_rawDesc = nil
	file_peek_proto_goTypes = nil
	file_peek_proto_depIdxs = nil
}
