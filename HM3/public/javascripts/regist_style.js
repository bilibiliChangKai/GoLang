$(document).ready(function() {
	var warning_string = {
		"name": "error:用户名为6~18位英文字母、数字或下划线，必须以英文字母开头！",
		"password1": "error:密码为6~12位数字、大小写字母、中划线、下划线！",
		"password2": "error:密码为6~12位数字、大小写字母、中划线、下划线！",
		"number": "error:学号为8位数字，不能以0开头！",
		"phone": "error:电话为11位数字，不能以0开头！",
		"email": "error:邮箱输入错误！",
		"repead": "error:两次密码不对！"
	};

	var data_string = {
		"name": "",
		"password1": "",
		"password2": "",
		"number": "",
		"phone": "",
		"email": "",

		"test_name": /^[a-zA-Z][a-zA-Z0-9_]{5,17}$/,
		"test_password1": /^[0-9a-zA-Z_\-]{6,12}$/,
		"test_password2": /^[0-9a-zA-Z_\-]{6,12}$/,
		"test_number": /^[1-9][0-9]{7}$/,
		"test_phone": /^[1-9][0-9]{10}$/,
		"test_email": /^[a-zA-Z_\-1-9][a-zA-Z_\-0-9]*@(([a-zA-Z_\-])+\.)+[a-zA-Z]{2,4}$/,

		clear: function() {
			$("input").attr("value", "");
			$("input").parent().next().removeClass('red green');
		},
		test_and_sent: function() {
			//if (this.local_test())
				this.online_test();
			//else alert("请按格式填好信息！");
		},
		local_test: function(name) { // 进行单一判断和commit整体判断
			if (arguments.length == 0) {
				for (var i = 0; i < 6; i++)
					if ($("input").eq(i).attr("value") == "")  return false;
				return  !($("input").parent().next().hasClass('red'));
			}

			if (name == "password2" && this['password1'] != this['password2']) {
				$("#" + name).parent().next().text(warning_string["repead"]).removeClass('green').addClass('red');
				return;
			}

			if (!this["test_" + name].test(this[name]))
				$("#" + name).parent().next().text(warning_string[name]).removeClass('green').addClass('red');
			else  $("#" + name).parent().next().text("pass").removeClass('red').addClass('green');
		},
		online_test: function() { // 进行网络判断
			var that = this;
			$.post("http://localhost:8000/test/regist", { // 先发送一个返回本地的url，由返回内容判断
					"name": data_string["name"],
					//"number": data_string["number"],
					//"phone": data_string["phone"],
					//"email": data_string["email"],
					//"password": data_string["password1"]
				}, function(data) { // 同步测试是否注册过
					try {
						console.log(data);
						if (data != "False") {
							var err = "已经有人注册过";
							if (data.indexOf("name") != -1)  err += "相同用户名！";
							if (data.indexOf("number") != -1)  err += "相同学号！";
							if (data.indexOf("phone") != -1)  err += "相同手机！";
							if (data.indexOf("email") != -1)  err += "相同邮箱！";
							throw err;
						}
						that.put_post('http://localhost:8000/sign_in_succeed', {
							"name": data_string["name"],
						});
					} catch(e) {alert(e); }
				}
			);
		},
		put_post: function(URL, PARAMS) { // 前面是URL地址，后面是传输过去的数据
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
		}
	};

	$("#reset").click(function(event) {
		data_string.clear();
	});

	$("input").blur(function(event) {
		data_string[$(this).attr("id")] = $(this).attr("value");
		if ($(this).attr("value") == "")
			$(this).parent().next().removeClass('red green');
		else  data_string.local_test($(this).attr("id"));
	});

	$("#commit").click(function(event) {
		try {
			data_string.test_and_sent();
		} catch(e) {alert(e); }
	});
});
