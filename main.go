
package main

import (
	"awesomeProject/algorithm_utils/MakeItTalk"
	"awesomeProject/algorithm_utils/actionReg"
	"awesomeProject/algorithm_utils/colorImage"
	"awesomeProject/algorithm_utils/fireDetect"
	"awesomeProject/algorithm_utils/firstOrder"
	"awesomeProject/algorithm_utils/realesrgan"
	"awesomeProject/algorithm_utils/stargan"
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
	Stargan    bool `json:"stargan"`
	Debvc      bool `json:"debvc"`
	Deoldify   bool `json:"deoldify"`
	FirstOrder bool `json:"firstOrder"`
	MakeItTalk   bool `json:"MakeItTalk"`
	FireDetect bool `json:"fireDetect"`
}
var model_stat model_states

func main() {
	//gin.SetMode(gin.DebugMode)
	r := gin.Default()

	//// 定义静态文件映射
	r.Static("/static","static")

	//r.LoadHTMLGlob("static/*.html")
	//r.LoadHTMLFiles(static_file+"/*.html")

	wait_time := 5;

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

	r.GET("/realesrgan", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/realesrgan.html")
	})

	r.GET("/actionReg", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/actionReg.html")
	})

	r.GET("/stargan", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/stargan.html")
	})

	r.GET("/colorImage", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/colorImg.html")
	})

	r.GET("/activateImg", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/activateImg.html")
	})

	r.GET("/fireDetect", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/fireDetect.html")
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
		// time.Sleep(time.Duration(wait_time)*time.Second);
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


	// ################################### define realesrgan map ##########################################

	realesregan_config := realesrgan.Laod_config()

	r.GET("/realesrgan/get_base_info", func(c *gin.Context) {
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

		var user_videos []string
		if err == nil {
			f_list,err := ioutil.ReadDir(realesregan_config.User_video_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_videos = append(user_videos, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(realesregan_config.User_video_dir, os.ModePerm)
			if err != nil {
				return
			}
		}

		c.JSON(200,gin.H{
			"img_list":user_imgs,
			"video_list":user_videos,
		})
	})

	// upload image
	r.POST("/realesrgan/upload_file/", func(c *gin.Context) {
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

	// upload videos
	r.POST("/realesrgan/upload_video/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		video_name := upfile.Filename
		if strings.HasSuffix(video_name,".mp4") {
			save_pth := path.Join(realesregan_config.User_video_dir,video_name)
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
	r.POST("/realesrgan/generate_img/", func(c *gin.Context) {
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
	actionReg_msg := actionReg.Msg
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
		actionReg_msg.Play_video = true
		err := c.BindJSON(&actionReg_msg)
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
		err = enc.Encode(actionReg_msg)
		if err != nil {
			println(err)
		}
		err = file.Close()
		if err != nil {
			println(err)
		}

		// 获取max_id
		current_video := actionReg_msg.Video_name
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
		if max_id - actionReg_msg.Current_id < 10{
			time.Sleep(5);
		}
		c.JSON(200,gin.H{"max_id":max_id})
	})

	// 播放结束
	r.GET("/actionReg/play_end/", func(c *gin.Context) {
		actionReg_msg.Play_video = false
		file, err := os.OpenFile(actionReg_config.Message_json,os.O_WRONLY|os.O_TRUNC,0666)
		if err != nil {
			println(err)
		}

		enc := json.NewEncoder(file)
		err = enc.Encode(actionReg_msg)
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


	// ################################### define stargan map ###########################################

	stargan_config := stargan.Laod_config();
	stargan_msg := stargan.Msg;
	// 获取基本信息
	r.GET("/stargan/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(stargan_config.User_img_dir)
		var user_imgs []string
		if err == nil {
			f_list,err := ioutil.ReadDir(stargan_config.User_img_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(stargan_config.User_img_dir, os.ModePerm)
			if err != nil {
				return
			}
		}
		c.JSON(200,gin.H{
			"img_list":user_imgs,
		})
	})

	// upload image
	r.POST("/stargan/upload_file/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(stargan_config.User_img_dir,img_name)
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

	// blend images
	r.POST("/stargan/blend_imgs/", func(c *gin.Context) {
		err := c.BindJSON(&stargan_msg)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(stargan_config.Message_json)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(stargan_config.Message_json,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(stargan_config.Message_json)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(stargan_msg)
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

	// ################################# define colorImage map ##############################################
	colorImage_config := colorImage.Laod_config();
	colorImage_msg := colorImage.Msg;
	// 获取基本信息
	r.GET("/colorImage/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(colorImage_config.User_img_dir)
		var user_imgs []string
		var ref_imgs []string

		if err == nil {
			f_list,err := ioutil.ReadDir(colorImage_config.User_img_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(colorImage_config.User_img_dir, os.ModePerm)
			if err != nil {
				return
			}
		}

		if err == nil {
			f_list,err := ioutil.ReadDir(colorImage_config.Ref_img_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					ref_imgs = append(ref_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(colorImage_config.Ref_img_dir, os.ModePerm)
			if err != nil {
				return
			}
		}

		c.JSON(200,gin.H{
			"ref_list":ref_imgs,
			"src_list":user_imgs,
		})
	})

	// upload src image
	r.POST("/colorImage/upload_src", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(colorImage_config.User_img_dir,img_name)
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

	// upload ref image
	r.POST("/colorImage/upload_ref", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(colorImage_config.Ref_img_dir,img_name)
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

	// blend images
	r.POST("/colorImage/generate_img/", func(c *gin.Context) {
		err := c.BindJSON(&colorImage_msg)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(colorImage_config.Message_json)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(colorImage_config.Message_json,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(colorImage_config.Message_json)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(colorImage_msg)
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


	// ######################################## define firstOrder ################################################

	firstOrder_config := firstOrder.Laod_config();
	firstOrder_msg := firstOrder.Msg;
	// 获取基本信息
	r.GET("/firstOrder/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(firstOrder_config.UserImgDir)
		var user_imgs []string
		if err == nil {
			f_list,err := ioutil.ReadDir(firstOrder_config.UserImgDir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(firstOrder_config.UserImgDir, os.ModePerm)
			if err != nil {
				return
			}
		}

		var user_videos []string
		if err == nil {
			f_list,err := ioutil.ReadDir(firstOrder_config.UserVideoDir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_videos = append(user_videos, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(firstOrder_config.UserVideoDir, os.ModePerm)
			if err != nil {
				return
			}
		}

		c.JSON(200,gin.H{
			"img_list":user_imgs,
			"video_list":user_videos,
		})
	})

	// upload src image
	r.POST("/firstOrder/upload_img/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(firstOrder_config.UserImgDir,img_name)
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

	// upload ref image
	r.POST("/firstOrder/upload_video/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		video_name := upfile.Filename
		log.Println(video_name)
		if strings.HasSuffix(video_name,".mp4") {
			save_pth := path.Join(firstOrder_config.UserVideoDir,video_name)
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

	// blend images
	r.POST("/firstOrder/generate_result/", func(c *gin.Context) {
		err := c.BindJSON(&firstOrder_msg)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(firstOrder_config.MessageJson)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(firstOrder_config.MessageJson,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(firstOrder_config.MessageJson)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(firstOrder_msg)
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


	// ############################################## define MakeItTalk ##############################################

	MakeItTalk_config := MakeItTalk.Laod_config();
	MakeItTalk_msg := MakeItTalk.Msg;
	// 获取基本信息
	r.GET("/MakeItTalk/get_base_info/", func(c *gin.Context) {
		_, err := os.Stat(MakeItTalk_config.UserImgDir)
		var user_imgs []string
		if err == nil {
			f_list,err := ioutil.ReadDir(MakeItTalk_config.UserImgDir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(MakeItTalk_config.UserImgDir, os.ModePerm)
			if err != nil {
				return
			}
		}

		var user_audios []string
		if err == nil {
			f_list,err := ioutil.ReadDir(MakeItTalk_config.UserAudioDir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_audios = append(user_audios, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(MakeItTalk_config.UserAudioDir, os.ModePerm)
			if err != nil {
				return
			}
		}

		c.JSON(200,gin.H{
			"img_list":user_imgs,
			"audio_list":user_audios,
		})
	})

	// upload src image
	r.POST("/MakeItTalk/upload_img/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(MakeItTalk_config.UserImgDir,img_name)
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

	// upload
	r.POST("/MakeItTalk/upload_audio/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".mp3") {
			save_pth := path.Join(MakeItTalk_config.UserAudioDir,img_name)
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

	// blend images
	r.POST("/MakeItTalk/generate_result/", func(c *gin.Context) {
		err := c.BindJSON(&MakeItTalk_msg)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(MakeItTalk_config.MessageJson)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(MakeItTalk_config.MessageJson,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(MakeItTalk_config.MessageJson)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(MakeItTalk_msg)
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

	// ############################################## define fireDetect ##############################################
	fireDetect_config := fireDetect.Laod_config()

	r.GET("/fireDetect/get_base_info", func(c *gin.Context) {
		_, err := os.Stat(fireDetect_config.User_img_dir)
		var user_imgs []string
		if err == nil {
			f_list,err := ioutil.ReadDir(fireDetect_config.User_img_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_imgs = append(user_imgs, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(fireDetect_config.User_img_dir, os.ModePerm)
			if err != nil {
				return
			}
		}

		var user_videos []string
		if err == nil {
			f_list,err := ioutil.ReadDir(fireDetect_config.User_video_dir)
			if err!=nil{
				log.Fatal(err)
			} else {
				for _,f := range f_list{
					user_videos = append(user_videos, f.Name())
				}
			}
		}

		if os.IsNotExist(err) {
			err := os.Mkdir(fireDetect_config.User_video_dir, os.ModePerm)
			if err != nil {
				return
			}
		}

		c.JSON(200,gin.H{
			"img_list":user_imgs,
			"video_list":user_videos,
		})
	})

	// upload image
	r.POST("/fireDetect/upload_file/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		img_name := upfile.Filename
		log.Println(img_name)
		if strings.HasSuffix(img_name,".jpeg") || strings.HasSuffix(img_name,".jpg") || strings.HasSuffix(img_name,".png"){
			save_pth := path.Join(fireDetect_config.User_img_dir,img_name)
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

	// upload videos
	r.POST("/fireDetect/upload_video/", func(c *gin.Context) {
		upfile, err := c.FormFile("file")
		if err != nil {
			return
		}
		video_name := upfile.Filename
		if strings.HasSuffix(video_name,".mp4") || strings.HasSuffix(video_name,".flv") || strings.HasSuffix(video_name,".avi"){
			save_pth := path.Join(fireDetect_config.User_video_dir,video_name)
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
	r.POST("/fireDetect/generate_img/", func(c *gin.Context) {
		err := c.BindJSON(&fireDetect.FiredetectMsg)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Printf("%v",&msg)

		// 写json文件
		_, err = os.Stat(fireDetect_config.FireDetect_message_path)
		var file *os.File
		if err == nil {
			file, err = os.OpenFile(fireDetect_config.FireDetect_message_path,os.O_WRONLY|os.O_TRUNC,0666)
			if err != nil {
				log.Println(err)
			}
		}else {
			file, err = os.Create(fireDetect_config.FireDetect_message_path)
			if err != nil {
				log.Println(err)
			}
		}
		enc := json.NewEncoder(file)
		err = enc.Encode(fireDetect.FiredetectMsg)
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


	err := r.Run(":43476")
	if err != nil {
		return 
	}
}




