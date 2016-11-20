# watson

This is an simple wrapper for Watson conversation service, based on
http://www.ibm.com/watson/developercloud/conversation/api/v1/


   - [x] Conversation Service

### Conversation Service

Credentials creation, you must to provide the credentials provide by watson conversational service.
```go

import "github.com/cristianchaparroa/watson/conversational"

passwordConversational 	:= os.Getenv("PASSWORD_WATSON_CONVERSATION")
userConversational 	:= os.Getenv("USERNAME_WATSON_CONVERSATION")
urlConversational 	:= os.Getenv("URL_WATSON_CONVERSATION")
workspaceId 		:= os.Getenv("WORKSPACE_ID_WATSON_CONVERSATION")


credentials := &watson.Credentials{
		Password: passwordConversational,
		Username: userConversational,
		Url: urlConversational}
```

Create a service instance with the credentials

```go
service := watson.NewConversationService(credentials,workspaceId)
```

Process messages with watson conversational,  there are two kind of messages
 - Opening message (first message in the conversation)
 - Message in the conversational (any message after the first message, keeps a conversation context)

Opening message example
```go
message := "string message to process"
m := watson.NewMessageRequest(message)

intent, err := service..NewMessage(m)
```

Example of message with conversation context

```go
contextString := `  "context": {
    "conversation_id": "f1ab5f76-f41b-47b4-a8dc-e1c32b925b79",
    "system": {
      "dialog_stack": [
        {
          "dialog_node": "root"
        }
      ],
      "dialog_turn_counter": 2,
      "dialog_request_counter": 2
    },
    "defaultCounter": 0
}`
cmsBytes := []byte(contextString)
var contextMessage watson.Context
err := json.Unmarshal(cmsBytes, &contextMessage)

if err != nil {
  ...
}

message := "string message to process"
m := watson.NewMessageRequest(message)
m.Context = contextMessage
intent, err := service..NewMessage(m)
```
