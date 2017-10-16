
<div id="site_content">
  <div class="sidebar">
    <h4>一个瓜皮打代码</h4>
    <h5>2016-09-01</h5>
    <p>不引战，不BB<br /></p>
    <h3>文章分类</h3>
    <ul>
      	{{range .category}}
      	<li><a href="/?cid={{.Id}}">{{.TypeName}}</a></li>
		{{end}}
    </ul>
    <h1>Search</h1>
      <p style="width:190px;">
        <input class="search" type="text" name="search_field" placeholder="Enter keywords....." />
        <input name="search" type="image" onclick="search()" style="border: 0; margin: 0 0 -9px 5px;" src="./static/css/search.png" alt="Search" title="Search" />
      </p>
  </div>
  <div id="content">
    <div class="detail">
        <h1>{{.Article.Title}}</h1>
	{{str2html .Article.Content}}
    </div>
  </div>
</div>
<div class="side" style="">
    <div style="border:1px solid #A9A9A9; width:600px;"></div>
    <p>评论区</p>
    <div style="border-bottom:1px solid #A9A9A9 ;width:600px;">
       <p style="padding:0px 0px 0px 0px; color:#d32;">123</p>
       <span>123123</span><br>
       <span style="font-size: 12px;margin-right: 8px;color: #999;">123123</span>
    </div>
    <div style="margin-top:30px;">
        <div class="form_settings" style="margin-bottom:10px;">
		{{if .Member}}
            <div>你好， 123</div>
		{{else}}
		<span>帐号：</span><input class="contact" type="text" name="username" value="" /></p>
            <p><span>密码：</span><input class="contact" type="password" name="password" value="" /></p>
            <button class="button" onclick="login()">登录</button>

		{{end}}
        </div>
	{{if .Member}}
        <textarea rows="8" cols="50" name="name"></textarea><br>
        <button class="button">发表评论</button>
	{{end}}
    </div>
</div>

<script>
	function search(keyword){
		var keyword = $("input[name=keyword]").val()
		window.location.href="/?keyword="+keyword
	}
	
	var loginLock = false
	function login(){
		if (loginLock){
			return ;
		}
		var username = $("input[name=username]").val();
		var password = $("input[name=password]").val();
		if(username == ""){
			alert("用户名不能为空")
			return ;
		}
		if(password == ""){
			alert("密码不能为空")
			return ;
		}
		loginLock = true;
		$.ajax({
	        type: "POST",
	        cache: false,
	        url: "/login",
	        data:{username: username, password: password},
	        dataType: 'json',
	        success: function (obj) {
	            loginLock = false;
				console.log(obj)
	        },
			error: function(obj){
				console.log(obj)
			}
	    });
	}
	
	
</script>
