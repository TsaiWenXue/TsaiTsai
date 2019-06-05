package src

type configPath string

const messagePath configPath = "message.json"

const (
	// cnn news used
	https     = "https:"
	cnnDomain = "https://edition.cnn.com"
	//color
	white          = "#FFFFFF"
	black          = "#000000"
	gray           = "#AAAAAA"
	waterBlue      = "#2C91CA"
	lightOrange    = "#FD0000"
	lightRed       = "#FD0000"
	lightGreenBlue = "#D8FFEE"
	// Info template used
	dennyTsai      = "Denny Tsai"
	characteristic = "Characteristic"
	habit          = "Habit"
	motto          = "Motto"
)

const (
	welcome   = "Hello~ I'm TsaiTsai. Let me introduce you my best friend, Denny Tsai! Please tap the following button to see more."
	helpReply = "TsaTsai:\n\n`Hello` -> Say hello to me\n`HP` -> Handsome photo of Denny\n`Project` -> Projects that Denny used to do\n\nYou can click `More` below to get more information about Denny!"
	// Alternative Text
	projectAltText = "Projects that Denny used to do"
	newsAltText    = "Worlds news"
	infoAltText    = "Denny Tsai info"
)

type specialWord string

const (
	hello         specialWord = "hello"
	hi            specialWord = "hi"
	handsomePhoto specialWord = "hp"
	help          specialWord = "help"
	project       specialWord = "project"
	news          specialWord = "news"
	aboutDenny    specialWord = "About Denny"
)

const (
	// line sticker package id
	brownConySally = "11537"
	chocoFriends   = "11538"
	univerStarBT21 = "11539"
)

var (
	// line sticker package and sticker map
	stickersPackageMap = []string{brownConySally, chocoFriends, univerStarBT21}
	stickersMap        = map[string][]string{
		brownConySally: []string{"52002734", "52002735", "52002736", "52002737", "52002738", "52002739", "52002740", "52002741", "52002742", "52002743",
			"52002744", "52002745", "52002746", "52002747", "52002748", "52002749", "52002750", "52002751", "52002752", "52002753",
			"52002754", "52002755", "52002756", "52002757", "52002758", "52002759", "52002760", "52002761", "52002762", "52002763",
			"52002764", "52002765", "52002766", "52002767", "52002768", "52002769", "52002770", "5200277", "52002778", "52002779"},
		chocoFriends: []string{"51626494", "51626495", "51626496", "51626497", "51626498", "51626499", "51626500", "51626501", "51626502", "51626503",
			"51626504", "51626505", "51626506", "51626507", "51626508", "51626509", "51626510", "51626511", "51626512", "51626513",
			"51626514", "51626515", "51626516", "51626517", "51626518", "51626519", "51626520", "51626521", "51626522", "51626523",
			"51626524", "51626525", "51626526", "51626527", "51626528", "51626529", "51626530", "51626531", "51626532", "51626533"},
		univerStarBT21: []string{"52114110", "52114111", "52114112", "52114113", "52114114", "52114115", "52114116", "52114117", "52114118", "52114119",
			"52114120", "52114121", "52114122", "52114123", "52114124", "52114125", "52114126", "52114127", "52114128", "52114129",
			"52114130", "52114131", "52114132", "52114133", "52114134", "52114135", "52114136", "52114137", "52114138", "52114139",
			"52114140", "52114141", "52114142", "52114143", "52114144", "52114145", "52114146", "52114147", "52114148", "52114149"},
	}
)
