var debug = require('debug')('signin:sign');

module.exports = function(db) {
	var connect = db.collection("signin_user"); // 链接mongodb数据库
	debug("connect setup by: ", connect);
	var hash_to_password = function(password) {
		var hash = "";
		for (var i = 0; i < password.length; i++) {
			var ascii = password[i].charCodeAt();
			hash += ascii * ascii;
		}
		return hash;
	};

	var store = {
		add_number: function(data, callback) { // 新加一个user
			data["password"] = hash_to_password(data["password"]);
			connect.insert(data, function(err, result) {
				if (err) {
					console.log("Error:" + err);
					return;
				}
				callback(result);
			});
		},
		find_one_type: function(target, type, callback) { // 通过一种数据寻找user
			var where = {};
			if (type == "password")  where[type] = hash_to_password(target);
			else  where[type] = target;
			connect.findOne(where, function(err, result) {
				if (err) {
					console.log("Error:" + err);
					return;
				}
				callback(result);
			});
		},	
		find_two_type: function(name, password, callback) { // 通过两种数据寻找user
			var where = {};
			where["name"] = name;
			where["password"] = hash_to_password(password);
			connect.findOne(where, function(err, result) {
				if (err) {
					console.log("Error:" + err);
					return;
				}
				callback(result);
			});
		}
	};	
	return store;
}