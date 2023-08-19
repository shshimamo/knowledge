import { App, AwsLambdaReceiver } from '@slack/bolt';
import axios from 'axios';

const awsLambdaReceiver = new AwsLambdaReceiver({
  signingSecret: process.env.SLACK_SIGNING_SECRET,
});

const app = new App({
  token: process.env.SLACK_BOT_TOKEN,
  receiver: awsLambdaReceiver,
});

const githubToken = process.env.GITHUB_TOKEN;
const owner = 'shshimamo';
const repo = 'knowledge';
const allowedEventType = [
  'build-push-all',
  'build-push-backend-auth',
  'build-push-backend-auth-migration',
  'build-push-backend-auth-all',
  'build-push-backend-main',
  'build-push-backend-main-migration',
  'build-push-backend-main-all',
  'build-push-frontend-main'
];

app.command('/gha', async ({ command, ack, respond }) => {
  await ack();

  // Get Event Type from command
  const eventType = command.text
  if (!allowedEventType.includes(eventType)) {
    return respond(`[ERROR] ${eventType} is Invalid. Valid EventTypes: ${allowedEventType.join(', ')}`);
  }

  try {
    const res = await axios.post(`https://api.github.com/repos/${owner}/${repo}/dispatches`,
      { event_type: eventType },
      {
        headers: {
          Accept: 'application/vnd.github.everest-preview+json',
          Authorization: `token ${githubToken}`
        }
      }
    )
    return respond(`[SUCCESS] eventType: ${eventType}, status: ${res.status}`);
  } catch (error) {
    return respond(`[ERROR] eventType:${eventType},  message:${error.message}`);
  }
});

export const handler = async (event: any, context: any, callback: any) => {
  const handler = await awsLambdaReceiver.start();
  return handler(event, context, callback);
}