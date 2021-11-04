package wav2lip

type config struct {
	UserAudioDir string `json:"user_audio_dir"`
	UserImgDir   string `json:"user_img_dir"`
	ResVideoDir  string `json:"res_video_dir"`
	MessageJson  string `json:"message_json"`
}
var config_file config

// send gan-nn param
type message struct {
	UserImg   string `json:"user_img"`
	UserAudio string `json:"user_audio"`
}

var Msg message

func Laod_config() config {
	config_file.UserImgDir = "static/algorithm/wav2lip/user_imgs"
	config_file.UserAudioDir = "static/algorithm/wav2lip/user_audios"
	config_file.ResVideoDir = "static/algorithm/wav2lip/res_videos"
	config_file.MessageJson = "algorithm_utils/wav2lip/message.json"
	return config_file
}