
package main

import (
	"awesomeProject/algorithm_utils/actionReg"
	"awesomeProject/algorithm_utils/realesrgan"
	"awesomeProject/algorithm_utils/stylegan"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type model_states struct {
	ActionReg  bool `json:"actionReg"`
	Realesrgan bool `json:"realesrgan"`
	Stylegan   bool `json:"stylegan"`
}
var model_stat model_states

func main() {
	//gin.SetMode(gin.DebugMode)
	r := gin.Default()

	//// 定义静态文件映射
	r.Static("/static","static")

	//r.LoadHTMLGlob("static/*.html")
	//r.LoadHTMLFiles(static_file+"/*.html")

	wait_time := 2;

	// ################################### define global map ############################################
	r.GET("/index", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "stylegan.html",nil)
		c.Redirect(http.StatusFound, "/static/index.html")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound,"/index")
	})

	r.GET("/stylegan", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/stylegan.html")
	})

	r.GET("/realesran", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/realesran.html")
	})

	r.GET("/actionReg", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/actionReg.html")
	})

	r.POST("/update_models_states", func(c *gin.Context) {
		err := c.BindJSON(&model_stat)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)
		model_states_json_path := "algorithm_utils/model_states.json"
		// 写json文件
		_, err = os.Stat(model_states_json_path)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(model_states_json_path,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(model_states_json_path)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(model_stat)
		if err != nil {
			log.Println(err)
		}
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(wait_time)*time.Second);
		c.JSON(200,gin.H{})
	})
	// ################################## define stylegan map #########################################
	// 读取配置的路径5

	stylegan_config := stylegan.Laod_config()

	// 获取基本信息
	r.GET("/stylegan/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(stylegan_config.User_img_dir)
		var user_imgs []string
		if err == nil {
			f_list,err := ioutil.ReadDir(stylegan_config.User_img_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(stylegan_config.User_img_dir, os.ModePerm)
			if err != nil {
				return 
			}
		}
		c.JSON(200,gin.H{
			"img_list":user_imgs,
		})
	})

	// upload image
	r.POST("/stylegan/upload_file/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(stylegan_config.User_img_dir,img_name)
			_, err = os.Stat(save_pth)
			if !os.IsNotExist(err) {
				err := os.Remove(save_pth)
				if err != nil {
					return
				}
			}

			err = c.SaveUploadedFile(upfile, save_pth)
			if err != nil {
				return
			}
			c.JSON(200,gin.H{})
		} else{
			c.JSON(304,gin.H{})
		}
	})

	// generate image
	r.POST("/stylegan/generate_img/", func(c *gin.Context) {
		err := c.BindJSON(&stylegan.Style_gan_msg)
		if err != nil {
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(stylegan_config.Style_gan_message_path)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(stylegan_config.Style_gan_message_path,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(stylegan_config.Style_gan_message_path)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(stylegan.Style_gan_msg)
		if err != nil {
			log.Println(err)
		}
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(wait_time)*time.Second);
		c.JSON(200,gin.H{})
	})


	// ################################### define realesran map ##########################################

	realesregan_config := realesrgan.Laod_config()

	r.GET("/realesran/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(realesregan_config.User_img_dir)
		var user_imgs []string
		if err == nil {
			f_list,err := ioutil.ReadDir(realesregan_config.User_img_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(realesregan_config.User_img_dir, os.ModePerm)
			if err != nil {
				return
			}
		}
		c.JSON(200,gin.H{
			"img_list":user_imgs,
		})
	})

	// upload image
	r.POST("/realesran/upload_file/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(realesregan_config.User_img_dir,img_name)
			_, err = os.Stat(save_pth)
			if !os.IsNotExist(err) {
				err := os.Remove(save_pth)
				if err != nil {
					return
				}
			}

			err = c.SaveUploadedFile(upfile, save_pth)
			if err != nil {
				return
			}
			c.JSON(200,gin.H{})
		} else{
			c.JSON(304,gin.H{})
		}
	})

	// generate image
	r.POST("/realesran/generate_img/", func(c *gin.Context) {
		err := c.BindJSON(&realesrgan.Realesrgan_msg)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(realesregan_config.Realesrgan_message_path)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(realesregan_config.Realesrgan_message_path,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(realesregan_config.Realesrgan_message_path)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(realesrgan.Realesrgan_msg)
		if err != nil {
			log.Println(err)
		}
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(wait_time)*time.Second);
		c.JSON(200,gin.H{})
	})


	// ################################### define action_reg map ###########################################

	actionReg_config := actionReg.Laod_config()
	msg := actionReg.Msg
	// 获取基本信息
	r.GET("/actionReg/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(actionReg_config.Videos_pth)
		var video_names []string
		video_names = append(video_names,"capture")
		if err == nil {
			f_list,err := ioutil.ReadDir(actionReg_config.Videos_pth)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					video_names = append(video_names, f.Name())
				}
			}
		}
		if os.IsNotExist(err) {
			err := os.Mkdir(actionReg_config.Videos_pth, os.ModePerm)
			if err != nil {
				return
			}
		}
		c.JSON(200,gin.H{
			"video_list":video_names,
		})
	})

	// 更新本地文件
	r.POST("/actionReg/update_play_info", func(c *gin.Context) {
		msg.Play_video = true
		err := c.BindJSON(&msg)
		if err != nil {
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(actionReg_config.Message_json)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(actionReg_config.Message_json,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				println(err)
			}
		}else {
			file, err = os.Create(actionReg_config.Message_json)
			if err != nil {
				println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(msg)
		if err != nil {
			println(err)
		}
		err = file.Close()
		if err != nil {
			println(err)
		}

		// 获取max_id
		current_video := msg.Video_name
		max_id := -1
		_, err = os.Stat(path.Join(actionReg_config.Predict_dir, current_video))
		if err != nil {
			println(err)
		}
		dir, err := ioutil.ReadDir(path.Join(actionReg_config.Predict_dir,current_video))
		if err != nil {
			println(err)
		}

		for _,f := range dir{
			if strings.HasSuffix(f.Name(),"_predict.json"){
				max_id++
			}
		}
		if max_id - msg.Current_id < 10{
			time.Sleep(5);
		}
		c.JSON(200,gin.H{"max_id":max_id})
	})

	// 播放结束
	r.GET("/actionReg/play_end/", func(c *gin.Context) {
		msg.Play_video = false
		file, err := os.OpenFile(actionReg_config.Message_json,os.O_WRONLY|os.O_TRUNC,0666)
		if err != nil {
			println(err)
		}

		enc := json.NewEncoder(file)
		err = enc.Encode(msg)
		if err != nil {
			println(err)
		}
		err = file.Close()
		if err != nil {
			println(err)
		}
		c.JSON(200,gin.H{})
	})

	// 上传文件
	r.POST("/actionReg/upload_file/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		video_name := upfile.Filename
		log.Println(video_name)
		if strings.HasSuffix(video_name,".mp4") || strings.HasSuffix(video_name,".avi") || strings.HasSuffix(video_name,".flv"){
			save_pth := path.Join(actionReg_config.Videos_pth,video_name)
			_, err = os.Stat(save_pth)
			if !os.IsNotExist(err) {
				err := os.Remove(save_pth)
				if err != nil {
					return
				}
			}

			err = c.SaveUploadedFile(upfile, save_pth)
			if err != nil {
				return
			}
			c.JSON(200,gin.H{})
		} else{
			c.JSON(304,gin.H{})
		}
	})

	// 获取数据分析的数据
	r.POST("/actionReg/get_analysis_data/", func(c *gin.Context) {
		data_json := make(map[string]interface{}) //注意该结构接受的内容
		err := c.BindJSON(&data_json)
		if err != nil {
			return
		}
		//log.Printf("%v",&data_json)
		img_id := int(data_json["id"].(float64))
		current_video := data_json["video"].(string)
		json_pth := path.Join(actionReg_config.Predict_dir,current_video,strconv.Itoa(img_id)+".json")
		_, err = os.Stat(json_pth)
		if os.IsNotExist(err){
			// c.JSON(200,gin.H{})
			log.Println(err)
			return
		}

		var json_file *os.File
		var data_analysis actionReg.Data_analysis_res

		json_file, err = os.Open(json_pth)
		if err != nil {
			log.Println("文件打开错误",err)
			return
		}
		defer func(json_file *os.File) {
			err := json_file.Close()
			if err != nil {
				return
			}
		}(json_file)
		dco := json.NewDecoder(json_file)
		err = dco.Decode(&data_analysis)
		if err != nil {
			log.Println("解析错误",err)
			return
		}
		//log.Println("data_res:",data_analysis.Action_analysis,data_analysis.Expression_analysis,data_analysis.Attention_analysis,data_analysis.Person_num)
		c.JSON(http.StatusOK,data_analysis)
	})


	// define stargan map
	//// blend images
	//r.POST("/stargan/convert_img/", func(c *gin.Context) {
	//	err := c.BindJSON(&star_gan_msg)
	//	if err != nil {
	//		return
	//	}
	//	//log.Printf("%v",&msg)
	//
	//	// 写json文件
	//	_, err = os.Stat(star_gan_message)
	//	var file *os.File
	//	if err == nil {
	//		file, err = os.OpenFile(star_gan_message,os.O_WRONLY|os.O_TRUNC,0666)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//	}else {
	//		file, err = os.Create(star_gan_message)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//	}
	//	enc := json.NewEncoder(file)
	//	err = enc.Encode(star_gan_msg)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	err = file.Close()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	time.Sleep(time.Duration(3)*time.Second);
	//	//for i:=0;i<100;i++{
	//	//	blend_img := "/static/blend.jpg"
	//	//	_, err :=os.Stat(blend_img)
	//	//	if err == nil {
	//	//		log.Println("generate dest image success!")
	//	//		break
	//	//	}
	//	//	time.Sleep(time.Duration(1)*time.Second);
	//	//}
	//	c.JSON(200,gin.H{})
	//})

	err := r.Run(":43476")
	if err != nil {
		return 
	}
}




