<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	{* if结构 *}
	{{if .IsDisplay}}
    	<em>{{.Content}}</em>
    {{else}}
    	<em>{{.Content2}}</em>
    {{end}}

	{* 遍历结构 *}
    {{range .Users}}
		{* len是外部的  所以用$ *}
    	{{.Username}} {{$.len}}<br>
    {{end}}
</body>
</html>