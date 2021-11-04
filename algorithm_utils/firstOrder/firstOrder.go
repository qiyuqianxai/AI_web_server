package firstOrder

type config struct {
	UserVideoDir string `json:"user_Video_Dir"`
	UserImgDir   string `json:"user_img_dir"`
	ResImgDir    string `json:"res_img_dir"`
	MessageJson  string `json:"message_json"`
}
var config_file config

// send gan-nn param
type message struct {
	UserImg    string `json:"user_img"`
	UserVideo  string `json:"user_video"`
	Relative   bool   `json:"relative"`
	AdaptScale bool   `json:"adapt_scale"`
}

var Msg message

func Laod_config() config {
	config_file.UserImgDir = "static/algorithm/firstOrder/user_imgs"
	config_file.UserVideoDir = "static/algorithm/firstOrder/user_videos"
	config_file.ResImgDir = "static/algorithm/firstOrder/res_videos"
	config_file.MessageJson = "algorithm_utils/firstOrder/message.json"
	return config_file
}