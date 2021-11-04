package deoldify

type config struct {
	Res_img_dir            string `json:"res_img_dir"`
	User_img_dir           string `json:"user_img_dir"`
	Message_json 		   string `json:"message_json"`
}
var config_file config

// send gan-nn param
type message struct {
	Artistic     string `json:"artistic"`
	UserImg      string `json:"user_img"`
	RenderFactor int    `json:"render_factor"`
}

var Msg message

func Laod_config() config {
	config_file.User_img_dir = "static/algorithm/deoldify/user_imgs"
	config_file.Res_img_dir = "static/algorithm/deoldify/res_imgs"
	config_file.Message_json = "algorithm_utils/deoldify/message.json"
	return config_file
}