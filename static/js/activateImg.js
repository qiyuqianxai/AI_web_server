// current src image
src_image = ""
src_image_2 = ""
user_audio = ""
user_video = ""

params = {
    "user_img": "",
    "user_audio": "",
    "amp_lip_x": 1.0,
    "amp_lip_y": 1.0,
    "amp_pos": 0
}
params2 = {
    "user_img": "trumps.png",
    "user_video": "zhiming.mp4",
    "relative": true,
    "adapt_scale": true
}

// ref = user
user_imgs_dir = "algorithm/MakeItTalk/user_imgs/" //第一张图

user_imgs_dir_2 = "algorithm/firstOrder/user_imgs/" //第二张图

user_radio_dir = "algorithm/MakeItTalk/user_audios/" //audio路径

user_video_dir = "algorithm/firstOrder/user_videos/" //video路径

result_video_dir = "algorithm/MakeItTalk/res_videos/" //第一个结果

result_video_dir_2 = "algorithm/firstOrder/res_videos/" //第二个结果

$(function() {

    // window.onbeforeunload = boot_model("stargan");
    //
    get_base_info_1();
    //
    get_base_info_2();
    // 加载按键信息
    set_click_response();

});
//限制输入
function clearNoNum(obj) {
    obj.value = obj.value.replace(/[^\d.]/g, ""); //清除“数字”和“.”以外的字符
    obj.value = obj.value.replace(/^\./g, ""); //验证第一个字符是数字
    obj.value = obj.value.replace(".", "$#$").replace(/\./g, "").replace("$#$", ".");
    obj.value = obj.value.replace(/^(\-)*(\d+)\.(\d\d).*$/, '$1$2.$3');
}

// load src_imgs list
function get_base_info_1() {
    $.ajax({
        url: "/MakeItTalk/get_base_info/",
        contentType: "application/json; charset=utf-8",
        type: "GET",
        cache: false,
        success: function(data) {
            //每次加载时重置一些参数

            var img_list = data['img_list']; //图片列表1
            var audio_list = data['audio_list'] //音频列表
            console.log(img_list)
            $("#user_imgs_1").empty();
            $("#user_imgs_1").append("<option value=''></option>");
            $.each(img_list, function(i, img_name) {
                $("#user_imgs_1").append("<option value=" + img_name + ">" + img_name + "</option>");
            });

            $('#user_imgs_1').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_img = $(this).find("option:selected").val();
                    if (selected_img !== src_image) {
                        var img = new Image()
                            // 改变图片的src
                        img.src = user_imgs_dir + selected_img
                            // 加载完成执行
                        img.onload = function() {
                            $('#user_img_1').attr("src", user_imgs_dir + selected_img);
                            var windowW = $(window).width() * 0.45; //获取当前窗口宽度
                            var windowH = $(window).height() * 0.51; //获取当前窗口高度
                            var realWidth = img.width; //获取图片真实宽度
                            var realHeight = img.height; //获取图片真实高度
                            var scale = Math.max(realWidth / windowW, realHeight / windowH); //缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                            console.log(realWidth, realHeight, windowW, windowH, scale)
                            $('#user_img_1').css({ "width": realWidth / scale, "height": realHeight / scale });
                            src_image = selected_img;
                        }

                    }
                    // console.log(current_video);
                }
            });


            $('#user_audios').empty();
            $("#user_audios").append("<option value=''></option>");
            $.each(audio_list, function(i, radio_name) {
                $("#user_audios").append("<option value=" + radio_name + ">" + radio_name + "</option>");
            });

            $("#user_audios").on('change', function(e) {
                if (e.originalEvent) {
                    let selected_audio = $(this).find("option:selected").val();
                    if (selected_audio !== user_audio) {
                        var radio_url = user_radio_dir + selected_audio;
                        var audio_div = document.getElementById("ref_audio");
                        var embed = audio_div.getElementsByTagName('embed')[0]
                        var hasembed = embed ? true : false;
                        if (hasembed) {
                            //清空父元素下的所有内容，保证始终都是有一个音频链接
                            audio_div.innerHTML = '';
                        }
                        var audio = document.createElement('embed');
                        audio.src = radio_url;
                        audio_div.appendChild(audio);
                        user_audio = selected_audio;
                    }
                }
                // console.log(current_video);
            });

        },
        error: function(data) {
            alert("数据加载出错，请联系管理员！");
            // top.location.reload();
        }
    });
}

function get_base_info_2() {
    $.ajax({
        url: "/firstOrder/get_base_info/",
        contentType: "application/json; charset=utf-8",
        type: "GET",
        cache: false,
        success: function(data) {
            //每次加载时重置一些参数

            var img_list = data['img_list']; //图片列表1
            var video_list = data['video_list'] //音频列表

            $('#user_imgs_2').empty();
            $("#user_imgs_2").append("<option value=''></option>");
            $.each(img_list, function(i, img_name) {
                $("#user_imgs_2").append("<option value=" + img_name + ">" + img_name + "</option>");
            });

            $('#user_imgs_2').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_img = $(this).find("option:selected").val();
                    if (selected_img !== src_image_2) {
                        var img = new Image()
                        // 改变图片的src
                        img.src = user_imgs_dir_2 + selected_img
                        // 加载完成执行
                        img.onload = function() {
                            $('#user_img_2').attr("src", user_imgs_dir_2 + selected_img);
                            var windowW = $(window).width() * 0.45; //获取当前窗口宽度
                            var windowH = $(window).height() * 0.51; //获取当前窗口高度
                            var realWidth = img.width; //获取图片真实宽度
                            var realHeight = img.height; //获取图片真实高度
                            var scale = Math.max(realWidth / windowW, realHeight / windowH); //缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                            console.log(realWidth, realHeight, windowW, windowH, scale)
                            $('#user_img_2').css({ "width": realWidth / scale, "height": realHeight / scale });
                            src_image_2 = selected_img;
                        }

                    }
                    // console.log(current_video);
                }
            });

            $('#user_videos').empty();
            $("#user_videos").append("<option value=''></option>");
            $.each(video_list, function(i, video_name) {
                $("#user_videos").append("<option value=" + video_name + ">" + video_name + "</option>");
            });

            $('#user_videos').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_video = $(this).find("option:selected").val();
                    if (selected_video !== user_video) {
                        var video_url = user_video_dir + selected_video;
                        var video_div = document.getElementById("ref_video");
                        var embed = video_div.getElementsByTagName('embed')[0];
                        var hasembed = embed ? true : false;
                        if (hasembed) {
                            //清空父元素下的所有内容，保证始终都是有一个音频链接
                            video_div.innerHTML = '';
                        }
                        var video = document.createElement('embed');
                        video.src = video_url;
                        video_div.appendChild(video);
                        user_video = selected_video;
                    }
                }
                // console.log(current_video);
            });

        },
        error: function(data) {
            alert("数据加载出错，请联系管理员！");
            // top.location.reload();
        }
    });
}

// 设置各个功能响应
function set_click_response() {
    // ai功能响应
    $('#blend').blur().on("click", function() {
        boot_model(["MakeItTalk"]);
        generate_1();
    });
    $('#blend_2').blur().on("click", function() {
        boot_model(["firstOrder"]);
        generate_2();
    });

    $('#img_file_1').blur().on("click", function() {
        upload_image_1();
    })
    $('#img_file_2').blur().on("click", function() {
        upload_image_2();
    })
    $('#audio_file').blur().on("click", function() {
        upload_audio();
    })
    $('#video_file').blur().on("click", function() {
        upload_video();
    })
}

// blend images
function generate_1() {
    params["user_img"] = src_image;
    params["user_audio"] = $("#user_audios").find("option:selected").val();
    params['amp_lip_x'] = parseFloat($("#amp_lip_x").val());
    params['amp_lip_y'] = parseFloat($("#amp_lip_y").val());
    params['amp_pos'] = parseFloat($("#amp_pos").val());
    var post_data = JSON.stringify(params);
    console.log(post_data)
    $.ajax({
        url: "/MakeItTalk/generate_result/",
        type: "POST",
        cache: false,
        data: post_data,
        success: function(data) {
            var blend_audio = result_video_dir +
                src_image.replace(".jpeg", "").replace(".jpg", "").replace(".png", "") + "_" +
                user_audio.replace(".mp3", "").replace(".wav", "").replace(".cda", "") + "_" +
                params['amp_lip_x'].toString() + params['amp_lip_y'].toString()+params['amp_pos'].toString()+ '.mp4'
            var blend_show_div = document.getElementById("blend_show_1");
            var embed = blend_show_div.getElementsByTagName("embed")[0];
            var hasembed = embed ? true : false;
            if (hasembed) {
                blend_show_div.innerHTML = '';
            }
            var audio = document.createElement("embed");
            audio.src = blend_audio;
            blend_show_div.appendChild(audio);
            alert("blend success!")
        },
        error: function(data) {
            alert("出现错误，请联系管理员！");
        }
    })
}
// blend images2
function generate_2() {
    params2["user_img"] = src_image_2;
    params2["user_video"] = $("#user_videos").find("option:selected").val();
    var post_data = JSON.stringify(params2);
    console.log(post_data)
    $.ajax({
        url: "/firstOrder/generate_result/",
        type: "POST",
        cache: false,
        data: post_data,
        success: function(data) {
            var blend_video = result_video_dir_2 + src_image_2.replace(".jpeg", "").replace(".jpg", "").replace(".png", "") + "_" + user_video.replace(".mp4", "").replace(".avi", "").replace(".wmv", "") + '.mp4'
            var blend_show_div = document.getElementById("blend_show_2");
            var embed = blend_show_div.getElementsByTagName("embed")[0];
            var hasembed = embed ? true : false;
            if (hasembed) {
                blend_show_div.innerHTML = '';
            }
            var video = document.createElement("embed");
            video.src = blend_video;
            blend_show_div.appendChild(video);
            alert("blend success!")
        },
        error: function(data) {
            alert("出现错误，请联系管理员！");
        }
    })
}

// upload images_1
function upload_image_1() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#img_file_1").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#img_file_1").files[0];
        var name = file.name;

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/MakeItTalk/upload_img/");
        //回调
        xhr.onreadystatechange = function() {
                if (xhr.readyState === 4 && xhr.status === 200) {
                    alert("上传成功！");
                    get_base_info_1();
                }
            }
            //获取上传的进度
        xhr.upload.onprogress = function(event) {
                if (event.lengthComputable) {
                    var percent = event.loaded / event.total * 100;
                    document.querySelector("#up_progress_1 .progress-item").style.width = percent + "%";
                }
            }
            //将formdata上传
        xhr.send(formdata);
    });
}
// upload images_2
function upload_image_2() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#img_file_2").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#img_file_2").files[0];
        var name = file.name;

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/firstOrder/upload_img/");
        //回调
        xhr.onreadystatechange = function() {
                if (xhr.readyState === 4 && xhr.status === 200) {
                    alert("上传成功！");
                    get_base_info_2();
                }
            }
            //获取上传的进度
        xhr.upload.onprogress = function(event) {
                if (event.lengthComputable) {
                    var percent = event.loaded / event.total * 100;
                    document.querySelector("#up_progress_3 .progress-item").style.width = percent + "%";
                }
            }
            //将formdata上传
        xhr.send(formdata);
    });
}
// upload audios
function upload_audio() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#audio_file").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#audio_file").files[0];
        var name = file.name;

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/MakeItTalk/upload_audio/");
        //回调
        xhr.onreadystatechange = function() {
                if (xhr.readyState === 4 && xhr.status === 200) {
                    alert("上传成功！");
                    get_base_info_1();
                }
            }
            //获取上传的进度
        xhr.upload.onprogress = function(event) {
                if (event.lengthComputable) {
                    var percent = event.loaded / event.total * 100;
                    document.querySelector("#up_progress_2 .progress-item").style.width = percent + "%";
                }
            }
            //将formdata上传
        xhr.send(formdata);
    });
}
// upload video
function upload_video() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#video_file").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#video_file").files[0];
        var name = file.name;

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/firstOrder/upload_video/");
        //回调
        xhr.onreadystatechange = function() {
                if (xhr.readyState === 4 && xhr.status === 200) {
                    alert("上传成功！");
                    get_base_info_2();
                }
            }
            //获取上传的进度
        xhr.upload.onprogress = function(event) {
                if (event.lengthComputable) {
                    var percent = event.loaded / event.total * 100;
                    document.querySelector("#up_progress_4 .progress-item").style.width = percent + "%";
                }
            }
            //将formdata上传
        xhr.send(formdata);
    });
}