<!DOCTYPE html>
<html>
<head>
    <title>tpl</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf8">
</head>
   <p>param1 : {{.param1}}</p><br>
   <p>param2 : {{.param2}}</p><br>
   <p>param3 : {{.param3}}</p><br>

   <form enctype="multipart/form-data" method="post" action="./anothersave">
    <input type="file" name="uploadname" />
    <input type="submit">
   </form>
</html>