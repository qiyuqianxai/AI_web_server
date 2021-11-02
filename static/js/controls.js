
all_model_states = {
    "actionReg":false,
    "realesrgan":false,
    "stylegan":false,
    "stargan":false
}

function boot_model(model_name){
    $.each(all_model_states,function (key,val) {
        all_model_states[key] = key === model_name;
    })
    update_models_states();
}

function stop_model() {
    $.each(all_model_states,function (key,val) {
        all_model_states[key] = false;
    })
    update_models_states();
}

function update_models_states(){
    var post_data = JSON.stringify(all_model_states)
    $.ajax({
        url: "/update_models_states/",
        type: "POST",
        cache: false,
        data: post_data,
        success: function(data) {
            console.log("update models states success!");
        },
        error: function(data) {
            alert("出现错误，请联系管理员！");
        }
    })
}