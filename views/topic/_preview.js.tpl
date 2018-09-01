
<script type="text/javascript">
  $(function () {
    $('form .form-group .edit').on('click', function(evt){
      evt.preventDefault();
      $('#content').show();
      $('.edit').addClass('cur');
      $('.preview').removeClass('cur');
      $('.content-preview').hide();
      return;
    });
    
    $('form .form-group .preview').on('click', function(evt){
      evt.preventDefault();
      $('#content').hide();
      $('.edit').removeClass('cur');
      $('.preview').addClass('cur');
      // 值
      content=$("#content").val();
      $.post("/topic/preview",{content:content},function(data, textStatus){
        $(".content-preview").html(data);
        $('.content-preview').show();
      });
    });
  });
</script>

<style type="text/css">
  .cur a{
    border:1px solid #ddd;
    border-radius:8px;
  }
</style>