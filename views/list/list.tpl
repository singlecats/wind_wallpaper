<html>
<ul>
    {{range .Str}}
    <li>
        <img src="{{.src}}" alt="">
        <label for="">{{.url}}</label>
        <label for="">{{.link}}</label>
    </li>
    {{end}}
</ul>
<label for="">{{.Page}}</label>
</html>
