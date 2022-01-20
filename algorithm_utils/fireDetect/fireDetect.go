package fireDetect

type fireDetect_info struct {
	Res_img_dir             string `json:"res_img_dir"`
	User_img_dir            string `json:"user_img_dir"`
	User_video_dir          string `json:"user_video_dir"`
	Res_video_dir           string `json:"res_video_dir"`
	FireDetect_message_path string `json:"fireDetect_Message_Path"`

}
var fireDetect_config fireDetect_info

// send gan-nn param
type fireDetect_message struct {
	UserImg          string `json:"user_img"`
	//Model 			 string `json:"model"`
}
var FiredetectMsg fireDetect_message

//Laod_config
func Laod_config() fireDetect_info {
	fireDetect_config.Res_img_dir = "static/algorithm/fireDetect/result_imgs"
	fireDetect_config.User_img_dir = "static/algorithm/fireDetect/user_imgs"
	fireDetect_config.FireDetect_message_path = "algorithm_utils/fireDetect/message.json"
	fireDetect_config.User_video_dir = "static/algorithm/fireDetect/user_videos"
	fireDetect_config.Res_video_dir = "static/algorithm/fireDetect/res_videos"
	return fireDetect_config
}

