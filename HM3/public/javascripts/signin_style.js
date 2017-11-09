$(document).ready(function() {
	var	put_post = function(URL, PARAMS) { // 前面是URL地址，后面是传输过去的数据
		var temp_form = document.createElement("form");
		temp_form.action = URL;
		temp_form.target = "_self";
		temp_form.method = "post";
		temp_form.style.display = "none";
		for (var i in PARAMS) {
			var opt = document.createElement("textarea");
			opt.name = i;
			opt.value = PARAMS[i];
			temp_form.appendChild(opt);
		}
		document.body.appendChild(temp_form);
		temp_form.submit();
	};

	var	put_get = function(URL) { // 前面是URL地址，后面是传输过去的数据
		var temp_form = document.createElement("form");
		temp_form.action = URL;
		temp_form.target = "_self";
		temp_form.method = "get";
		temp_form.style.display = "none";
		document.body.appendChild(temp_form);
		temp_form.submit();
	};

	$("#regist").click(function(event) { // 请求跳转到注册页面
		put_get('http://localhost:8000/regist');
	});

	$("#signin").click(function(event) {
		$.post('http://localhost:8000/text/signin', {
				"name": $("#name").attr("value"),
				"password": $("#password").attr("value")
			},
			function(data, textStatus, xhr) {
				try {
					if (data == "false")  throw "错误的用户名或者密码!";
					put_post('sign_out_succeed', {
						"name": $("#name").attr("value")
					});
				} catch(e) {alert(e); }
			}
		);
	});
});
