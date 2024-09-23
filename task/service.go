package task

import (
	context "context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Service struct {
	tasks           map[string]*Task
	EnvironmentName string
}

func NewService() *Service {
	return &Service{
		tasks: map[string]*Task{},
	}
}

// Get(context.Context, *Task) (*Execution, error)
// func (s *Service) ListRunning(*Filter, ) (err error) {
// 	return
// }

func (s *Service) Start(ctx context.Context, req *Task) (execution *Execution, err error) {
	fmt.Println(req)
	var aws_cfg aws.Config
	var dynamodb_client *dynamodb.Client
	var defn = &Definition{}
	execution = &Execution{
		State:  State_PENDING,
		Type:   fmt.Sprintf("exec::%s::%d", s.EnvironmentName, time.Now().Unix()),
		STDOUT: []byte("Task execution..."),
	}
	if aws_cfg, err = config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
		opts.Region = "us-east-1"
		return nil
	}); err != nil {
		return
	}
	dynamodb_client = dynamodb.NewFromConfig(aws_cfg)
	if err = defn.From(ctx, *dynamodb_client, req.Name); err != nil {
		log.Println("No task definition found")
		execution.State = State_FAILED
		execution.STDOUT = []byte(fmt.Sprintf("No task definition found for task name %s", req.Name))
		return
	}

	if err = execution.New(ctx, dynamodb_client); err != nil {
		execution.State = State_FAILED
		return
	}

	return
}

//	func (s *Service) Execute(req *Task, stream Tasks_ExecuteServer) (err error) {
//		fmt.Println(req)
//		var aws_cfg aws.Config
//		var dynamodb_client *dynamodb.Client
//		var defn_resp *dynamodb.GetItemOutput
//		var defn = &Definition{}
//		var ctx = context.Background()
//		var execution = &Execution{
//			State:  State_PENDING,
//			Type:   fmt.Sprintf("exec::%s::%d", s.EnvironmentName, time.Now().Unix()),
//			STDOUT: []byte("Task execution..."),
//		}
//		if aws_cfg, err = config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
//			opts.Region = "us-east-1"
//			return nil
//		}); err != nil {
//			return
//		}
//		dynamodb_client = dynamodb.NewFromConfig(aws_cfg)
//		//If the definition is not found, return an error
//		if defn_resp, err = dynamodb_client.GetItem(ctx, &dynamodb.GetItemInput{
//			TableName: aws.String("tasks"),
//			Key: map[string]types.AttributeValue{
//				"name": &types.AttributeValueMemberS{
//					Value: req.Name,
//				},
//				"type": &types.AttributeValueMemberS{
//					Value: "defn",
//				},
//			},
//			ConsistentRead: aws.Bool(true),
//		}); err != nil {
//			log.Println("No task definition found")
//			execution.State = State_FAILED
//			execution.STDOUT = []byte(fmt.Sprintf("No task definition found for task name %s", req.Name))
//			stream.Send(execution)
//			return
//		}
//		if err = attributevalue.UnmarshalMap(defn_resp.Item, &defn); err != nil {
//			log.Println("Error unmarshalling task definition")
//			return
//		}
//		log.Printf("Task command '%s' with args '%s'\n", defn.Cmd, defn.Args)
//		execution.New(ctx, dynamodb_client)
//		stream.Send(execution)
//		var cmd = exec.Command(defn.Cmd, defn.Args...)
//		if execution.STDOUT, err = cmd.Output(); err != nil {
//			log.Println(err)
//			execution.State = State_FAILED
//			execution.STDOUT = append(execution.STDOUT, []byte(fmt.Sprintf("\nTask execution failed: %s", err))...)
//		}
//		execution.State = State_FINISHED
//		stream.Send(execution)
//		return
//	}
func (e *Execution) AsTaskExecution() (te *TaskExecution) {
	te = &TaskExecution{
		State: e.State,
		Type:  e.Type,

		STDOUT: e.STDOUT,
	}
	if e.Created != nil {
		te.CreatedAt = e.Created.AsTime()
	}
	if e.Started != nil {
		te.CreatedAt = e.Started.AsTime()
	}
	if e.Ended != nil {
		t := e.Ended.AsTime()
		te.EndedAt = &t
	}
	return
}
func (e *Execution) New(ctx context.Context, client *dynamodb.Client) (err error) {
	e.State = State_PENDING
	var attribute_values map[string]types.AttributeValue
	if attribute_values, err = attributevalue.MarshalMap(e.AsTaskExecution()); err != nil {
		return fmt.Errorf("failed to marshal Record, %w", err)
	}
	log.Println(attribute_values)
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String("tasks"),
		Item:                attribute_values,
		ConditionExpression: aws.String("attribute_not_exists(name) AND attribute_not_exists(type)"),
	})

	return
}
