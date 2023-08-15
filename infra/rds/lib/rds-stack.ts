import * as cdk from 'aws-cdk-lib';
import * as ec2 from 'aws-cdk-lib/aws-ec2';
import * as rds from 'aws-cdk-lib/aws-rds';
import { Construct } from 'constructs';

interface RdsStackProps extends cdk.StackProps {
  vpcId: string,
  eksNodeSGId: string,
  dbPassword: string,
  myIp: string
}

export class RdsStack extends cdk.Stack {

  constructor(scope: Construct, id: string, props: RdsStackProps) {
    super(scope, id, props);

    // VPC
    const vpc = ec2.Vpc.fromLookup(this, 'EKS_VPC', {
      vpcId: props.vpcId,
    })

    // セキュリティグループ
    const rdsSG = new ec2.SecurityGroup(this, 'RDS_SG',
      {
        vpc: vpc,
        securityGroupName: 'rds_sg',
      }
    );
    const ekdNodeSG = ec2.SecurityGroup.fromSecurityGroupId(this, 'EKS_node_SG',
      props.eksNodeSGId,
      {
        mutable: false
      }
    );
    rdsSG.addIngressRule(
      ekdNodeSG,
      ec2.Port.tcp(5432),
      'allow db connection from node group'
    );
    rdsSG.addIngressRule(
      ec2.Peer.ipv4(`${props.myIp}/32`),
      ec2.Port.tcp(5432),
      'allow db connection from my ip'
    );

    // RDS
    const rdsInstance = new rds.DatabaseInstance(this, 'DBInstance', {
      engine: rds.DatabaseInstanceEngine.postgres({
        version: rds.PostgresEngineVersion.VER_14_7,
      }),
      instanceType: ec2.InstanceType.of(
        ec2.InstanceClass.T3,
        ec2.InstanceSize.MICRO
      ),
      vpc: vpc,
      vpcSubnets: {
        subnetType: ec2.SubnetType.PUBLIC,
      },
      publiclyAccessible: true,
      securityGroups: [rdsSG],
      credentials: rds.Credentials.fromPassword(
        'postgres',
        cdk.SecretValue.unsafePlainText(props.dbPassword)
      ),
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      deletionProtection: false,
    });
  }
}
