package colorImage

type config struct {
	Ref_img_dir  string `json:"ref_img_dir"`
	User_img_dir string `json:"user_img_dir"`
	Res_img_dir  string `json:"res_img_dir"`
	Message_json string `json:"message_json"`
}
var config_file config

// send gan-nn param
type message struct {
	UserImg      string `json:"user_img"`
	RefImg       string `json:"ref_img"`
	RenderFactor int    `json:"render_factor"`
	Artistic     bool   `json:"artistic"`
}

var Msg message

func Laod_config() config {
	config_file.User_img_dir = "static/algorithm/colorImage/user_imgs"
	config_file.Ref_img_dir = "static/algorithm/colorImage/ref_imgs"
	config_file.Res_img_dir = "static/algorithm/colorImage/res_imgs"
	config_file.Message_json = "algorithm_utils/colorImage/message.json"
	return config_file
}