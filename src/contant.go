package src

type configPath string

const (
	messagePath configPath = "message.json"
)

type replyMsg string

const (
	welcome      replyMsg = "Hello~ I'm TsaiTsai. Let me introduce you my best friend, Denny Tsai! Please tap the following button to see more."
	defaultReply replyMsg = "Sorry~ I can't understand what you are saying. Please type `help` for more infomation."
	helpReply    replyMsg = "`HP` -> handsome photo"
)

type specialWord string

const (
	handsomePhoto specialWord = "HP"
	help          specialWord = "help"
)
