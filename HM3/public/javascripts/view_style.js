$(document).ready(function() {
	var put_get = function(URL) { // 前面是URL地址，后面是传输过去的数据
		var temp_form = document.createElement("form");
		temp_form.action = URL;
		temp_form.target = "_self";
		temp_form.method = "get";
		temp_form.style.display = "none";
		document.body.appendChild(temp_form);
		temp_form.submit();
	}

	$("button").click(function(event) {
		put_get('http://localhost:8000/signin');
	});
});
