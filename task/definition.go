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
	"github.com/dubbikins/envy"
)

type Config struct {
	TableName string `env:"TASKS_TABLE_NAME" default:"tasks"`
}
type Definition struct {
	Name      string    `dynamodbav:"name" json:"name"`
	Type      string    `dynamodbav:"type" json:"type"`
	Cmd       string    `dynamodbav:"cmd" json:"cmd"`
	Args      []string  `dynamodbav:"args" json:"args"`
	CreatedAt time.Time `dynamodbav:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `dynamodbav:"updated_at,omitempty" json:"updated_at"`
}

func (defn *Definition) Unmarshal(reader io.Reader) (err error) {
	var data []byte
	if data, err = io.ReadAll(reader); err != nil {
		return
	}
	fmt.Println(string(data))
	return json.Unmarshal(data, defn)
}
func (defn *Definition) From(ctx context.Context, dynamodb_client dynamodb.Client, name string) (err error) {
	var resp *dynamodb.GetItemOutput
	var cfg *Config
	if cfg, err = envy.New(envy.FromEnvironment[Config]); err != nil {
		return
	}
	//If the definition is not found, return an error
	if resp, err = dynamodb_client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(cfg.TableName),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{
				Value: name,
			},
			"type": &types.AttributeValueMemberS{
				Value: "defn",
			},
		},
		ConsistentRead: aws.Bool(true),
	}); err != nil {
		return
	}
	if err = attributevalue.UnmarshalMap(resp.Item, defn); err != nil {
		log.Println("Error unmarshalling task definition")
		return
	}
	log.Printf("Task command '%s' with args '%s'\n", defn.Cmd, defn.Args)
	return
}
func (defn *Definition) Save(ctx context.Context, dynamodb_client *dynamodb.Client) (err error) {
	var resp *dynamodb.PutItemOutput
	var cfg *Config
	var item map[string]types.AttributeValue
	if cfg, err = envy.New(envy.FromEnvironment[Config]); err != nil {
		return
	}
	defn.CreatedAt = time.Now()
	defn.UpdatedAt = time.Now()
	if item, err = attributevalue.MarshalMap(defn); err != nil {
		return
	}
	fmt.Println(defn)
	fmt.Println(item)
	//If the definition is not found, return an error
	if resp, err = dynamodb_client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(cfg.TableName),

		Item: item,
	}); err != nil {
		return
	}
	log.Println(resp.ResultMetadata)
	log.Println("Task created successfully")
	return
}

func (defn *Definition) CreateGlobalTable(ctx context.Context, dynamodb_client *dynamodb.Client) (err error) {
	var resp *dynamodb.CreateTableOutput
	var cfg *Config
	if cfg, err = envy.New(envy.FromEnvironment[Config]); err != nil {
		return
	}

	//If the definition is not found, return an error
	if resp, err = dynamodb_client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(cfg.TableName),
		
		
	}); err != nil {
		return
	}
	log.Println(resp.ResultMetadata)
	log.Println("Global Task table created successfully")
	return
}

type TaskDefinitions []*Definition

func (t *TaskDefinitions) List(ctx context.Context, client *dynamodb.Client) (result TaskDefinitions, err error) {
	var cfg *Config
	if cfg, err = envy.New(envy.FromEnvironment[Config]); err != nil {
		return
	}
	result = TaskDefinitions{}
	var request = &dynamodb.ScanInput{
		TableName: aws.String(cfg.TableName),
	}
	var scanPaginator = dynamodb.NewScanPaginator(client, request)
	log.Println("Begin Scanning")
	for scanPaginator.HasMorePages() {
		var tasksDefns = []*Definition{}
		var resp *dynamodb.ScanOutput
		log.Println("Fetching Next Page")
		if resp, err = scanPaginator.NextPage(ctx); err != nil {
			log.Println(err)
			return
		}

		if err = attributevalue.UnmarshalListOfMaps(resp.Items, &tasksDefns); err != nil {
			log.Println(err)
			return
		}

		result = append(result, tasksDefns...)
	}
	log.Println(result)
	return
}
