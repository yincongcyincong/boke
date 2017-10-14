
<div id="site_content">
  <div class="sidebar">
    <h1>PHP</h1>
    <h5>2010-01-01</h5>
    <p>PHP是世界上最好的语言<br /><a href="#">Read more</a></p>
    <h1>文章分类</h1>
    <ul>
	{{range .category}}
      <li><a href="#">{{.TypeName}}</a></li>
	{{end}}
    </ul>
    <h1>Search</h1>
    <form method="post" action="#" id="search_form">
      <p>
        <input class="search" type="text" name="search_field" placeholder="Enter keywords....." />
        <input name="search" type="image" style="border: 0; margin: 0 0 -9px 5px;" src="./static/css/search.png" alt="Search" title="Search" />
      </p>
    </form>
  </div>
  <div id="content">
    <div>
	{{range .article }}
        <a href="/page?aid={{.Id}}">
            <h1>{{.Title}}</h1>
            <p>{{.Content}}</p>
        </a>
	{{if $.isMaster}}
        <button class="button">编辑</button>
	{{end}}
        <span style="font-size: 12px;margin-right: 8px;color: #999; float:right;">{{.Ctime}}</span>
	{{end}}
    </div>
	<ul id="page"></ul>
  </div>
</div>
<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
<script type="text/javascript">
  $(function () {
    $("#tab_{{.S}}").addClass("active");
    $("#page").bootstrapPaginator({
      currentPage: '{{.Page.PageNo}}',
      totalPages: '{{.Page.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        var s = {{.S}};
        if (s > 0) {
          window.location.href = "/?p=" + page + "&s={{.S}}"
        } else {
          window.location.href = "/?p=" + page
        }
      }
    });
  });
</script>
