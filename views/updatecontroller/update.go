  <div id="site_content" style="height:1000px;">
      <div class="form_settings" style="margin-bottom:10px;">
          <p><span>文章类别：</span>
          <select class="contact"  name="your_name" style="height:33px;">
              	{{range .category}}
				<option value="{{.Id}}">{{.Title}}</option>
				{{end}}
          </select></p>
          <p><span>密码：</span><input class="contact" type="text" name="your_email" value="" /></p>
          <div><button class="button" onclick="subm()">提交</button></div>
      </div>
      <textarea id="article"></textarea>
  </div>

<script type="text/javascript" charset="utf-8" src="./static/js/ueditor/utf8-php/ueditor.config.js"></script>
<script type="text/javascript" charset="utf-8" src="./static/js/ueditor/utf8-php/ueditor.all.js"> </script>
<script type="text/javascript" charset="utf-8" src="./static/js/ueditor/utf8-php/lang/zh-cn/zh-cn.js"></script>
<script type="text/javascript">
  	var editor_a = new baidu.editor.ui.Editor();
  	editor_a.render( 'article' ); //此处的参数值为编辑器的id值
	
	function subm(){
		$.ajax({
	        type: "POST",
	        cache: false,
	        url: "/login",
	        data:$("form").serialize(),
	        dataType: 'json',
	        success: function (obj) {
			    if (obj.Code == "200"){
					alert("操作成功")
					window.location.href="/"
			    } else{
					alert(obj.Content)
			    }
	        },
			error: function(obj){
				console.log(obj)
			}
	    });
	}

</script>
