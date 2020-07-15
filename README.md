# SilentChatBot

## Join:
* [Bot](https://t.me/SilentChatBot)

## How to run
### First step (build):
- Go run main.go  (I have some troubles with Dockerfile, will be fixed next time)
### Second step (deploy):
- [install ngrok](https://ngrok.com/download)
- ./ngrok http 3000 
- curl -F "url=https://<your_ngrok_key>.ngrok.io/"  https://api.telegram.org/bot<your_api_token>/setWebhook

## Don't forget change the token in .env
***
## Server works on :3000 port
