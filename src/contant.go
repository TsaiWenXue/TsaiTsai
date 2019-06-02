package src

type configPath string

const (
	messagePath configPath = "message.json"

	welcome      = "Hello~ I'm TsaiTsai. Let me introduce you my best friend, Denny Tsai! Please tap the following button to see more."
	defaultReply = "Sorry~ I can't understand what you are saying. Please type `help` for more infomation."
	helpReply    = "`HP` -> handsome photo\n`project` -> Projects that Denny used to do"

	projectAltText = "Projects that Denny used to do"
)

type specialWord string

const (
	handsomePhoto specialWord = "HP"
	help          specialWord = "help"
	project       specialWord = "project"
)
