// current src image
src_image = ""
ref_image = ""

params = {
    "user_img":"77d6aa752a15d04427f1ac4b6e4864aa.jpeg",
    "ref_img":"",
    "render_factor": 20,
    "artistic": true
}

//
user_imgs_dir = "algorithm/colorImage/user_imgs/"
ref_imgs_dir = "algorithm/colorImage/ref_imgs/"
result_imgs_dir = "algorithm/colorImage/res_imgs/"

$(function() {

    window.onbeforeunload = boot_model(["deoldify"]);

    // 加载按键信息
    set_click_response();
    //
    get_base_info();
});

// load src_imgs list
function get_base_info() {
    $.ajax({
        url: "/colorImage/get_base_info/",
        contentType: "application/json; charset=utf-8",
        type: "GET",
        cache: false,
        success: function(data) {
            //每次加载时重置一些参数

            var src_list = data['src_list']; //视频列表

            $('#user_imgs').empty();
            $.each(src_list, function(i, img_name) {
                $("#user_imgs").append("<option value=" + img_name + ">" + img_name + "</option>");
            });
            $("#user_imgs").append("<option value=''></option>");
            $('#user_imgs').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_img = $(this).find("option:selected").val();
                    if (selected_img !== src_image) {
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
                            src_image = selected_img;
                        }

                    }
                    // console.log(current_video);
                }
            });

            var ref_list = data['ref_list']
            console.log(ref_list);
            $('#ref_imgs').empty();
            $.each(ref_list, function(i, img_name) {
                $("#ref_imgs").append("<option value=" + img_name + ">" + img_name + "</option>");
            });
            $("#ref_imgs").append("<option value=''></option>");
            $('#ref_imgs').on('change', function(e) {
                if (e.originalEvent) {
                    let selected_img = $(this).find("option:selected").val();
                    if (selected_img !== ref_image) {
                        var img = new Image()
                            // 改变图片的src
                        img.src = ref_imgs_dir + selected_img
                            // 加载完成执行
                        img.onload = function() {
                            $('#ref_img').attr("src", ref_imgs_dir + selected_img);
                            var windowW = $(window).width() * 0.45; //获取当前窗口宽度
                            var windowH = $(window).height() * 0.51; //获取当前窗口高度
                            var realWidth = img.width; //获取图片真实宽度
                            var realHeight = img.height; //获取图片真实高度
                            var scale = Math.max(realWidth / windowW, realHeight / windowH); //缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                            console.log(realWidth, realHeight, windowW, windowH, scale)
                            $('#ref_img').css({ "width": realWidth / scale, "height": realHeight / scale });
                            ref_image = selected_img;
                        }

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
    $('#blend').blur().on("click", function() {
        blend_img();
    });

    $('#upload_src').blur().on("click", function() {
        upload_image('upload_src');
    });

    $('#upload_ref').blur().on("click", function() {
        upload_image('upload_ref');
    })


}

// blend images
function blend_img() {
    params["user_img"] = src_image;
    params["ref_img"] = ref_image;
    // params['model'] = $('#user_model').find("option:selected").val()
    params['render_factor'] = parseInt($("#render_factor").val())
    params['artistic'] = parseInt($("#artistic").val()) === 1
    var model = $('#user_model').find("option:selected").val();
    boot_model([model]);
    var post_data = JSON.stringify(params);
    $.ajax({
        url: "/colorImage/generate_img/",
        type: "POST",
        cache: false,
        data: post_data,
        success: function(data) {
            var blend_img;
            if(model === "debvc")
                blend_img = result_imgs_dir + src_image.replace(".jpeg", "").replace(".jpg", "").replace(".png", "") + "_" + ref_image.replace(".jpeg", "").replace(".jpg", "").replace(".png", "") + '.png'
            else
                blend_img = result_imgs_dir + "deoldify_"+$("#render_factor").val().toString()+$("#artistic").val().toString()+"_"+ src_image;
            var img = new Image()
                // 改变图片的src
            img.src = blend_img
                // 加载完成执行
            img.onload = function() {
                $('#blend_img').attr("src", blend_img);
                var windowW = $(window).width() * 0.45; //获取当前窗口宽度
                var windowH = $(window).height() * 0.51; //获取当前窗口高度
                var realWidth = img.width; //获取图片真实宽度
                var realHeight = img.height; //获取图片真实高度
                var scale = Math.max(realWidth / windowW, realHeight / windowH); //缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                console.log(realWidth, realHeight, windowW, windowH, scale)
                $('#blend_img').css({ "width": realWidth / scale, "height": realHeight / scale });
            }
            alert("blend success!")
        },
        error: function(data) {
            alert("出现错误，请联系管理员！");
        }
    })
}

// upload images
function upload_image(target) {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#"+target+"").addEventListener("change", function() {
        //获取到选中的文件
        var file = document.querySelector("#"+target+"").files[0];

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file", file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post", "/colorImage/"+target);
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