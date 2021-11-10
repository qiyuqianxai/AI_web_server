// current src image
current_image = ""
current_video = ""
direction_val = {
    "user_img":"2.jpg",
    "face_enhance":true,
    "model":"RealESRGAN_x2plus.pth"
}

user_imgs_dir = "algorithm/realesrgan/user_imgs/" //上传图片地址

fake_imgs_dir = "algorithm/realesrgan/result_imgs/" //生成图片地址

user_video_dir = "algorithm/realesrgan/user_videos/"

res_video_dir = "algorithm/realesrgan/res_videos/"

$(function() {

    window.onbeforeunload = boot_model(["realesrgan"]);


    var html_str = "";

    html_str += '<div style="width:35%; float: left;border-bottom: #1b1e21 ;justify-content: center;align-items: center" >\n' +
        '                            <button class="btn-primary" style="margin-left:1%;margin-top:1%;margin-bottom:1%;font-family: 微软雅黑;font-size: large;float: left;width:70%">select network</button>\n' +
        '                            <select id="network" name="network">' +
        '<option value="RealESRGAN_x2plus.pth">RealESRGAN_x2plus</option>' +
        '<option value="RealESRGAN_x4plus.pth">RealESRGAN_x4plus</option>' +
        '<option value="RealESRGAN_x4plus_anime_6B.pth">RealESRGAN_x4plus_anime_6B</option>' +
        '</select>\n' +
        '                        </div>'
    html_str += '<div style="width:35%; float: left;border-bottom: #1b1e21 ;justify-content: center;align-items: center" >\n' +
        '                            <button class="btn-primary" style="margin-left:1%;margin-top:1%;margin-bottom:1%;font-family: 微软雅黑;font-size: large;float: left;width:70%">face enhance</button><Br/><Br/>' +
        '                            <input type="radio" name="face_enhance" value="1" checked="checked"><span>是</span>\n' +
        '                            <input type="radio" name="face_enhance" value="0"><span>否</span>\n' +
        '                        </div>'
    html_str +=
        '                            <div id="circle-btn" style="width:30%; float: left;justify-content: center;align-items: center" >\n' +
        '                                <button id="image_generate" class="btn-warning" style="height:70%; width:100%; margin-left:5%; margin-top:1%; margin-bottom:1%;font-family: 微软雅黑; font-size: large;float: left; border:medium none; border-radius: 5px;">&nbsp;generate&nbsp;&nbsp;➤&nbsp;</button>\n' +
        '                            </div>'
    $('#gan-param-set').html(html_str);

    html_str = "";

    html_str += '<div style="width:35%; float: left;border-bottom: #1b1e21 ;justify-content: center;align-items: center" >\n' +
        '                            <button class="btn-primary" style="margin-left:1%;margin-top:1%;margin-bottom:1%;font-family: 微软雅黑;font-size: large;float: left;width:70%">select network</button>\n' +
        '                            <select id="network_2" name="network">' +
        '<option value="RealESRGAN_x2plus.pth">RealESRGAN_x2plus</option>' +
        '<option value="RealESRGAN_x4plus.pth">RealESRGAN_x4plus</option>' +
        '<option value="RealESRGAN_x4plus_anime_6B.pth">RealESRGAN_x4plus_anime_6B</option>' +
        '</select>\n' +
        '                        </div>'
    html_str += '<div style="width:35%; float: left;border-bottom: #1b1e21 ;justify-content: center;align-items: center" >\n' +
        '                            <button class="btn-primary" style="margin-left:1%;margin-top:1%;margin-bottom:1%;font-family: 微软雅黑;font-size: large;float: left;width:70%">face enhance</button><Br/><Br/>' +
        '                            <input type="radio" name="face_enhance" value="1" checked="checked"><span>是</span>\n' +
        '                            <input type="radio" name="face_enhance" value="0"><span>否</span>\n' +
        '                        </div>'
    html_str +=
        '                            <div style="width:30%; float: left;justify-content: center;align-items: center" >\n' +
        '                                <button id="video_generate" class="btn-warning" style="height:70%; width:100%; margin-left:5%; margin-top:1%; margin-bottom:1%;font-family: 微软雅黑; font-size: large;float: left; border:medium none; border-radius: 5px;">&nbsp;generate&nbsp;&nbsp;➤&nbsp;</button>\n' +
        '                            </div>'
    $('#gan-param-set-2').html(html_str);
    // 加载按键信息
    set_click_response();
    get_base_info();
});

// load src_imgs list
function get_base_info() {
    $.ajax({
        url: "/realesran/get_base_info/",
        contentType: "application/json; charset=utf-8",
        type: "GET",
        cache: false,
        success: function(data) {
            //每次加载时重置一些参数

            var img_list = data['img_list'];
            $('#current-image').empty();

            $.each(img_list, function(i, img_name) {
                $("#current-image").append("<option value=" + img_name + ">" + img_name + "</option>");
            });
            $("#current-image").append("<option value=''></option>");
            $('#current-image').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_img = $(this).find("option:selected").val();
                    if (selected_img !== current_image) {
                        var img = new Image()
                            // 改变图片的src
                        img.src = user_imgs_dir + selected_img
                            // 加载完成执行
                        img.onload = function() {
                            $('#user_img').attr("src", user_imgs_dir + selected_img);
                            var windowW = $(window).width() * 0.45; //获取当前窗口宽度
                            var windowH = $(window).height() * 0.51; //获取当前窗口高度
                            var realWidth = img.width; //获取图片真实宽度
                            var realHeight = img.height; //获取图片真实高度
                            var scale = Math.max(realWidth / windowW, realHeight / windowH); //缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                            console.log(realWidth, realHeight, windowW, windowH, scale)
                            $('#user_img').css({ "width": realWidth / scale, "height": realHeight / scale });
                            current_image = selected_img;
                        }

                    }
                    // console.log(current_video);
                }
            });

            var video_list = data['video_list'];
            $('#current-video').empty();

            $.each(video_list, function(i, vid_name) {
                $("#current-video").append("<option value=" + vid_name + ">" + vid_name + "</option>");
            });
            $("#current-video").append("<option value=''></option>");
            $('#current-video').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_video = $(this).find("option:selected").val();
                    if (selected_video !== current_video) {
                        var video_url = user_video_dir + selected_video;
                        var video_div = document.getElementById("user_video");
                        var embed = video_div.getElementsByTagName('embed')[0];
                        var hasembed = embed ? true : false;
                        if (hasembed) {
                            //清空父元素下的所有内容，保证始终都是有一个音频链接
                            video_div.innerHTML = '';
                        }
                        var video = document.createElement('embed');
                        video.src = video_url;
                        video_div.appendChild(video);
                        current_video = selected_video;
                    }
                    // console.log(current_video);
                }
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

    $('#upload_image').blur().on("click", function() {
        upload_image();
    })

    $('#image_generate').blur().on("click", function() {
        generate_image();
    })

    $('#upload_video').blur().on("click", function() {
        upload_video();
    })

    $('#video_generate').blur().on("click", function() {
        generate_video();
    })

}


// upload images
function upload_image() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#file").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#file").files[0];
        var name = file.name;

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/realesran/upload_file/");
        //回调
        xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    alert("上传成功！");
                    get_base_info();
                }
            }
            //获取上传的进度
        xhr.upload.onprogress = function(event) {
                if (event.lengthComputable) {
                    var percent = event.loaded / event.total * 100;
                    document.querySelector("#up_progress .progress-item").style.width = percent + "%";
                }
            }
            //将formdata上传
        xhr.send(formdata);
    });
}

// generate images
function generate_image() {
    // $('.face_param').each(function() {
    //     direction_val[$(this).attr('id')] = $(this).val()
    // })
    var network_selected = $('#network').find("option:selected").val()
    direction_val['model'] = network_selected
    var enhance_selected = $("input[type='radio']:checked").val();
    if (enhance_selected==="1")
        direction_val['face_enhance'] = true
    else
        direction_val['face_enhance'] = false
    direction_val['user_img'] = current_image
    var post_data = JSON.stringify(direction_val)

    $.ajax({
        url: "/realesran/generate_img/",
        type: "POST",
        cache: false,
        data: post_data,
        success: function(data) {
            var fake_img = fake_imgs_dir + network_selected + '_' + enhance_selected + "_" + current_image

            $('#fake_img').attr("src", fake_img);
            var img = new Image()
                // 改变图片的src
            img.src = fake_img
                // 加载完成执行
            img.onload = function() {
                var windowW = $(window).width() * 0.45; //获取当前窗口宽度
                var windowH = $(window).height() * 0.51; //获取当前窗口高度
                var realWidth = img.width; //获取图片真实宽度
                var realHeight = img.height; //获取图片真实高度
                var scale = Math.max(realWidth / windowW, realHeight / windowH); //缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                // console.log(realWidth,realHeight,windowW,windowH,scale)
                $('#fake_img').css({ "width": realWidth / scale, "height": realHeight / scale });
            }
            alert("generate success!")
        },
        error: function(data) {
            alert("出现错误，请联系管理员！");
        }
    })

}

// upload videos
function upload_video() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#video_file").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#video_file").files[0];

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/realesran/upload_video/");
        //回调
        xhr.onreadystatechange = function() {
            if (xhr.readyState === 4 && xhr.status === 200) {
                alert("上传成功！");
                get_base_info();
            }
        }
        //获取上传的进度
        xhr.upload.onprogress = function(event) {
            if (event.lengthComputable) {
                var percent = event.loaded / event.total * 100;
                document.querySelector("#up_progress_video .progress-item").style.width = percent + "%";
            }
        }
        //将formdata上传
        xhr.send(formdata);
    });
}

// generate videos
function generate_video() {
     var network_selected = $('#network_2').find("option:selected").val();
    direction_val['model'] = network_selected;
    var enhance_selected = $("input[type='radio']:checked").val();
    direction_val['face_enhance'] = enhance_selected==="1"
    direction_val['user_img'] = current_video
    var post_data = JSON.stringify(direction_val)

    $.ajax({
        url: "/realesran/generate_img/",
        type: "POST",
        cache: false,
        data: post_data,
        success: function(data) {
            var blend_video = res_video_dir + network_selected + '_' + enhance_selected +"_" +current_video
            var blend_show_div = document.getElementById("res_video");
            var embed = blend_show_div.getElementsByTagName("embed")[0];
            var hasembed = embed ? true : false;
            if (hasembed) {
                blend_show_div.innerHTML = '';
            }
            var video = document.createElement("embed");
            video.src = blend_video;
            blend_show_div.appendChild(video);
            alert("generate success!")
        },
        error: function(data) {
            alert("出现错误，请联系管理员！");
        }
    })

}