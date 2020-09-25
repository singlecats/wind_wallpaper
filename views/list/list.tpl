<html>
<ul>
    {{range .Str}}
        <li>
            <img src="{{.src}}" alt="">
            <label for="">{{.url}}</label>
        </li>
    {{end}}
</ul>
<label for="">{{.Page}}</label>
</html>
