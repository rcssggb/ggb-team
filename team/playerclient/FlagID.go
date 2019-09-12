package playerclient

// FlagID is the type representing each marker flag in the field
type FlagID byte

const (
	CenterFlag FlagID = iota + 0

	CenterTopFlag

	CenterBotFlag

	LeftGoalFlag

	LeftGoalTopFlag

	LeftGoalBotFlag

	RightGoalFlag

	RightGoalTopFlag

	RightGoalBotFlag

	LeftPenaltyCenterFlag

	LeftPenaltyTopFlag

	LeftPenaltyBotFlag

	RightPenaltyCenterFlag

	RightPenaltyTopFlag

	RightPenaltyBotFlag

	LeftTopFlag

	LeftBotFlag

	RightTopFlag

	RightBotFlag

	TopLeft50Flag

	TopLeft40Flag

	TopLeft30Flag

	TopLeft20Flag

	TopLeft10Flag

	Top0Flag

	TopRight10Flag

	TopRight20Flag

	TopRight30Flag

	TopRight40Flag

	TopRight50Flag

	BotLeft50Flag

	BotLeft40Flag

	BotLeft30Flag

	BotLeft20Flag

	BotLeft10Flag

	Bot0Flag

	BotRight10Flag

	BotRight20Flag

	BotRight30Flag

	BotRight40Flag

	BotRight50Flag

	LeftTop30Flag

	LeftTop20Flag

	LeftTop10Flag

	Left0Flag

	LeftBot10Flag

	LeftBot20Flag

	LeftBot30Flag

	RightTop30Flag

	RightTop20Flag

	RightTop10Flag

	Right0Flag

	RightBot10Flag

	RightBot20Flag

	RightBot30Flag
)
