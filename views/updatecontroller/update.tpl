  <div id="site_content" style="height:700px; background:#f6f6f0;">
<form>
      <div class="form_settings" style="margin-bottom:10px;">
          <p><span>文章类别：</span>
          <select class="contact"  name="categoryId" style="height:33px;">
              	{{range .category}}
				<option value="{{.Id}}">{{.TypeName}}</option>
				{{end}}
          </select></p>
          <p><span>文章名称：</span><input class="contact" type="text" name="title" value="" /></p>
      </div>
      <textarea id="article" name="content"></textarea>
</form>
          <div style="margin-top:100px;"><button class="button" onclick="subm()">提交</button></div>
  </div>
<script type="text/javascript" charset="utf-8" src="./static/js/ueditor/ueditor.config.js"></script>
<script type="text/javascript" charset="utf-8" src="./static/js/ueditor/ueditor.all.js"> </script>
<script type="text/javascript" charset="utf-8" src="./static/js/ueditor/lang/zh-cn/zh-cn.js"></script>
<script type="text/javascript">
  	var editor_a = new baidu.editor.ui.Editor();
  	editor_a.render( 'article' ); //此处的参数值为编辑器的id值
	
	function subm(){
		$.ajax({
	        type: "POST",
	        cache: false,
	        url: "/doUpdate",
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
