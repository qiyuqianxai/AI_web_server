package stargan

type config struct {
	Ref_img_dir            string `json:"ref_img_dir"`
	User_img_dir           string `json:"user_img_dir"`
	Message_json 		   string `json:"message_json"`
}
var config_file config

// send gan-nn param
type message struct {
	User_img   string `json:"user_img"`
	Ref_img    string `json:"ref_img"`
}

var Msg message

func Laod_config() config {
	config_file.User_img_dir = "static/algorithm/stargan/user_imgs"
	config_file.Ref_img_dir = "static/algorithm/stargan/user_imgs"
	config_file.Message_json = "algorithm_utils/stargan/message.json"
	return config_file
}