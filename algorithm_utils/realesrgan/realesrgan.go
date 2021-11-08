package realesrgan

type realesrgan_info struct {
	Result_img_dir      string `json:"result_img_dir"`
	User_img_dir      	string `json:"user_img_dir"`
	User_video_dir 		string `json:"user_video_dir"`
	Res_video_dir 		string `json:"res_video_dir"`
	Realesrgan_message_path string `json:"realesrgan_Message_Path"`

}
var Realesrgan_config realesrgan_info

// send gan-nn param
type realesrgan_message struct {
	UserImg          string `json:"user_img"`
	Face_enhance     bool `json:"face_enhance"`
	Model 			 string `json:"model"`
}
var Realesrgan_msg realesrgan_message

//Laod_config
func Laod_config() realesrgan_info {
	Realesrgan_config.Result_img_dir = "static/algorithm/realesrgan/result_imgs"
	Realesrgan_config.User_img_dir = "static/algorithm/realesrgan/user_imgs"
	Realesrgan_config.Realesrgan_message_path = "algorithm_utils/realesrgan/message.json"
	Realesrgan_config.User_video_dir = "static/algorithm/realesrgan/user_videos"
	Realesrgan_config.Res_video_dir = "static/algorithm/realesrgan/res_videos"
	return Realesrgan_config
}
