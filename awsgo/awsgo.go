package awsgo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const defaultRegion = "us-east-1"

var Ctx context.Context
var Cfg aws.Config
var err error

func InitAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(defaultRegion))
	if err != nil {
		panic("Error al cargar la configuración .aws/config " + err.Error())
	}
}
