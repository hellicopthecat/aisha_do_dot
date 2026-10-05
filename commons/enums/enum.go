package enums

type Social string

const (
	SocialGoogle Social = "GOOGLE"
	SocialKakao  Social = "KAKAO"
)

type OneDepth string

const (
	OneDepthAll      OneDepth = "ALL"
	OneDepthNormal   OneDepth = "NORMAL"
	OneDepthDaily    OneDepth = "DAILY"
	OneDepthWorkout  OneDepth = "WORKOUT"
	OneDepthBusiness OneDepth = "BUSINESS"
	OneDepthTravel   OneDepth = "TRAVEL"
	OneDepthFood     OneDepth = "FOOD"
)

type Access string

const (
	AccessAccess Access = "ACCESS"
	AccessDenied Access = "DENIED"
)

type AccessRole string

const (
	AccessOnlyRead AccessRole = "READ"
	AccessEditable AccessRole = "EDITABLE"
)
