package queue

const (
	EFFECT_ZERO = iota
	EFFECT_ONE
)
const (
	WRITE_ZERO = iota
	WRITE_ONE
)
const (
	TIME_WAIT = 100
)

// CreateTopic create a topic
func CreateTopic(sTopic string) string {
	return "list_" + sTopic + "_topic"
}

// CreateSetBodyKey create topic content
func CreateSetBodyKey(sBodyKey string) string {
	return "set_" + sBodyKey + "_body"
}
