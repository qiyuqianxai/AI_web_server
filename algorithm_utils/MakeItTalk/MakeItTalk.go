package MakeItTalk

type config struct {
	UserAudioDir string `json:"user_audio_dir"`
	UserImgDir   string `json:"user_img_dir"`
	ResVideoDir  string `json:"res_video_dir"`
	MessageJson  string `json:"message_json"`
}
var config_file config

// send gan-nn param
type message struct {
	UserImg   string  `json:"user_img"`
	UserAudio string  `json:"user_audio"`
	AmpLipX   float32 `json:"amp_lip_x"`
	AmpLipY   float32 `json:"amp_lip_y"`
	AmpPos    float32 `json:"amp_pos"`
}

var Msg message

func Laod_config() config {
	config_file.UserImgDir = "static/algorithm/MakeItTalk/user_imgs"
	config_file.UserAudioDir = "static/algorithm/MakeItTalk/user_audios"
	config_file.ResVideoDir = "static/algorithm/MakeItTalk/res_videos"
	config_file.MessageJson = "algorithm_utils/MakeItTalk/message.json"
	return config_file
}