var USER = initUser();

$().ready(function() {
    const nav = $("#navbar-nav");
    if (USER) {
        const user = USER;
        nav.append("<a class=\"nav-item nav-link\" href='/users/" + user.id + "'>" + user.name + "</a>")
        nav.append("<a class=\"nav-item nav-link\" href='/settings'>设置</a>")
        nav.append("<a class=\"nav-item nav-link main-color\" href=\"/share\">分享好货</a>")
    } else {
        nav.append("<a class=\"nav-item nav-link\" href=\"/login\">登录 / 注册</a>")
    }
});

function initUser() {
    const content = localStorage.getItem("user");
    if (content) {
        return JSON.parse(content);
    }
    return null;
}

function $user() {
    if (USER) {
        return USER
    }
    window.location = "/login"
}

function isLogin() {
    return USER !== undefined && USER !== null
}

function updateLocalstorageUsername(name) {
    const content = localStorage.getItem("user");
    if (content) {
        let u = JSON.parse(content);
        u.name = name;
        localStorage.setItem("user", JSON.stringify(u))
    }
}

// api
const api = {
    favorite: function (data, onSuccess) {
        $.ajax({
            type: "PUT",
            url: "/api/v1/haohuos/" + data + "/favorite",
            headers: {"Authorization": "Bearer " + $user().token},
            contentType: "application/json",
            cache: false,
            success: onSuccess
        });
    },
    comment: function (id, data, success, error) {
        $.ajax({
            type: "PUT",
            url: "/api/v1/haohuos/" + id + "/comment",
            headers: {"Authorization": "Bearer " + $user()},
            contentType: "application/json",
            data: JSON.stringify(data),
            success: success,
            error: error
        });
    },
};

const onStarClick = function(id) {
     api.favorite(id, function (ret) {
        $("#" + id).html("<i class='fa fa-star'></i> " + ret.favorites);
        $.notify("已加入收藏", "success");
     });
};