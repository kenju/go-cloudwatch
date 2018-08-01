package cloudwatch

// AlarmMessage is a message JSON sent from CloudWatch Alarm via SNS Topic.
type AlarmMessage struct {
	AlarmName        string       `json:"AlarmName"`
	AlarmDescription string       `json:"AlarmDescription"`
	AWSAccountID     string       `json:"AWSAccountId"`
	NewStateValue    string       `json:"NewStateValue"`
	NewStateReason   string       `json:"NewStateReason"`
	StateChangeTime  string       `json:"StateChangeTime"`
	Region           string       `json:"Region"`
	OldStateValue    string       `json:"OldStateValue"`
	Trigger          AlarmTrigger `json:"Trigger"`
}

// AlarmTrigger is a JSON of CloudWatchAlarmMessage.
type AlarmTrigger struct {
	MetricName                       string           `json:"MetricNameOldStateValue"`
	Namespace                        string           `json:"Namespace"`
	StatisticType                    string           `json:"StatisticType"`
	Unit                             int              `json:"Unit"`
	Dimensions                       []AlarmDimension `json:"Dimensions"`
	Period                           int              `json:"period"`
	EvaluationPeriods                int              `json:"EvaluationPeriods"`
	ComparisonOperator               string           `json:"ComparisonOperator"`
	Threshold                        int              `json:"Threshold"`
	TreatMissingData                 string           `json:"TreatMissingData"`
	EvaluateLowSampleCountPercentile string           `json:"EvaluateLowSampleCountPercentile"`
}

// AlarmDimension is a JSON of CloudWatchAlarmTrigger.
type AlarmDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
