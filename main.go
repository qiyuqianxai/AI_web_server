
package main

import (
	"awesomeProject/algorithm_utils/stylegan"
	_ "awesomeProject/algorithm_utils/stylegan"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	//gin.SetMode(gin.DebugMode)
	r := gin.Default()

	//// 定义静态文件映射
	r.Static("/static","static")

	//r.LoadHTMLGlob("static/*.html")
	//r.LoadHTMLFiles(static_file+"/*.html")


	// define global map
	r.GET("/index", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "stylegan.html",nil)
		c.Redirect(http.StatusFound, "/static/stylegan.html")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound,"/index")
	})

	r.GET("/animal_gan", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/animal_gan.html")
	})


	// define stylegan map
	// 读取配置的路径
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
		time.Sleep(time.Duration(3)*time.Second);
		c.JSON(200,gin.H{})
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




