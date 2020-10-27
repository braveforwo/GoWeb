prettyPrint();
layui.use(['layedit'], function () {
    var form = layui.form();
    var $ = layui.jquery;
    var layedit = layui.layedit;

    //评论和留言的编辑器
    var editIndex = layedit.build('remarkEditor', {
        height: 150,
        tool: ['face', '|', 'left', 'center', 'right', '|', 'link'],
    });

    $("#commitComment").click(function (){
        comments={}
        comments["comment"]=layedit.getText(editIndex)
        comments["articleId"]=parseInt($("#articleId").val())
        $.ajax({
            type: "POST",
            url: "/commitComment",
            dataType: "json",
            contentType: "application/json;charset=utf-8",
            data:JSON.stringify(comments),
            success:function (message) {
                layer.msg("评论成功",{icon:1})
            },
            error:function (message) {
                console.log(message)
                layer.msg(message.responseJSON.msg,{icon:5})
            }
        });
    })
        return false;
});