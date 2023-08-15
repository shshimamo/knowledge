#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { RdsStack } from '../lib/rds-stack';

const app = new cdk.App();

const vpcId = process.env.CDK_VPC_ID;
if (!vpcId) {
  throw new Error("Environment variable CDK_VPC_ID is not defined");
}
const eksNodeSGId = process.env.CDK_EKS_NODE_SG_ID;
if (!eksNodeSGId) {
  throw new Error("Environment variable CDK_EKS_NODE_SG_ID is not defined");
}
const dbPassword = process.env.CDK_DB_PASSWORD;
if (!dbPassword) {
  throw new Error("Environment variable CDK_DB_PASSWORD is not defined");
}
const myIp = process.env.CDK_MY_IP;
if (!myIp) {
  throw new Error("Environment variable CDK_MY_IP is not defined");
}

new RdsStack(app, 'RdsStack', {
  env: {account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION},
  vpcId,
  eksNodeSGId,
  dbPassword,
  myIp,
});