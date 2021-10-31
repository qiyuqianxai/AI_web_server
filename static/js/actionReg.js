// 当前使用功能
// 人体检测
body_detect = true
// 人脸检测
face_detect = false
// 人体关键点检测
body_keypoints_detect = false
//人脸关键点检测
face_keypoints_detect = false
// 识别
action_reg = false
expression_reg = false
attention_reg = false
// 人体跟踪
person_track = false
// 数据分析
dense_analysis = false
// 展示的当前的图片的id
img_id = 0
max_id = -1

// 播放视频
play_vedio = false
// 播放速度
normal_rate = 120
play_rate = normal_rate

// 当前视频名
current_video = "capture"

// 确认次数
ensure_count = 0

predictions_dir = "algorithm/actionReg/predictions/"

$(function () {
    stop_play();
    // 加载按键信息
    set_click_response();
    // 获取视频列表等信息
    get_base_info()
    // 初始化读取读取显示图片图片
    show_pic();

    // 展示当前播放信息
    // show_play_info();
    // 数据分析
    // data_analysis();

    window.onbeforeunload = stop_play;
    window.onended = stop_play;
});


// 加载网页基本信息
function get_base_info() {
    $.ajax({
        url:"/actionReg/get_base_info/",
        contentType: "application/json; charset=utf-8",
        type:"GET",
        cache:false,
        success:function(data){
            //每次加载时重置一些参数
            var json_data=data;
            var video_list = json_data['video_list'];//视频列表
            $('#current-video').empty();
            // 加载视频
            $.each(video_list, function (i, video_name) {
                $("#current-video").append("<option value=" + i + ">"+video_name+"</option>");
            });
            $('#current-video').on('change', function(e){
                if (e.originalEvent) {
                    let video_name = $(this).find("option:selected").text();
                    if(video_name !== current_video)
                    {
                        $('#show_img').attr("src", "images/cover.jpg").css({"width":"100%","height":"100%"});
                        play_vedio = false;
                        current_video = video_name;
                        max_id = -1;
                        img_id = 0;
                        play_rate = normal_rate;
                        update_play_info();
                    }
                    // console.log(current_video);
                }
            });

        },
        error:function(data){
            alert("数据加载出错，请联系管理员！");
            // top.location.reload();
        }
    });
}

// 设置各个功能响应
function set_click_response() {
    // ai功能响应
    $('#body_detect').blur().on("click",function () {
        if (!play_vedio)
            body_detect = true;
    });

    $('#face_detect').blur().on("click",function () {
        if (!play_vedio)
            face_detect = true;


    });
    $('#body_keypoints_detect').blur().on("click",function () {
        if (!play_vedio)
        {
            body_detect = true;
            body_keypoints_detect = true;
        }
    });

    $('#face_keypoints_detect').blur().on("click",function () {
        if (!play_vedio)
        {
            face_detect = true;
            face_keypoints_detect = true;
        }
    });

    $('#person_track').blur().on("click",function () {
        if (!play_vedio)
        {
            body_detect = true;
            person_track = true;
        }

    });

    $('#action_reg').blur().on("click",function () {
        if (!play_vedio)
        {
            body_detect = true;
            action_reg = true;
        }
    });

    $('#expression_reg').blur().on("click",function () {
        if (!play_vedio)
        {
            face_detect = true
            expression_reg = true;
        }
    });
    $('#attention_reg').blur().on("click",function () {
        if (!play_vedio)
        {
            face_detect = true
            attention_reg = true;
        }
    });

    $('#dense_analysis').blur().on("click",function () {
        if (!play_vedio)
            dense_analysis = true;
    })

    $('.ai_feature').blur().on("click", function () {
        if (!play_vedio)
        {
            update_play_info();
            show_cur_function();
        }

    })

    // 播放操作响应
    $('#play_video').blur().on("click",function () {
        if (!play_vedio)
        {
            play_vedio = true;
            play_rate = normal_rate;
            update_play_info();
        }

    });

    $('#play_stop').blur().on("click",function () {
       play_vedio = false;
    });

    $('#play_advance').blur().on("click",function () {
        if (img_id+play_rate<max_id)
            img_id += play_rate;
        else
            img_id = max_id;
    });

    $('#play_back').blur().on("click",function () {
        if(img_id - play_rate > 0)
            img_id -= play_rate;
        else
            img_id = 0;
    });

    $('#play_accelerate').blur().on("click",function () {
        if( current_video === "capture")
        {
            alert("抱歉，播放录像模式下无法加速或减速！");
        }
        else {
            if(play_rate > 10 ) play_rate -= 10;
            else alert("已经加速到最大值！");
        }
    });

    $('#play_reduce').blur().on("click",function () {
        if( current_video === "capture")
        {
            alert("抱歉，播放录像模式下无法加速或减速！");
        }
        else {
            if(play_rate < 300) play_rate += 10;
            else alert("已经减速到最小值！");
        }
    })

    $('#play_end').blur().on("click",stop_play);

    $('#upload_video').blur().on("click",function () {
        upload_video();
    })
}

function stop_play() {
    play_vedio = false;
    max_id = -1;
    img_id = 0;
    play_rate = normal_rate;
    ensure_count = 0;
    $.ajax({
        url: "/actionReg/play_end/",
        type: "GET",
        cache:false,
        success: function (data) {
            console.log("video shutdowm!")
        },
        error: function (data) {
            alert("出现错误，请联系管理员！");
        }
    })
    $('#show_img').attr("src", "images/cover.jpg").css({"width":"100%","height":"100%"});
}

function update_play_info() {
    var data = JSON.stringify({
        id: img_id,//当前index
        video_name: current_video,// 当前视频名称
        // 需要的ai功能
        req_face_detect:face_detect,
        req_body_detect:body_detect,
        req_body_keypoints:body_keypoints_detect,
        req_face_keypoints:face_keypoints_detect,
        req_action_reg:action_reg,
        req_expression_reg:expression_reg,
        req_attention_reg:attention_reg,
        person_track:person_track,
      });
    $.ajax({
        url: "/actionReg/update_play_info/",
        type: "POST",
        cache:false,
        data:data,
        success: function (data) {
            var json_data= data;
            max_id = json_data["max_id"];
            // console.log(max_id);
        },
        error: function (data) {
            alert("出现错误，请联系管理员！");
        }
    })
}

// 展示播放内容
function show_pic() {
  if(play_vedio){
      if (ensure_count > 200)
      {
          alert("当前视频播放完成！")
          play_vedio = false;
          ensure_count = 0;
          max_id = -1;
          img_id = 0;
          play_rate = normal_rate;
          $('#show_img').attr("src", "images/cover.jpg").css({"width":"100%","height":"100%"});
      }
      else {
              if(img_id > max_id){
                  update_play_info();
                  ensure_count++;
                }
              else {
                  $('#show_img').attr("src", predictions_dir+ current_video+ "/" + (img_id).toString() + ".jpg").css({"width":"100%","height":"100%"});
                  var windowW = $(window).width();//获取当前窗口宽度
                  var windowH = $(window).height()*0.8;//获取当前窗口高度
                  var realWidth = $('#show_img').width();//获取图片真实宽度
                  var realHeight = $('#show_img').height();//获取图片真实高度
                  var scale = Math.max(realWidth/windowW,realHeight/windowH);//缩放尺寸，当图片真实宽度和高度大于窗口宽度和高度时进行缩放
                  console.log(realWidth,realHeight,windowW,windowH,scale)
                  $('#show_img').css({"width":realWidth/scale,"height":realHeight/scale});

                  img_id++;
                  ensure_count = 0;
                  data_analysis();
              }
      }
  }
  setTimeout("show_pic()", play_rate);
}

// 展示当前已启动的功能
function show_cur_function() {
    var html_s = "";
    if(face_detect) html_s += '<button value="face_detect" class="cur_features btn-block">人脸检测</button>';
    if(body_detect) html_s += '<button value="body_detect" class="cur_features btn-block">人体检测</button>';
    if(person_track) html_s += '<button value="person_track" class="cur_features btn-block">人体跟踪</button>';
    if(body_keypoints_detect) html_s += '<button value="body_keypoints_detect" class="cur_features btn-block">人体关键点检测</button>';
    if(face_keypoints_detect) html_s += '<button value="face_keypoints_detect" class="cur_features btn-block">人脸关键点检测</button>';
    if(action_reg) html_s += '<button value="action_reg" class="cur_features btn-block">行为识别</button>';
    if(expression_reg) html_s += '<button value="expression_reg" class="cur_features btn-block">表情识别</button>';
    if(attention_reg) html_s += '<button value="attention_reg" class="cur_features btn-block">注意力识别</button>';
    if(dense_analysis) html_s += '<button value="dense_analysis" class="cur_features btn-block">数据分析</button>';
    // 播放中不允许更改功能
    if(!play_vedio)
        $('#current-function').html(html_s);
    $('.cur_features').blur().on("click",function () {
        if(!play_vedio)
        {
            var id_name = $(this).attr("value");
            // console.log(id_name);
            if(id_name === "dense_analysis")
            {
                $(".data_pie").html("");
            }
            eval(id_name+"="+"false");
            show_cur_function();
        }
    })

}

// 上传视频到服务器
function upload_video() {
    //首先监听input框的变动，选中一个新的文件会触发change事件
    document.querySelector("#file").addEventListener("change",function () {
        //获取到选中的文件
        var file = document.querySelector("#file").files[0];
        var name = file.name;

        //创建formdata对象
        var formdata = new FormData();
        formdata.append("file",file);
        //创建xhr，使用ajax进行文件上传
        var xhr = new XMLHttpRequest();
        xhr.open("post","/actionReg/upload_file/");
        //回调
        xhr.onreadystatechange = function () {
            if (xhr.readyState==4 && xhr.status==200){
                alert("上传成功！");
                get_base_info();
            }
        }
        //获取上传的进度
        xhr.upload.onprogress = function (event) {
            if(event.lengthComputable){
                var percent = event.loaded/event.total *100;
                document.querySelector("#up_progress .progress-item").style.width = percent+"%";
            }
        }
        //将formdata上传
        xhr.send(formdata);
    });
}

// 展示数据分析的结果
function data_analysis() {
    // console.log("data analysis...")
    if(play_vedio) {
        var data=JSON.stringify({
          id: img_id,//当前index
          video: current_video,// 当前视频名称
        });
        $.ajax({
            url: "/actionReg/get_analysis_data/",
            type: "POST",
            data: data,
            cache:false,
            success: function (data) {
                var json_data = data;
                // console.log(json_data);
                if(json_data.hasOwnProperty("action_analysis")){
                    var action_data = json_data["action_analysis"];
                    var expresstion_data = json_data["expression_analysis"];
                    var attention_data = json_data["attention_analysis"];
                    var person_num = json_data["person_num"];
                    var i;
                    let html_s = "<p>视频名称："+current_video+"</p>"
                        + "<p>播放速度："+play_rate.toString()+"</p>"
                        + "<p>当前帧数："+img_id.toString()+"</p>"
                        + "<p>当前帧率："+(1000/play_rate).toFixed(1).toString()+"</p>"
                        + "<p>检测到人数:"+person_num.toString()+"</p>"

                    for (i=0;i<action_data.length;i++)
                    {
                        action_data[i][1] = parseInt(action_data[i][1]);
                        // html_s += "<p>"+action_data[i][0]+":"+action_data[i][1].toString()+"</p>"
                    }
                    for (i=0;i<expresstion_data.length;i++)
                    {
                        expresstion_data[i][1] = parseInt(expresstion_data[i][1]);
                        // html_s += "<p>"+expresstion_data[i][0]+":"+expresstion_data[i][1].toString()+"</p>"
                    }
                    for (i=0;i<attention_data.length;i++)
                    {
                        attention_data[i][1] = parseInt(attention_data[i][1]);
                        // html_s += "<p>"+attention_data[i][0]+":"+attention_data[i][1].toString()+"</p>"
                    }
                    $("#play_info").html(html_s)
                    if (dense_analysis && img_id%100===0){
                        draw_pie_chart(action_data,"action");
                        draw_pie_chart(expresstion_data,"expression");
                        draw_pie_chart(attention_data,"attention");
                    }
                }
            },
            error: function (data) {
                // alert("出现错误，请联系管理员！");
                console.log(data)
            }
        });
    }
}

function draw_pie_chart(task_data,target) {
    var chart = {
       plotBackgroundColor: null,
       plotBorderWidth: null,
       plotShadow: false
   };
   var title = {
      text: target+"分布"
   };
   var tooltip = {
      pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
   };
   var plotOptions = {
      pie: {
         allowPointSelect: true,
         cursor: 'pointer',
         dataLabels: {
            enabled: true,
            format: '<b>{point.name}%</b>: {point.percentage:.1f} %',
            style: {
               color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
            }
         }
      }
   };
   var series= [{
      type: 'pie',
      name: target+"占比",
      data: task_data,

   }];

   var json = {};
   json.chart = chart;
   json.title = title;
   json.tooltip = tooltip;
   json.series = series;
   json.plotOptions = plotOptions;
   $('#'+target+'_data_analysis').highcharts(json);
}

function show_play_info() {
    let html_s = "<p>视频名称："+current_video+"</p>"
        + "<p>播放速度："+play_rate.toString()+"</p>"
        + "<p>当前帧数："+img_id.toString()+"</p>"
        + "<p>当前帧率："+(1000/play_rate).toFixed(1).toString()+"</p>"
    $("#play_info").html(html_s)
    // setTimeout("show_play_info()",play_rate);
}