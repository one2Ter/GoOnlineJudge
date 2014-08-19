{{define "content"}}
<h1>Rejudge</h1>

<form accept-charset="UTF-8" id="search_form">
<select id="type" name="type">
<option value="Pid">Problem ID</option>
<option value="Sid">Solution ID</option>
</select>
<input id="id" name="id" size="20" type="text">
<input name="commit" type="submit" value="Rejudge">
</form>

<script type="text/javascript">
$('#search_form').submit( function(e) {
	e.preventDefault();
	var id = $('#id').val();
	var type = $('#type').val();

	if (id == ""){
		alert("ID must not be empty!")
	}
	
	$.ajax({
		type:'POST',
		url:'/admin/problem?rejudge/type?'+type+'/id?'+id,
		data:$(this).serialize(),
		error:function(response){
			var json = eval('('+response.responseText+')');
			if(json.pid != null) {
				alert(json.pid);
			}
		},
		success:function(response){
			alert("Rejudge Complete")
			window.location.reload();
		}
	});
});
</script>
{{end}}