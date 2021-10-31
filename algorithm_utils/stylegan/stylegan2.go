package stylegan

type stylegan_info struct {
	Fake_img_dir      string `json:"fake_img_dir"`
	User_img_dir      string `json:"user_img_dir"`
	Style_gan_message_path string `json:"style_gan_message"`
}
var Stylegan_config stylegan_info

// send gan-nn param
type style_gan_message struct {
	UserImg          string `json:"user_img"`
	Seed             string `json:"seed"`
	Age             string `json:"age"`
	AngleHorizontal string `json:"angle_horizontal"`
	AnglePitch      string `json:"angle_pitch"`
	Beauty          string `json:"beauty"`
	EmotionAngry    string `json:"emotion_angry"`
	EmotionDisgust  string `json:"emotion_disgust"`
	EmotionEasy     string `json:"emotion_easy"`
	EmotionFear     string `json:"emotion_fear"`
	EmotionHappy    string `json:"emotion_happy"`
	EmotionSad      string `json:"emotion_sad"`
	EmotionSurprise string `json:"emotion_surprise"`
	EyesOpen        string `json:"eyes_open"`
	FaceShape       string `json:"face_shape"`
	Gender          string `json:"gender"`
	Glasses         string `json:"glasses"`
	Height          string `json:"height"`
	RaceBlack       string `json:"race_black"`
	RaceWhite       string `json:"race_white"`
	RaceYellow      string `json:"race_yellow"`
	Smile           string `json:"smile"`
	Width           string `json:"width"`
	Cartoon 	    string `json:"cartoon"`
}
var Style_gan_msg style_gan_message

//Laod_config
func Laod_config() stylegan_info {
	Stylegan_config.Fake_img_dir = "static/algorithm/stylegan/fake_imgs"
	Stylegan_config.User_img_dir = "static/algorithm/stylegan/user_imgs"
	Stylegan_config.Style_gan_message_path = "algorithm_utils/stylegan/message.json"
	return Stylegan_config
}
