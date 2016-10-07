package watson

type MessageRequest struct {
	Context Context `json:"context,omitempty"`
	Input   Input   `json:"input"`
}

func NewMessageRequest(text string) *MessageRequest {
	return &MessageRequest{Input: Input{Text: text}}
}

type Context struct {
	ConversationID string `json:"conversation_id,omitempty"`
	System         System `json:"system,omitempty"`
}

type System struct {
	DialogRequestCounter int           `json:"dialog_request_counter,omitempty"`
	DialogStack          []interface{} `json:"dialog_stack,omitempty"`
	DialogTurnCounter    int           `json:"dialog_turn_counter,omitempty"`
}

type Input struct {
	Text string `json:"text"`
}

// MessageResponse
type MessageResponse struct {
	Context  Context          `json:"context"`
	Entities []EntityResponse `json:"entities"`
	Input    InputResponse    `json:"input"`
	Intents  []Intent         `json:"intents"`
	Output   OutputResponse   `json:"output"`
}

type EntityResponse struct {
	Entity   string `json:"entity"`
	Location []int  `json:"location"`
	Value    string `json:"value"`
}

type InputResponse struct {
	Text string `json:"text"`
}

type Intent struct {
	Confidence float64 `json:"confidence"`
	Intent     string  `json:"intent"`
}

type OutputResponse struct {
	NodesVisited []string `json:"nodes_visited"`
	Text         []string `json:"text"`
}
