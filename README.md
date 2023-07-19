
# Why use programming languages for infrastructure as code? Modularity.

A functionality is defined that receives a post request and stores the request. \
However, this functionality is implemented in four different ways to show the modular changes in the infrastructure with dev and prod configurations. \
 You can deploy this functionality through a pipeline too.


## Workflows

![Alt text](/media/lambdatable.png "LambdaTable")
![Alt text](/media/lambdabucket.png "LambdaBucket")
![Alt text](/media/fargatebucket.png "FargateBucket")
![Alt text](/media/fargatetable.png "FargateTable")


## Take a look


Clone the project
```bash
 git clone https://github.com/orionverso/echoapp
```
Go to the project directory

```bash
  cd echoapp
```
Choose one stack in the main function
```go
package main

import (
	"castor/stack/computesave"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	
    // STACKS DIRECTLY
	//computesave.NewFargateWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveBlockDataProps_DEV)
	// computesave.NewFargateWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveObjectProps_DEV)
	//	computesave.NewApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveBlockDataProps_DEV)
	computesave.NewApiWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveObjectProps_DEV)


	app.Synth(nil)
}

```

Generate cloudformation template

```bash
  ./cdk-synth-cross-account.sh <dev-region> <dev-account> <prod-region> <prod-account> --all
```
Deploy via script

```bash
  ./cdk-deploy-cross-account.sh <dev-region> <dev-account> <prod-region> <prod-account> --profile <aws-user-with-admin-permision> --all
```

Expected output
```bash
 
Outputs:
ComputeSavePipelineStateless1C16B3AD.ComputeApiGatewayWithLambdaProxyEndpointF7889D50 = https://<api-id>.execute-api.<region>.amazonaws.com/dev/
```
Check funcionality

```bash
curl https://<api-id>.execute-api.<region>.amazonaws.com/dev/ \
-X POST \
-d "$RANDOM"

Thank you for take a look. I am from Lambda writing to S3. See you 
```

let's just change the storage

```bash

  // STACKS DIRECTLY
//computesave.NewFargateWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveBlockDataProps_DEV)
	// computesave.NewFargateWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveObjectProps_DEV)
	computesave.NewApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveBlockDataProps_DEV)
	//computesave.NewApiWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveObjectProps_DEV)
```

Compare implementation
```bash
./cdk-diff-cross-account.sh <dev-region> <dev-account> <prod-region> <prod-account> --all

[-] AWS::S3::Bucket Stateless/Save/Bucket SaveBucket5C9EDB8F destroy
[-] AWS::S3::BucketPolicy Stateless/Save/Bucket/Policy SaveBucketPolicy255FE3E4 destroy
[-] Custom::S3AutoDeleteObjects Stateless/Save/Bucket/AutoDeleteObjectsCustomResource SaveBucketAutoDeleteObjectsCustomResourceE04EFD88 destroy
[-] AWS::IAM::Role Stateless/Custom::S3AutoDeleteObjectsCustomResourceProvider/Role CustomS3AutoDeleteObjectsCustomResourceProviderRole3B1BD092 destroy
[-] AWS::Lambda::Function Stateless/Custom::S3AutoDeleteObjectsCustomResourceProvider/Handler CustomS3AutoDeleteObjectsCustomResourceProviderHandler9D90184F destroy
[+] AWS::DynamoDB::Table Stateless/Save/Table SaveTable126A1A49
[~] AWS::IAM::Policy Stateless/Compute/FunctionWithSqsDestinations/LambdaFunction/ServiceRole/DefaultPolicy ComputeFunctionWithSqsDestinationsLambdaFunctionServiceRoleDefaultPolicy71D75A6C

Only the storage will change
```

let's just change the writer

```bash
  // STACKS DIRECTLY
    //computesave.NewFargateWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveBlockDataProps_DEV)
	computesave.NewFargateWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveObjectProps_DEV)
	//computesave.NewApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveBlockDataProps_DEV)
	//computesave.NewApiWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveObjectProps_DEV)
```

Compare implementation
```bash
./cdk-diff-cross-account.sh <dev-region> <dev-account> <prod-region> <prod-account> --all
...
[-] AWS::ApiGateway::Method Stateless/Compute/ApiGatewayWithLambdaProxy/Default/{proxy+}/ANY ComputeApiGatewayWithLambdaProxyproxyANY8A71D8E5 destroy
[-] AWS::Lambda::Permission Stateless/Compute/ApiGatewayWithLambdaProxy/Default/ANY/ApiPermission.ComputeSavePipelineStatelessComputeApiGatewayWithLambdaProxy84E1CEFD.ANY.. ComputeApiGatewayWithLambdaProxyANYApiPermissionComputeSavePipelineStatelessComputeApiGatewayWithLambdaProxy84E1CEFDANYF9700948 destroy
[-] AWS::Lambda::Permission Stateless/Compute/ApiGatewayWithLambdaProxy/Default/ANY/ApiPermission.Test.ComputeSavePipelineStatelessComputeApiGatewayWithLambdaProxy84E1CEFD.ANY.. ComputeApiGatewayWithLambdaProxyANYApiPermissionTestComputeSavePipelineStatelessComputeApiGatewayWithLambdaProxy84E1CEFDANY187A068C destroy
[-] AWS::ApiGateway::Method Stateless/Compute/ApiGatewayWithLambdaProxy/Default/ANY ComputeApiGatewayWithLambdaProxyANYC16C1EF9 destroy
[+] AWS::ElasticLoadBalancingV2::LoadBalancer Stateless/Compute/ApplicationLoadBalancedFargateService/LB ComputeApplicationLoadBalancedFargateServiceLB4A1ADA1C
[+] AWS::EC2::SecurityGroup Stateless/Compute/ApplicationLoadBalancedFargateService/LB/SecurityGroup ComputeApplicationLoadBalancedFargateServiceLBSecurityGroup03E0505C
[+] AWS::EC2::SecurityGroupEgress Stateless/Compute/ApplicationLoadBalancedFargateService/LB/SecurityGroup/to ComputeSavePipelineStatelessComputeApplicationLoadBalancedFargateServiceSecurityGroup7040C965:80 ComputeApplicationLoadBalancedFargateServiceLBSecurityGrouptoComputeSavePipelineStatelessComputeApplicationLoadBalancedFargateServiceSecurityGroup7040C96580AEE328B8
...

Only the writer will change
You can imagine the next implementation
```








## Pipelines bonus

The Fargate workloads requires a docker image in AWS Ecr. As the changes are modular, the same pipeline can be updated to serve all implementation versions.

![Alt text](/media/pipelambdabucket.png "PipeLambdaBucket")
![Alt text](/media/pipelambdatable.png "PipeLambdaTable")
![Alt text](/media/pipefargatetable.png "PipeFargateBucket")
![Alt text](/media/pipefargatebucket.png "PipeFargateTable")

If you are interest to execute the example cross account, you must bootstrapping first.

https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping.html

You can fork this repository to trigger the pipeline on push in your own repo.

```bash

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	
   // PIPELINES
	// pipeline.NewPipelineDeployFargateWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployFargateWriteToSaveBlockDataProps_DEV)
	pipeline.NewPipelineDeployFargateWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployFargateWriteToSaveObjectProps_DEV)
	// pipeline.NewPipelineDeployApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployApiWriteToSaveBlockDataProps_DEV)
	// pipeline.NewPipelineDeployApiWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployApiWriteToSaveObjectProps_DEV)



	app.Synth(nil)
}
```

## Conclusion

This project was a way to show myself how powerful building in the cloud by AWS cdk can be. \
Imagine that you had to rebuild an infrastructure to change the encryption of a bucket. It would be unmaintainable.\
However here we can create construct patterns , reuse and change them modularly.\
Finally, If your main cloud provider is AWS, I think that AWS cdk is a more than considerable option to build dynamic infrastructure that constantly adapts to new requirements.\
Can general purpose programming languages define infrastructure?  It's just the beginning.
