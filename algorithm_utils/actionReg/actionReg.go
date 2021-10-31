package actionReg

type config struct {
	Predict_dir  string `json:"predict_dir"`
	Videos_pth   string `json:"videos_pth"`
	Message_json string `json:"message_json"`
}

var config_file config

type message struct {
	Current_id int `json:"id"`
	Play_video 		bool
	Video_name      string `json:"video_name"`
	// 需要的ai功能
	Req_face_detect    bool `json:"req_face_detect"`
	Req_body_detect    bool `json:"req_body_detect"`
	Req_body_keypoints bool `json:"req_body_keypoints"`
	Req_face_keypoints bool `json:"req_face_keypoints"`
	Req_expression_reg bool `json:"req_expression_reg"`
	Req_action_reg     bool `json:"req_action_reg"`
	Person_track       bool `json:"person_track"`
	Req_attention_reg  bool `json:"req_attention_reg"`
}

var Msg message


type Data_analysis_res struct {
	Action_analysis     [][]string `json:"action_analysis"`
	Expression_analysis [][]string `json:"expression_analysis"`
	Attention_analysis  [][]string `json:"attention_analysis"`
	Person_num          int        `json:"person_num"`
}


func Laod_config() config {
	config_file.Videos_pth = "static/algorithm/actionReg/videos"
	config_file.Predict_dir = "static/algorithm/actionReg/predictions"
	config_file.Message_json = "algorithm_utils/actionReg/message.json"
	return config_file
}
