<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <a href="/">WhereSmile</a> / 发布话题
      </div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form method="post" action="/topic/create">
          <div class="form-group">
            <label for="title">标题</label>
            <input type="text" class="form-control" id="title" name="title" placeholder="标题">
          </div>
          <div class="form-group">
            <ul class="col-md-12 list-inline">
              <li class="edit"><a href="#" tabindex="-1"><i class="glyphicon glyphicon-edit"></i> 编辑</a></li>
              <li class="preview"><a href="#" title="预览" tabindex="-1"><i class="glyphicon glyphicon-eye-open preview"></i> 预览</a></li>
              <!-- <li class="wide"><a href="/wide/playground" tabindex="-1" target="_blank" title="通过Wide编辑代码"><i class="glyphicon glyphicon-cloud"></i> Wide</a></li> -->
            </ul>
            <textarea name="content" id="content" rows="15" class="form-control" placeholder="支持Markdown语法哦~"></textarea>
            <div class="content-preview"></div>
          </div>
          <div class="form-group">
            <label for="title">版块</label>
            <select name="sid" id="sid" class="form-control">
              {{range .Sections}}
                <option value="{{.Id}}">{{.Name}}</option>
              {{end}}
            </select>
          </div>
          <button type="submit" class="btn btn-default">发布</button>
        </form>
      </div>
    </div>
  </div>
</div>

<script type="text/javascript">
  $(function () {
    $('form .form-group .edit').on('click', function(evt){
      evt.preventDefault();
      $('#content').show();
      $('.content-preview').hide();
      return;
    });
    
    $('form .form-group .preview').on('click', function(evt){
      evt.preventDefault();
      $('#content').hide();
      // 值
      content=$("#content").val();
      $.post("/topic/preview",{content:content},function(data, textStatus){
        $(".content-preview").html(data);
        $('.content-preview').show();
      });
    });
  });
</script>