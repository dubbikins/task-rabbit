package task

import (
	context "context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TaskExecution struct {
	Name       string     `json:"name" dynamodbav:"name"`
	Type       string     `json:"type" dynamodbav:"type"`
	State      State      `json:"state" dynamodbav:"state"`
	CreatedAt  time.Time  `json:"created_at" dynamodbav:"created_at"`
	InstanceID string     `json:"instance_id" dynamodbav:"instance_id"`
	StartedAt  *time.Time `json:"started_at" dynamodbav:"started_at"`
	EndedAt    *time.Time `json:"ended_at,omitempty" dynamodbav:"ended"`
	STDOUT     []byte     `json:"stdout,omitempty" dynamodbav:"stdout"`
}

func (e *TaskExecution) Save(ctx context.Context, client *dynamodb.Client) (err error) {
	e.State = State_PENDING

	e.CreatedAt = time.Now()
	e.Type = fmt.Sprintf("exec::%s::%d", e.InstanceID, e.CreatedAt.Unix())
	var attribute_values map[string]types.AttributeValue
	if attribute_values, err = attributevalue.MarshalMap(e); err != nil {
		return fmt.Errorf("failed to marshal Record, %w", err)
	}
	log.Println(attribute_values)
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String("tasks"),
		Item:                attribute_values,
		ConditionExpression: aws.String("attribute_not_exists(#name) AND attribute_not_exists(#type)"),
		ExpressionAttributeNames: map[string]string{
			"#name": "name",
			"#type": "type",
		},
	})

	return
}

func (e *TaskExecution) Unmarshal(reader io.Reader) (err error) {
	var data []byte
	if data, err = io.ReadAll(reader); err != nil {
		return
	}
	fmt.Println(string(data))
	return json.Unmarshal(data, e)
}

func (e *Execution) Update(ctx context.Context, client *dynamodb.Client) (err error) {
	var attribute_values map[string]types.AttributeValue
	if attribute_values, err = attributevalue.MarshalMap(e); err != nil {
		return fmt.Errorf("failed to marshal Record, %w", err)
	}
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String("tasks"),
		Item:                attribute_values,
		ConditionExpression: aws.String("attribute_exists(name) AND attribute_exists(type)"),
	})
	return
}

type TaskExecutions []*TaskExecution

func (e *TaskExecutions) Find(ctx context.Context, client *dynamodb.Client) (err error) {
	var resp *dynamodb.QueryOutput
	var ok bool
	var name string
	if name, ok = ctx.Value(TASK_NAME_CONTEXT_KEY).(string); !ok {
		log.Println("ERROR: 'unable to retrieve task name from context'")
		return fmt.Errorf("unable to retrieve task name from context")
	}
	if resp, err = client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String("tasks"),
		KeyConditionExpression: aws.String("#TASKNAME = :name AND begins_with(#TASKTYPE, :type)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{
				Value: name,
			},
			":type": &types.AttributeValueMemberS{
				Value: "exec",
			},
		},
		ExpressionAttributeNames: map[string]string{
			"#TASKNAME": "name",
			"#TASKTYPE": "type",
		},
	}); err != nil {
		return
	}
	if err = attributevalue.UnmarshalListOfMaps(resp.Items, e); err != nil {
		return
	}
	return
}

// func (e *Executions) Marshal(ctx context.Context, w *http.ResponseWriter) (err error) {
// 	var data []byte
// 	if data, err = json.Marshal(e); err != nil {
// 	}
// 	return
// }
